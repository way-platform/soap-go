package soap

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"time"
)

// Client represents a generic SOAP HTTP client that can handle any type of SOAP request.
// It works with Envelope types and provides a clean abstraction over HTTP transport.
type Client struct {
	config clientConfig
}

// ClientOption configures a Client using the functional options pattern.
// Can be used both during client creation and per-call.
type ClientOption func(*clientConfig)

// clientConfig holds the configuration for a Client.
type clientConfig struct {
	endpoint          string
	httpClient        *http.Client
	addXMLDeclaration bool
	maxRetries        int
	timeout           time.Duration
	interceptors      []func(http.RoundTripper) http.RoundTripper
	checkRetry        func(context.Context, error, *http.Request, *http.Response) bool
}

// newClientConfig creates a new clientConfig with default values.
func newClientConfig() clientConfig {
	return clientConfig{
		endpoint:          "",
		addXMLDeclaration: true,
		maxRetries:        3,
		timeout:           30 * time.Second,
	}
}

// WithEndpoint sets the default SOAP endpoint URL.
// This can be overridden per call if needed in the future.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *clientConfig) {
		c.endpoint = endpoint
	}
}

// WithHTTPClient sets the base HTTP client whose transport is used as the
// innermost layer of the transport chain (interceptors and retry wrap it).
// If the client has no Transport, [http.DefaultTransport] is used.
// The client's Timeout is ignored — use [WithTimeout] instead.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.httpClient = client
	}
}

// WithDebug enables debug output, dumping HTTP requests and responses to
// stderr. This is a convenience shorthand for:
//
//	soap.WithHTTPClient(&http.Client{
//	    Transport: &soap.DebugTransport{Enabled: &alwaysTrue, Next: http.DefaultTransport},
//	})
//
// For lazy evaluation (e.g. a --debug flag parsed after client construction),
// use [WithHTTPClient] with a [DebugTransport] directly.
func WithDebug() ClientOption {
	enabled := true
	return WithHTTPClient(&http.Client{
		Transport: &DebugTransport{
			Enabled: &enabled,
			Next:    http.DefaultTransport,
		},
	})
}

// WithXMLDeclaration controls whether XML declaration is automatically added to requests.
// Defaults to true. Set to false if your SOAP service doesn't expect or allow XML declarations.
func WithXMLDeclaration(include bool) ClientOption {
	return func(c *clientConfig) {
		c.addXMLDeclaration = include
	}
}

// WithTimeout sets the timeout for the SOAP client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *clientConfig) {
		c.timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries for a request.
// Defaults to 3.
func WithMaxRetries(retries int) ClientOption {
	return func(c *clientConfig) {
		c.maxRetries = retries
	}
}

// WithInterceptor adds a request interceptor for the Client.
func WithInterceptor(interceptor func(http.RoundTripper) http.RoundTripper) ClientOption {
	return func(c *clientConfig) {
		c.interceptors = append(c.interceptors, interceptor)
	}
}

// WithCheckRetry sets a custom retry check function.
// If not provided, uses DefaultCheckRetry for generic HTTP retry logic.
func WithCheckRetry(checkRetry func(context.Context, error, *http.Request, *http.Response) bool) ClientOption {
	return func(c *clientConfig) {
		c.checkRetry = checkRetry
	}
}

// with returns a new clientConfig with the given options applied.
// This enables per-request configuration overrides.
func (c clientConfig) with(opts ...ClientOption) clientConfig {
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// NewClient creates a new SOAP client with the specified options.
// Returns an error if the configuration is invalid.
func NewClient(opts ...ClientOption) (*Client, error) {
	config := newClientConfig()
	for _, opt := range opts {
		opt(&config)
	}
	return &Client{config: config}, nil
}

// Call executes a SOAP request with the provided action, envelope, and call-specific options.
func (c *Client) Call(
	ctx context.Context,
	action string,
	requestEnvelope *Envelope,
	opts ...ClientOption,
) (*Envelope, error) {
	config := c.config.with(opts...)
	xmlData, err := xml.Marshal(requestEnvelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}
	if config.addXMLDeclaration {
		xmlData = addXMLDeclaration(xmlData)
	}
	bodyReader := bytes.NewReader(xmlData)
	return c.doRequest(ctx, action, bodyReader, config)
}

// doRequest performs a single SOAP request.
func (c *Client) doRequest(
	ctx context.Context,
	action string,
	body io.Reader,
	config clientConfig,
) (*Envelope, error) {
	if config.endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	req, err := http.NewRequestWithContext(ctx, "POST", config.endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	if action != "" {
		req.Header.Set("SOAPAction", action)
	}
	req.Header.Set("User-Agent", getUserAgent())
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpClient := c.httpClient(config)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var env Envelope
	if xmlErr := xml.Unmarshal(respBody, &env); xmlErr != nil {
		// Not a valid SOAP envelope, but we might still have a useful HTTP error
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, &Error{
				StatusCode:   resp.StatusCode,
				ResponseBody: respBody,
			}
		}
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", xmlErr)
	}
	fault := checkForSOAPFault(&env)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 || fault != nil {
		return nil, &Error{
			StatusCode:   resp.StatusCode,
			ResponseBody: respBody,
			Envelope:     &env,
			Fault:        fault,
		}
	}
	return &env, nil
}

// httpClient creates a new HTTP client with the given configuration.
func (c *Client) httpClient(cfg clientConfig) *http.Client {
	// Use injected client's transport as the base, falling back to default.
	var transport http.RoundTripper
	if cfg.httpClient != nil {
		transport = cfg.httpClient.Transport
	}
	if transport == nil {
		transport = http.DefaultTransport
	}
	// Add middleware transport if middlewares are configured.
	if len(cfg.interceptors) > 0 {
		transport = &interceptorTransport{
			interceptors: cfg.interceptors,
			next:         transport,
		}
	}
	// Add retry transport if retry count > 0.
	if cfg.maxRetries > 0 {
		transport = &retryTransport{
			maxRetries:  cfg.maxRetries,
			next:        transport,
			shouldRetry: cfg.checkRetry,
		}
	}
	return &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}
}

// addXMLDeclaration adds an XML declaration to the beginning of XML data if it doesn't already have one.
func addXMLDeclaration(xmlData []byte) []byte {
	if len(xmlData) > 5 && string(xmlData[:5]) == "<?xml" {
		return xmlData
	}
	return append([]byte(xml.Header), xmlData...)
}

// checkForSOAPFault checks if the response envelope contains a SOAP fault.
func checkForSOAPFault(envelope *Envelope) *Fault {
	if envelope == nil || len(envelope.Body.Content) == 0 {
		return nil
	}
	var fault Fault
	if err := xml.Unmarshal(envelope.Body.Content, &fault); err != nil {
		return nil
	}
	if fault.XMLName.Local == "Fault" && fault.FaultCode != "" && fault.FaultString != "" {
		return &fault
	}
	return nil
}

func getUserAgent() string {
	userAgent := "WayPlatformSOAPGo"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		userAgent += "/" + info.Main.Version
	}
	return userAgent
}

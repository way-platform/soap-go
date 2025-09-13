package soap

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
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
	httpClient        http.Client
	endpoint          string
	debug             bool
	addXMLDeclaration bool
	maxRetries        int
	retryDelay        time.Duration
}

// newClientConfig creates a new clientConfig with default values.
func newClientConfig() clientConfig {
	return clientConfig{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		endpoint:          "",
		debug:             false,
		addXMLDeclaration: true,
		maxRetries:        3,
		retryDelay:        2 * time.Second,
	}
}

// WithEndpoint sets the default SOAP endpoint URL.
// This can be overridden per call if needed in the future.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *clientConfig) {
		c.endpoint = endpoint
	}
}

// WithDebug enables debug mode, which will dump HTTP requests and responses to stderr.
// This works the same as the debug mode in the soapcall package.
func WithDebug(debug bool) ClientOption {
	return func(c *clientConfig) {
		c.debug = debug
	}
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
		c.httpClient.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries for a request.
// Defaults to 3.
func WithMaxRetries(retries int) ClientOption {
	return func(c *clientConfig) {
		c.maxRetries = retries
	}
}

// WithRetryDelay sets the initial delay for the exponential backoff strategy.
// Defaults to 2 seconds.
func WithRetryDelay(delay time.Duration) ClientOption {
	return func(c *clientConfig) {
		c.retryDelay = delay
	}
}

// NewClient creates a new SOAP client with the specified options.
// Returns an error if the configuration is invalid.
func NewClient(opts ...ClientOption) (*Client, error) {
	config := newClientConfig()
	for _, opt := range opts {
		opt(&config)
	}
	return &Client{
			config: config,
		},
		nil
}

// Call executes a SOAP request with the provided action, envelope, and call-specific options.
func (c *Client) Call(ctx context.Context, action string, requestEnvelope *Envelope, opts ...ClientOption) (*Envelope, error) {
	config := c.config
	for _, opt := range opts {
		opt(&config)
	}
	xmlData, err := xml.Marshal(requestEnvelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}
	if config.addXMLDeclaration {
		xmlData = addXMLDeclaration(xmlData)
	}
	var lastErr error
	var responseEnvelope *Envelope
	for i := 0; i <= config.maxRetries; i++ {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		bodyReader := bytes.NewReader(xmlData)
		var resp *http.Response
		responseEnvelope, resp, err = c.doRequest(ctx, action, bodyReader, &config)
		if err == nil {
			return responseEnvelope, nil
		}
		if !checkRetry(err) {
			return nil, err
		}
		lastErr = err
		wait := backoff(config.retryDelay, config.retryDelay*10, i, resp)
		if err := sleepWithContext(ctx, wait); err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("request failed after %d retries: %w", config.maxRetries, lastErr)
}

// doRequest performs a single SOAP request.
func (c *Client) doRequest(ctx context.Context, action string, body io.Reader, config *clientConfig) (*Envelope, *http.Response, error) {
	if config.endpoint == "" {
		return nil, nil, fmt.Errorf("endpoint is required")
	}
	req, err := http.NewRequestWithContext(ctx, "POST", config.endpoint, body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	if action != "" {
		req.Header.Set("SOAPAction", action)
	}
	req.Header.Set("User-Agent", getUserAgent())
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	if config.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to dump request for debug: %w", err)
		}
		var output bytes.Buffer
		output.Grow(len(dump) * 2)
		for _, line := range bytes.Split(dump, []byte("\n")) {
			output.WriteString("> ")
			output.Write(line)
			output.WriteByte('\n')
		}
		_, _ = os.Stderr.Write(output.Bytes())
	}
	resp, err := config.httpClient.Do(req)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()
	if config.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, resp, fmt.Errorf("failed to dump response for debug: %w", err)
		}
		var output bytes.Buffer
		output.Grow(len(dump) * 2)
		for _, line := range bytes.Split(dump, []byte("\n")) {
			output.WriteString("< ")
			output.Write(line)
			output.WriteByte('\n')
		}
		_, _ = os.Stderr.Write(output.Bytes())
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to read response body: %w", err)
	}
	var env Envelope
	if xmlErr := xml.Unmarshal(respBody, &env); xmlErr != nil {
		// Not a valid SOAP envelope, but we might still have a useful HTTP error
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, resp, &Error{
				StatusCode:   resp.StatusCode,
				ResponseBody: respBody,
			}
		}
		return nil, resp, fmt.Errorf("failed to unmarshal SOAP response: %w", xmlErr)
	}
	fault := checkForSOAPFault(&env)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 || fault != nil {
		return &env, resp, &Error{
			StatusCode:   resp.StatusCode,
			ResponseBody: respBody,
			Envelope:     &env,
			Fault:        fault,
		}
	}
	return &env, resp, nil
}

// checkRetry determines if an error should be retried.
func checkRetry(err error) bool {
	if err == nil {
		return false
	}
	var soapErr *Error
	if errors.As(err, &soapErr) {
		// Retry on 5xx server errors
		if soapErr.StatusCode >= 500 && soapErr.StatusCode <= 599 {
			return true
		}
		// Retry on specific 4xx codes that indicate temporary issues
		if soapErr.StatusCode == http.StatusTooManyRequests || // 429
			soapErr.StatusCode == 420 { // 420 Enhance Your Calm
			return true
		}
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

// backoff calculates the time to wait before the next retry, with exponential backoff and jitter.
func backoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
				if seconds, err := strconv.Atoi(retryAfter); err == nil {
					return time.Duration(seconds) * time.Second
				}
				if retryTime, err := time.Parse(time.RFC1123, retryAfter); err == nil {
					return time.Until(retryTime)
				}
			}
		}
	}
	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
	return sleep + jitter
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

// sleepWithContext sleeps for the specified duration, but can be interrupted by context cancellation.
func sleepWithContext(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

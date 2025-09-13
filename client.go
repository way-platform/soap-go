package soap

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
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
	timeout           time.Duration
	debug             bool
	addXMLDeclaration bool
}

// newClientConfig creates a new clientConfig with default values.
func newClientConfig() clientConfig {
	return clientConfig{
		httpClient:        http.Client{},
		endpoint:          "",
		timeout:           10 * time.Second,
		debug:             false,
		addXMLDeclaration: true,
	}
}

// WithHTTPClient sets a custom HTTP client for the SOAP client.
// If not provided, http.DefaultClient is used.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.httpClient = *client
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
		c.timeout = timeout
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
	}, nil
}

// Call executes a SOAP request with the provided action and envelope.
// The action parameter is used to set the SOAPAction header.
// Call-specific options can override client defaults.
func (c *Client) Call(ctx context.Context, action string, requestEnvelope *Envelope, opts ...ClientOption) (*Envelope, error) {
	config := c.config
	for _, opt := range opts {
		opt(&config)
	}
	if config.endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	xmlData, err := xml.Marshal(requestEnvelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}
	if config.addXMLDeclaration {
		xmlData = addXMLDeclaration(xmlData)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", config.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	if action != "" {
		req.Header.Set("SOAPAction", action)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	if config.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, fmt.Errorf("failed to dump request for debug: %w", err)
		}
		var output bytes.Buffer
		output.Grow(2 * len(dump))
		for line := range bytes.Lines(dump) {
			output.WriteString("> ")
			output.Write(line)
		}
		output.WriteByte('\n')
		_, _ = os.Stderr.Write(output.Bytes())
	}
	resp, err := config.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()
	if config.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, fmt.Errorf("failed to dump response for debug: %w", err)
		}
		var output bytes.Buffer
		output.Grow(2 * len(dump))
		for line := range bytes.Lines(dump) {
			output.WriteString("< ")
			output.Write(line)
		}
		output.WriteByte('\n')
		_, _ = os.Stderr.Write(output.Bytes())
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	// Note: SOAP faults are typically returned with HTTP 200 or 500, but we let the caller
	// handle SOAP faults by examining the returned envelope
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}
	var responseEnvelope Envelope
	if err := xml.Unmarshal(respBody, &responseEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}
	if fault := checkForSOAPFault(&responseEnvelope); fault != nil {
		return &responseEnvelope, fault
	}
	return &responseEnvelope, nil
}

// addXMLDeclaration adds an XML declaration to the beginning of XML data if it doesn't already have one.
func addXMLDeclaration(xmlData []byte) []byte {
	if len(xmlData) > 5 && string(xmlData[:5]) == "<?xml" {
		return xmlData
	}
	return append([]byte(xml.Header), xmlData...)
}

// checkForSOAPFault checks if the response envelope contains a SOAP fault.
// Returns the fault as an error if found, nil otherwise.
func checkForSOAPFault(envelope *Envelope) error {
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

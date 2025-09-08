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
)

// Client represents a generic SOAP HTTP client that can handle any type of SOAP request.
// It works with Envelope types and provides a clean abstraction over HTTP transport.
type Client struct {
	httpClient     *http.Client
	endpoint       string
	debug          bool
	xmlDeclaration bool
}

// ClientOption configures a Client using the functional options pattern.
// Can be used both during client creation and per-call.
type ClientOption func(*clientConfig)

// clientConfig holds the configuration for a Client.
type clientConfig struct {
	httpClient     *http.Client
	endpoint       string
	debug          bool
	xmlDeclaration bool
	headers        map[string]string
}

// WithHTTPClient sets a custom HTTP client for the SOAP client.
// If not provided, http.DefaultClient is used.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.httpClient = client
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
		c.xmlDeclaration = include
	}
}

// WithHeader sets a custom header for requests.
// Can be used for SOAPAction or any other custom headers.
func WithHeader(key, value string) ClientOption {
	return func(c *clientConfig) {
		if c.headers == nil {
			c.headers = make(map[string]string)
		}
		c.headers[key] = value
	}
}

// WithSOAPAction is a convenience function for setting the SOAPAction header.
func WithSOAPAction(action string) ClientOption {
	return WithHeader("SOAPAction", action)
}

// NewClient creates a new SOAP client with the specified options.
// Returns an error if the configuration is invalid.
func NewClient(opts ...ClientOption) (*Client, error) {
	config := &clientConfig{
		httpClient:     http.DefaultClient,
		endpoint:       "",
		debug:          false,
		xmlDeclaration: true, // Default to including XML declaration
		headers:        make(map[string]string),
	}

	for _, opt := range opts {
		opt(config)
	}

	return &Client{
		httpClient:     config.httpClient,
		endpoint:       config.endpoint,
		debug:          config.debug,
		xmlDeclaration: config.xmlDeclaration,
	}, nil
}

// Call executes a SOAP request with the provided action and envelope.
// The action parameter is used to set the SOAPAction header.
// Call-specific options can override client defaults.
func (c *Client) Call(ctx context.Context, action string, requestEnvelope *Envelope, opts ...ClientOption) (*Envelope, error) {
	// Copy client configuration for this call
	config := &clientConfig{
		httpClient:     c.httpClient,
		endpoint:       c.endpoint,
		debug:          c.debug,
		xmlDeclaration: c.xmlDeclaration,
		headers:        make(map[string]string),
	}

	// Set SOAPAction from the action parameter
	if action != "" {
		config.headers["SOAPAction"] = action
	}

	// Apply call-specific options
	for _, opt := range opts {
		opt(config)
	}

	// Validate endpoint
	if config.endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}

	// Marshal the request envelope to XML
	xmlData, err := xml.Marshal(requestEnvelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Add XML declaration if enabled
	if config.xmlDeclaration {
		xmlData = addXMLDeclaration(xmlData)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", config.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set default headers
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	// Set headers from configuration
	for key, value := range config.headers {
		req.Header.Set(key, value)
	}

	// Debug: dump request
	if config.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, fmt.Errorf("failed to dump request for debug: %w", err)
		}
		fmt.Fprintf(os.Stderr, "=== SOAP Request ===\n%s\n", dump)
	}

	// Execute request
	resp, err := config.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Debug: dump response
	if config.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, fmt.Errorf("failed to dump response for debug: %w", err)
		}
		fmt.Fprintf(os.Stderr, "=== SOAP Response ===\n%s\n", dump)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	// Note: SOAP faults are typically returned with HTTP 200 or 500, but we let the caller
	// handle SOAP faults by examining the returned envelope
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response envelope
	responseEnvelope, err := parseSOAPResponse(respBody)
	if err != nil {
		return nil, err
	}

	// Check for SOAP faults in the response
	if fault := checkForSOAPFault(responseEnvelope); fault != nil {
		return responseEnvelope, fault
	}

	return responseEnvelope, nil
}

// AddXMLDeclaration adds an XML declaration to the beginning of XML data if it doesn't already have one.
// This is exported so it can be used by other packages like the CLI.
func AddXMLDeclaration(xmlData []byte) []byte {
	// Check if XML declaration is already present
	if len(xmlData) > 5 && string(xmlData[:5]) == "<?xml" {
		return xmlData
	}

	// Add standard XML declaration
	return append([]byte(xml.Header), xmlData...)
}

// addXMLDeclaration is kept for internal use to maintain the same interface.
func addXMLDeclaration(xmlData []byte) []byte {
	return AddXMLDeclaration(xmlData)
}

// parseSOAPResponse parses a SOAP response, handling both prefixed and non-prefixed namespace formats.
func parseSOAPResponse(respBody []byte) (*Envelope, error) {
	var responseEnvelope Envelope
	if err := xml.Unmarshal(respBody, &responseEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}
	return &responseEnvelope, nil
}

// checkForSOAPFault checks if the response envelope contains a SOAP fault.
// Returns the fault as an error if found, nil otherwise.
func checkForSOAPFault(envelope *Envelope) error {
	if envelope == nil || len(envelope.Body.Content) == 0 {
		return nil
	}

	// Try to unmarshal the body content as a SOAP fault
	var fault Fault
	if err := xml.Unmarshal(envelope.Body.Content, &fault); err != nil {
		// Not a fault or unmarshaling failed - not necessarily an error
		return nil
	}

	// Check if this is actually a fault by verifying required fields
	if fault.XMLName.Local == "Fault" && fault.FaultCode != "" && fault.FaultString != "" {
		return &fault
	}

	return nil
}

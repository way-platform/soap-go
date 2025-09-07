package soapcall

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

// Client represents a SOAP HTTP client.
type Client struct {
	config     Config
	httpClient *http.Client
	endpoint   string
	headers    map[string]string
}

// Config holds configuration for the SOAP client.
type Config struct {
	Endpoint   string
	Timeout    time.Duration
	Headers    map[string]string
	Insecure   bool
	SOAPAction string
	Debug      bool
}

// NewClient creates a new SOAP client with the given configuration.
func NewClient(config Config) *Client {
	transport := &http.Transport{}

	if config.Insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	headers := make(map[string]string)

	// Set default headers
	headers["Content-Type"] = "text/xml; charset=utf-8"

	// Set SOAPAction header if provided
	if config.SOAPAction != "" {
		headers["SOAPAction"] = config.SOAPAction
	}

	// Add custom headers
	for k, v := range config.Headers {
		headers[k] = v
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
		endpoint:   config.Endpoint,
		headers:    headers,
	}
}

// Call makes a SOAP request with the provided XML payload.
func (c *Client) Call(ctx context.Context, xmlPayload []byte) ([]byte, error) {
	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Execute request
	if c.config.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(os.Stderr, "%s", dump)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	if c.config.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(os.Stderr, "%s", dump)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// ParseHeaders parses header strings in the format "key:value" or "key=value".
func ParseHeaders(headerStrings []string) (map[string]string, error) {
	headers := make(map[string]string)

	for _, headerStr := range headerStrings {
		var key, value string

		// Try colon separator first
		if parts := strings.SplitN(headerStr, ":", 2); len(parts) == 2 {
			key = strings.TrimSpace(parts[0])
			value = strings.TrimSpace(parts[1])
		} else if parts := strings.SplitN(headerStr, "=", 2); len(parts) == 2 {
			// Try equals separator
			key = strings.TrimSpace(parts[0])
			value = strings.TrimSpace(parts[1])
		} else {
			return nil, fmt.Errorf("invalid header format: %s (expected key:value or key=value)", headerStr)
		}

		if key == "" {
			return nil, fmt.Errorf("empty header key in: %s", headerStr)
		}

		headers[key] = value
	}

	return headers, nil
}

// AddXMLDeclaration adds an XML declaration to the beginning of XML data if it doesn't already have one.
func AddXMLDeclaration(xmlData []byte) []byte {
	// Check if XML declaration is already present
	if len(xmlData) > 5 && string(xmlData[:5]) == "<?xml" {
		return xmlData
	}

	// Add standard XML declaration
	return append([]byte(xml.Header), xmlData...)
}

// EnsureXMLDeclaration ensures XML data has proper XML declaration with UTF-8 encoding.
// This is useful for SOAP services that require explicit encoding declaration.
func EnsureXMLDeclaration(xmlData []byte) []byte {
	return EnsureXMLDeclarationWithEncoding(xmlData, "UTF-8")
}

// EnsureXMLDeclarationWithEncoding ensures XML data has proper XML declaration with specified encoding.
func EnsureXMLDeclarationWithEncoding(xmlData []byte, encoding string) []byte {
	// Check if XML declaration is already present
	if len(xmlData) > 5 && string(xmlData[:5]) == "<?xml" {
		return xmlData
	}

	// Create custom XML declaration with specified encoding
	declaration := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"%s\"?>\n", encoding)
	return append([]byte(declaration), xmlData...)
}

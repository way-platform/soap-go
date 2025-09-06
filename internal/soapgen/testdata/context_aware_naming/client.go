package context_aware_naming

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ClientOption configures a Client.
type ClientOption func(*clientConfig)

// clientConfig holds the configuration for a Client.
type clientConfig struct {
	httpClient *http.Client
	endpoint   string
}

// WithHTTPClient sets a custom HTTP client for the SOAP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.httpClient = client
	}
}

// WithEndpoint sets the SOAP endpoint URL.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *clientConfig) {
		c.endpoint = endpoint
	}
}

// Client is a SOAP client for this service.
type Client struct {
	httpClient *http.Client
	endpoint   string
}

// NewClient creates a new SOAP client.
func NewClient(opts ...ClientOption) (*Client, error) {
	config := &clientConfig{
		httpClient: http.DefaultClient,
		endpoint:   "http://example.com/context-naming",
	}

	for _, opt := range opts {
		opt(config)
	}

	// Validate that we have an endpoint
	if config.endpoint == "" {
		return nil, fmt.Errorf("SOAP endpoint is required")
	}

	return &Client{
		httpClient: config.httpClient,
		endpoint:   config.endpoint,
	}, nil
}

// ProcessUserData executes the ProcessUserData SOAP operation.
func (c *Client) ProcessUserData(ctx context.Context, req *UserData) (*UserData, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: reqXML},
	}

	// Marshal envelope to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "urn:ProcessUserData")

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response
	var respEnvelope soapEnvelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	// Extract response from SOAP body
	var result UserData
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessRequest executes the ProcessRequest SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *ProcessRequest) (*ProcessRequest, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: reqXML},
	}

	// Marshal envelope to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "urn:ProcessRequest")

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response
	var respEnvelope soapEnvelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	// Extract response from SOAP body
	var result ProcessRequest
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// GetSystemInfo executes the GetSystemInfo SOAP operation.
func (c *Client) GetSystemInfo(ctx context.Context, req *SystemInfo) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: reqXML},
	}

	// Marshal envelope to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "urn:GetSystemInfo")

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response
	var respEnvelope soapEnvelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	// Extract response from SOAP body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// UpdateUserData executes the UpdateUserData SOAP operation.
func (c *Client) UpdateUserData(ctx context.Context, req *UserData) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: reqXML},
	}

	// Marshal envelope to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "urn:UpdateUserData")

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response
	var respEnvelope soapEnvelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	// Extract response from SOAP body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ValidateProcessRequest executes the ValidateProcessRequest SOAP operation.
func (c *Client) ValidateProcessRequest(ctx context.Context, req *ProcessRequest) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: reqXML},
	}

	// Marshal envelope to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SOAP envelope: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(xmlData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "urn:ValidateProcessRequest")

	// Execute request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response
	var respEnvelope soapEnvelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}

	// Extract response from SOAP body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// soapEnvelope represents a SOAP envelope.
type soapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    soapBody `xml:"soap:Body"`
}

// soapBody represents a SOAP body.
type soapBody struct {
	Content []byte `xml:",innerxml"`
}

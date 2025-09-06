package numbered_suffix_collisions

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
		endpoint:   "http://example.com/numbered-suffix-collisions",
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

// ProcessRequest executes the ProcessRequest SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *Request) (*interface{}, error) {
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
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessRequestLower executes the ProcessRequestLower SOAP operation.
func (c *Client) ProcessRequestLower(ctx context.Context, req *Request) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessRequestLower")

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

// ProcessRequestUpper executes the ProcessRequestUpper SOAP operation.
func (c *Client) ProcessRequestUpper(ctx context.Context, req *REQUEST) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessRequestUpper")

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

// ProcessData executes the ProcessData SOAP operation.
func (c *Client) ProcessData(ctx context.Context, req *Data) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessData")

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

// ProcessDataLower executes the ProcessDataLower SOAP operation.
func (c *Client) ProcessDataLower(ctx context.Context, req *Data) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessDataLower")

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

// ProcessDataUpper executes the ProcessDataUpper SOAP operation.
func (c *Client) ProcessDataUpper(ctx context.Context, req *DATA) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessDataUpper")

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

// ProcessExtremeCase executes the ProcessExtremeCase SOAP operation.
func (c *Client) ProcessExtremeCase(ctx context.Context, req *ExtremeCaseElement) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ProcessExtremeCase")

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

// HandleRequest executes the HandleRequest SOAP operation.
func (c *Client) HandleRequest(ctx context.Context, req *Request) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:HandleRequest")

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

// ValidateRequest executes the ValidateRequest SOAP operation.
func (c *Client) ValidateRequest(ctx context.Context, req *Request) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:ValidateRequest")

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

// SubmitRequest executes the SubmitRequest SOAP operation.
func (c *Client) SubmitRequest(ctx context.Context, req *Request) (*interface{}, error) {
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
	httpReq.Header.Set("SOAPAction", "urn:SubmitRequest")

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

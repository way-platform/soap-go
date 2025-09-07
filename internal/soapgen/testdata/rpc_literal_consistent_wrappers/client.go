package rpc_literal_consistent_wrappers

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	soap "github.com/way-platform/soap-go"
	"io"
	"net/http"
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
		endpoint:   "http://example.com/rpc-test",
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

// Authenticate executes the Authenticate SOAP operation.
func (c *Client) Authenticate(ctx context.Context, req *AuthenticateWrapper) (*AuthenticateResponseWrapper, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	reqEnvelope := &soap.Envelope{
		XMLNS: soap.Namespace,
		Body:  soap.Body{Content: reqXML},
	}
	xmlData, err := xml.Marshal(&reqEnvelope)
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
	httpReq.Header.Set("SOAPAction", "http://example.com/rpc-literal-test/Authenticate")

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

	// Unmarshal SOAP envelope
	var respEnvelope soap.Envelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP envelope: %w", err)
	}

	// Unmarshal response body
	var result AuthenticateResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// FetchData executes the FetchData SOAP operation.
func (c *Client) FetchData(ctx context.Context, req *FetchDataWrapper) (*FetchDataResponseWrapper, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope
	reqEnvelope := &soap.Envelope{
		XMLNS: soap.Namespace,
		Body:  soap.Body{Content: reqXML},
	}
	xmlData, err := xml.Marshal(&reqEnvelope)
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
	httpReq.Header.Set("SOAPAction", "http://example.com/rpc-literal-test/FetchData")

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

	// Unmarshal SOAP envelope
	var respEnvelope soap.Envelope
	if err := xml.Unmarshal(respBody, &respEnvelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP envelope: %w", err)
	}

	// Unmarshal response body
	var result FetchDataResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

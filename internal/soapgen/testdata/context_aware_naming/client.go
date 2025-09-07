package context_aware_naming

import (
	"context"
	"encoding/xml"
	"fmt"
	soap "github.com/way-platform/soap-go"
)

// ClientOption configures a Client.
type ClientOption = soap.ClientOption

// Client is a SOAP client for this service.
type Client struct {
	*soap.Client
}

// NewClient creates a new SOAP client.
func NewClient(opts ...ClientOption) (*Client, error) {
	// Prepend default endpoint from WSDL to user options
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/context-naming"),
	}, opts...)

	// Create underlying SOAP client
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}

	return &Client{
		Client: soapClient,
	}, nil
}

// ProcessUserData executes the ProcessUserData SOAP operation.
func (c *Client) ProcessUserData(ctx context.Context, req *UserDataWrapper) (*UserDataWrapper, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessUserData", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result UserDataWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessRequest executes the ProcessRequest SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *ProcessRequestWrapper) (*ProcessRequestWrapper, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result ProcessRequestWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// GetSystemInfo executes the GetSystemInfo SOAP operation.
func (c *Client) GetSystemInfo(ctx context.Context, req *SystemInfoWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:GetSystemInfo", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// UpdateUserData executes the UpdateUserData SOAP operation.
func (c *Client) UpdateUserData(ctx context.Context, req *UserDataWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:UpdateUserData", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ValidateProcessRequest executes the ValidateProcessRequest SOAP operation.
func (c *Client) ValidateProcessRequest(ctx context.Context, req *ProcessRequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ValidateProcessRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

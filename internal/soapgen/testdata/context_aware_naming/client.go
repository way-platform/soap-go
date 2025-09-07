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
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/context-naming"),
	}, opts...)
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
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	respEnvelope, err := c.Call(ctx, "urn:ProcessUserData", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result UserDataWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

// ProcessRequest executes the ProcessRequest SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *ProcessRequestWrapper) (*ProcessRequestWrapper, error) {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	respEnvelope, err := c.Call(ctx, "urn:ProcessRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result ProcessRequestWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

// GetSystemInfo executes the GetSystemInfo one-way SOAP operation.
func (c *Client) GetSystemInfo(ctx context.Context, req *SystemInfoWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:GetSystemInfo", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// UpdateUserData executes the UpdateUserData one-way SOAP operation.
func (c *Client) UpdateUserData(ctx context.Context, req *UserDataWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:UpdateUserData", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ValidateProcessRequest executes the ValidateProcessRequest one-way SOAP operation.
func (c *Client) ValidateProcessRequest(ctx context.Context, req *ProcessRequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ValidateProcessRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

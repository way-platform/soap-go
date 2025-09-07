package comprehensive_wrapper_naming

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
		soap.WithEndpoint("http://example.com/test"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// Login executes the Login SOAP operation.
func (c *Client) Login(ctx context.Context, req *LoginWrapper) (*LoginResponseWrapper, error) {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	respEnvelope, err := c.Call(ctx, "http://example.com/test/Login", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result LoginResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

// GetUserInfo executes the GetUserInfo one-way SOAP operation.
func (c *Client) GetUserInfo(ctx context.Context, req *GetUserInfoWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "http://example.com/test/GetUserInfo", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

package inline_enums

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
		soap.WithEndpoint("http://example.com/inlineenums"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// GetServerProperties executes the GetServerProperties SOAP operation.
func (c *Client) GetServerProperties(ctx context.Context, req *GetServerPropertiesRequestWrapper, opts ...ClientOption) (*GetServerPropertiesResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://example.com/inlineenums/GetServerProperties", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result GetServerPropertiesResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

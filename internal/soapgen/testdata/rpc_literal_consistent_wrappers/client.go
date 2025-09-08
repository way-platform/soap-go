package rpc_literal_consistent_wrappers

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
		soap.WithEndpoint("http://example.com/rpc-test"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// Authenticate executes the Authenticate SOAP operation.
func (c *Client) Authenticate(ctx context.Context, req *AuthenticateWrapper, opts ...ClientOption) (*AuthenticateResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://example.com/rpc-literal-test/Authenticate", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result AuthenticateResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

// FetchData executes the FetchData SOAP operation.
func (c *Client) FetchData(ctx context.Context, req *FetchDataWrapper, opts ...ClientOption) (*FetchDataResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://example.com/rpc-literal-test/FetchData", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result FetchDataResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

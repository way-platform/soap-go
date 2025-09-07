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
	// Prepend default endpoint from WSDL to user options
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/rpc-test"),
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

// Authenticate executes the Authenticate SOAP operation.
func (c *Client) Authenticate(ctx context.Context, req *AuthenticateWrapper) (*AuthenticateResponseWrapper, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "http://example.com/rpc-literal-test/Authenticate", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
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

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "http://example.com/rpc-literal-test/FetchData", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result FetchDataResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

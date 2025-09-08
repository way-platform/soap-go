package case_insensitive_collisions

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
		soap.WithEndpoint("http://example.com/collisions"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// TestUserRequest executes the TestUserRequest one-way SOAP operation.
func (c *Client) TestUserRequest(ctx context.Context, req *UserRequestWrapper, opts ...ClientOption) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:TestUserRequest", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// TestUserRequestLower executes the TestUserRequestLower one-way SOAP operation.
func (c *Client) TestUserRequestLower(ctx context.Context, req *UserRequestWrapper, opts ...ClientOption) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:TestUserRequestLower", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// TestGetFleetResponse executes the TestGetFleetResponse one-way SOAP operation.
func (c *Client) TestGetFleetResponse(ctx context.Context, req *GetFleetResponseWrapper, opts ...ClientOption) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:TestGetFleetResponse", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// TestGetFleetResponseLower executes the TestGetFleetResponseLower one-way SOAP operation.
func (c *Client) TestGetFleetResponseLower(ctx context.Context, req *GetFleetResponseWrapper, opts ...ClientOption) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:TestGetFleetResponseLower", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

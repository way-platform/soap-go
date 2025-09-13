package simple_content_extension

import (
	"context"
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
		soap.WithEndpoint("http://localhost/TestService"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// TestOperation executes the TestOperation one-way SOAP operation.
func (c *Client) TestOperation(ctx context.Context, req *StatesContainerWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "http://tempuri.org/TestOperation", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

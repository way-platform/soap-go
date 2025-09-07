package attribute_element_collisions

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
		soap.WithEndpoint("http://example.com/field-collisions"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// ProcessDownloadRequest executes the ProcessDownloadRequest one-way SOAP operation.
func (c *Client) ProcessDownloadRequest(ctx context.Context, req *DownloadRequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessDownloadRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessConfigData executes the ProcessConfigData one-way SOAP operation.
func (c *Client) ProcessConfigData(ctx context.Context, req *ConfigDataWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessConfigData", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

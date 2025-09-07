package numbered_suffix_collisions

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
		soap.WithEndpoint("http://example.com/numbered-suffix-collisions"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// ProcessRequest executes the ProcessRequest one-way SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *RequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessRequestLower executes the ProcessRequestLower one-way SOAP operation.
func (c *Client) ProcessRequestLower(ctx context.Context, req *RequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessRequestLower", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessRequestUpper executes the ProcessRequestUpper one-way SOAP operation.
func (c *Client) ProcessRequestUpper(ctx context.Context, req *REQUESTWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessRequestUpper", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessData executes the ProcessData one-way SOAP operation.
func (c *Client) ProcessData(ctx context.Context, req *DataWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessData", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessDataLower executes the ProcessDataLower one-way SOAP operation.
func (c *Client) ProcessDataLower(ctx context.Context, req *DataWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessDataLower", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessDataUpper executes the ProcessDataUpper one-way SOAP operation.
func (c *Client) ProcessDataUpper(ctx context.Context, req *DATAWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessDataUpper", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessExtremeCase executes the ProcessExtremeCase one-way SOAP operation.
func (c *Client) ProcessExtremeCase(ctx context.Context, req *ExtremeCaseElementWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessExtremeCase", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// HandleRequest executes the HandleRequest one-way SOAP operation.
func (c *Client) HandleRequest(ctx context.Context, req *RequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:HandleRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ValidateRequest executes the ValidateRequest one-way SOAP operation.
func (c *Client) ValidateRequest(ctx context.Context, req *RequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ValidateRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// SubmitRequest executes the SubmitRequest one-way SOAP operation.
func (c *Client) SubmitRequest(ctx context.Context, req *RequestWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:SubmitRequest", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

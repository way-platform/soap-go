package numbered_suffix_collisions

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
func (c *Client) ProcessRequest(ctx context.Context, req *RequestWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessRequest", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessRequestLower executes the ProcessRequestLower one-way SOAP operation.
func (c *Client) ProcessRequestLower(ctx context.Context, req *RequestWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessRequestLower", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessRequestUpper executes the ProcessRequestUpper one-way SOAP operation.
func (c *Client) ProcessRequestUpper(ctx context.Context, req *REQUESTWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessRequestUpper", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessData executes the ProcessData one-way SOAP operation.
func (c *Client) ProcessData(ctx context.Context, req *DataWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessData", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessDataLower executes the ProcessDataLower one-way SOAP operation.
func (c *Client) ProcessDataLower(ctx context.Context, req *DataWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessDataLower", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessDataUpper executes the ProcessDataUpper one-way SOAP operation.
func (c *Client) ProcessDataUpper(ctx context.Context, req *DATAWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessDataUpper", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessExtremeCase executes the ProcessExtremeCase one-way SOAP operation.
func (c *Client) ProcessExtremeCase(ctx context.Context, req *ExtremeCaseElementWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ProcessExtremeCase", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// HandleRequest executes the HandleRequest one-way SOAP operation.
func (c *Client) HandleRequest(ctx context.Context, req *RequestWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:HandleRequest", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ValidateRequest executes the ValidateRequest one-way SOAP operation.
func (c *Client) ValidateRequest(ctx context.Context, req *RequestWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:ValidateRequest", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// SubmitRequest executes the SubmitRequest one-way SOAP operation.
func (c *Client) SubmitRequest(ctx context.Context, req *RequestWrapper, opts ...ClientOption) error {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	_, err = c.Call(ctx, "urn:SubmitRequest", reqEnvelope, opts...)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

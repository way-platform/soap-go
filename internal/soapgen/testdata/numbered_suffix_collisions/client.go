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
	// Prepend default endpoint from WSDL to user options
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/numbered-suffix-collisions"),
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

// ProcessRequest executes the ProcessRequest SOAP operation.
func (c *Client) ProcessRequest(ctx context.Context, req *RequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessRequestLower executes the ProcessRequestLower SOAP operation.
func (c *Client) ProcessRequestLower(ctx context.Context, req *RequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessRequestLower", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessRequestUpper executes the ProcessRequestUpper SOAP operation.
func (c *Client) ProcessRequestUpper(ctx context.Context, req *REQUESTWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessRequestUpper", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessData executes the ProcessData SOAP operation.
func (c *Client) ProcessData(ctx context.Context, req *DataWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessData", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessDataLower executes the ProcessDataLower SOAP operation.
func (c *Client) ProcessDataLower(ctx context.Context, req *DataWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessDataLower", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessDataUpper executes the ProcessDataUpper SOAP operation.
func (c *Client) ProcessDataUpper(ctx context.Context, req *DATAWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessDataUpper", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ProcessExtremeCase executes the ProcessExtremeCase SOAP operation.
func (c *Client) ProcessExtremeCase(ctx context.Context, req *ExtremeCaseElementWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessExtremeCase", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// HandleRequest executes the HandleRequest SOAP operation.
func (c *Client) HandleRequest(ctx context.Context, req *RequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:HandleRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// ValidateRequest executes the ValidateRequest SOAP operation.
func (c *Client) ValidateRequest(ctx context.Context, req *RequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ValidateRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

// SubmitRequest executes the SubmitRequest SOAP operation.
func (c *Client) SubmitRequest(ctx context.Context, req *RequestWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:SubmitRequest", reqEnvelope)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}

	// Unmarshal response body
	var result interface{}
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

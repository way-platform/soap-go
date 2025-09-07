package complex_rawxml_scenarios

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
		soap.WithEndpoint("http://example.com/rawxml-scenarios"),
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

// ProcessFlexibleDocument executes the ProcessFlexibleDocument SOAP operation.
func (c *Client) ProcessFlexibleDocument(ctx context.Context, req *FlexibleDocumentWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessFlexibleDocument", reqEnvelope)
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

// ProcessDynamicContent executes the ProcessDynamicContent SOAP operation.
func (c *Client) ProcessDynamicContent(ctx context.Context, req *DynamicContentWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessDynamicContent", reqEnvelope)
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

// ProcessMixedDocument executes the ProcessMixedDocument SOAP operation.
func (c *Client) ProcessMixedDocument(ctx context.Context, req *MixedDocumentWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessMixedDocument", reqEnvelope)
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

// ProcessPerformanceReport executes the ProcessPerformanceReport SOAP operation.
func (c *Client) ProcessPerformanceReport(ctx context.Context, req *PerformanceReportWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessPerformanceReport", reqEnvelope)
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

// ProcessUntypedElement executes the ProcessUntypedElement SOAP operation.
func (c *Client) ProcessUntypedElement(ctx context.Context, req *UntypedElementWrapper) (*interface{}, error) {
	// Marshal request to XML
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create SOAP envelope with request body
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)

	// Make SOAP call
	respEnvelope, err := c.Call(ctx, "urn:ProcessUntypedElement", reqEnvelope)
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

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
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/rawxml-scenarios"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// ProcessFlexibleDocument executes the ProcessFlexibleDocument one-way SOAP operation.
func (c *Client) ProcessFlexibleDocument(ctx context.Context, req *FlexibleDocumentWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessFlexibleDocument", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessDynamicContent executes the ProcessDynamicContent one-way SOAP operation.
func (c *Client) ProcessDynamicContent(ctx context.Context, req *DynamicContentWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessDynamicContent", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessMixedDocument executes the ProcessMixedDocument one-way SOAP operation.
func (c *Client) ProcessMixedDocument(ctx context.Context, req *MixedDocumentWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessMixedDocument", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessPerformanceReport executes the ProcessPerformanceReport one-way SOAP operation.
func (c *Client) ProcessPerformanceReport(ctx context.Context, req *PerformanceReportWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessPerformanceReport", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

// ProcessUntypedElement executes the ProcessUntypedElement one-way SOAP operation.
func (c *Client) ProcessUntypedElement(ctx context.Context, req *UntypedElementWrapper) error {
	reqXML, err := xml.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	reqEnvelope := soap.NewEnvelopeWithBody(reqXML)
	_, err = c.Call(ctx, "urn:ProcessUntypedElement", reqEnvelope)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}
	return nil
}

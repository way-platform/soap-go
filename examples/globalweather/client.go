package globalweather

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
		soap.WithEndpoint("http://www.webservicex.com/globalweather.asmx"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// GetWeather Get weather report for all major cities around the world.
func (c *Client) GetWeather(ctx context.Context, req *GetWeatherWrapper, opts ...ClientOption) (*GetWeatherResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://www.webserviceX.NET/GetWeather", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result GetWeatherResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

// GetCitiesByCountry Get all major                 cities by country name(full / part).
func (c *Client) GetCitiesByCountry(ctx context.Context, req *GetCitiesByCountryWrapper, opts ...ClientOption) (*GetCitiesByCountryResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://www.webserviceX.NET/GetCitiesByCountry", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result GetCitiesByCountryResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}

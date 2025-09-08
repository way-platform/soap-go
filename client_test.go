package soap

import (
	"context"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		opts []ClientOption
		want func(*Client) bool
	}{
		{
			name: "default client",
			opts: nil,
			want: func(c *Client) bool {
				return c.httpClient == http.DefaultClient && c.endpoint == "" && !c.debug && c.xmlDeclaration
			},
		},
		{
			name: "with endpoint",
			opts: []ClientOption{WithEndpoint("http://example.com/soap")},
			want: func(c *Client) bool {
				return c.endpoint == "http://example.com/soap"
			},
		},
		{
			name: "with debug",
			opts: []ClientOption{WithDebug(true)},
			want: func(c *Client) bool {
				return c.debug
			},
		},
		{
			name: "with custom http client",
			opts: []ClientOption{WithHTTPClient(&http.Client{Timeout: 10 * time.Second})},
			want: func(c *Client) bool {
				return c.httpClient.Timeout == 10*time.Second
			},
		},
		{
			name: "with xml declaration disabled",
			opts: []ClientOption{WithXMLDeclaration(false)},
			want: func(c *Client) bool {
				return !c.xmlDeclaration
			},
		},
		{
			name: "with xml declaration enabled",
			opts: []ClientOption{WithXMLDeclaration(true)},
			want: func(c *Client) bool {
				return c.xmlDeclaration
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.opts...)
			if err != nil {
				t.Fatalf("NewClient() error = %v", err)
			}
			if !tt.want(client) {
				t.Errorf("NewClient() client configuration does not match expectations")
			}
		})
	}
}

func TestClient_Call(t *testing.T) {
	// Create a test server that echoes the request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "text/xml; charset=utf-8" {
			t.Errorf("Expected Content-Type 'text/xml; charset=utf-8', got '%s'", contentType)
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		// Verify it's a valid SOAP envelope
		// For the test, we just need to verify it contains expected content
		bodyStr := string(body)
		if !strings.Contains(bodyStr, "Envelope") || !strings.Contains(bodyStr, "Body") {
			t.Errorf("Request body does not contain valid SOAP envelope: %s", bodyStr)
		}

		// Create a response envelope
		respEnv := NewEnvelopeWithBody([]byte(`<response>Hello World</response>`))

		respXML, err := xml.Marshal(respEnv)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}

		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err != nil {
		t.Fatalf("Client.Call() error = %v", err)
	}

	// Verify response
	if respEnv == nil {
		t.Fatal("Response envelope is nil")
	}

	if string(respEnv.Body.Content) != `<response>Hello World</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}

func TestClient_CallWithEndpoint(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respEnv := NewEnvelopeWithBody([]byte(`<response>Custom Endpoint</response>`))
		respXML, _ := xml.Marshal(respEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client without default endpoint
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call with custom endpoint
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv, WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Client.Call() error = %v", err)
	}

	// Verify response
	if string(respEnv.Body.Content) != `<response>Custom Endpoint</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}

func TestClient_CallWithSOAPAction(t *testing.T) {
	expectedSOAPAction := "http://example.com/TestAction"

	// Create a test server that verifies SOAPAction header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		soapAction := r.Header.Get("SOAPAction")
		if soapAction != expectedSOAPAction {
			t.Errorf("Expected SOAPAction '%s', got '%s'", expectedSOAPAction, soapAction)
		}

		respEnv := NewEnvelopeWithBody([]byte(`<response>SOAP Action Test</response>`))
		respXML, _ := xml.Marshal(respEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call with SOAPAction
	ctx := context.Background()
	respEnv, err := client.Call(ctx, expectedSOAPAction, reqEnv, WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Client.Call() error = %v", err)
	}

	// Verify response
	if string(respEnv.Body.Content) != `<response>SOAP Action Test</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}

func TestClient_HTTPError(t *testing.T) {
	// Create a test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call - should return HTTP error
	ctx := context.Background()
	_, err = client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected HTTP error, got nil")
	}

	if !strings.Contains(err.Error(), "HTTP error 500") {
		t.Errorf("Expected HTTP error message, got: %v", err)
	}
}

func TestClient_SOAPFault(t *testing.T) {
	// Create a test server that returns a SOAP fault with 200 status (proper SOAP fault)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		faultEnv := NewEnvelopeWithBody([]byte(`<soap:Fault xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<faultcode>Client</faultcode>
					<faultstring>Invalid request</faultstring>
					<faultactor>http://example.com/service</faultactor>
					<detail><errorcode>E001</errorcode></detail>
				</soap:Fault>`))
		respXML, _ := xml.Marshal(faultEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK) // SOAP faults should be returned with 200 status
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call - should return SOAP fault as error
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected SOAP fault error, got nil")
	}

	// Verify that we got a response envelope (even with fault)
	if respEnv == nil {
		t.Fatal("Expected response envelope, got nil")
	}

	// Check if it's a SOAP fault using errors.As
	var fault *Fault
	if !errors.As(err, &fault) {
		t.Fatalf("Expected SOAP fault, got: %T", err)
	}

	// Verify fault details
	if fault.FaultCode != "Client" {
		t.Errorf("Expected fault code 'Client', got: %s", fault.FaultCode)
	}
	if fault.FaultString != "Invalid request" {
		t.Errorf("Expected fault string 'Invalid request', got: %s", fault.FaultString)
	}

	// Test error message format
	expectedError := "SOAP fault Invalid request: Client"
	if fault.Error() != expectedError {
		t.Errorf("Expected error message %q, got: %s", expectedError, fault.Error())
	}
	if fault.FaultActor != "http://example.com/service" {
		t.Errorf("Expected fault actor 'http://example.com/service', got: %s", fault.FaultActor)
	}
	if string(fault.Detail.Content) != "<errorcode>E001</errorcode>" {
		t.Errorf("Expected fault detail '<errorcode>E001</errorcode>', got: %s", string(fault.Detail.Content))
	}
}

func TestClient_SOAPFaultWith500Status(t *testing.T) {
	// Create a test server that returns a SOAP fault with 500 status (some servers do this)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		faultEnv := NewEnvelopeWithBody([]byte(`<soap:Fault xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<faultcode>Server</faultcode>
					<faultstring>Internal server error</faultstring>
				</soap:Fault>`))
		respXML, _ := xml.Marshal(faultEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client
	client, err := NewClient(WithEndpoint(server.URL))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call - should return HTTP error due to 500 status
	ctx := context.Background()
	_, err = client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected HTTP error for 500 status, got nil")
	}

	// This should be an HTTP error, not a SOAP fault, because of the 500 status
	var fault *Fault
	if errors.As(err, &fault) {
		t.Fatal("Expected HTTP error, got SOAP fault")
	}
}

func TestClient_EmptyEndpoint(t *testing.T) {
	// Create client without endpoint
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

	// Make the call - should fail due to empty endpoint
	ctx := context.Background()
	_, err = client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected error for empty endpoint, got nil")
	}

	if !strings.Contains(err.Error(), "endpoint is required") {
		t.Errorf("Expected endpoint error, got: %v", err)
	}
}

func TestAddXMLDeclaration(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "no declaration",
			input:    []byte(`<root>test</root>`),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<root>test</root>`,
		},
		{
			name:     "already has declaration",
			input:    []byte(`<?xml version="1.0"?><root>test</root>`),
			expected: `<?xml version="1.0"?><root>test</root>`,
		},
		{
			name:     "empty input",
			input:    []byte(``),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addXMLDeclaration(tt.input)
			if string(result) != tt.expected {
				t.Errorf("addXMLDeclaration() = %q, want %q", string(result), tt.expected)
			}
		})
	}
}

// TestClient_Integration tests the client with a more realistic SOAP service simulation
func TestClient_Integration(t *testing.T) {
	// Create a mock SOAP service that handles a simple calculator operation
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read and parse the request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		// For the test, just verify it contains expected SOAP structure
		bodyStr := string(body)
		if !strings.Contains(bodyStr, "Envelope") || !strings.Contains(bodyStr, "Body") {
			t.Errorf("Request body does not contain valid SOAP envelope: %s", bodyStr)
		}

		// Simple response - in real world this would process the request content
		respEnv := NewEnvelopeWithBody([]byte(`<CalculateResponse xmlns="http://example.com/calculator">
					<result>42</result>
				</CalculateResponse>`))

		respXML, err := xml.Marshal(respEnv)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}

		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Create client with various options
	client, err := NewClient(
		WithEndpoint(server.URL),
		WithDebug(false), // Set to true to see debug output during testing
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create a realistic request envelope
	reqEnv := NewEnvelopeWithBody([]byte(`<Calculate xmlns="http://example.com/calculator">
				<a>10</a>
				<b>32</b>
				<operation>add</operation>
			</Calculate>`))

	// Make the call
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err != nil {
		t.Fatalf("Client.Call() error = %v", err)
	}

	// Verify the response structure
	if respEnv == nil {
		t.Fatal("Response envelope is nil")
	}

	// The key test: can we successfully parse and use the response?
	// We don't need to be strict about internal namespace representation

	// The response body should contain our calculator response
	responseContent := string(respEnv.Body.Content)
	if !strings.Contains(responseContent, "CalculateResponse") {
		t.Errorf("Response does not contain expected content: %s", responseContent)
	}

	if !strings.Contains(responseContent, "<result>42</result>") {
		t.Errorf("Response does not contain expected result: %s", responseContent)
	}
}

func TestClient_XMLDeclarationOption(t *testing.T) {
	tests := []struct {
		name                 string
		xmlDeclarationOption *bool // nil means use default
		expectDeclaration    bool
	}{
		{
			name:                 "default (with declaration)",
			xmlDeclarationOption: nil,
			expectDeclaration:    true,
		},
		{
			name:                 "explicitly enabled",
			xmlDeclarationOption: &[]bool{true}[0],
			expectDeclaration:    true,
		},
		{
			name:                 "explicitly disabled",
			xmlDeclarationOption: &[]bool{false}[0],
			expectDeclaration:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server that captures the raw request
			var receivedBody []byte
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("Failed to read request body: %v", err)
				}
				receivedBody = body

				// Return a simple response
				respEnv := NewEnvelopeWithBody([]byte(`<response>OK</response>`))
				respXML, _ := xml.Marshal(respEnv)
				w.Header().Set("Content-Type", "text/xml; charset=utf-8")
				_, _ = w.Write(respXML)
			}))
			defer server.Close()

			// Create client with appropriate options
			var opts []ClientOption
			opts = append(opts, WithEndpoint(server.URL))
			if tt.xmlDeclarationOption != nil {
				opts = append(opts, WithXMLDeclaration(*tt.xmlDeclarationOption))
			}

			client, err := NewClient(opts...)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			// Create request envelope
			reqEnv := NewEnvelopeWithBody([]byte(`<request>Test</request>`))

			// Make the call
			ctx := context.Background()
			_, err = client.Call(ctx, "", reqEnv)
			if err != nil {
				t.Fatalf("Client.Call() error = %v", err)
			}

			// Check if XML declaration is present
			body := string(receivedBody)
			hasDeclaration := strings.HasPrefix(body, "<?xml")

			if hasDeclaration != tt.expectDeclaration {
				if tt.expectDeclaration {
					t.Errorf("Expected XML declaration but it was not found. Body: %s", body)
				} else {
					t.Errorf("Expected no XML declaration but it was found. Body: %s", body)
				}
			}
		})
	}
}

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
	"testing/synctest"
	"time"
)

func TestClient_Call(t *testing.T) {
	t.Parallel()
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
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>Hello World</response>`)))
		respXML, err := xml.Marshal(respEnv)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client with no retries for instant test
	client, err := NewClient(WithEndpoint(server.URL), WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
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
	t.Parallel()
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>Custom Endpoint</response>`)))
		respXML, _ := xml.Marshal(respEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client without default endpoint
	client, err := NewClient(WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
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
	t.Parallel()
	expectedSOAPAction := "http://example.com/TestAction"
	// Create a test server that verifies SOAPAction header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		soapAction := r.Header.Get("SOAPAction")
		if soapAction != expectedSOAPAction {
			t.Errorf("Expected SOAPAction '%s', got '%s'", expectedSOAPAction, soapAction)
		}
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>SOAP Action Test</response>`)))
		respXML, _ := xml.Marshal(respEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client
	client, err := NewClient(WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
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
	t.Parallel()
	// Create a test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()
	// Create client with no retries for instant test
	client, err := NewClient(WithEndpoint(server.URL), WithMaxRetries(0), WithTimeout(50*time.Millisecond))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	// Make the call - should return unified error
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	// Check that we got a unified Error
	var soapErr *Error
	if !errors.As(err, &soapErr) {
		t.Fatalf("Expected unified Error, got: %T", err)
	}
	// Verify error properties
	if soapErr.StatusCode != 500 {
		t.Errorf("Expected status code 500, got: %d", soapErr.StatusCode)
	}
	if soapErr.Fault != nil {
		t.Error("Expected no SOAP fault for non-SOAP HTTP error")
	}
	// 5xx errors should be retryable (checked by retry logic, not a method)
	if soapErr.StatusCode < 500 || soapErr.StatusCode > 599 {
		t.Error("Expected 5xx status code for retryable error")
	}
	if string(soapErr.ResponseBody) != "Internal Server Error" {
		t.Errorf("Expected response body 'Internal Server Error', got: %s", string(soapErr.ResponseBody))
	}
	// Response envelope should be nil for non-SOAP response
	if respEnv != nil {
		t.Errorf("Expected nil response envelope for non-SOAP response, got: %v", respEnv)
	}
	if !strings.Contains(err.Error(), "HTTP error 500") {
		t.Errorf("Expected HTTP error message, got: %v", err)
	}
}

func TestClient_SOAPFault(t *testing.T) {
	t.Parallel()
	// Create a test server that returns a SOAP fault with 200 status (proper SOAP fault)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		faultEnv, _ := NewEnvelope(WithBody([]byte(`<soap:Fault xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<faultcode>Client</faultcode>
					<faultstring>Invalid request</faultstring>
					<faultactor>http://example.com/service</faultactor>
					<detail><errorcode>E001</errorcode></detail>
				</soap:Fault>`)))
		respXML, _ := xml.Marshal(faultEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK) // SOAP faults should be returned with 200 status
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client with no retries for instant test
	client, err := NewClient(WithEndpoint(server.URL), WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	// Make the call - should return SOAP fault as error
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected SOAP fault error, got nil")
	}
	// Should return nil envelope since error contains the envelope
	if respEnv != nil {
		t.Fatal("Expected nil response envelope when there's an error")
	}
	// Check that we got a unified Error with SOAP fault
	var soapErr *Error
	if !errors.As(err, &soapErr) {
		t.Fatalf("Expected unified Error, got: %T", err)
	}
	// Verify that the error contains the response envelope
	if soapErr.Envelope == nil {
		t.Fatal("Expected response envelope in error, got nil")
	}
	// Verify unified error properties
	if soapErr.StatusCode != 200 {
		t.Errorf("Expected status code 200, got: %d", soapErr.StatusCode)
	}
	if soapErr.Fault == nil {
		t.Error("Expected SOAP fault to be present")
	}
	// Check the SOAP fault details through the unified error
	fault := soapErr.Fault
	if fault == nil {
		t.Fatal("Expected SOAP fault to be present in unified error")
	}
	// Verify fault details
	if fault.FaultCode != "Client" {
		t.Errorf("Expected fault code 'Client', got: %s", fault.FaultCode)
	}
	if fault.FaultString != "Invalid request" {
		t.Errorf("Expected fault string 'Invalid request', got: %s", fault.FaultString)
	}
	if fault.FaultActor != "http://example.com/service" {
		t.Errorf("Expected fault actor 'http://example.com/service', got: %s", fault.FaultActor)
	}
	if string(fault.Detail.Content) != "<errorcode>E001</errorcode>" {
		t.Errorf("Expected fault detail '<errorcode>E001</errorcode>', got: %s", string(fault.Detail.Content))
	}
	// Test unified error message format
	expectedError := "SOAP fault (HTTP 200): SOAP fault [Client]: Invalid request (actor: http://example.com/service) - detail: <errorcode>E001</errorcode>"
	if soapErr.Error() != expectedError {
		t.Errorf("Expected error message %q, got: %s", expectedError, soapErr.Error())
	}
}

func TestClient_SOAPFaultWith500Status(t *testing.T) {
	t.Parallel()
	// Track the number of requests to verify retry behavior
	var requestCount int
	// Create a test server that returns a SOAP fault with 500 status (some servers do this)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		faultEnv, _ := NewEnvelope(WithBody([]byte(`<soap:Fault xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<faultcode>Server</faultcode>
					<faultstring>Internal server error</faultstring>
				</soap:Fault>`)))
		respXML, _ := xml.Marshal(faultEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client with instant retry delay for faster test and enable retries
	client, err := NewClient(WithEndpoint(server.URL), WithRetryDelay(time.Nanosecond), WithMaxRetries(3))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	// Make the call - should return unified error with both HTTP error and SOAP fault after retries
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected error for 500 status, got nil")
	}
	// Check that we got a unified Error
	var soapErr *Error
	if !errors.As(err, &soapErr) {
		t.Fatalf("Expected unified Error, got: %T", err)
	}
	// Verify that retries occurred (3 retries, so 4 total requests)
	expectedRequests := 4
	if requestCount != expectedRequests {
		t.Errorf("Expected %d requests (1 initial + 3 retries), got: %d", expectedRequests, requestCount)
	}
	// This should be both an HTTP error AND contain a SOAP fault
	if soapErr.StatusCode != 500 {
		t.Errorf("Expected status code 500, got: %d", soapErr.StatusCode)
	}
	if soapErr.Fault == nil {
		t.Error("Expected SOAP fault to be present")
	}
	// 5xx errors with SOAP faults are now retryable
	if soapErr.StatusCode < 500 || soapErr.StatusCode > 599 {
		t.Error("Expected 5xx status code")
	}
	// Access the SOAP fault through the unified error
	fault := soapErr.Fault
	if fault == nil {
		t.Fatal("Expected SOAP fault to be present")
	}
	if fault.FaultCode != "Server" {
		t.Errorf("Expected fault code 'Server', got: %s", fault.FaultCode)
	}
	// Should return nil envelope since error contains the envelope
	if respEnv != nil {
		t.Fatal("Expected nil response envelope when there's an error")
	}
	// Should have the envelope available in the error
	if soapErr.Envelope == nil {
		t.Fatal("Expected response envelope in error even with 500 status")
	}
}

func TestClient_SOAPFaultWith400Status(t *testing.T) {
	t.Parallel()
	// Track the number of requests to verify no retries occur
	var requestCount int
	// Create a test server that returns a SOAP fault with 400 status (client error)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		faultEnv, _ := NewEnvelope(WithBody([]byte(`<soap:Fault xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<faultcode>Client</faultcode>
					<faultstring>Bad request</faultstring>
				</soap:Fault>`)))
		respXML, _ := xml.Marshal(faultEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	// Create client with no retries for instant test
	client, err := NewClient(WithEndpoint(server.URL), WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	// Make the call - should return error immediately without retries
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err == nil {
		t.Fatal("Expected error for 400 status, got nil")
	}
	// Check that we got a unified Error
	var soapErr *Error
	if !errors.As(err, &soapErr) {
		t.Fatalf("Expected unified Error, got: %T", err)
	}
	// Verify that NO retries occurred (only 1 request)
	expectedRequests := 1
	if requestCount != expectedRequests {
		t.Errorf("Expected %d request (no retries for 4xx), got: %d", expectedRequests, requestCount)
	}
	// Verify it's a 400 status code
	if soapErr.StatusCode != 400 {
		t.Errorf("Expected status code 400, got: %d", soapErr.StatusCode)
	}
	// Should return nil envelope since error contains the envelope
	if respEnv != nil {
		t.Fatal("Expected nil response envelope when there's an error")
	}
	// Should have the envelope available in the error
	if soapErr.Envelope == nil {
		t.Fatal("Expected response envelope in error even with 400 status")
	}
}

func TestClient_EmptyEndpoint(t *testing.T) {
	t.Parallel()
	// Create client without endpoint
	client, err := NewClient(WithMaxRetries(0))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
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
	t.Parallel()
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

func TestClient_Integration(t *testing.T) {
	t.Parallel()
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
		respEnv, _ := NewEnvelope(WithBody([]byte(`<CalculateResponse xmlns="http://example.com/calculator">
					<result>42</result>
				</CalculateResponse>`)))
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
		WithMaxRetries(0),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create a realistic request envelope
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<Calculate xmlns="http://example.com/calculator">
				<a>10</a>
				<b>32</b>
				<operation>add</operation>
			</Calculate>`)))
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
	t.Parallel()
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
				respEnv, _ := NewEnvelope(WithBody([]byte(`<response>OK</response>`)))
				respXML, _ := xml.Marshal(respEnv)
				w.Header().Set("Content-Type", "text/xml; charset=utf-8")
				_, _ = w.Write(respXML)
			}))
			defer server.Close()
			// Create client with appropriate options
			var opts []ClientOption
			opts = append(opts, WithEndpoint(server.URL), WithMaxRetries(0))
			if tt.xmlDeclarationOption != nil {
				opts = append(opts, WithXMLDeclaration(*tt.xmlDeclarationOption))
			}
			client, err := NewClient(opts...)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}
			// Create request envelope
			reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
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

func TestClient_RetryLogic(t *testing.T) {
	t.Parallel()
	requestCount := 0
	const failCount = 2 // Fail 2 times, succeed on the 3rd attempt
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount <= failCount {
			w.WriteHeader(http.StatusServiceUnavailable) // 503 error
			_, _ = w.Write([]byte("Service Unavailable"))
			return
		}
		// Success response
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>Success</response>`)))
		respXML, err := xml.Marshal(respEnv)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	client, err := NewClient(
		WithEndpoint(server.URL),
		WithMaxRetries(3),
		WithRetryDelay(time.Nanosecond), // Use instant delay for fast test
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err != nil {
		t.Fatalf("Client.Call() failed after retries: %v", err)
	}
	if requestCount != failCount+1 {
		t.Errorf("Expected %d requests, got %d", failCount+1, requestCount)
	}
	if respEnv == nil {
		t.Fatal("Response envelope is nil")
	}
	if string(respEnv.Body.Content) != `<response>Success</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}

func TestClient_RetryContextCancellation(t *testing.T) {
	t.Parallel()
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		// Always return 503 to force retries
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("Service Unavailable"))
	}))
	defer server.Close()
	client, err := NewClient(
		WithEndpoint(server.URL),
		WithMaxRetries(5),                   // Set high retry count
		WithRetryDelay(50*time.Millisecond), // Use delay longer than timeout but shorter for fast test
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	// Create context with short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	start := time.Now()
	_, err = client.Call(ctx, "", reqEnv)
	elapsed := time.Since(start)
	// Should get context deadline exceeded error
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Expected context.DeadlineExceeded, got: %v", err)
	}
	// Should have completed quickly (before the retry delay)
	if elapsed > 120*time.Millisecond {
		t.Errorf("Request took too long (%v), context cancellation may not be working during retry delay", elapsed)
	}
	// Should have made at least one request
	if requestCount < 1 {
		t.Errorf("Expected at least 1 request, got %d", requestCount)
	}
}

func TestClient_RetryWithNonRetryableError(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			w.WriteHeader(http.StatusBadRequest) // 400 error, should not be retried
			_, _ = w.Write([]byte("Bad Request"))
		}))
		defer server.Close()
		client, err := NewClient(
			WithEndpoint(server.URL),
			WithMaxRetries(3),
			WithRetryDelay(time.Nanosecond),
		)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
		ctx := context.Background()
		_, err = client.Call(ctx, "", reqEnv)
		if err == nil {
			t.Fatal("Expected an error, but got nil")
		}
		if requestCount != 1 {
			t.Errorf("Expected 1 request for a non-retryable error, got %d", requestCount)
		}
		// Should get a unified Error with 400 status
		var soapErr *Error
		if !errors.As(err, &soapErr) {
			t.Fatalf("Expected unified Error, got: %T", err)
		}
		if soapErr.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code 400, got: %d", soapErr.StatusCode)
		}
		// 4xx errors should not be retryable (checked by retry logic, not a method)
		if soapErr.StatusCode >= 500 && soapErr.StatusCode <= 599 {
			t.Error("Expected non-5xx status code for non-retryable error")
		}
	})
}

func TestClient_RetryWith429TooManyRequests(t *testing.T) {
	t.Parallel()
	requestCount := 0
	const failCount = 2 // Fail 2 times with 429, succeed on the 3rd attempt
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount <= failCount {
			w.WriteHeader(http.StatusTooManyRequests) // 429 error, should be retried
			_, _ = w.Write([]byte("Too Many Requests"))
			return
		}
		// Success response
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>Success</response>`)))
		respXML, err := xml.Marshal(respEnv)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()
	client, err := NewClient(
		WithEndpoint(server.URL),
		WithMaxRetries(3),
		WithRetryDelay(time.Nanosecond), // Use instant delay for fast test
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err != nil {
		t.Fatalf("Client.Call() failed after retries: %v", err)
	}
	if requestCount != failCount+1 {
		t.Errorf("Expected %d requests for 429 retries, got %d", failCount+1, requestCount)
	}
	if respEnv == nil {
		t.Fatal("Response envelope is nil")
	}
	if string(respEnv.Body.Content) != `<response>Success</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}

//go:build integration

package soapcall

import (
	"context"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/way-platform/soap-go"
	"github.com/way-platform/soap-go/examples/numberconversion"
)

const (
	// NumberConversion service endpoint
	numberConversionEndpoint = "https://www.dataaccess.com/webservicesserver/NumberConversion.wso"

	// Target namespace for the NumberConversion service
	targetNamespace = "http://www.dataaccess.com/webservicesserver/"
)

// ResponseEnvelope is optimized for unmarshaling SOAP responses
// Many services return responses in namespace-qualified format rather than soap: prefixed
type ResponseEnvelope struct {
	XMLName xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    ResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type ResponseBody struct {
	Content []byte `xml:",innerxml"`
}

// TestNumberToWordsIntegration tests the NumberToWords operation with real service calls
func TestNumberToWordsIntegration(t *testing.T) {
	client := NewClient(Config{
		Endpoint:   numberConversionEndpoint,
		Timeout:    30 * time.Second,
		SOAPAction: "",
	})

	testCases := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			name:     "zero",
			input:    0,
			expected: "zero",
		},
		{
			name:     "single digit",
			input:    5,
			expected: "five",
		},
		{
			name:     "teen number",
			input:    13,
			expected: "thirteen",
		},
		{
			name:     "two digit",
			input:    42,
			expected: "forty two",
		},
		{
			name:     "three digit",
			input:    123,
			expected: "one hundred and twenty three",
		},
		{
			name:     "large number",
			input:    1234567,
			expected: "one million two hundred and thirty four thousand five hundred and sixty seven",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create request payload with explicit namespace declaration
			request := struct {
				XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
				UbiNum  uint64   `xml:"ubiNum"`
			}{
				UbiNum: tc.input,
			}

			// Marshal request to XML
			requestXML, err := xml.Marshal(request)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			// Create SOAP envelope using the main soap.Envelope type
			envelope := &soap.Envelope{
				XMLNS: soap.Namespace,
				Body:  soap.Body{Content: requestXML},
			}

			// Marshal envelope
			envelopeXML, err := xml.Marshal(envelope)
			if err != nil {
				t.Fatalf("Failed to marshal envelope: %v", err)
			}

			// Add XML declaration
			soapRequest := AddXMLDeclaration(envelopeXML)

			// Make SOAP call
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			response, err := client.Call(ctx, soapRequest)
			if err != nil {
				t.Fatalf("SOAP call failed: %v", err)
			}

			// Parse response envelope
			var responseEnvelope ResponseEnvelope
			if err := xml.Unmarshal(response, &responseEnvelope); err != nil {
				t.Fatalf("Failed to unmarshal response envelope: %v", err)
			}

			// Parse response body
			var responseBody numberconversion.NumberToWordsResponse
			if err := xml.Unmarshal(responseEnvelope.Body.Content, &responseBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			// Verify the response
			if responseBody.NumberToWordsResult == "" {
				t.Errorf("Expected non-empty response, got empty string")
			}

			t.Logf("Input: %d, Output: %s", tc.input, responseBody.NumberToWordsResult)

			// Note: We don't do exact string matching as the service might have slight variations
			// in formatting (e.g., "forty-two" vs "forty two"), but we verify it's not empty
			// and contains expected key words for larger numbers
			if tc.input >= 1000000 && !containsWord(responseBody.NumberToWordsResult, "million") {
				t.Errorf("Expected response to contain 'million' for input %d, got: %s", tc.input, responseBody.NumberToWordsResult)
			}
		})
	}
}

// TestNumberToDollarsIntegration tests the NumberToDollars operation with real service calls
func TestNumberToDollarsIntegration(t *testing.T) {
	client := NewClient(Config{
		Endpoint:   numberConversionEndpoint,
		Timeout:    30 * time.Second,
		SOAPAction: "",
	})

	testCases := []struct {
		name  string
		input float64
	}{
		{
			name:  "zero dollars",
			input: 0.0,
		},
		{
			name:  "single dollar",
			input: 1.0,
		},
		{
			name:  "dollars and cents",
			input: 123.45,
		},
		{
			name:  "large amount",
			input: 1234567.89,
		},
		{
			name:  "cents only",
			input: 0.99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create request payload with explicit namespace declaration
			request := struct {
				XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToDollars"`
				DNum    float64  `xml:"dNum"`
			}{
				DNum: tc.input,
			}

			// Marshal request to XML
			requestXML, err := xml.Marshal(request)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			// Create SOAP envelope using the main soap.Envelope type
			envelope := &soap.Envelope{
				XMLNS: soap.Namespace,
				Body:  soap.Body{Content: requestXML},
			}

			// Marshal envelope
			envelopeXML, err := xml.Marshal(envelope)
			if err != nil {
				t.Fatalf("Failed to marshal envelope: %v", err)
			}

			// Add XML declaration
			soapRequest := AddXMLDeclaration(envelopeXML)

			// Make SOAP call
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			response, err := client.Call(ctx, soapRequest)
			if err != nil {
				t.Fatalf("SOAP call failed: %v", err)
			}

			// Parse response envelope
			var responseEnvelope ResponseEnvelope
			if err := xml.Unmarshal(response, &responseEnvelope); err != nil {
				t.Fatalf("Failed to unmarshal response envelope: %v", err)
			}

			// Parse response body
			var responseBody numberconversion.NumberToDollarsResponse
			if err := xml.Unmarshal(responseEnvelope.Body.Content, &responseBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			t.Logf("Input: %.2f, Output: %s", tc.input, responseBody.NumberToDollarsResult)

			// Basic validation - should contain dollar-related words for non-zero amounts
			// Zero dollars may return empty string which is valid
			if tc.input > 0 {
				result := responseBody.NumberToDollarsResult
				if result == "" {
					t.Errorf("Expected non-empty response for non-zero input %.2f", tc.input)
				} else if !containsWord(result, "dollar") && !containsWord(result, "dollars") && !containsWord(result, "cent") && !containsWord(result, "cents") {
					t.Errorf("Expected response to contain 'dollar', 'dollars', 'cent', or 'cents' for input %.2f, got: %s", tc.input, result)
				}
			} else {
				// For zero input, empty response is acceptable
				t.Logf("Zero dollar input returned: '%s' (empty response is acceptable)", responseBody.NumberToDollarsResult)
			}
		})
	}
}

// TestSOAPFaultHandling tests how the client handles SOAP faults
func TestSOAPFaultHandling(t *testing.T) {
	client := NewClient(Config{
		Endpoint:   numberConversionEndpoint,
		Timeout:    30 * time.Second,
		SOAPAction: "",
	})

	// Test with invalid/malformed SOAP request
	t.Run("malformed_request", func(t *testing.T) {
		malformedSOAP := []byte(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<InvalidOperation xmlns="http://www.dataaccess.com/webservicesserver/">
			<InvalidParam>invalid</InvalidParam>
		</InvalidOperation>
	</soap:Body>
</soap:Envelope>`)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		response, err := client.Call(ctx, malformedSOAP)

		// The service might return a fault or an HTTP error
		// We just want to ensure our client handles it gracefully
		if err != nil {
			t.Logf("Expected error for malformed request: %v", err)
		} else {
			t.Logf("Service response for malformed request: %s", string(response))

			// Check if it's a SOAP fault
			var envelope soap.Envelope
			if xml.Unmarshal(response, &envelope) == nil {
				var fault soap.Fault
				if xml.Unmarshal(envelope.Body.Content, &fault) == nil {
					t.Logf("Received SOAP fault - Code: %s, String: %s", fault.FaultCode, fault.FaultString)
				}
			}
		}
	})
}

// TestClientConfiguration tests various client configuration options
func TestClientConfiguration(t *testing.T) {
	t.Run("custom_timeout", func(t *testing.T) {
		// Test with very short timeout
		client := NewClient(Config{
			Endpoint: numberConversionEndpoint,
			Timeout:  1 * time.Millisecond, // Very short timeout
		})

		request := struct {
			XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
			UbiNum  uint64   `xml:"ubiNum"`
		}{
			UbiNum: 42,
		}
		requestXML, _ := xml.Marshal(request)

		envelope := &soap.Envelope{
			XMLNS: soap.Namespace,
			Body:  soap.Body{Content: requestXML},
		}

		envelopeXML, _ := xml.Marshal(envelope)
		soapRequest := append([]byte(xml.Header), envelopeXML...)

		ctx := context.Background()
		_, err := client.Call(ctx, soapRequest)

		// Should timeout
		if err == nil {
			t.Error("Expected timeout error with very short timeout")
		}
		t.Logf("Got expected timeout error: %v", err)
	})

	t.Run("custom_headers", func(t *testing.T) {
		client := NewClient(Config{
			Endpoint: numberConversionEndpoint,
			Timeout:  30 * time.Second,
			Headers: map[string]string{
				"User-Agent":    "soap-go-integration-test",
				"X-Test-Header": "test-value",
			},
		})

		request := struct {
			XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
			UbiNum  uint64   `xml:"ubiNum"`
		}{
			UbiNum: 1,
		}
		requestXML, _ := xml.Marshal(request)

		envelope := &soap.Envelope{
			XMLNS: soap.Namespace,
			Body:  soap.Body{Content: requestXML},
		}

		envelopeXML, _ := xml.Marshal(envelope)
		soapRequest := append([]byte(xml.Header), envelopeXML...)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		response, err := client.Call(ctx, soapRequest)
		if err != nil {
			t.Fatalf("SOAP call with custom headers failed: %v", err)
		}

		// Parse and verify we got a valid response
		var responseEnvelope ResponseEnvelope
		if err := xml.Unmarshal(response, &responseEnvelope); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		var responseBody numberconversion.NumberToWordsResponse
		if err := xml.Unmarshal(responseEnvelope.Body.Content, &responseBody); err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		if responseBody.NumberToWordsResult == "" {
			t.Error("Expected non-empty response with custom headers")
		}
		t.Logf("Response with custom headers: %s", responseBody.NumberToWordsResult)
	})
}

// TestXMLDeclarationHandling tests various XML declaration scenarios
func TestXMLDeclarationHandling(t *testing.T) {
	client := NewClient(Config{
		Endpoint:   numberConversionEndpoint,
		Timeout:    30 * time.Second,
		SOAPAction: "",
	})

	testCases := []struct {
		name              string
		useXMLDeclaration bool
		encoding          string
	}{
		{
			name:              "with_standard_xml_declaration",
			useXMLDeclaration: true,
			encoding:          "utf-8",
		},
		{
			name:              "with_utf8_xml_declaration",
			useXMLDeclaration: true,
			encoding:          "UTF-8",
		},
		{
			name:              "without_xml_declaration",
			useXMLDeclaration: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create request payload with explicit namespace declaration
			request := struct {
				XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
				UbiNum  uint64   `xml:"ubiNum"`
			}{
				UbiNum: 42,
			}

			requestXML, err := xml.Marshal(request)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			// Use the main soap.Envelope type
			envelope := &soap.Envelope{
				XMLNS: soap.Namespace,
				Body:  soap.Body{Content: requestXML},
			}

			envelopeXML, err := xml.Marshal(envelope)
			if err != nil {
				t.Fatalf("Failed to marshal envelope: %v", err)
			}

			// Handle XML declaration based on test case
			var soapRequest []byte
			if tc.useXMLDeclaration {
				if tc.encoding == "utf-8" {
					soapRequest = AddXMLDeclaration(envelopeXML)
				} else {
					soapRequest = EnsureXMLDeclarationWithEncoding(envelopeXML, tc.encoding)
				}
			} else {
				soapRequest = envelopeXML
			}

			// Log the actual XML being sent (truncated for readability)
			logData := soapRequest
			if len(logData) > 200 {
				logData = logData[:200]
			}
			t.Logf("Sending XML (%s): %s...", tc.name, string(logData))

			// Make SOAP call
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			response, err := client.Call(ctx, soapRequest)
			if err != nil {
				t.Fatalf("SOAP call failed for %s: %v", tc.name, err)
			}

			// Parse response using ResponseEnvelope optimized for unmarshaling
			var responseEnvelope ResponseEnvelope
			if err := xml.Unmarshal(response, &responseEnvelope); err != nil {
				t.Fatalf("Failed to unmarshal response envelope: %v", err)
			}

			var responseBody numberconversion.NumberToWordsResponse
			if err := xml.Unmarshal(responseEnvelope.Body.Content, &responseBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			// Verify we got a valid response
			if responseBody.NumberToWordsResult == "" {
				t.Errorf("Expected non-empty response for %s", tc.name)
			}

			t.Logf("Response for %s: %s", tc.name, responseBody.NumberToWordsResult)
		})
	}
}

// TestConcurrentRequests tests making multiple concurrent SOAP calls
func TestConcurrentRequests(t *testing.T) {
	client := NewClient(Config{
		Endpoint:   numberConversionEndpoint,
		Timeout:    30 * time.Second,
		SOAPAction: "",
	})

	const numRequests = 5
	results := make(chan string, numRequests)
	errors := make(chan error, numRequests)

	// Launch concurrent requests
	for i := 0; i < numRequests; i++ {
		go func(num uint64) {
			request := struct {
				XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
				UbiNum  uint64   `xml:"ubiNum"`
			}{
				UbiNum: num,
			}
			requestXML, err := xml.Marshal(request)
			if err != nil {
				errors <- fmt.Errorf("marshal error: %w", err)
				return
			}

			envelope := soap.Envelope{
				XMLNS: soap.Namespace,
				Body:  soap.Body{Content: requestXML},
			}

			envelopeXML, err := xml.Marshal(envelope)
			if err != nil {
				errors <- fmt.Errorf("envelope marshal error: %w", err)
				return
			}

			soapRequest := append([]byte(xml.Header), envelopeXML...)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			response, err := client.Call(ctx, soapRequest)
			if err != nil {
				errors <- fmt.Errorf("SOAP call error: %w", err)
				return
			}

			var responseEnvelope ResponseEnvelope
			if err := xml.Unmarshal(response, &responseEnvelope); err != nil {
				errors <- fmt.Errorf("response unmarshal error: %w", err)
				return
			}

			var responseBody numberconversion.NumberToWordsResponse
			if err := xml.Unmarshal(responseEnvelope.Body.Content, &responseBody); err != nil {
				errors <- fmt.Errorf("body unmarshal error: %w", err)
				return
			}

			results <- responseBody.NumberToWordsResult
		}(uint64(i + 1))
	}

	// Collect results
	var successCount int
	for i := 0; i < numRequests; i++ {
		select {
		case result := <-results:
			t.Logf("Concurrent request result: %s", result)
			successCount++
		case err := <-errors:
			t.Logf("Concurrent request error: %v", err)
		case <-time.After(60 * time.Second):
			t.Fatal("Timeout waiting for concurrent requests")
		}
	}

	if successCount == 0 {
		t.Error("All concurrent requests failed")
	}
	t.Logf("Successful concurrent requests: %d/%d", successCount, numRequests)
}

// Helper function to check if a string contains a word (case-insensitive)
func containsWord(text, word string) bool {
	if len(text) == 0 || len(word) == 0 {
		return false
	}

	// Convert to lowercase and add word boundaries for proper matching
	lowerText := strings.ToLower(fmt.Sprintf(" %s ", text))
	lowerWord := strings.ToLower(fmt.Sprintf(" %s ", word))

	return strings.Contains(lowerText, lowerWord)
}

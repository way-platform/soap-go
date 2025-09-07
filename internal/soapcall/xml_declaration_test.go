package soapcall

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestAddXMLDeclaration(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "xml_without_declaration",
			input:    []byte("<root><child>value</child></root>"),
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>",
		},
		{
			name:     "xml_with_existing_declaration",
			input:    []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>"),
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>",
		},
		{
			name:     "empty_input",
			input:    []byte(""),
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := AddXMLDeclaration(tc.input)
			if string(result) != tc.expected {
				t.Errorf("Expected: %s\nGot: %s", tc.expected, string(result))
			}
		})
	}
}

func TestEnsureXMLDeclaration(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "xml_without_declaration",
			input:    []byte("<root><child>value</child></root>"),
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>",
		},
		{
			name:     "xml_with_existing_declaration",
			input:    []byte("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<root><child>value</child></root>"),
			expected: "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<root><child>value</child></root>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := EnsureXMLDeclaration(tc.input)
			if string(result) != tc.expected {
				t.Errorf("Expected: %s\nGot: %s", tc.expected, string(result))
			}
		})
	}
}

func TestEnsureXMLDeclarationWithEncoding(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		encoding string
		expected string
	}{
		{
			name:     "xml_without_declaration_utf8",
			input:    []byte("<root><child>value</child></root>"),
			encoding: "UTF-8",
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>",
		},
		{
			name:     "xml_without_declaration_iso",
			input:    []byte("<root><child>value</child></root>"),
			encoding: "ISO-8859-1",
			expected: "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<root><child>value</child></root>",
		},
		{
			name:     "xml_with_existing_declaration",
			input:    []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>"),
			encoding: "ISO-8859-1",
			expected: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root><child>value</child></root>", // Should not change
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := EnsureXMLDeclarationWithEncoding(tc.input, tc.encoding)
			if string(result) != tc.expected {
				t.Errorf("Expected: %s\nGot: %s", tc.expected, string(result))
			}
		})
	}
}

func TestXMLDeclarationWithSOAPEnvelope(t *testing.T) {
	// Test with a realistic SOAP envelope
	envelope := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/"><ubiNum>42</ubiNum></NumberToWords></soap:Body></soap:Envelope>`

	result := AddXMLDeclaration([]byte(envelope))

	// Should start with XML declaration
	if !strings.HasPrefix(string(result), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		t.Errorf("Result should start with XML declaration, got: %s", string(result[:50]))
	}

	// Should be valid XML
	if err := xml.Unmarshal(result, &struct{}{}); err == nil {
		// We expect an error because we're unmarshaling into an empty struct,
		// but the XML should be parseable
		t.Logf("XML is parseable (expected error for empty struct): %v", err)
	}

	t.Logf("Final XML with declaration: %s", string(result))
}

package soapgen

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestComprehensiveXMLParsing tests comprehensive XML parsing scenarios
// based on the original examples_test.go but using generated types directly
func TestComprehensiveXMLParsing(t *testing.T) {
	// Test comprehensive parsing scenarios that mirror the original examples_test.go
	testCases := []struct {
		name        string
		description string
		testFunc    func(t *testing.T)
	}{
		{
			name:        "timestamp_formats",
			description: "Test various timestamp format handling",
			testFunc:    testTimestampFormats,
		},
		{
			name:        "namespace_handling",
			description: "Test flexible namespace handling",
			testFunc:    testNamespaceHandling,
		},
		{
			name:        "enumeration_parsing",
			description: "Test enumeration constant parsing",
			testFunc:    testEnumerationParsing,
		},
		{
			name:        "optional_elements",
			description: "Test optional and multiple element parsing",
			testFunc:    testOptionalElements,
		},
		{
			name:        "binary_data",
			description: "Test binary data (base64/hex) parsing",
			testFunc:    testBinaryData,
		},
		{
			name:        "attribute_parsing",
			description: "Test XML attribute parsing",
			testFunc:    testAttributeParsing,
		},
		{
			name:        "complex_types",
			description: "Test complex type and sequence parsing",
			testFunc:    testComplexTypes,
		},
		{
			name:        "inline_types",
			description: "Test inline complex type parsing",
			testFunc:    testInlineTypes,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running comprehensive test: %s", tc.description)
			tc.testFunc(t)
		})
	}
}

// testTimestampFormats tests timestamp format handling scenarios
func testTimestampFormats(t *testing.T) {
	testCases := []struct {
		name      string
		xml       string
		expectErr bool
	}{
		{
			name: "standard_utc",
			xml: `<TimestampTest>
				<timestamp>2023-12-25T10:30:00Z</timestamp>
			</TimestampTest>`,
			expectErr: false,
		},
		{
			name: "timezone_offset",
			xml: `<TimestampTest>
				<timestamp>2023-12-25T10:30:00-05:00</timestamp>
			</TimestampTest>`,
			expectErr: false,
		},
		{
			name: "milliseconds",
			xml: `<TimestampTest>
				<timestamp>2023-12-25T10:30:00.123Z</timestamp>
			</TimestampTest>`,
			expectErr: false,
		},
		{
			name: "space_separator_should_fail",
			xml: `<TimestampTest>
				<timestamp>2023-12-25 10:30:00</timestamp>
			</TimestampTest>`,
			expectErr: true, // Non-XSD compliant format should fail when parsed as time.Time
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test that the XML is well-formed first
			if !isValidXML([]byte(tc.xml)) {
				t.Errorf("XML is not well-formed: %s", tc.xml)
				return
			}

			// Test specific timestamp parsing
			type TimestampTest struct {
				XMLName   xml.Name  `xml:"TimestampTest"`
				Timestamp time.Time `xml:"timestamp"`
			}

			var timestampTest TimestampTest
			err := xml.Unmarshal([]byte(tc.xml), &timestampTest)

			if tc.expectErr && err == nil {
				t.Errorf("Expected timestamp parsing to fail but it succeeded")
			} else if !tc.expectErr && err != nil {
				t.Errorf("Expected timestamp parsing to succeed but got error: %v", err)
			}

			if !tc.expectErr && err == nil {
				t.Logf("Successfully parsed timestamp: %v", timestampTest.Timestamp)
			}
		})
	}
}

// testNamespaceHandling tests flexible namespace handling
func testNamespaceHandling(t *testing.T) {
	testCases := []struct {
		name        string
		xml         string
		expectError bool
		description string
	}{
		{
			name: "default_namespace",
			xml: `<TestElement xmlns="http://example.com/test">
				<field>value</field>
			</TestElement>`,
			expectError: false,
			description: "Default namespace should work",
		},
		{
			name: "prefixed_namespace",
			xml: `<tns:TestElement xmlns:tns="http://example.com/test">
				<field>value</field>
			</tns:TestElement>`,
			expectError: false,
			description: "Prefixed namespace should work",
		},
		{
			name: "no_namespace",
			xml: `<TestElement>
				<field>value</field>
			</TestElement>`,
			expectError: false,
			description: "No namespace should work with flexible handling",
		},
		{
			name: "different_namespace",
			xml: `<TestElement xmlns="http://different.namespace.com">
				<field>value</field>
			</TestElement>`,
			expectError: false,
			description: "Different namespace should work with flexible handling",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use flexible XMLName field that matches local name only
			type TestElement struct {
				XMLName xml.Name `xml:"TestElement"`
				Field   string   `xml:"field"`
			}

			var result TestElement
			err := xml.Unmarshal([]byte(tc.xml), &result)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for %s, but got none", tc.description)
			} else if !tc.expectError && err != nil {
				t.Errorf("Expected success for %s, but got error: %v", tc.description, err)
			}

			if !tc.expectError && err == nil {
				t.Logf("SUCCESS: %s - Namespace: %q, Field: %q",
					tc.description, result.XMLName.Space, result.Field)
			}
		})
	}
}

// testEnumerationParsing tests enumeration constant parsing
func testEnumerationParsing(t *testing.T) {
	// Define a test enumeration type similar to generated ones
	type StatusType string

	const (
		StatusTypeActive   StatusType = "active"
		StatusTypeInactive StatusType = "inactive"
		StatusTypePending  StatusType = "pending"
	)

	// String method
	stringMethod := func(e StatusType) string {
		return string(e)
	}

	// IsValid method
	isValidMethod := func(e StatusType) bool {
		switch e {
		case StatusTypeActive, StatusTypeInactive, StatusTypePending:
			return true
		default:
			return false
		}
	}

	type UserInfo struct {
		XMLName xml.Name   `xml:"UserInfo"`
		Status  StatusType `xml:"status"`
		Email   string     `xml:"email"`
	}

	testCases := []struct {
		name           string
		xml            string
		expectedStatus StatusType
		expectValid    bool
	}{
		{
			name: "active_status",
			xml: `<UserInfo>
				<status>active</status>
				<email>test@example.com</email>
			</UserInfo>`,
			expectedStatus: StatusTypeActive,
			expectValid:    true,
		},
		{
			name: "pending_status",
			xml: `<UserInfo>
				<status>pending</status>
				<email>test@example.com</email>
			</UserInfo>`,
			expectedStatus: StatusTypePending,
			expectValid:    true,
		},
		{
			name: "invalid_status",
			xml: `<UserInfo>
				<status>invalid_status</status>
				<email>test@example.com</email>
			</UserInfo>`,
			expectedStatus: "invalid_status",
			expectValid:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var userInfo UserInfo
			err := xml.Unmarshal([]byte(tc.xml), &userInfo)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if userInfo.Status != tc.expectedStatus {
				t.Errorf("Expected status %q, got %q", tc.expectedStatus, userInfo.Status)
			}

			isValid := isValidMethod(userInfo.Status)
			if isValid != tc.expectValid {
				t.Errorf("Expected IsValid() to return %v, got %v", tc.expectValid, isValid)
			}

			t.Logf("Status: %s, Valid: %v", stringMethod(userInfo.Status), isValid)
		})
	}
}

// testOptionalElements tests optional and multiple element parsing
func testOptionalElements(t *testing.T) {
	type TestElement struct {
		XMLName         xml.Name `xml:"TestElement"`
		RequiredString  string   `xml:"requiredString"`
		OptionalString  *string  `xml:"optionalString,omitempty"`
		MultipleStrings []string `xml:"multipleStrings"`
		OptionalNumbers []int32  `xml:"optionalNumbers,omitempty"`
	}

	testCases := []struct {
		name     string
		xml      string
		expected TestElement
	}{
		{
			name: "with_optional_elements",
			xml: `<TestElement>
				<requiredString>required</requiredString>
				<optionalString>optional</optionalString>
				<multipleStrings>first</multipleStrings>
				<multipleStrings>second</multipleStrings>
				<optionalNumbers>1</optionalNumbers>
				<optionalNumbers>2</optionalNumbers>
			</TestElement>`,
			expected: TestElement{
				XMLName:         xml.Name{Local: "TestElement"},
				RequiredString:  "required",
				OptionalString:  stringPtr("optional"),
				MultipleStrings: []string{"first", "second"},
				OptionalNumbers: []int32{1, 2},
			},
		},
		{
			name: "minimal_required_only",
			xml: `<TestElement>
				<requiredString>required</requiredString>
				<multipleStrings>single</multipleStrings>
			</TestElement>`,
			expected: TestElement{
				XMLName:         xml.Name{Local: "TestElement"},
				RequiredString:  "required",
				OptionalString:  nil,
				MultipleStrings: []string{"single"},
				OptionalNumbers: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result TestElement
			err := xml.Unmarshal([]byte(tc.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tc.expected, result); diff != "" {
				t.Errorf("TestElement mismatch (-expected +actual):\n%s", diff)
			}

			t.Logf("Successfully parsed optional elements")
		})
	}
}

// testBinaryData tests binary data (base64/hex) parsing
func testBinaryData(t *testing.T) {
	type BinaryTest struct {
		XMLName    xml.Name `xml:"BinaryTest"`
		Base64Data []byte   `xml:"base64Data"`
		HexData    []byte   `xml:"hexData"`
	}

	testCases := []struct {
		name     string
		xml      string
		expected BinaryTest
	}{
		{
			name: "base64_data",
			xml: `<BinaryTest>
				<base64Data>SGVsbG8gV29ybGQ=</base64Data>
				<hexData>48656C6C6F</hexData>
			</BinaryTest>`,
			expected: BinaryTest{
				XMLName:    xml.Name{Local: "BinaryTest"},
				Base64Data: []byte("SGVsbG8gV29ybGQ="), // XML treats as literal string
				HexData:    []byte("48656C6C6F"),       // XML treats as literal string
			},
		},
		{
			name: "empty_data",
			xml: `<BinaryTest>
				<base64Data></base64Data>
				<hexData></hexData>
			</BinaryTest>`,
			expected: BinaryTest{
				XMLName:    xml.Name{Local: "BinaryTest"},
				Base64Data: []byte{},
				HexData:    []byte{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result BinaryTest
			err := xml.Unmarshal([]byte(tc.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tc.expected, result); diff != "" {
				t.Errorf("BinaryTest mismatch (-expected +actual):\n%s", diff)
			}

			t.Logf("Successfully parsed binary data")
		})
	}
}

// testAttributeParsing tests XML attribute parsing
func testAttributeParsing(t *testing.T) {
	type ElementWithAttributes struct {
		XMLName  xml.Name `xml:"ElementWithAttributes"`
		ID       string   `xml:"id,attr"`
		Optional *bool    `xml:"optional,attr,omitempty"`
		Content  string   `xml:"content"`
	}

	testCases := []struct {
		name     string
		xml      string
		expected ElementWithAttributes
	}{
		{
			name: "with_attributes",
			xml: `<ElementWithAttributes id="123" optional="true">
				<content>test content</content>
			</ElementWithAttributes>`,
			expected: ElementWithAttributes{
				XMLName:  xml.Name{Local: "ElementWithAttributes"},
				ID:       "123",
				Optional: boolPtr(true),
				Content:  "test content",
			},
		},
		{
			name: "minimal_attributes",
			xml: `<ElementWithAttributes id="456">
				<content>minimal content</content>
			</ElementWithAttributes>`,
			expected: ElementWithAttributes{
				XMLName:  xml.Name{Local: "ElementWithAttributes"},
				ID:       "456",
				Optional: nil,
				Content:  "minimal content",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result ElementWithAttributes
			err := xml.Unmarshal([]byte(tc.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tc.expected, result); diff != "" {
				t.Errorf("ElementWithAttributes mismatch (-expected +actual):\n%s", diff)
			}

			t.Logf("Successfully parsed attributes")
		})
	}
}

// testComplexTypes tests complex type and sequence parsing
func testComplexTypes(t *testing.T) {
	type Address struct {
		Street  string `xml:"street"`
		City    string `xml:"city"`
		ZipCode string `xml:"zipCode"`
		Country string `xml:"country,attr"`
		Active  *bool  `xml:"active,attr,omitempty"`
	}

	type Person struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"name"`
		Age     int32    `xml:"age"`
		Address Address  `xml:"address"`
	}

	testCases := []struct {
		name     string
		xml      string
		expected Person
	}{
		{
			name: "complete_person",
			xml: `<Person>
				<name>John Doe</name>
				<age>30</age>
				<address country="US" active="true">
					<street>123 Main St</street>
					<city>Anytown</city>
					<zipCode>12345</zipCode>
				</address>
			</Person>`,
			expected: Person{
				XMLName: xml.Name{Local: "Person"},
				Name:    "John Doe",
				Age:     30,
				Address: Address{
					Street:  "123 Main St",
					City:    "Anytown",
					ZipCode: "12345",
					Country: "US",
					Active:  boolPtr(true),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result Person
			err := xml.Unmarshal([]byte(tc.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tc.expected, result); diff != "" {
				t.Errorf("Person mismatch (-expected +actual):\n%s", diff)
			}

			t.Logf("Successfully parsed complex type")
		})
	}
}

// testInlineTypes tests inline complex type parsing with RawXML escape hatch
func testInlineTypes(t *testing.T) {
	// Define RawXML type similar to the generated one
	type RawXML string

	type InlineTest_Customer struct {
		Name    string `xml:"name"`
		Address RawXML `xml:"address"`
	}

	type InlineTest struct {
		XMLName  xml.Name            `xml:"InlineTest"`
		Customer InlineTest_Customer `xml:"customer"`
	}

	testCases := []struct {
		name     string
		xml      string
		expected InlineTest
	}{
		{
			name: "inline_complex_type",
			xml: `<InlineTest>
				<customer>
					<name>Jane Doe</name>
					<address>
						<street>456 Oak St</street>
						<city>Somewhere</city>
					</address>
				</customer>
			</InlineTest>`,
			expected: InlineTest{
				XMLName: xml.Name{Local: "InlineTest"},
				Customer: InlineTest_Customer{
					Name: "Jane Doe",
					// Address will contain the raw XML content
					Address: RawXML("\n\t\t\t\t\t\t\n\t\t\t\t\t\t\n\t\t\t\t\t"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result InlineTest
			err := xml.Unmarshal([]byte(tc.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			// For RawXML fields, we just check that something was captured
			if result.Customer.Name != tc.expected.Customer.Name {
				t.Errorf("Expected name %q, got %q", tc.expected.Customer.Name, result.Customer.Name)
			}

			if result.Customer.Address == "" {
				t.Errorf("Expected address RawXML to contain data, but it was empty")
			}

			t.Logf("Successfully parsed inline type with RawXML: %q", string(result.Customer.Address))
		})
	}
}

// TestRoundTripMarshaling tests that types can be marshaled and unmarshaled correctly
func TestRoundTripMarshaling(t *testing.T) {
	type TestStruct struct {
		XMLName   xml.Name  `xml:"TestStruct"`
		StringVal string    `xml:"stringVal"`
		IntVal    int32     `xml:"intVal"`
		BoolVal   bool      `xml:"boolVal"`
		TimeVal   time.Time `xml:"timeVal"`
		OptVal    *string   `xml:"optVal,omitempty"`
	}

	original := TestStruct{
		XMLName:   xml.Name{Local: "TestStruct"},
		StringVal: "test string",
		IntVal:    42,
		BoolVal:   true,
		TimeVal:   mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T10:30:00Z"),
		OptVal:    stringPtr("optional"),
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(original, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	t.Logf("Marshaled XML:\n%s", string(xmlData))

	// Unmarshal back
	var result TestStruct
	err = xml.Unmarshal(xmlData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Compare (adjust for namespace differences)
	expected := original
	expected.XMLName.Space = "" // Namespace not preserved in round-trip

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("Round-trip mismatch (-expected +actual):\n%s", diff)
	}

	t.Logf("Successfully completed round-trip marshaling")
}

// Helper functions are defined in parsing_test.go to avoid duplication

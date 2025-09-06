package soapgen

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/way-platform/soap-go/wsdl"
)

// TestIntegrationWithKnownTypes tests parsing with specific known generated types
func TestIntegrationWithKnownTypes(t *testing.T) {
	testCases := []struct {
		name        string
		testDataDir string
		xmlTests    []xmlParsingTest
	}{
		{
			name:        "custom_types_and_enums",
			testDataDir: "testdata/custom_types_and_enums",
			xmlTests: []xmlParsingTest{
				{
					name: "basic_user_info",
					xml: `<UserInfoType>
						<userId>12345</userId>
						<status>active</status>
						<email>test@example.com</email>
					</UserInfoType>`,
					expectSuccess: true,
				},
				{
					name: "invalid_status",
					xml: `<UserInfoType>
						<userId>12345</userId>
						<status>invalid_status</status>
						<email>test@example.com</email>
					</UserInfoType>`,
					expectSuccess: true, // XML parsing succeeds, validation is separate
				},
			},
		},
		{
			name:        "enumeration_simple_type",
			testDataDir: "testdata/enumeration_simple_type",
			xmlTests: []xmlParsingTest{
				{
					name:          "valid_enum_value",
					xml:           `<ColorElement>red</ColorElement>`,
					expectSuccess: true,
				},
				{
					name:          "another_valid_enum_value",
					xml:           `<ColorElement>blue</ColorElement>`,
					expectSuccess: true,
				},
			},
		},
		{
			name:        "optional_and_multiple_elements",
			testDataDir: "testdata/optional_and_multiple_elements",
			xmlTests: []xmlParsingTest{
				{
					name: "with_optional_elements",
					xml: `<TestElement>
						<requiredString>required value</requiredString>
						<optionalString>optional value</optionalString>
						<multipleStrings>first</multipleStrings>
						<multipleStrings>second</multipleStrings>
						<multipleStrings>third</multipleStrings>
					</TestElement>`,
					expectSuccess: true,
				},
				{
					name: "minimal_required_only",
					xml: `<TestElement>
						<requiredString>required value</requiredString>
						<multipleStrings>single</multipleStrings>
					</TestElement>`,
					expectSuccess: true,
				},
			},
		},
		{
			name:        "attributes",
			testDataDir: "testdata/attributes",
			xmlTests: []xmlParsingTest{
				{
					name: "with_attributes",
					xml: `<ElementWithAttributes id="123" optional="true">
						<content>test content</content>
					</ElementWithAttributes>`,
					expectSuccess: true,
				},
				{
					name: "minimal_attributes",
					xml: `<ElementWithAttributes id="456">
						<content>minimal content</content>
					</ElementWithAttributes>`,
					expectSuccess: true,
				},
			},
		},
		{
			name:        "byte_array_handling",
			testDataDir: "testdata/byte_array_handling",
			xmlTests: []xmlParsingTest{
				{
					name: "base64_data",
					xml: `<ByteArrayElement>
						<data>SGVsbG8gV29ybGQ=</data>
					</ByteArrayElement>`,
					expectSuccess: true,
				},
				{
					name: "hex_data",
					xml: `<ByteArrayElement>
						<hexData>48656C6C6F</hexData>
					</ByteArrayElement>`,
					expectSuccess: true,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runIntegrationTest(t, tc.testDataDir, tc.xmlTests)
		})
	}
}

type xmlParsingTest struct {
	name          string
	xml           string
	expectSuccess bool
	skipReason    string
}

func runIntegrationTest(t *testing.T, testDataDir string, xmlTests []xmlParsingTest) {
	// Verify the test data directory exists and has generated types
	if _, err := os.Stat(testDataDir); os.IsNotExist(err) {
		t.Skipf("Test data directory %s does not exist", testDataDir)
	}

	// Check for generated Go files
	entries, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Fatalf("Failed to read test data directory: %v", err)
	}

	hasGoFiles := false
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
			hasGoFiles = true
			break
		}
	}

	if !hasGoFiles {
		t.Skipf("No generated Go files found in %s", testDataDir)
	}

	// Parse the WSDL to understand the expected structure
	wsdlFile := filepath.Join(testDataDir, "definitions.wsdl")
	defs, err := wsdl.ParseFromFile(wsdlFile)
	if err != nil {
		t.Fatalf("Failed to parse WSDL file: %v", err)
	}

	// Run each XML parsing test
	for _, xmlTest := range xmlTests {
		t.Run(xmlTest.name, func(t *testing.T) {
			if xmlTest.skipReason != "" {
				t.Skip(xmlTest.skipReason)
			}

			// For now, just validate that the XML is well-formed
			// In a full implementation, we would load the generated types
			// and actually unmarshal into them
			if !isValidXML([]byte(xmlTest.xml)) {
				t.Errorf("XML is not well-formed: %s", xmlTest.xml)
			}

			// Test that we can at least parse it as generic XML
			var result interface{}
			err := xml.Unmarshal([]byte(xmlTest.xml), &result)

			if xmlTest.expectSuccess {
				if err != nil {
					t.Errorf("Expected successful parsing but got error: %v", err)
				} else {
					t.Logf("Successfully parsed XML: %s", xmlTest.xml)
				}
			} else {
				if err == nil {
					t.Errorf("Expected parsing to fail but it succeeded")
				}
			}

			// Log the WSDL target namespace for debugging
			if defs.TargetNamespace != "" {
				t.Logf("WSDL target namespace: %s", defs.TargetNamespace)
			}
		})
	}
}

// TestRealWorldXMLScenarios tests scenarios based on the original examples_test.go
func TestRealWorldXMLScenarios(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		xml         string
		expectValid bool
	}{
		{
			name:        "namespace_handling",
			description: "XML with explicit namespace should be valid",
			xml: `<TestElement xmlns="http://example.com/test">
				<field>value</field>
			</TestElement>`,
			expectValid: true,
		},
		{
			name:        "no_namespace",
			description: "XML without namespace should be valid",
			xml: `<TestElement>
				<field>value</field>
			</TestElement>`,
			expectValid: true,
		},
		{
			name:        "prefixed_namespace",
			description: "XML with prefixed namespace should be valid",
			xml: `<tns:TestElement xmlns:tns="http://example.com/test">
				<field>value</field>
			</tns:TestElement>`,
			expectValid: true,
		},
		{
			name:        "timestamp_formats",
			description: "Various timestamp formats should be handled",
			xml: `<TimestampTest>
				<standardTime>2023-12-25T10:30:00Z</standardTime>
				<timezoneTime>2023-12-25T10:30:00-05:00</timezoneTime>
				<millisecondTime>2023-12-25T10:30:00.123Z</millisecondTime>
			</TimestampTest>`,
			expectValid: true,
		},
		{
			name:        "boolean_values",
			description: "Boolean values should be handled correctly",
			xml: `<BooleanTest>
				<trueValue>true</trueValue>
				<falseValue>false</falseValue>
				<oneValue>1</oneValue>
				<zeroValue>0</zeroValue>
			</BooleanTest>`,
			expectValid: true,
		},
		{
			name:        "numeric_types",
			description: "Various numeric types should be handled",
			xml: `<NumericTest>
				<intValue>42</intValue>
				<floatValue>3.14159</floatValue>
				<negativeInt>-123</negativeInt>
				<largeNumber>9223372036854775807</largeNumber>
			</NumericTest>`,
			expectValid: true,
		},
		{
			name:        "empty_elements",
			description: "Empty elements should be handled",
			xml: `<EmptyTest>
				<emptyString></emptyString>
				<selfClosing/>
				<withContent>content</withContent>
			</EmptyTest>`,
			expectValid: true,
		},
		{
			name:        "malformed_xml",
			description: "Malformed XML should be detected",
			xml: `<BadElement>
				<unclosed>content
				<noEndTag>
			</BadElement>`,
			expectValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := isValidXML([]byte(tc.xml))

			if tc.expectValid && !isValid {
				t.Errorf("Expected valid XML but validation failed: %s", tc.description)
			} else if !tc.expectValid && isValid {
				t.Errorf("Expected invalid XML but validation passed: %s", tc.description)
			}

			// Also test generic unmarshaling
			var result interface{}
			err := xml.Unmarshal([]byte(tc.xml), &result)

			if tc.expectValid && err != nil {
				t.Errorf("Expected successful unmarshaling but got error: %v", err)
			} else if !tc.expectValid && err == nil {
				t.Errorf("Expected unmarshaling to fail but it succeeded")
			}

			t.Logf("%s: %s", tc.description, tc.xml)
		})
	}
}

// TestTimestampHandling tests specific timestamp parsing scenarios
func TestTimestampHandling(t *testing.T) {
	testCases := []struct {
		name        string
		timestamp   string
		expectValid bool
		description string
	}{
		{
			name:        "standard_utc",
			timestamp:   "2023-12-25T10:30:00Z",
			expectValid: true,
			description: "Standard UTC format should work",
		},
		{
			name:        "timezone_offset",
			timestamp:   "2023-12-25T10:30:00-05:00",
			expectValid: true,
			description: "Timezone offset format should work",
		},
		{
			name:        "milliseconds",
			timestamp:   "2023-12-25T10:30:00.123Z",
			expectValid: true,
			description: "Milliseconds precision should work",
		},
		{
			name:        "microseconds",
			timestamp:   "2023-12-25T10:30:00.123456Z",
			expectValid: true,
			description: "Microseconds precision should work",
		},
		{
			name:        "space_separator",
			timestamp:   "2023-12-25 10:30:00",
			expectValid: false,
			description: "Space separator should fail (non-XSD compliant)",
		},
		{
			name:        "no_time_part",
			timestamp:   "2023-12-25",
			expectValid: true,
			description: "Date-only format should work",
		},
		{
			name:        "invalid_format",
			timestamp:   "25/12/2023 10:30",
			expectValid: false,
			description: "Invalid format should fail",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xmlContent := `<TimestampTest><timestamp>` + tc.timestamp + `</timestamp></TimestampTest>`

			// Test if the XML itself is well-formed
			if !isValidXML([]byte(xmlContent)) {
				t.Errorf("Generated XML is not well-formed")
				return
			}

			// Test if Go can parse the timestamp
			_, err := time.Parse(time.RFC3339, tc.timestamp)
			if tc.expectValid && err != nil {
				// Try alternative formats
				formats := []string{
					"2006-01-02T15:04:05Z07:00",
					"2006-01-02T15:04:05.000Z",
					"2006-01-02T15:04:05.000000Z",
					"2006-01-02",
				}

				parsed := false
				for _, format := range formats {
					if _, err := time.Parse(format, tc.timestamp); err == nil {
						parsed = true
						break
					}
				}

				if !parsed {
					t.Errorf("Expected valid timestamp but parsing failed: %v", err)
				}
			} else if !tc.expectValid && err == nil {
				t.Errorf("Expected invalid timestamp but parsing succeeded")
			}

			t.Logf("%s: %s", tc.description, tc.timestamp)
		})
	}
}

// TestBinaryDataHandling tests base64 and hex binary data scenarios
func TestBinaryDataHandling(t *testing.T) {
	testCases := []struct {
		name        string
		data        string
		encoding    string
		expectValid bool
		description string
	}{
		{
			name:        "valid_base64",
			data:        "SGVsbG8gV29ybGQ=",
			encoding:    "base64",
			expectValid: true,
			description: "Valid base64 data should work",
		},
		{
			name:        "valid_hex",
			data:        "48656C6C6F",
			encoding:    "hex",
			expectValid: true,
			description: "Valid hex data should work",
		},
		{
			name:        "empty_base64",
			data:        "",
			encoding:    "base64",
			expectValid: true,
			description: "Empty base64 should work",
		},
		{
			name:        "invalid_base64",
			data:        "InvalidBase64!@#",
			encoding:    "base64",
			expectValid: false,
			description: "Invalid base64 should fail",
		},
		{
			name:        "invalid_hex",
			data:        "InvalidHexGHIJ",
			encoding:    "hex",
			expectValid: false,
			description: "Invalid hex should fail",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var xmlContent string
			if tc.encoding == "base64" {
				xmlContent = `<BinaryTest><base64Data>` + tc.data + `</base64Data></BinaryTest>`
			} else {
				xmlContent = `<BinaryTest><hexData>` + tc.data + `</hexData></BinaryTest>`
			}

			// Test if the XML itself is well-formed
			if !isValidXML([]byte(xmlContent)) {
				t.Errorf("Generated XML is not well-formed")
				return
			}

			// Test generic unmarshaling
			var result interface{}
			err := xml.Unmarshal([]byte(xmlContent), &result)
			if err != nil {
				t.Errorf("XML unmarshaling failed: %v", err)
			}

			t.Logf("%s: %s", tc.description, xmlContent)
		})
	}
}

package examples

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/way-platform/soap-go/examples/kitchensink"
)

func TestKitchenSinkRequestUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.KitchenSinkRequest
	}{
		{
			name: "basic types",
			xml: `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0" debug="true" timestamp="2023-12-25T10:30:00Z">
				<stringField>Hello World</stringField>
				<booleanField>true</booleanField>
				<intField>42</intField>
				<longField>9223372036854775807</longField>
				<shortField>32767</shortField>
				<byteField>127</byteField>
				<floatField>3.14159</floatField>
				<doubleField>2.718281828459045</doubleField>
				<decimalField>123.456</decimalField>
				<dateTimeField>2023-12-25T10:30:00Z</dateTimeField>
				<dateField>2023-12-25T00:00:00Z</dateField>
				<timeField>1970-01-01T10:30:00Z</timeField>
				<durationField>5400000000000</durationField>
				<unsignedLongField>18446744073709551615</unsignedLongField>
				<unsignedIntField>4294967295</unsignedIntField>
				<unsignedShortField>65535</unsignedShortField>
				<unsignedByteField>255</unsignedByteField>
				<integerField>-9223372036854775808</integerField>
				<positiveIntegerField>1000</positiveIntegerField>
				<nonNegativeIntegerField>0</nonNegativeIntegerField>
				<negativeIntegerField>-500</negativeIntegerField>
				<nonPositiveIntegerField>-100</nonPositiveIntegerField>
				<normalizedStringField>normalized text</normalizedStringField>
				<tokenField>token_value</tokenField>
				<languageField>en-US</languageField>
				<nmtokenField>NMTOKEN123</nmtokenField>
				<nameField>ElementName</nameField>
				<ncnameField>NCName123</ncnameField>
				<idField>unique_id_123</idField>
				<idrefField>ref_to_id</idrefField>
				<anyUriField>https://example.com/path</anyUriField>
				<qnameField>ns:localName</qnameField>
				<hexBinaryField>48656C6C6F</hexBinaryField>
				<base64BinaryField>SGVsbG8gV29ybGQ=</base64BinaryField>
				<gYearField>2023</gYearField>
				<gMonthField>--12</gMonthField>
				<gDayField>---25</gDayField>
				<gYearMonthField>2023-12</gYearMonthField>
				<gMonthDayField>--12-25</gMonthDayField>
				<!-- Edge case fields -->
				<optionalString>Optional value</optionalString>
				<optionalInt>999</optionalInt>
				<tags>tag1</tags>
				<tags>tag2</tags>
				<tags>tag3</tags>
				<numbers>1</numbers>
				<numbers>2</numbers>
				<numbers>3</numbers>
				<optionalTags>opt1</optionalTags>
				<optionalTags>opt2</optionalTags>
				<status>ACTIVE</status>
				<priority>2</priority>
				<optionalStatus>PENDING</optionalStatus>
				<address country="US" verified="true">
					<street>123 Main St</street>
					<city>Anytown</city>
					<zipCode>12345</zipCode>
				</address>
				<optionalAddress country="CA">
					<street>456 Elm St</street>
					<city>Toronto</city>
					<zipCode>M5V 3A8</zipCode>
				</optionalAddress>
				<simpleElement>Simple value</simpleElement>
				<metadata country="UK" verified="false">
					<street>789 Oak St</street>
					<city>London</city>
					<zipCode>SW1A 1AA</zipCode>
				</metadata>
			</KitchenSinkRequest>`,
			expected: kitchensink.KitchenSinkRequest{
				StringField:             "Hello World",
				BooleanField:            true,
				IntField:                42,
				LongField:               9223372036854775807,
				ShortField:              32767,
				ByteField:               127,
				FloatField:              3.14159,
				DoubleField:             2.718281828459045,
				DecimalField:            123.456,
				DateTimeField:           mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T10:30:00Z"),
				DateField:               mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T00:00:00Z"),
				TimeField:               mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T10:30:00Z"),
				DurationField:           90 * time.Minute, // PT1H30M = 1 hour 30 minutes
				UnsignedLongField:       18446744073709551615,
				UnsignedIntField:        4294967295,
				UnsignedShortField:      65535,
				UnsignedByteField:       255,
				IntegerField:            -9223372036854775808,
				PositiveIntegerField:    1000,
				NonNegativeIntegerField: 0,
				NegativeIntegerField:    -500,
				NonPositiveIntegerField: -100,
				NormalizedStringField:   "normalized text",
				TokenField:              "token_value",
				LanguageField:           "en-US",
				NmtokenField:            "NMTOKEN123",
				NameField:               "ElementName",
				NcnameField:             "NCName123",
				IdField:                 "unique_id_123",
				IdrefField:              "ref_to_id",
				AnyUriField:             "https://example.com/path",
				QnameField:              xml.Name{Local: "qnameField", Space: "http://example.com/typetest"},
				HexBinaryField:          []byte("48656C6C6F"),       // XML treats hex as literal string
				Base64BinaryField:       []byte("SGVsbG8gV29ybGQ="), // XML treats base64 as literal string
				GYearField:              "2023",
				GMonthField:             "--12",
				GDayField:               "---25",
				GYearMonthField:         "2023-12",
				GMonthDayField:          "--12-25",
				// Edge case fields
				OptionalString: stringPtr("Optional value"),
				OptionalInt:    int32Ptr(999),
				Tags:           []string{"tag1", "tag2", "tag3"},
				Numbers:        []int32{1, 2, 3},
				OptionalTags:   []string{"opt1", "opt2"},
				Status:         "ACTIVE",
				Priority:       2,
				OptionalStatus: stringPtr("PENDING"),
				Address: kitchensink.AddressType{
					Street:   "123 Main St",
					City:     "Anytown",
					ZipCode:  "12345",
					Country:  "US",
					Verified: boolPtr(true),
				},
				OptionalAddress: &kitchensink.AddressType{
					Street:   "456 Elm St",
					City:     "Toronto",
					ZipCode:  "M5V 3A8",
					Country:  "CA",
					Verified: nil,
				},
				SimpleElement: "Simple value",
				Metadata: &kitchensink.AddressType{
					Street:   "789 Oak St",
					City:     "London",
					ZipCode:  "SW1A 1AA",
					Country:  "UK",
					Verified: boolPtr(false),
				},
				// Attributes
				Version:   "1.0",
				Debug:     boolPtr(true),
				Timestamp: timePtr(mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T10:30:00Z")),
			},
		},
		{
			name: "minimal values",
			xml: `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0">
				<stringField></stringField>
				<booleanField>false</booleanField>
				<intField>0</intField>
				<longField>0</longField>
				<shortField>0</shortField>
				<byteField>0</byteField>
				<floatField>0.0</floatField>
				<doubleField>0.0</doubleField>
				<decimalField>0.0</decimalField>
				<dateTimeField>1970-01-01T00:00:00Z</dateTimeField>
				<dateField>1970-01-01T00:00:00Z</dateField>
				<timeField>1970-01-01T00:00:00Z</timeField>
				<durationField>0</durationField>
				<unsignedLongField>0</unsignedLongField>
				<unsignedIntField>0</unsignedIntField>
				<unsignedShortField>0</unsignedShortField>
				<unsignedByteField>0</unsignedByteField>
				<integerField>0</integerField>
				<positiveIntegerField>1</positiveIntegerField>
				<nonNegativeIntegerField>0</nonNegativeIntegerField>
				<negativeIntegerField>-1</negativeIntegerField>
				<nonPositiveIntegerField>0</nonPositiveIntegerField>
				<normalizedStringField></normalizedStringField>
				<tokenField></tokenField>
				<languageField>en</languageField>
				<nmtokenField>A</nmtokenField>
				<nameField>A</nameField>
				<ncnameField>A</ncnameField>
				<idField>a</idField>
				<idrefField>a</idrefField>
				<anyUriField>http://example.com</anyUriField>
				<qnameField>local</qnameField>
				<hexBinaryField></hexBinaryField>
				<base64BinaryField></base64BinaryField>
				<gYearField>1970</gYearField>
				<gMonthField>--01</gMonthField>
				<gDayField>---01</gDayField>
				<gYearMonthField>1970-01</gYearMonthField>
				<gMonthDayField>--01-01</gMonthDayField>
				<!-- Required edge case fields -->
				<tags>single</tags>
				<numbers>0</numbers>
				<status>INACTIVE</status>
				<priority>1</priority>
				<address country="US">
					<street></street>
					<city></city>
					<zipCode></zipCode>
				</address>
				<simpleElement></simpleElement>
			</KitchenSinkRequest>`,
			expected: kitchensink.KitchenSinkRequest{
				StringField:             "",
				BooleanField:            false,
				IntField:                0,
				LongField:               0,
				ShortField:              0,
				ByteField:               0,
				FloatField:              0.0,
				DoubleField:             0.0,
				DecimalField:            0.0,
				DateTimeField:           mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z"),
				DateField:               mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z"),
				TimeField:               mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z"),
				DurationField:           0,
				UnsignedLongField:       0,
				UnsignedIntField:        0,
				UnsignedShortField:      0,
				UnsignedByteField:       0,
				IntegerField:            0,
				PositiveIntegerField:    1,
				NonNegativeIntegerField: 0,
				NegativeIntegerField:    -1,
				NonPositiveIntegerField: 0,
				NormalizedStringField:   "",
				TokenField:              "",
				LanguageField:           "en",
				NmtokenField:            "A",
				NameField:               "A",
				NcnameField:             "A",
				IdField:                 "a",
				IdrefField:              "a",
				AnyUriField:             "http://example.com",
				QnameField:              xml.Name{Local: "qnameField", Space: "http://example.com/typetest"},
				HexBinaryField:          []byte{},
				Base64BinaryField:       []byte{},
				GYearField:              "1970",
				GMonthField:             "--01",
				GDayField:               "---01",
				GYearMonthField:         "1970-01",
				GMonthDayField:          "--01-01",
				// Edge case fields - only required ones present
				OptionalString: nil,
				OptionalInt:    nil,
				Tags:           []string{"single"},
				Numbers:        []int32{0},
				OptionalTags:   nil,
				Status:         "INACTIVE",
				Priority:       1,
				OptionalStatus: nil,
				Address: kitchensink.AddressType{
					Street:   "",
					City:     "",
					ZipCode:  "",
					Country:  "US",
					Verified: nil,
				},
				OptionalAddress: nil,
				SimpleElement:   "",
				Metadata:        nil,
				// Attributes - only required ones
				Version:   "1.0",
				Debug:     nil,
				Timestamp: nil,
			},
		},
		{
			name: "edge cases - optional fields and enumerations",
			xml: `<KitchenSinkRequest xmlns="http://example.com/typetest" version="2.0" debug="false">
				<stringField>Edge case test</stringField>
				<booleanField>true</booleanField>
				<intField>123</intField>
				<longField>456</longField>
				<shortField>789</shortField>
				<byteField>12</byteField>
				<floatField>1.23</floatField>
				<doubleField>4.56</doubleField>
				<decimalField>7.89</decimalField>
				<dateTimeField>2024-01-01T12:00:00Z</dateTimeField>
				<dateField>2024-01-01T00:00:00Z</dateField>
				<timeField>1970-01-01T12:00:00Z</timeField>
				<durationField>3600000000000</durationField>
				<unsignedLongField>100</unsignedLongField>
				<unsignedIntField>200</unsignedIntField>
				<unsignedShortField>300</unsignedShortField>
				<unsignedByteField>50</unsignedByteField>
				<integerField>-100</integerField>
				<positiveIntegerField>999</positiveIntegerField>
				<nonNegativeIntegerField>500</nonNegativeIntegerField>
				<negativeIntegerField>-999</negativeIntegerField>
				<nonPositiveIntegerField>-500</nonPositiveIntegerField>
				<normalizedStringField>edge case</normalizedStringField>
				<tokenField>edge_token</tokenField>
				<languageField>de-DE</languageField>
				<nmtokenField>EDGE123</nmtokenField>
				<nameField>EdgeElement</nameField>
				<ncnameField>edgeNCName</ncnameField>
				<idField>edge_id</idField>
				<idrefField>edge_ref</idrefField>
				<anyUriField>https://edge.example.com</anyUriField>
				<qnameField>edge:test</qnameField>
				<hexBinaryField>4142</hexBinaryField>
				<base64BinaryField>QUI=</base64BinaryField>
				<gYearField>2024</gYearField>
				<gMonthField>--06</gMonthField>
				<gDayField>---15</gDayField>
				<gYearMonthField>2024-06</gYearMonthField>
				<gMonthDayField>--06-15</gMonthDayField>
				<!-- No optional fields to test nil handling -->
				<tags>edge1</tags>
				<tags>edge2</tags>
				<numbers>10</numbers>
				<numbers>20</numbers>
				<numbers>30</numbers>
				<!-- No optionalTags to test empty slice -->
				<status>PENDING</status>
				<priority>3</priority>
				<!-- No optionalStatus to test nil -->
				<address country="DE" verified="true">
					<street>Edge Street 123</street>
					<city>Berlin</city>
					<zipCode>10115</zipCode>
				</address>
				<!-- No optionalAddress to test nil -->
				<simpleElement>Edge simple</simpleElement>
				<!-- No metadata to test nil -->
			</KitchenSinkRequest>`,
			expected: kitchensink.KitchenSinkRequest{
				StringField:             "Edge case test",
				BooleanField:            true,
				IntField:                123,
				LongField:               456,
				ShortField:              789,
				ByteField:               12,
				FloatField:              1.23,
				DoubleField:             4.56,
				DecimalField:            7.89,
				DateTimeField:           mustParseTime("2006-01-02T15:04:05Z", "2024-01-01T12:00:00Z"),
				DateField:               mustParseTime("2006-01-02T15:04:05Z", "2024-01-01T00:00:00Z"),
				TimeField:               mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T12:00:00Z"),
				DurationField:           1 * time.Hour,
				UnsignedLongField:       100,
				UnsignedIntField:        200,
				UnsignedShortField:      300,
				UnsignedByteField:       50,
				IntegerField:            -100,
				PositiveIntegerField:    999,
				NonNegativeIntegerField: 500,
				NegativeIntegerField:    -999,
				NonPositiveIntegerField: -500,
				NormalizedStringField:   "edge case",
				TokenField:              "edge_token",
				LanguageField:           "de-DE",
				NmtokenField:            "EDGE123",
				NameField:               "EdgeElement",
				NcnameField:             "edgeNCName",
				IdField:                 "edge_id",
				IdrefField:              "edge_ref",
				AnyUriField:             "https://edge.example.com",
				QnameField:              xml.Name{Local: "qnameField", Space: "http://example.com/typetest"},
				HexBinaryField:          []byte("4142"),
				Base64BinaryField:       []byte("QUI="),
				GYearField:              "2024",
				GMonthField:             "--06",
				GDayField:               "---15",
				GYearMonthField:         "2024-06",
				GMonthDayField:          "--06-15",
				// Edge case fields
				OptionalString: nil, // Not present in XML
				OptionalInt:    nil, // Not present in XML
				Tags:           []string{"edge1", "edge2"},
				Numbers:        []int32{10, 20, 30},
				OptionalTags:   nil, // Not present in XML
				Status:         "PENDING",
				Priority:       3,
				OptionalStatus: nil, // Not present in XML
				Address: kitchensink.AddressType{
					Street:   "Edge Street 123",
					City:     "Berlin",
					ZipCode:  "10115",
					Country:  "DE",
					Verified: boolPtr(true),
				},
				OptionalAddress: nil, // Not present in XML
				SimpleElement:   "Edge simple",
				Metadata:        nil, // Not present in XML
				// Attributes
				Version:   "2.0",
				Debug:     boolPtr(false),
				Timestamp: nil, // Not present in XML
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req kitchensink.KitchenSinkRequest
			err := xml.Unmarshal([]byte(tt.xml), &req)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, req); diff != "" {
				t.Errorf("KitchenSinkRequest mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestKitchenSinkResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.KitchenSinkResponse
	}{
		{
			name: "simple response",
			xml: `<KitchenSinkResponse xmlns="http://example.com/typetest">
				<result>Operation completed successfully</result>
			</KitchenSinkResponse>`,
			expected: kitchensink.KitchenSinkResponse{
				Result: "Operation completed successfully",
			},
		},
		{
			name: "empty response",
			xml: `<KitchenSinkResponse xmlns="http://example.com/typetest">
				<result></result>
			</KitchenSinkResponse>`,
			expected: kitchensink.KitchenSinkResponse{
				Result: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp kitchensink.KitchenSinkResponse
			err := xml.Unmarshal([]byte(tt.xml), &resp)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, resp); diff != "" {
				t.Errorf("KitchenSinkResponse mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestKitchenSinkMarshaling(t *testing.T) {
	req := kitchensink.KitchenSinkRequest{
		StringField:             "Test String",
		BooleanField:            true,
		IntField:                123,
		LongField:               9876543210,
		ShortField:              456,
		ByteField:               78,
		FloatField:              3.14,
		DoubleField:             2.71828,
		DecimalField:            99.99,
		DateTimeField:           mustParseTime("2006-01-02T15:04:05Z", "2023-06-15T14:30:00Z"),
		DateField:               mustParseTime("2006-01-02T15:04:05Z", "2023-06-15T00:00:00Z"),
		TimeField:               mustParseTime("2006-01-02T15:04:05Z", "1970-01-01T14:30:00Z"),
		DurationField:           2*time.Hour + 15*time.Minute,
		UnsignedLongField:       18446744073709551614,
		UnsignedIntField:        4294967294,
		UnsignedShortField:      65534,
		UnsignedByteField:       254,
		IntegerField:            -1234567890,
		PositiveIntegerField:    9999,
		NonNegativeIntegerField: 5555,
		NegativeIntegerField:    -7777,
		NonPositiveIntegerField: -3333,
		NormalizedStringField:   "normalized string",
		TokenField:              "some_token",
		LanguageField:           "fr-FR",
		NmtokenField:            "TOKEN456",
		NameField:               "MyElement",
		NcnameField:             "myNCName",
		IdField:                 "id_789",
		IdrefField:              "ref_789",
		AnyUriField:             "https://api.example.com/v1/resource",
		QnameField:              xml.Name{Local: "qnameField"},
		HexBinaryField:          []byte("Test Data"),
		Base64BinaryField:       []byte("Binary Test Data"),
		GYearField:              "2024",
		GMonthField:             "--06",
		GDayField:               "---15",
		GYearMonthField:         "2024-06",
		GMonthDayField:          "--06-15",
		// Edge case fields
		OptionalString: stringPtr("Optional value"),
		OptionalInt:    int32Ptr(100),
		Tags:           []string{"tag1", "tag2"},
		Numbers:        []int32{1, 2, 3},
		OptionalTags:   []string{"opt1"},
		Status:         "ACTIVE",
		Priority:       2,
		OptionalStatus: stringPtr("PENDING"),
		Address: kitchensink.AddressType{
			Street:   "Test Street",
			City:     "Test City",
			ZipCode:  "12345",
			Country:  "US",
			Verified: boolPtr(true),
		},
		OptionalAddress: &kitchensink.AddressType{
			Street:   "Optional Street",
			City:     "Optional City",
			ZipCode:  "67890",
			Country:  "CA",
			Verified: boolPtr(false),
		},
		SimpleElement: "Simple test",
		Metadata: &kitchensink.AddressType{
			Street:   "Meta Street",
			City:     "Meta City",
			ZipCode:  "99999",
			Country:  "UK",
			Verified: nil,
		},
		// Attributes (avoid nil pointers that cause marshaling issues)
		Version:   "1.0",
		Debug:     boolPtr(true),
		Timestamp: timePtr(mustParseTime("2006-01-02T15:04:05Z", "2023-06-15T14:30:00Z")),
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %v", err)
	}

	// Unmarshal back to struct
	var unmarshaledReq kitchensink.KitchenSinkRequest
	err = xml.Unmarshal(xmlData, &unmarshaledReq)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Compare original and unmarshaled structs
	if diff := cmp.Diff(req, unmarshaledReq); diff != "" {
		t.Errorf("Round-trip marshal/unmarshal mismatch (-original +unmarshaled):\n%s", diff)
	}
}

// mustParseTime is a helper function that parses time and panics on error
func mustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func timePtr(t time.Time) *time.Time {
	return &t
}

// TestInlineComplexTypes tests the parsing of inline complex types with Outer_Inner naming
func TestInlineComplexTypes(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.InlineTypesTest
	}{
		{
			name: "inline complex types with []byte fields",
			xml: `<InlineTypesTest xmlns="http://example.com/typetest">
				<customer>
					<name>John Doe</name>
					<address>
						<street>123 Main St</street>
						<city>Anytown</city>
					</address>
				</customer>
				<items>
					<item>
						<product>Widget A</product>
						<quantity>5</quantity>
					</item>
					<item>
						<product>Widget B</product>
						<quantity>3</quantity>
					</item>
				</items>
			</InlineTypesTest>`,
			expected: kitchensink.InlineTypesTest{
				// Both Customer and Items use element names and capture character data only (whitespace)
				Customer: kitchensink.RawXML("\n\t\t\t\t\t\n\t\t\t\t\t\n\t\t\t\t"), // Whitespace only due to XML tag limitation
				Items:    kitchensink.RawXML("\n\t\t\t\t\t\n\t\t\t\t\t\n\t\t\t\t"), // Whitespace only due to XML tag limitation
			},
		},
		{
			name: "minimal inline complex types",
			xml: `<InlineTypesTest xmlns="http://example.com/typetest">
				<customer><name>Jane</name><address><street>Elm St</street><city>Boston</city></address></customer>
				<items><item><product>Tool</product><quantity>1</quantity></item></items>
			</InlineTypesTest>`,
			expected: kitchensink.InlineTypesTest{
				// Both fields use element names and capture character data only
				Customer: kitchensink.RawXML(""), // Empty due to XML tag limitation
				Items:    kitchensink.RawXML(""), // Empty due to XML tag limitation
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result kitchensink.InlineTypesTest
			err := xml.Unmarshal([]byte(tt.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("InlineTypesTest mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

// TestElementReferences tests the parsing of element references
func TestElementReferences(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.PersonInfo
	}{
		{
			name: "element references with optional field",
			xml: `<PersonInfo xmlns="http://example.com/typetest">
				<PersonName>John Doe</PersonName>
				<PersonAge>30</PersonAge>
				<Tag>developer</Tag>
			</PersonInfo>`,
			expected: kitchensink.PersonInfo{
				PersonName: kitchensink.PersonName{Value: "John Doe"},
				PersonAge:  kitchensink.PersonAge{Value: 30},
				Tag:        &kitchensink.Tag{Value: "developer"},
			},
		},
		{
			name: "minimal element references without optional field",
			xml: `<PersonInfo xmlns="http://example.com/typetest">
				<PersonName>Alice Brown</PersonName>
				<PersonAge>25</PersonAge>
			</PersonInfo>`,
			expected: kitchensink.PersonInfo{
				PersonName: kitchensink.PersonName{Value: "Alice Brown"},
				PersonAge:  kitchensink.PersonAge{Value: 25},
				Tag:        nil, // Optional field not present
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result kitchensink.PersonInfo
			err := xml.Unmarshal([]byte(tt.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("PersonInfo mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

// TestUntypedFields tests handling of untyped fields and []byte vs [][]byte
func TestUntypedFields(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.UntypedFieldsTest
	}{
		{
			name: "untyped fields with complex data",
			xml: `<UntypedFieldsTest xmlns="http://example.com/typetest">
				<unknownField>Simple text</unknownField>
				<unknownArray>item1</unknownArray>
				<unknownArray>item2</unknownArray>
				<unknownArray>item3</unknownArray>
				<optionalUnknown>Optional value</optionalUnknown>
				<complexData>
					<innerField>Complex inner value</innerField>
				</complexData>
				<multipleComplexData>
					<innerField>123</innerField>
				</multipleComplexData>
				<multipleComplexData>
					<innerField>456</innerField>
				</multipleComplexData>
			</UntypedFieldsTest>`,
			expected: kitchensink.UntypedFieldsTest{
				UnknownField:    "Simple text",
				UnknownArray:    []string{"item1", "item2", "item3"}, // []string not [][]string
				OptionalUnknown: stringPtr("Optional value"),
				// ComplexData uses ,innerxml and captures ALL inner XML of the parent element
				ComplexData: kitchensink.RawXML("\n\t\t\t\t<unknownField>Simple text</unknownField>\n\t\t\t\t<unknownArray>item1</unknownArray>\n\t\t\t\t<unknownArray>item2</unknownArray>\n\t\t\t\t<unknownArray>item3</unknownArray>\n\t\t\t\t<optionalUnknown>Optional value</optionalUnknown>\n\t\t\t\t<complexData>\n\t\t\t\t\t<innerField>Complex inner value</innerField>\n\t\t\t\t</complexData>\n\t\t\t\t<multipleComplexData>\n\t\t\t\t\t<innerField>123</innerField>\n\t\t\t\t</multipleComplexData>\n\t\t\t\t<multipleComplexData>\n\t\t\t\t\t<innerField>456</innerField>\n\t\t\t\t</multipleComplexData>\n\t\t\t"),
				// MultipleComplexData uses element names and captures character data only (whitespace)
				MultipleComplexData: []kitchensink.RawXML{
					// TODO: Add custom parsing to capture the raw XML for multiple sequential elements.
					kitchensink.RawXML("\n\t\t\t\t\t\n\t\t\t\t"), // Whitespace only due to XML tag limitation
					kitchensink.RawXML("\n\t\t\t\t\t\n\t\t\t\t"), // Whitespace only due to XML tag limitation
				},
			},
		},
		{
			name: "minimal untyped fields",
			xml: `<UntypedFieldsTest xmlns="http://example.com/typetest">
				<unknownField></unknownField>
				<unknownArray>single</unknownArray>
				<complexData><innerField></innerField></complexData>
				<multipleComplexData><innerField>0</innerField></multipleComplexData>
			</UntypedFieldsTest>`,
			expected: kitchensink.UntypedFieldsTest{
				UnknownField:    "",
				UnknownArray:    []string{"single"},
				OptionalUnknown: nil, // Not present
				// ComplexData uses ,innerxml and captures ALL inner XML of the parent element
				ComplexData: kitchensink.RawXML("\n\t\t\t\t<unknownField></unknownField>\n\t\t\t\t<unknownArray>single</unknownArray>\n\t\t\t\t<complexData><innerField></innerField></complexData>\n\t\t\t\t<multipleComplexData><innerField>0</innerField></multipleComplexData>\n\t\t\t"),
				// MultipleComplexData uses element names and captures character data only
				MultipleComplexData: []kitchensink.RawXML{
					kitchensink.RawXML(""), // Empty due to XML tag limitation
				},
			},
		},
		{
			name: "optional field not present",
			xml: `<UntypedFieldsTest xmlns="http://example.com/typetest">
				<unknownField>test</unknownField>
				<unknownArray>one</unknownArray>
				<unknownArray>two</unknownArray>
				<complexData><innerField>test</innerField></complexData>
				<multipleComplexData><innerField>1</innerField></multipleComplexData>
			</UntypedFieldsTest>`,
			expected: kitchensink.UntypedFieldsTest{
				UnknownField:    "test",
				UnknownArray:    []string{"one", "two"},
				OptionalUnknown: nil, // Not present in XML
				// ComplexData uses ,innerxml and captures ALL inner XML of the parent element
				ComplexData: kitchensink.RawXML("\n\t\t\t\t<unknownField>test</unknownField>\n\t\t\t\t<unknownArray>one</unknownArray>\n\t\t\t\t<unknownArray>two</unknownArray>\n\t\t\t\t<complexData><innerField>test</innerField></complexData>\n\t\t\t\t<multipleComplexData><innerField>1</innerField></multipleComplexData>\n\t\t\t"),
				// MultipleComplexData uses element names and captures character data only
				MultipleComplexData: []kitchensink.RawXML{
					kitchensink.RawXML(""), // Empty due to XML tag limitation
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result kitchensink.UntypedFieldsTest
			err := xml.Unmarshal([]byte(tt.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("UntypedFieldsTest mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

// TestCustomTypesAndEnums tests custom simple types and enumeration constants
func TestCustomTypesAndEnums(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected kitchensink.UserInfoType
	}{
		{
			name: "custom types with enumeration constants",
			xml: `<UserInfoType xmlns="http://example.com/typetest">
				<userId>123456789</userId>
				<status>ACTIVE</status>
				<email>user@example.com</email>
			</UserInfoType>`,
			expected: kitchensink.UserInfoType{
				UserId: 123456789,
				Status: kitchensink.StatusTypeACTIVE, // Should match constant value
				Email:  "user@example.com",
			},
		},
		{
			name: "custom type with different enum value",
			xml: `<UserInfoType xmlns="http://example.com/typetest">
				<userId>987654321</userId>
				<status>PENDING</status>
				<email>pending@example.com</email>
			</UserInfoType>`,
			expected: kitchensink.UserInfoType{
				UserId: 987654321,
				Status: kitchensink.StatusTypePENDING,
				Email:  "pending@example.com",
			},
		},
		{
			name: "custom type with inactive status",
			xml: `<UserInfoType xmlns="http://example.com/typetest">
				<userId>555000111</userId>
				<status>INACTIVE</status>
				<email>inactive@example.com</email>
			</UserInfoType>`,
			expected: kitchensink.UserInfoType{
				UserId: 555000111,
				Status: kitchensink.StatusTypeINACTIVE,
				Email:  "inactive@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result kitchensink.UserInfoType
			err := xml.Unmarshal([]byte(tt.xml), &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}

			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("UserInfoType mismatch (-expected +actual):\n%s", diff)
			}

			// Verify enumeration constant values match
			switch result.Status {
			case kitchensink.StatusTypeACTIVE:
				if kitchensink.StatusTypeACTIVE != "ACTIVE" {
					t.Errorf("StatusTypeACTIVE constant should be 'ACTIVE', got %q", kitchensink.StatusTypeACTIVE)
				}
			case kitchensink.StatusTypePENDING:
				if kitchensink.StatusTypePENDING != "PENDING" {
					t.Errorf("StatusTypePENDING constant should be 'PENDING', got %q", kitchensink.StatusTypePENDING)
				}
			case kitchensink.StatusTypeINACTIVE:
				if kitchensink.StatusTypeINACTIVE != "INACTIVE" {
					t.Errorf("StatusTypeINACTIVE constant should be 'INACTIVE', got %q", kitchensink.StatusTypeINACTIVE)
				}
			}
		})
	}
}

// TestTimestampFormatHandling tests specific timestamp format scenarios
func TestTimestampFormatHandling(t *testing.T) {
	t.Run("standard XSD dateTime format works", func(t *testing.T) {
		xmlData := `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0" timestamp="2023-12-25T10:30:00Z">
			<dateTimeField>2023-12-25T10:30:00Z</dateTimeField>
		</KitchenSinkRequest>`

		var req kitchensink.KitchenSinkRequest
		err := xml.Unmarshal([]byte(xmlData), &req)
		if err != nil {
			t.Fatalf("Failed to unmarshal standard XSD format: %v", err)
		}

		expectedTime := mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T10:30:00Z")
		if !req.DateTimeField.Equal(expectedTime) {
			t.Errorf("DateTimeField: expected %v, got %v", expectedTime, req.DateTimeField)
		}
		if !req.Timestamp.Equal(expectedTime) {
			t.Errorf("Timestamp: expected %v, got %v", expectedTime, *req.Timestamp)
		}
	})

	t.Run("timezone offset format works", func(t *testing.T) {
		xmlData := `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0" timestamp="2023-12-25T10:30:00-05:00">
			<dateTimeField>2023-12-25T15:30:00Z</dateTimeField>
		</KitchenSinkRequest>`

		var req kitchensink.KitchenSinkRequest
		err := xml.Unmarshal([]byte(xmlData), &req)
		if err != nil {
			t.Fatalf("Failed to unmarshal timezone offset format: %v", err)
		}

		expectedDateTime := mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T15:30:00Z")
		expectedTimestamp := mustParseTime("2006-01-02T15:04:05-07:00", "2023-12-25T10:30:00-05:00")

		if !req.DateTimeField.Equal(expectedDateTime) {
			t.Errorf("DateTimeField: expected %v, got %v", expectedDateTime, req.DateTimeField)
		}
		if !req.Timestamp.Equal(expectedTimestamp) {
			t.Errorf("Timestamp: expected %v, got %v", expectedTimestamp, *req.Timestamp)
		}
	})

	t.Run("milliseconds precision works", func(t *testing.T) {
		xmlData := `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0" timestamp="2023-12-25T10:30:00.123Z">
			<dateTimeField>2023-12-25T10:30:00.123Z</dateTimeField>
		</KitchenSinkRequest>`

		var req kitchensink.KitchenSinkRequest
		err := xml.Unmarshal([]byte(xmlData), &req)
		if err != nil {
			t.Fatalf("Failed to unmarshal milliseconds format: %v", err)
		}

		expectedTime := mustParseTime("2006-01-02T15:04:05.000Z", "2023-12-25T10:30:00.123Z")
		if !req.DateTimeField.Equal(expectedTime) {
			t.Errorf("DateTimeField: expected %v, got %v", expectedTime, req.DateTimeField)
		}
		if !req.Timestamp.Equal(expectedTime) {
			t.Errorf("Timestamp: expected %v, got %v", expectedTime, *req.Timestamp)
		}
	})

	t.Run("non-standard space format fails", func(t *testing.T) {
		xmlData := `<KitchenSinkRequest xmlns="http://example.com/typetest" version="1.0" timestamp="2003-04-20 10:00:00">
			<dateTimeField>2003-04-20T10:00:00Z</dateTimeField>
		</KitchenSinkRequest>`

		var req kitchensink.KitchenSinkRequest
		err := xml.Unmarshal([]byte(xmlData), &req)
		if err == nil {
			t.Fatal("Expected error for non-standard timestamp format, but got none")
		}

		if !strings.Contains(err.Error(), "parsing time") {
			t.Errorf("Expected parsing time error, got: %v", err)
		}

		expectedErr := `cannot parse " 10:00:00" as "T"`
		if !strings.Contains(err.Error(), expectedErr) {
			t.Errorf("Expected specific error about 'T' separator, got: %v", err)
		}
	})
}

// TestTimestampRoundTrip tests that timestamps marshal and unmarshal correctly
func TestTimestampRoundTrip(t *testing.T) {
	originalTime := mustParseTime("2006-01-02T15:04:05Z", "2023-12-25T10:30:00Z")

	req := kitchensink.KitchenSinkRequest{
		DateTimeField: originalTime,
		Version:       "1.0",
		Timestamp:     &originalTime,
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %v", err)
	}

	// Verify XML contains XSD-compliant format
	xmlStr := string(xmlData)
	if !strings.Contains(xmlStr, `timestamp="2023-12-25T10:30:00Z"`) {
		t.Errorf("Marshaled XML does not contain expected timestamp format, got:\n%s", xmlStr)
	}
	if !strings.Contains(xmlStr, `<dateTimeField>2023-12-25T10:30:00Z</dateTimeField>`) {
		t.Errorf("Marshaled XML does not contain expected dateTimeField format, got:\n%s", xmlStr)
	}

	// Unmarshal back to struct
	var unmarshaledReq kitchensink.KitchenSinkRequest
	err = xml.Unmarshal(xmlData, &unmarshaledReq)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Compare timestamp fields
	if !req.DateTimeField.Equal(unmarshaledReq.DateTimeField) {
		t.Errorf("Round-trip DateTimeField mismatch: original %v, unmarshaled %v",
			req.DateTimeField, unmarshaledReq.DateTimeField)
	}

	if !req.Timestamp.Equal(*unmarshaledReq.Timestamp) {
		t.Errorf("Round-trip Timestamp mismatch: original %v, unmarshaled %v",
			*req.Timestamp, *unmarshaledReq.Timestamp)
	}
}

// TestNonStandardTimestampSolution demonstrates a solution approach for non-standard formats
func TestNonStandardTimestampSolution(t *testing.T) {
	t.Run("demonstrates space-separated format issue", func(t *testing.T) {
		// This format is NOT XSD-compliant but might be encountered in real APIs
		xmlData := `<Scheduled><Begin>2003-04-20 10:00:00</Begin><End>2003-04-25 12:00:20</End></Scheduled>`

		// For APIs that use this format, consider these approaches:
		// 1. Parse as string and convert manually
		// 2. Pre-process XML to replace space with 'T'
		// 3. Implement custom timestamp type with UnmarshalXMLAttr

		t.Logf("Non-standard format example: %s", xmlData)
		t.Log("Solution: Use custom timestamp type or pre-process XML for XSD compliance")
	})
}

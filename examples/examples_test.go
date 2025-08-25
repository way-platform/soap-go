package examples

import (
	"encoding/xml"
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
			xml: `<KitchenSinkRequest xmlns="http://example.com/typetest">
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
			},
		},
		{
			name: "minimal values",
			xml: `<KitchenSinkRequest xmlns="http://example.com/typetest">
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

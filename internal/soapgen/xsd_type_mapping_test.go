package soapgen

import (
	"testing"

	"github.com/way-platform/soap-go/xsd"
)

func TestMapXSDTypeToGo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		typ      xsd.Type
		expected string
	}{
		// String types
		{xsd.String, "string"},
		{xsd.NormalizedString, "string"},
		{xsd.Token, "string"},
		{xsd.NCName, "string"},
		{xsd.ID, "string"},
		{xsd.AnyURI, "string"},

		// String slice types
		{xsd.NMTOKENS, "[]string"},
		{xsd.IDREFS, "[]string"},

		// Boolean
		{xsd.Boolean, "bool"},

		// IEEE 754 floating point types
		{xsd.Float, "float64"},
		{xsd.Double, "float64"},
		// Arbitrary-precision decimal maps to string
		{xsd.Decimal, "string"},

		// Signed integers
		{xsd.Byte, "int8"},
		{xsd.Short, "int16"},
		{xsd.Int, "int32"},
		{xsd.Long, "int64"},
		{xsd.Integer, "int64"},

		// Unsigned integers
		{xsd.UnsignedByte, "uint8"},
		{xsd.UnsignedShort, "uint16"},
		{xsd.UnsignedInt, "uint32"},
		{xsd.UnsignedLong, "uint64"},

		// Time types
		{xsd.DateTime, "time.Time"},
		{xsd.Time, "time.Time"},
		{xsd.Date, "time.Time"},
		{xsd.Duration, "string"},
		{xsd.GYear, "string"},

		// Binary types
		{xsd.HexBinary, "[]byte"},
		{xsd.Base64Binary, "[]byte"},

		// Special types
		{xsd.QName, "xml.Name"},
		{xsd.NOTATION, "string"},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := mapXSDTypeToGo(tt.typ); got != tt.expected {
				t.Errorf("mapXSDTypeToGo() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetRequiredImports(t *testing.T) {
	t.Parallel()
	tests := []struct {
		typ      xsd.Type
		expected []string
	}{
		{xsd.String, nil},
		{xsd.Boolean, nil},
		{xsd.Integer, nil},
		{xsd.DateTime, []string{"time"}},
		{xsd.Time, []string{"time"}},
		{xsd.Date, []string{"time"}},
		{xsd.Duration, nil},
		{xsd.QName, []string{"encoding/xml"}},
		{xsd.HexBinary, nil},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			got := getRequiredImports(tt.typ)
			if len(got) != len(tt.expected) {
				t.Errorf("getRequiredImports() = %v, want %v", got, tt.expected)
				return
			}
			for i, imp := range got {
				if imp != tt.expected[i] {
					t.Errorf("getRequiredImports()[%d] = %v, want %v", i, imp, tt.expected[i])
				}
			}
		})
	}
}

func TestToGoTypeName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected string
	}{
		{"", "string"},
		{"simpleType", "SimpleType"},
		{"complex_type", "ComplexType"},
		{"kebab-case", "KebabCase"},
		{"dot.separated", "DotSeparated"},
		{"namespace:localName", "NamespaceLocalname"},
		{"mixed_case-example.test", "MixedCaseExampleTest"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := toGoTypeName(tt.input); got != tt.expected {
				t.Errorf("toGoTypeName() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Test that IEEE 754 float types map to float64, and decimal maps to string.
func TestFloatingPointPrecision(t *testing.T) {
	t.Parallel()
	for _, typ := range []xsd.Type{xsd.Float, xsd.Double} {
		t.Run(string(typ), func(t *testing.T) {
			if got := mapXSDTypeToGo(typ); got != "float64" {
				t.Errorf("Expected IEEE 754 type %s to map to float64, got %s", typ, got)
			}
		})
	}
	t.Run(string(xsd.Decimal), func(t *testing.T) {
		if got := mapXSDTypeToGo(xsd.Decimal); got != "string" {
			t.Errorf("Expected xs:decimal to map to string (arbitrary precision), got %s", got)
		}
	})
}

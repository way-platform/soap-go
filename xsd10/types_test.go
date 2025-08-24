package xsd10_test

import (
	"testing"

	"github.com/way-platform/soap-go/xsd10"
)

func TestType_IsPrimitive(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.String, true},
		{xsd10.Boolean, true},
		{xsd10.Decimal, true},
		{xsd10.Float, true},
		{xsd10.Double, true},
		{xsd10.DateTime, true},
		{xsd10.QName, true},
		{xsd10.Integer, false},      // derived from decimal
		{xsd10.Token, false},        // derived from normalizedString
		{xsd10.NCName, false},       // derived from Name
		{xsd10.UnsignedLong, false}, // derived from nonNegativeInteger
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsPrimitive(); got != tt.expected {
				t.Errorf("IsPrimitive() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_IsDerived(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.String, false},      // primitive
		{xsd10.Boolean, false},     // primitive
		{xsd10.Integer, true},      // derived from decimal
		{xsd10.Token, true},        // derived from normalizedString
		{xsd10.NCName, true},       // derived from Name
		{xsd10.UnsignedLong, true}, // derived from nonNegativeInteger
		{xsd10.Long, true},         // derived from integer
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsDerived(); got != tt.expected {
				t.Errorf("IsDerived() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_IsNumeric(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.String, false},
		{xsd10.Boolean, false},
		{xsd10.Decimal, true},
		{xsd10.Float, true},
		{xsd10.Double, true},
		{xsd10.Integer, true},
		{xsd10.Long, true},
		{xsd10.Int, true},
		{xsd10.Short, true},
		{xsd10.Byte, true},
		{xsd10.UnsignedLong, true},
		{xsd10.UnsignedInt, true},
		{xsd10.UnsignedShort, true},
		{xsd10.UnsignedByte, true},
		{xsd10.DateTime, false},
		{xsd10.Date, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsNumeric(); got != tt.expected {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_IsString(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.String, true},
		{xsd10.NormalizedString, true},
		{xsd10.Token, true},
		{xsd10.NCName, true},
		{xsd10.ID, true},
		{xsd10.IDREF, true},
		{xsd10.Boolean, false},
		{xsd10.Integer, false},
		{xsd10.DateTime, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsString(); got != tt.expected {
				t.Errorf("IsString() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_IsTemporal(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.DateTime, true},
		{xsd10.Time, true},
		{xsd10.Date, true},
		{xsd10.Duration, true},
		{xsd10.GYear, true},
		{xsd10.GMonth, true},
		{xsd10.GDay, true},
		{xsd10.String, false},
		{xsd10.Integer, false},
		{xsd10.Boolean, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsTemporal(); got != tt.expected {
				t.Errorf("IsTemporal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_IsBinary(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected bool
	}{
		{xsd10.HexBinary, true},
		{xsd10.Base64Binary, true},
		{xsd10.String, false},
		{xsd10.Integer, false},
		{xsd10.Boolean, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsBinary(); got != tt.expected {
				t.Errorf("IsBinary() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_ToGoType(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected string
	}{
		// String types
		{xsd10.String, "string"},
		{xsd10.NormalizedString, "string"},
		{xsd10.Token, "string"},
		{xsd10.NCName, "string"},
		{xsd10.ID, "string"},
		{xsd10.AnyURI, "string"},

		// String slice types
		{xsd10.NMTOKENS, "[]string"},
		{xsd10.IDREFS, "[]string"},

		// Boolean
		{xsd10.Boolean, "bool"},

		// Floating point
		{xsd10.Float, "float32"},
		{xsd10.Double, "float64"},
		{xsd10.Decimal, "float64"},

		// Signed integers
		{xsd10.Byte, "int8"},
		{xsd10.Short, "int16"},
		{xsd10.Int, "int32"},
		{xsd10.Long, "int64"},
		{xsd10.Integer, "int64"},

		// Unsigned integers
		{xsd10.UnsignedByte, "uint8"},
		{xsd10.UnsignedShort, "uint16"},
		{xsd10.UnsignedInt, "uint32"},
		{xsd10.UnsignedLong, "uint64"},

		// Time types
		{xsd10.DateTime, "time.Time"},
		{xsd10.Time, "time.Time"},
		{xsd10.Date, "time.Time"},
		{xsd10.Duration, "time.Duration"},
		{xsd10.GYear, "string"},

		// Binary types
		{xsd10.HexBinary, "[]byte"},
		{xsd10.Base64Binary, "[]byte"},

		// Special types
		{xsd10.QName, "xml.Name"},
		{xsd10.NOTATION, "string"},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.ToGoType(); got != tt.expected {
				t.Errorf("ToGoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_RequiresImport(t *testing.T) {
	tests := []struct {
		typ      xsd10.Type
		expected []string
	}{
		{xsd10.String, nil},
		{xsd10.Boolean, nil},
		{xsd10.Integer, nil},
		{xsd10.DateTime, []string{"time"}},
		{xsd10.Time, []string{"time"}},
		{xsd10.Date, []string{"time"}},
		{xsd10.Duration, []string{"time"}},
		{xsd10.QName, []string{"encoding/xml"}},
		{xsd10.HexBinary, nil},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			got := tt.typ.RequiresImport()
			if len(got) != len(tt.expected) {
				t.Errorf("RequiresImport() = %v, want %v", got, tt.expected)
				return
			}
			for i, imp := range got {
				if imp != tt.expected[i] {
					t.Errorf("RequiresImport()[%d] = %v, want %v", i, imp, tt.expected[i])
				}
			}
		})
	}
}

func TestParseType(t *testing.T) {
	tests := []struct {
		input    string
		expected xsd10.Type
	}{
		{"string", xsd10.String},
		{"xs:string", xsd10.String},
		{"xsd:string", xsd10.String},
		{"s:string", xsd10.String},
		{"boolean", xsd10.Boolean},
		{"xs:boolean", xsd10.Boolean},
		{"integer", xsd10.Integer},
		{"xs:integer", xsd10.Integer},
		{"unsignedLong", xsd10.UnsignedLong},
		{"xs:unsignedLong", xsd10.UnsignedLong},
		{"decimal", xsd10.Decimal},
		{"xs:decimal", xsd10.Decimal},
		{"dateTime", xsd10.DateTime},
		{"xs:dateTime", xsd10.DateTime},
		{"QName", xsd10.QName},
		{"xs:QName", xsd10.QName},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := xsd10.ParseType(tt.input); got != tt.expected {
				t.Errorf("ParseType(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestType_String(t *testing.T) {
	typ := xsd10.String
	if got := typ.String(); got != "string" {
		t.Errorf("String() = %v, want %v", got, "string")
	}
}

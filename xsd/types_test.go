package xsd_test

import (
	"testing"

	"github.com/way-platform/soap-go/xsd"
)

func TestType_IsPrimitive(t *testing.T) {
	tests := []struct {
		typ      xsd.Type
		expected bool
	}{
		{xsd.String, true},
		{xsd.Boolean, true},
		{xsd.Decimal, true},
		{xsd.Float, true},
		{xsd.Double, true},
		{xsd.DateTime, true},
		{xsd.QName, true},
		{xsd.Integer, false},      // derived from decimal
		{xsd.Token, false},        // derived from normalizedString
		{xsd.NCName, false},       // derived from Name
		{xsd.UnsignedLong, false}, // derived from nonNegativeInteger
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
		typ      xsd.Type
		expected bool
	}{
		{xsd.String, false},      // primitive
		{xsd.Boolean, false},     // primitive
		{xsd.Integer, true},      // derived from decimal
		{xsd.Token, true},        // derived from normalizedString
		{xsd.NCName, true},       // derived from Name
		{xsd.UnsignedLong, true}, // derived from nonNegativeInteger
		{xsd.Long, true},         // derived from integer
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
		typ      xsd.Type
		expected bool
	}{
		{xsd.String, false},
		{xsd.Boolean, false},
		{xsd.Decimal, true},
		{xsd.Float, true},
		{xsd.Double, true},
		{xsd.Integer, true},
		{xsd.Long, true},
		{xsd.Int, true},
		{xsd.Short, true},
		{xsd.Byte, true},
		{xsd.UnsignedLong, true},
		{xsd.UnsignedInt, true},
		{xsd.UnsignedShort, true},
		{xsd.UnsignedByte, true},
		{xsd.DateTime, false},
		{xsd.Date, false},
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
		typ      xsd.Type
		expected bool
	}{
		{xsd.String, true},
		{xsd.NormalizedString, true},
		{xsd.Token, true},
		{xsd.NCName, true},
		{xsd.ID, true},
		{xsd.IDREF, true},
		{xsd.Boolean, false},
		{xsd.Integer, false},
		{xsd.DateTime, false},
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
		typ      xsd.Type
		expected bool
	}{
		{xsd.DateTime, true},
		{xsd.Time, true},
		{xsd.Date, true},
		{xsd.Duration, true},
		{xsd.GYear, true},
		{xsd.GMonth, true},
		{xsd.GDay, true},
		{xsd.String, false},
		{xsd.Integer, false},
		{xsd.Boolean, false},
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
		typ      xsd.Type
		expected bool
	}{
		{xsd.HexBinary, true},
		{xsd.Base64Binary, true},
		{xsd.String, false},
		{xsd.Integer, false},
		{xsd.Boolean, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			if got := tt.typ.IsBinary(); got != tt.expected {
				t.Errorf("IsBinary() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseType(t *testing.T) {
	tests := []struct {
		input    string
		expected xsd.Type
	}{
		{"string", xsd.String},
		{"xs:string", xsd.String},
		{"xsd:string", xsd.String},
		{"s:string", xsd.String},
		{"boolean", xsd.Boolean},
		{"xs:boolean", xsd.Boolean},
		{"integer", xsd.Integer},
		{"xs:integer", xsd.Integer},
		{"unsignedLong", xsd.UnsignedLong},
		{"xs:unsignedLong", xsd.UnsignedLong},
		{"decimal", xsd.Decimal},
		{"xs:decimal", xsd.Decimal},
		{"dateTime", xsd.DateTime},
		{"xs:dateTime", xsd.DateTime},
		{"QName", xsd.QName},
		{"xs:QName", xsd.QName},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := xsd.ParseType(tt.input); got != tt.expected {
				t.Errorf("ParseType(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestType_String(t *testing.T) {
	typ := xsd.String
	if got := typ.String(); got != "string" {
		t.Errorf("String() = %v, want %v", got, "string")
	}
}

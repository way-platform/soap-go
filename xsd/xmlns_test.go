package xsd_test

import (
	"strings"
	"testing"

	"github.com/way-platform/soap-go/xsd"
)

func TestSchemaNamespacePrefixMap(t *testing.T) {
	t.Parallel()

	schemaXML := `<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           xmlns:ns1="http://example.com/types/v1"
           xmlns:ns2="http://example.com/core/v1"
           xmlns:tns="http://example.com/service/v1"
           targetNamespace="http://example.com/service/v1"
           elementFormDefault="qualified">
  <xs:import namespace="http://example.com/types/v1"/>
  <xs:import namespace="http://example.com/core/v1"/>
  <xs:complexType name="Foo">
    <xs:sequence>
      <xs:element name="bar" type="ns1:Bar"/>
      <xs:element name="baz" type="ns2:Baz"/>
    </xs:sequence>
  </xs:complexType>
</xs:schema>`

	schema, err := xsd.Parse(strings.NewReader(schemaXML))
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	nsMap := schema.NamespacePrefixMap()

	tests := []struct {
		prefix string
		uri    string
	}{
		{"xs", "http://www.w3.org/2001/XMLSchema"},
		{"ns1", "http://example.com/types/v1"},
		{"ns2", "http://example.com/core/v1"},
		{"tns", "http://example.com/service/v1"},
	}

	for _, tt := range tests {
		got, ok := nsMap[tt.prefix]
		if !ok {
			t.Errorf("xmlns:%s not captured", tt.prefix)
			continue
		}
		if got != tt.uri {
			t.Errorf("xmlns:%s = %q, want %q", tt.prefix, got, tt.uri)
		}
	}

	if len(nsMap) != 4 {
		t.Errorf("expected 4 namespace prefixes, got %d", len(nsMap))
	}
}

func TestSchemaNamespacePrefixMap_Empty(t *testing.T) {
	t.Parallel()

	schemaXML := `<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
</xs:schema>`

	schema, err := xsd.Parse(strings.NewReader(schemaXML))
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	nsMap := schema.NamespacePrefixMap()
	// Should have at least the xs prefix
	if _, ok := nsMap["xs"]; !ok {
		t.Error("expected xs prefix to be captured")
	}
}

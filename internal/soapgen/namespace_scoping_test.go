package soapgen

import (
	"strings"
	"testing"

	"github.com/way-platform/soap-go/xsd"
)

func TestNsPrefixedName(t *testing.T) {
	t.Parallel()

	g := NewGenerator(nil, Config{
		NamespacePrefixes: map[string]string{
			"http://example.com/core/v1":    "Core",
			"http://example.com/billing/v1": "CB",
		},
	})

	tests := []struct {
		ns   string
		name string
		want string
	}{
		{"http://example.com/core/v1", "FlexAttr", "Core_FlexAttr"},
		{"http://example.com/billing/v1", "Invoice", "CB_Invoice"},
		{"http://example.com/unknown/v1", "Foo", "Foo"},   // unmapped namespace
		{"", "Bar", "Bar"},                                  // empty namespace
	}

	for _, tt := range tests {
		got := g.nsPrefixedName(tt.ns, tt.name)
		if got != tt.want {
			t.Errorf("nsPrefixedName(%q, %q) = %q, want %q", tt.ns, tt.name, got, tt.want)
		}
	}
}

func TestNsPrefixedName_Disabled(t *testing.T) {
	t.Parallel()

	g := NewGenerator(nil, Config{})

	got := g.nsPrefixedName("http://example.com/core/v1", "FlexAttr")
	if got != "FlexAttr" {
		t.Errorf("expected plain name when scoping disabled, got %q", got)
	}
}

func TestNamespaceScopingEnabled(t *testing.T) {
	t.Parallel()

	g1 := NewGenerator(nil, Config{})
	if g1.namespaceScopingEnabled() {
		t.Error("expected false when no prefixes configured")
	}

	g2 := NewGenerator(nil, Config{
		NamespacePrefixes: map[string]string{"http://example.com": "EX"},
	})
	if !g2.namespaceScopingEnabled() {
		t.Error("expected true when prefixes configured")
	}
}

func TestResolveNsScopedGoName(t *testing.T) {
	t.Parallel()

	schema := &xsd.Schema{
		TargetNamespace: "http://example.com/service/v1",
	}
	// ExtraAttrs would normally be populated by XML parsing; not needed for these tests
	// since we're testing the generator/context methods directly

	g := NewGenerator(nil, Config{
		NamespacePrefixes: map[string]string{
			"http://example.com/service/v1": "SVC",
			"http://example.com/core/v1":    "Core",
		},
	})

	ctx := newSchemaContext(schema, g)

	// Test scopedGoTypeName — for types declared in the current schema
	got := ctx.scopedGoTypeName("ProductOrder")
	if got != "SVC_ProductOrder" {
		t.Errorf("scopedGoTypeName = %q, want SVC_ProductOrder", got)
	}

	// Test currentNsPrefix
	prefix := ctx.currentNsPrefix()
	if prefix != "SVC" {
		t.Errorf("currentNsPrefix = %q, want SVC", prefix)
	}
}

func TestResolveNsScopedGoName_NoScoping(t *testing.T) {
	t.Parallel()

	schema := &xsd.Schema{
		TargetNamespace: "http://example.com/service/v1",
	}

	g := NewGenerator(nil, Config{})
	ctx := newSchemaContext(schema, g)

	got := ctx.scopedGoTypeName("ProductOrder")
	if got != "ProductOrder" {
		t.Errorf("expected plain name without scoping, got %q", got)
	}

	got = ctx.resolveNsScopedGoName("ProductOrder")
	if got != "ProductOrder" {
		t.Errorf("expected plain name without scoping, got %q", got)
	}
}

func TestMapXSDTypeToGoWithContext_NamespaceScoping(t *testing.T) {
	t.Parallel()

	// Create a schema with a complexType and xmlns declarations
	schemaXML := `<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           xmlns:tns="http://example.com/service/v1"
           xmlns:core="http://example.com/core/v1"
           targetNamespace="http://example.com/service/v1">
  <xs:import namespace="http://example.com/core/v1"/>
  <xs:complexType name="MyType">
    <xs:sequence>
      <xs:element name="value" type="xs:string"/>
    </xs:sequence>
  </xs:complexType>
</xs:schema>`

	schema, err := xsd.Parse(strings.NewReader(schemaXML))
	if err != nil {
		t.Fatalf("failed to parse schema: %v", err)
	}

	g := NewGenerator(nil, Config{
		NamespacePrefixes: map[string]string{
			"http://example.com/service/v1": "SVC",
			"http://example.com/core/v1":    "Core",
		},
	})

	ctx := newSchemaContext(schema, g)

	// Local type resolution (tns:MyType or just MyType)
	got := mapXSDTypeToGoWithContext("tns:MyType", ctx)
	if got != "SVC_MyType" {
		t.Errorf("tns:MyType resolved to %q, want SVC_MyType", got)
	}

	// Cross-namespace type resolution (core:SomeType — not in this schema)
	got = mapXSDTypeToGoWithContext("core:SomeType", ctx)
	if got != "Core_SomeType" {
		t.Errorf("core:SomeType resolved to %q, want Core_SomeType", got)
	}

	// XSD builtins should NOT be prefixed
	got = mapXSDTypeToGoWithContext("xs:string", ctx)
	if strings.Contains(got, "_") {
		t.Errorf("xs:string should not be namespace-prefixed, got %q", got)
	}
}

package soapgen

import (
	"sort"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/xsd"
)

// generateInlineComplexTypes generates structs for all inline complex types found in elements and named complex types
func generateInlineComplexTypes(g *codegen.File, ctx *SchemaContext, elements []*xsd.Element) {
	registry := newAnonymousTypeRegistry()

	// FIRST PASS: Register all inline types without generating struct definitions
	// This ensures all types are available for field generation

	// Register inline types within top-level elements
	for _, element := range elements {
		registerInlineTypesFromElement(element, "", ctx, registry)
	}

	// Register inline types within named complex types
	for _, complexType := range ctx.complexTypes {
		registerInlineTypesFromComplexType(complexType, complexType.Name, ctx, registry)
	}

	// SECOND PASS: Generate the actual struct definitions
	// Now that all types are registered, field generation can reference them
	hasTypes := false

	// Generate inline types within top-level elements
	for _, element := range elements {
		if generated := generateInlineTypesFromElement(g, element, "", ctx, registry); generated && !hasTypes {
			g.P("// Inline complex types")
			g.P()
			hasTypes = true
		}
	}

	// Generate inline types within named complex types
	// Sort complex type names for deterministic output
	var complexTypeNames []string
	for name := range ctx.complexTypes {
		complexTypeNames = append(complexTypeNames, name)
	}
	sort.Strings(complexTypeNames)

	for _, name := range complexTypeNames {
		complexType := ctx.complexTypes[name]
		if generated := generateInlineTypesFromComplexType(g, complexType, complexType.Name, ctx, registry); generated && !hasTypes {
			g.P("// Inline complex types")
			g.P()
			hasTypes = true
		}
	}

	// Generate RawXML wrapper types for multiple RawXML fields
	generateRawXMLWrapperTypes(g, ctx)

	if hasTypes {
		g.P()
	}
}

// generateInlineTypesFromElement recursively generates inline complex types from an element
func generateInlineTypesFromElement(g *codegen.File, element *xsd.Element, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) bool {
	generated := false

	if element.ComplexType != nil && element.ComplexType.Sequence != nil {
		for _, field := range element.ComplexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate inline complex type using Outer_Inner naming
				typeName := registry.generateTypeName(element.Name, field.Name)
				generateInlineComplexTypeStruct(g, typeName, field.ComplexType, ctx)
				generated = true

				// Register this type so we can reference it later
				ctx.anonymousTypes[typeName] = true

				// Recursively check for nested inline types
				if generateInlineTypesFromComplexType(g, field.ComplexType, typeName, ctx, registry) {
					generated = true
				}
			}
		}
	}

	return generated
}

// generateInlineTypesFromComplexType recursively generates inline complex types from a complex type
func generateInlineTypesFromComplexType(g *codegen.File, complexType *xsd.ComplexType, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) bool {
	generated := false

	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate inline complex type using Outer_Inner naming
				typeName := registry.generateTypeName(parentName, field.Name)
				generateInlineComplexTypeStruct(g, typeName, field.ComplexType, ctx)
				generated = true

				// Register this type so we can reference it later
				ctx.anonymousTypes[typeName] = true

				// Recursively check for nested inline types
				if generateInlineTypesFromComplexType(g, field.ComplexType, typeName, ctx, registry) {
					generated = true
				}
			}
		}
	}

	return generated
}

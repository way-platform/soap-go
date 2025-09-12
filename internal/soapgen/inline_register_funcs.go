package soapgen

import (
	"github.com/way-platform/soap-go/xsd"
)

// registerInlineTypesFromElement recursively registers inline complex types from an element without generating struct definitions
func registerInlineTypesFromElement(element *xsd.Element, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) {
	if element.ComplexType != nil && element.ComplexType.Sequence != nil {
		for _, field := range element.ComplexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate type name using the same logic as the generation pass
				// Don't use the registry to avoid conflicts - just compute the name directly
				typeName := toGoName(element.Name) + "_" + toGoName(field.Name)
				ctx.anonymousTypes[typeName] = true

				// Recursively register nested inline types
				registerInlineTypesFromComplexType(field.ComplexType, typeName, ctx, registry)
			}
		}
	}
}

// registerInlineTypesFromComplexType recursively registers inline complex types from a complex type without generating struct definitions
func registerInlineTypesFromComplexType(complexType *xsd.ComplexType, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) {
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate type name using the same logic as the generation pass
				// Don't use the registry to avoid conflicts - just compute the name directly
				typeName := toGoName(parentName) + "_" + toGoName(field.Name)
				ctx.anonymousTypes[typeName] = true

				// Recursively register nested inline types
				registerInlineTypesFromComplexType(field.ComplexType, typeName, ctx, registry)
			}
		}
	}
}

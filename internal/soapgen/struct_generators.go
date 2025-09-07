package soapgen

import (
	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/xsd"
)

// generateInlineComplexTypeStruct generates a struct for an inline complex type
func generateInlineComplexTypeStruct(g *codegen.File, typeName string, complexType *xsd.ComplexType, ctx *SchemaContext) {
	// Add comment
	g.P("// ", typeName, " represents an inline complex type")

	// Start struct declaration
	g.P("type ", typeName, " struct {")

	// Create field registry to track field name collisions
	fieldRegistry := newFieldRegistry()

	hasFields := false

	// Generate fields from the sequence
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, 1, typeName, fieldRegistry) {
				hasFields = true
			}
		}

		// Handle xs:any elements in the sequence
		for _, anyElement := range complexType.Sequence.Any {
			if generateAnyFieldWithFieldRegistry(g, &anyElement, ctx, 1, fieldRegistry) {
				hasFields = true
			}
		}
	}

	// Generate fields from attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeFieldWithFieldRegistry(g, &attr, ctx, fieldRegistry) {
			hasFields = true
		}
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined")
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructFromElement generates a Go struct from an XSD element
func generateStructFromElement(g *codegen.File, element *xsd.Element, ctx *SchemaContext, registry *TypeRegistry) {
	structName := toGoName(element.Name)
	generateStandardStructWithName(g, element, ctx, structName)
}

// generateStructFromElementWithWrapper generates a Go struct from an XSD element with wrapper naming
func generateStructFromElementWithWrapper(g *codegen.File, element *xsd.Element, ctx *SchemaContext, registry *TypeRegistry) {
	// Generate wrapper-style name
	structName := toGoName(element.Name) + "Wrapper"
	generateStandardStructWithName(g, element, ctx, structName)
}

// generateStandardStructWithName generates a standard struct with a custom struct name
func generateStandardStructWithName(g *codegen.File, element *xsd.Element, ctx *SchemaContext, structName string) {
	// Add comment
	g.P("// ", structName, " represents the ", element.Name, " element")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	// Add XMLName field for namespace handling
	generateXMLNameField(g, element, ctx)

	// Track if we've added any fields
	hasFields := true // XMLName counts as a field

	// Create field registry to track field name collisions
	fieldRegistry := newFieldRegistry()

	// Handle simple type elements (e.g., <element name="foo" type="xsd:string"/>)
	if element.Type != "" && element.ComplexType == nil {
		// This is a simple type element, generate a Value field
		goType := mapXSDTypeToGoWithContext(element.Type, ctx)
		goType = convertToQualifiedType(goType, g)

		// Handle complex type references - use the Go type name for complex types only
		if complexType := ctx.resolveComplexType(element.Type); complexType != nil {
			goType = toGoName(extractLocalName(element.Type))
		}

		g.P("\tValue ", goType, " `xml:\",chardata\"`")
		hasFields = true
	}

	if element.ComplexType != nil {
		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			// Count RawXML fields to determine XML tag behavior
			rawXMLCount := 0
			for _, field := range element.ComplexType.Sequence.Elements {
				if field.Type == "" && field.ComplexType != nil {
					rawXMLCount++
				}
			}
			for range element.ComplexType.Sequence.Any {
				rawXMLCount++
			}

			// Generate fields
			for _, field := range element.ComplexType.Sequence.Elements {
				if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, rawXMLCount, element.Name, fieldRegistry) {
					hasFields = true
				}
			}

			// Handle xs:any elements in the sequence
			for _, anyElement := range element.ComplexType.Sequence.Any {
				if generateAnyFieldWithFieldRegistry(g, &anyElement, ctx, rawXMLCount, fieldRegistry) {
					hasFields = true
				}
			}
		}

		// Handle attributes
		for _, attr := range element.ComplexType.Attributes {
			if generateAttributeFieldWithFieldRegistry(g, &attr, ctx, fieldRegistry) {
				hasFields = true
			}
		}

		// Handle complex content extensions
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, field := range ext.Sequence.Elements {
					if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, 1, element.Name, fieldRegistry) {
						hasFields = true
					}
				}
			}

			// Handle extension attributes
			for _, attr := range ext.Attributes {
				if generateAttributeFieldWithFieldRegistry(g, &attr, ctx, fieldRegistry) {
					hasFields = true
				}
			}
		}
	}

	// If no fields were generated beyond XMLName, add a placeholder comment
	if !hasFields {
		g.P("\t// No additional fields defined")
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructFromComplexType generates a Go struct from a named complex type
func generateStructFromComplexType(g *codegen.File, complexType *xsd.ComplexType, ctx *SchemaContext) {
	structName := toGoName(complexType.Name)

	// Add comment
	g.P("// ", structName, " represents the ", complexType.Name, " complex type")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	// Create field registry to track field name collisions
	fieldRegistry := newFieldRegistry()

	hasFields := false

	// Handle sequence elements
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, 1, complexType.Name, fieldRegistry) {
				hasFields = true
			}
		}

		// Handle xs:any elements in the sequence
		for _, anyElement := range complexType.Sequence.Any {
			if generateAnyFieldWithFieldRegistry(g, &anyElement, ctx, 1, fieldRegistry) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeFieldWithFieldRegistry(g, &attr, ctx, fieldRegistry) {
			hasFields = true
		}
	}

	// Handle complex content extensions
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil {
			for _, field := range ext.Sequence.Elements {
				if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, 1, complexType.Name, fieldRegistry) {
					hasFields = true
				}
			}
		}

		// Handle extension attributes
		for _, attr := range ext.Attributes {
			if generateAttributeFieldWithFieldRegistry(g, &attr, ctx, fieldRegistry) {
				hasFields = true
			}
		}
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined")
	}

	// Close struct
	g.P("}")
	g.P()
}

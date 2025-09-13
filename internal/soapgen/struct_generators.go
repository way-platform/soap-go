package soapgen

import (
	"strings"

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
		if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, typeName) {
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
		// Check if this references a complex type
		if complexType := ctx.resolveComplexType(element.Type); complexType != nil {
			// This element references a complex type - embed the complex type's fields directly
			if embedComplexTypeFields(g, complexType, ctx, fieldRegistry, element.Name) {
				hasFields = true
			}
		} else {
			// This is a simple type element, generate a Value field
			goType := mapXSDTypeToGoWithContext(element.Type, ctx)
			goType = convertToQualifiedType(goType, g)
			g.P("\tValue ", goType, " `xml:\",chardata\"`")
			hasFields = true
		}
	} else if element.SimpleType != nil && element.ComplexType == nil {
		// Handle inline simple type elements (e.g., elements with inline enumerations)
		// Check if this inline simple type has been generated as an enum type
		inlineEnumTypeName := ctx.getInlineEnumTypeName(element.Name, element.Name)
		if inlineEnumTypeName != "" {
			// Use the generated inline enum type
			g.P("\tValue ", inlineEnumTypeName, " `xml:\",chardata\"`")
		} else {
			// Fallback to mapping the base type
			baseType := "string" // Default fallback
			if element.SimpleType.Restriction != nil && element.SimpleType.Restriction.Base != "" {
				baseType = mapXSDTypeToGoWithContext(element.SimpleType.Restriction.Base, ctx)
				baseType = convertToQualifiedType(baseType, g)
			}
			g.P("\tValue ", baseType, " `xml:\",chardata\"`")
		}
		hasFields = true
	}

	if element.ComplexType != nil {
		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			// Count RawXML fields to determine XML tag behavior
			rawXMLCount := 0
			for _, field := range element.ComplexType.Sequence.Elements {
				if field.Type == "" && field.ComplexType != nil {
					// Only count if this inline complex type will actually become RawXML
					if shouldUseRawXMLForComplexType(field.ComplexType) {
						rawXMLCount++
					}
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
			if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, element.Name) {
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
				if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, element.Name) {
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
		if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, complexType.Name) {
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
			if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, complexType.Name) {
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

// generateRawXMLWrapperTypes generates wrapper types for RawXML fields that need their own ,innerxml
func generateRawXMLWrapperTypes(g *codegen.File, ctx *SchemaContext) {
	if ctx == nil || ctx.anonymousTypes == nil {
		return
	}

	generated := make(map[string]bool) // Track generated types to avoid duplicates
	hasTypes := false

	for typeName := range ctx.anonymousTypes {
		// Check if this is a RawXML wrapper type (has RAWXML_ prefix)
		if strings.HasPrefix(typeName, "RAWXML_") && !generated[typeName] {
			// Remove the RAWXML_ prefix to get the actual type name
			actualTypeName := strings.TrimPrefix(typeName, "RAWXML_")

			if !hasTypes {
				g.P("// RawXML wrapper types")
				g.P()
				hasTypes = true
			}

			// Generate a simple wrapper type with a single RawXML field
			g.P("// ", actualTypeName, " represents an inline complex type")
			g.P("type ", actualTypeName, " struct {")
			g.P("\tContent RawXML `xml:\",innerxml\"`")
			g.P("}")
			g.P()
			generated[typeName] = true
		}
	}
}

// embedComplexTypeFields embeds the fields from a complex type into the current struct
func embedComplexTypeFields(g *codegen.File, complexType *xsd.ComplexType, ctx *SchemaContext, fieldRegistry *FieldRegistry, parentElementName string) bool {
	hasFields := false

	// Handle sequence elements
	if complexType.Sequence != nil {
		// Count RawXML fields to determine XML tag behavior
		rawXMLCount := 0
		for _, field := range complexType.Sequence.Elements {
			if field.Type == "" && field.ComplexType != nil {
				// Only count if this inline complex type will actually become RawXML
				if shouldUseRawXMLForComplexType(field.ComplexType) {
					rawXMLCount++
				}
			}
		}
		for range complexType.Sequence.Any {
			rawXMLCount++
		}

		// Generate fields
		for _, field := range complexType.Sequence.Elements {
			// Use the complex type name as parent for anonymous type lookups
			complexTypeName := complexType.Name
			if complexTypeName == "" {
				complexTypeName = parentElementName
			}
			if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, rawXMLCount, complexTypeName, fieldRegistry) {
				hasFields = true
			}
		}

		// Handle xs:any elements in the sequence
		for _, anyElement := range complexType.Sequence.Any {
			if generateAnyFieldWithFieldRegistry(g, &anyElement, ctx, rawXMLCount, fieldRegistry) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, parentElementName) {
			hasFields = true
		}
	}

	// Handle complex content extensions
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil {
			// Count RawXML fields in extension
			rawXMLCount := 0
			for _, field := range ext.Sequence.Elements {
				if field.Type == "" && field.ComplexType != nil {
					// Only count if this inline complex type will actually become RawXML
					if shouldUseRawXMLForComplexType(field.ComplexType) {
						rawXMLCount++
					}
				}
			}
			for range ext.Sequence.Any {
				rawXMLCount++
			}

			for _, field := range ext.Sequence.Elements {
				// Use the complex type name as parent for anonymous type lookups
				complexTypeName := complexType.Name
				if complexTypeName == "" {
					complexTypeName = parentElementName
				}
				if generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g, &field, ctx, rawXMLCount, complexTypeName, fieldRegistry) {
					hasFields = true
				}
			}
		}

		// Handle extension attributes
		for _, attr := range ext.Attributes {
			if generateAttributeFieldWithParentName(g, &attr, ctx, fieldRegistry, parentElementName) {
				hasFields = true
			}
		}
	}

	return hasFields
}

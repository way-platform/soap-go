package soapgen

import (
	"sort"
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

	hasFields := false

	// Generate fields from the sequence
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypes(g, &field, ctx) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeField(g, &attr, ctx) {
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
func generateStructFromElement(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	structName := toGoName(element.Name)
	if structName == "" {
		return // Skip elements without valid names
	}

	// Add comment
	g.P("// ", structName, " represents the ", element.Name, " element")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	// Track if we've added any fields
	hasFields := false

	// Generate fields from the complex type or simple type
	if element.ComplexType != nil {
		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			for _, field := range element.ComplexType.Sequence.Elements {
				if generateStructFieldWithInlineTypes(g, &field, ctx) {
					hasFields = true
				}
			}
		}

		// Handle attributes
		for _, attr := range element.ComplexType.Attributes {
			if generateAttributeField(g, &attr, ctx) {
				hasFields = true
			}
		}

		// Handle complex content extensions
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, field := range ext.Sequence.Elements {
					if generateStructFieldWithInlineTypes(g, &field, ctx) {
						hasFields = true
					}
				}
			}
			// Handle extension attributes
			for _, attr := range ext.Attributes {
				if generateAttributeField(g, &attr, ctx) {
					hasFields = true
				}
			}
		}
	} else if element.Type != "" {
		// Handle elements with simple types - create a Value field
		goType := mapXSDTypeToGoWithContext(element.Type, ctx)
		g.P("\tValue ", goType, " `xml:\",chardata\"`")
		hasFields = true
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined for this element")
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructFieldWithInlineTypes generates a Go struct field with support for inline complex types
func generateStructFieldWithInlineTypes(g *codegen.File, element *xsd.Element, ctx *SchemaContext) bool {
	// Handle element references
	if element.Ref != "" {
		// Resolve the reference
		referencedElement := ctx.resolveElementRef(element.Ref)
		if referencedElement != nil {
			// Use the referenced element's name for the field
			fieldName := toGoName(referencedElement.Name)
			goType := toGoName(referencedElement.Name) // Reference the generated type
			xmlName := referencedElement.Name

			if fieldName != "" {
				// Handle optional elements
				if element.MinOccurs == "0" {
					goType = "*" + goType
				}
				// Handle multiple elements
				if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
					// For []byte (raw XML capture), don't create [][]byte - keep as []byte
					// For most other types, create slice of the type
					if goType != "[]byte" && goType != "*[]byte" {
						goType = "[]" + strings.TrimPrefix(goType, "*")
					}
				}

				// For []byte fields, use standard XML tags to capture element content
				g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, "\"`")
				return true
			}
		}
		return false
	}

	// Skip elements without names
	if element.Name == "" {
		return false
	}

	fieldName := toGoName(element.Name)
	if fieldName == "" {
		return false
	}

	// Determine the Go type
	var goType string
	if element.Type != "" {
		goType = mapXSDTypeToGoWithContext(element.Type, ctx)
	} else if element.ComplexType != nil {
		// For inline complex types without explicit type generation, use []byte to capture raw XML
		// This allows consumers to parse the XML content later when stronger typing is implemented
		goType = "[]byte"
	} else {
		goType = "string" // fallback
	}

	xmlName := element.Name

	// Handle optional elements
	if element.MinOccurs == "0" {
		if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") {
			goType = "*" + goType
		}
	}

	// Handle multiple elements
	if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
		// For []byte (raw XML capture), don't create [][]byte - keep as []byte to capture all XML content
		if goType != "[]byte" && goType != "*[]byte" {
			goType = "[]" + strings.TrimPrefix(goType, "*")
		}
	}

	// Generate the field with XML tag
	// For []byte fields, we want to capture the element content as-is
	// Note: Using standard xml:"elementName" tags for []byte fields will capture text content
	// For complex element content, consumers can manually unmarshal the []byte content
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, "\"`")
	return true
}

// generateAttributeField generates a Go struct field from an XSD attribute
func generateAttributeField(g *codegen.File, attr *xsd.Attribute, ctx *SchemaContext) bool {
	if attr.Name == "" {
		return false
	}

	fieldName := toGoName(attr.Name)
	if fieldName == "" {
		return false
	}

	goType := mapXSDTypeToGoWithContext(attr.Type, ctx)
	xmlName := attr.Name

	// Handle optional attributes
	if attr.Use != "required" {
		if !strings.HasPrefix(goType, "*") {
			goType = "*" + goType
		}
	}

	// Generate the field with XML attribute tag
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, ",attr\"`")
	return true
}

// generateSimpleTypeConstants generates Go constants for enumeration simple types
func generateSimpleTypeConstants(g *codegen.File, ctx *SchemaContext) {
	hasEnums := false

	// Sort simple type names for deterministic output
	var names []string
	for name := range ctx.simpleTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		simpleType := ctx.simpleTypes[name]
		if simpleType.Restriction != nil && len(simpleType.Restriction.Enumerations) > 0 {
			if !hasEnums {
				g.P("// Enumeration constants")
				g.P()
				hasEnums = true
			}

			generateEnumConstants(g, simpleType)
		}
	}

	if hasEnums {
		g.P()
	}
}

// generateEnumConstants generates Go constants for a single enumeration type
func generateEnumConstants(g *codegen.File, simpleType *xsd.SimpleType) {
	typeName := toGoName(simpleType.Name)

	g.P("// ", typeName, " enumeration values")
	g.P("const (")

	for _, enum := range simpleType.Restriction.Enumerations {
		constName := typeName + toGoName(enum.Value)
		g.P("\t", constName, " = \"", enum.Value, "\"")
	}

	g.P(")")
	g.P()
}

// generateComplexTypes generates Go structs for named complex types
func generateComplexTypes(g *codegen.File, ctx *SchemaContext) {
	hasTypes := false

	// Sort complex type names for deterministic output
	var names []string
	for name := range ctx.complexTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		complexType := ctx.complexTypes[name]
		if !hasTypes {
			g.P("// Complex types")
			g.P()
			hasTypes = true
		}

		generateStructFromComplexType(g, complexType, ctx)
	}

	if hasTypes {
		g.P()
	}
}

// generateStructFromComplexType generates a Go struct from an XSD complexType definition
func generateStructFromComplexType(g *codegen.File, complexType *xsd.ComplexType, ctx *SchemaContext) {
	structName := toGoName(complexType.Name)
	if structName == "" {
		return
	}

	// Add comment
	g.P("// ", structName, " represents the ", complexType.Name, " complex type")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	hasFields := false

	// Generate fields from the sequence
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypes(g, &field, ctx) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeField(g, &attr, ctx) {
			hasFields = true
		}
	}

	// Handle complex content extensions
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil {
			for _, field := range ext.Sequence.Elements {
				if generateStructFieldWithInlineTypes(g, &field, ctx) {
					hasFields = true
				}
			}
		}
		// Handle extension attributes
		for _, attr := range ext.Attributes {
			if generateAttributeField(g, &attr, ctx) {
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

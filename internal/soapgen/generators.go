package soapgen

import (
	"sort"
	"strings"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/xsd"
)

// buildXMLTag constructs an XML struct tag with appropriate omitempty behavior
func buildXMLTag(xmlName string, isOptional bool, isAttribute bool) string {
	parts := []string{xmlName}

	if isAttribute {
		parts = append(parts, "attr")
	}

	if isOptional {
		parts = append(parts, "omitempty")
	}

	return strings.Join(parts, ",")
}

// generateXMLNameField generates an XMLName field for flexible namespace handling
func generateXMLNameField(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	// Use local name only for flexible namespace handling
	// This allows the generated code to work with SOAP APIs that deviate from WSDL specs
	// by having different namespaces than expected while still validating element names
	elementName := element.Name
	g.P("\tXMLName xml.Name `xml:\"", elementName, "\"`")
}

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

// generateAnyField generates a RawXML field for xs:any elements (escape hatch for untyped content)
func generateAnyField(g *codegen.File, anyElement *xsd.Any, ctx *SchemaContext, singleRawXMLCount int) bool {
	// Generate a field name based on namespace or use a generic name
	fieldName := "Content"
	if anyElement.Namespace != "" && anyElement.Namespace != "##any" {
		// Use namespace to create a meaningful field name
		fieldName = "AnyContent"
	}

	// xs:any elements are the perfect use case for RawXML escape hatch
	goType := "RawXML"

	// Determine XML tag strategy for xs:any elements:
	// - Use ,innerxml when this is the only untyped content (true escape hatch)
	// - Use element name when there are multiple untyped fields
	var xmlTag string
	if singleRawXMLCount == 1 {
		// Single xs:any element gets ,innerxml to capture ALL inner content
		// This is the ideal escape hatch scenario - captures everything regardless of maxOccurs
		xmlTag = ",innerxml"
		// Don't use arrays for escape hatch - ,innerxml captures everything as one blob
		// goType remains "RawXML"
	} else {
		// Multiple xs:any elements use element names and can be arrays
		if anyElement.MaxOccurs == "unbounded" || (anyElement.MaxOccurs != "" && anyElement.MaxOccurs != "1") {
			goType = "[]" + goType
		}
		xmlTag = buildXMLTag("anyContent", anyElement.MinOccurs == "0", false)
	}

	// Handle optional occurrence (only for non-innerxml cases)
	if anyElement.MinOccurs == "0" && !strings.HasPrefix(goType, "[]") && xmlTag != ",innerxml" {
		goType = "*" + goType
	}

	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
	return true
}

// generateStructFromElement generates a Go struct from an XSD element
func generateStructFromElement(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	structName := toGoName(element.Name)
	if structName == "" {
		return // Skip elements without valid names
	}

	// Always use standard struct generation
	// The hybrid approach handles single vs multiple RawXML fields automatically

	// Standard struct generation for simple cases
	generateStandardStruct(g, element, ctx)
}

// generateStructFromElementWithWrapper generates a Go struct from an XSD element with wrapper naming
func generateStructFromElementWithWrapper(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	structName := toGoName(element.Name)
	if structName == "" {
		return // Skip elements without valid names
	}

	// Apply wrapper naming convention
	wrapperStructName := structName + "Wrapper"

	// Generate wrapper struct with modified name
	generateStandardStructWithName(g, element, ctx, wrapperStructName)
}

// generateStandardStruct generates a standard struct without custom parsing
func generateStandardStruct(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	structName := toGoName(element.Name)
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

	// Generate fields from the complex type or simple type
	if element.ComplexType != nil {
		// Count RawXML fields in this struct to determine proper XML tag strategy
		// Only count single RawXML fields (not arrays) for innerxml decision
		singleRawXMLCount := 0

		// Count sequence elements that will become single RawXML fields
		// Don't count fields that will get generated struct types
		if element.ComplexType.Sequence != nil {
			for _, field := range element.ComplexType.Sequence.Elements {
				if field.ComplexType != nil {
					// Check if this field will get a generated struct type
					typeName := toGoName(element.Name) + "_" + toGoName(field.Name)
					if !ctx.anonymousTypes[typeName] {
						// Only count single fields (not arrays) that will remain as RawXML
						if field.MaxOccurs == "" || field.MaxOccurs == "1" {
							singleRawXMLCount++
						}
					}
				}
			}

			// Count xs:any elements which always become RawXML fields
			for _, anyField := range element.ComplexType.Sequence.Any {
				// xs:any elements are escape hatch candidates - count all of them
				// regardless of maxOccurs, since we can use ,innerxml to capture everything
				singleRawXMLCount++
				// Note: We ignore maxOccurs here because for true escape hatch scenarios,
				// we want to use ,innerxml to capture all content as one blob
				_ = anyField
			}
		}

		// Count extension sequence elements that will become single RawXML fields
		// Don't count fields that will get generated struct types
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, field := range ext.Sequence.Elements {
					if field.ComplexType != nil {
						// Check if this field will get a generated struct type
						typeName := toGoName(element.Name) + "_" + toGoName(field.Name)
						if !ctx.anonymousTypes[typeName] {
							// Only count single fields (not arrays) that will remain as RawXML
							if field.MaxOccurs == "" || field.MaxOccurs == "1" {
								singleRawXMLCount++
							}
						}
					}
				}

				// Count xs:any elements in extensions which always become RawXML fields
				for _, anyField := range ext.Sequence.Any {
					// xs:any elements are escape hatch candidates - count all of them
					// regardless of maxOccurs, since we can use ,innerxml to capture everything
					singleRawXMLCount++
					// Note: We ignore maxOccurs here because for true escape hatch scenarios,
					// we want to use ,innerxml to capture all content as one blob
					_ = anyField
				}
			}
		}

		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			for _, field := range element.ComplexType.Sequence.Elements {
				if generateStructFieldWithInlineTypesAndContextAndParent(g, &field, ctx, singleRawXMLCount, element.Name) {
					hasFields = true
				}
			}

			// Handle xs:any elements - these become RawXML fields for escape hatch scenarios
			for _, anyField := range element.ComplexType.Sequence.Any {
				if generateAnyField(g, &anyField, ctx, singleRawXMLCount) {
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
					if generateStructFieldWithInlineTypesAndContextAndParent(g, &field, ctx, singleRawXMLCount, element.Name) {
						hasFields = true
					}
				}

				// Handle xs:any elements in extensions
				for _, anyField := range ext.Sequence.Any {
					if generateAnyField(g, &anyField, ctx, singleRawXMLCount) {
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

// generateStructFieldWithInlineTypesAndContext generates a Go struct field with support for inline complex types
func generateStructFieldWithInlineTypesAndContext(g *codegen.File, element *xsd.Element, ctx *SchemaContext, singleRawXMLCount int) bool {
	return generateStructFieldWithInlineTypesAndContextAndParent(g, element, ctx, singleRawXMLCount, "")
}

// generateStructFieldWithInlineTypesAndContextAndParent generates a Go struct field with support for inline complex types
func generateStructFieldWithInlineTypesAndContextAndParent(g *codegen.File, element *xsd.Element, ctx *SchemaContext, singleRawXMLCount int, parentElementName string) bool {
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
				xmlTag := buildXMLTag(xmlName, element.MinOccurs == "0", false)
				g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
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
		// Check if we have a generated struct type for this inline complex type
		// Use the same naming convention as the inline type generator
		if parentElementName != "" {
			typeName := toGoName(parentElementName) + "_" + toGoName(element.Name)
			if ctx.anonymousTypes[typeName] {
				// Use the generated struct type instead of RawXML
				goType = typeName
			} else {
				// Fallback to RawXML for truly untyped content
				goType = "RawXML"
			}
		} else {
			// No parent context available, fallback to RawXML
			goType = "RawXML"
		}
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
	// Use ,innerxml only when there's a single RawXML field in the struct (not for struct types)
	// Otherwise use element names for proper XML structure parsing
	var xmlTag string
	if goType == "RawXML" && singleRawXMLCount == 1 {
		// Use ,innerxml only for single RawXML fields to capture complete inner content
		// Don't add omitempty to innerxml tags as they capture all inner content
		xmlTag = ",innerxml"
	} else {
		// For struct types, multiple RawXML fields, or []RawXML, use element name for proper parsing
		// Add omitempty for optional fields
		xmlTag = buildXMLTag(xmlName, element.MinOccurs == "0", false)
	}
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
	return true
}

// generateStructFieldWithInlineTypes generates a Go struct field with support for inline complex types
func generateStructFieldWithInlineTypes(g *codegen.File, element *xsd.Element, ctx *SchemaContext) bool {
	// Default to rawXMLCount = 0 (use element names) for backwards compatibility
	return generateStructFieldWithInlineTypesAndContext(g, element, ctx, 0)
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
	xmlTag := buildXMLTag(xmlName, attr.Use != "required", true)
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
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
				g.P("// Enumeration types")
				g.P()
				hasEnums = true
			}

			generateEnumType(g, simpleType)
		}
	}

	if hasEnums {
		g.P()
	}
}

// generateEnumType generates Go enum type with constants and methods for a single enumeration type
func generateEnumType(g *codegen.File, simpleType *xsd.SimpleType) {
	typeName := toGoName(simpleType.Name)

	// Generate the enum type definition
	g.P("// ", typeName, " represents an enumeration type")
	g.P("type ", typeName, " string")
	g.P()

	// Generate the constants with typed values
	g.P("// ", typeName, " enumeration values")
	g.P("const (")

	var enumValues []string
	for _, enum := range simpleType.Restriction.Enumerations {
		constName := typeName + toGoName(enum.Value)
		g.P("\t", constName, " ", typeName, " = \"", enum.Value, "\"")
		enumValues = append(enumValues, constName)
	}

	g.P(")")
	g.P()

	// Generate String method
	g.P("// String returns the string representation of ", typeName)
	g.P("func (e ", typeName, ") String() string {")
	g.P("\treturn string(e)")
	g.P("}")
	g.P()

	// Generate IsValid method
	g.P("// IsValid returns true if the ", typeName, " value is valid")
	g.P("func (e ", typeName, ") IsValid() bool {")
	g.P("\tswitch e {")
	g.P("\tcase ", strings.Join(enumValues, ", "), ":")
	g.P("\t\treturn true")
	g.P("\tdefault:")
	g.P("\t\treturn false")
	g.P("\t}")
	g.P("}")
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

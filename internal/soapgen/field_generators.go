package soapgen

import (
	"strings"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/xsd"
)

// generateAnyFieldWithFieldRegistry generates a RawXML field for xs:any elements with collision detection
func generateAnyFieldWithFieldRegistry(g *codegen.File, anyElement *xsd.Any, ctx *SchemaContext, singleRawXMLCount int, fieldRegistry *FieldRegistry) bool {
	// Generate a field name based on namespace or use a generic name
	fieldName := "Content"
	if anyElement.Namespace != "" && anyElement.Namespace != "##any" {
		// Use namespace-specific field name, handling special namespace prefixes
		ns := anyElement.Namespace
		if strings.HasPrefix(ns, "##") {
			// Handle special namespace values like ##other, ##local, ##targetNamespace
			switch ns {
			case "##other":
				fieldName = "OtherContent"
			case "##local":
				fieldName = "LocalContent"
			case "##targetNamespace":
				fieldName = "TargetNamespaceContent"
			default:
				// Strip ## prefix and use the rest
				fieldName = toGoName(strings.TrimPrefix(ns, "##")) + "Content"
			}
		} else {
			fieldName = toGoName(ns) + "Content"
		}
	}

	// Use field registry for collision detection if available
	if fieldRegistry != nil {
		fieldName = fieldRegistry.generateUniqueFieldName(fieldName, false)
	}

	goType := "RawXML"

	// Determine XML tag behavior:
	// - Use innerxml for single RawXML fields to capture all untyped content
	// - Use element name when there are multiple untyped fields
	var xmlTag string
	if singleRawXMLCount == 1 {
		xmlTag = ",innerxml"
	} else {
		// For multiple RawXML fields, use a specific element name
		elementName := "content"
		if anyElement.Namespace != "" {
			elementName = strings.ToLower(anyElement.Namespace) + "Content"
		}
		xmlTag = buildXMLTag(elementName, false, false)
	}

	// Handle optional occurrence (only for non-innerxml cases)
	if anyElement.MinOccurs == "0" && !strings.HasPrefix(goType, "[]") && xmlTag != ",innerxml" {
		goType = "*" + goType
	}

	// Generate the field
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")

	return true
}

// generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry generates a Go struct field with support for inline complex types and field collision detection
func generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistry(g *codegen.File, element *xsd.Element, ctx *SchemaContext, singleRawXMLCount int, parentElementName string, fieldRegistry *FieldRegistry) bool {
	return generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistryInternal(g, element, ctx, singleRawXMLCount, parentElementName, fieldRegistry, false)
}

// generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistryInternal is the internal implementation
func generateStructFieldWithInlineTypesAndContextAndParentAndFieldRegistryInternal(g *codegen.File, element *xsd.Element, ctx *SchemaContext, singleRawXMLCount int, parentElementName string, fieldRegistry *FieldRegistry, isAttribute bool) bool {
	// Handle element references
	if element.Ref != "" {
		if referencedElement := ctx.resolveElementRef(element.Ref); referencedElement != nil {
			// Use the referenced element's name for the field
			var fieldName string
			if fieldRegistry != nil {
				fieldName = fieldRegistry.generateUniqueFieldName(referencedElement.Name, false)
			} else {
				fieldName = toGoName(referencedElement.Name)
			}

			// For element references, use the element's struct type name, not the underlying type
			goType := toGoName(referencedElement.Name)
			xmlName := referencedElement.Name

			// Handle optional elements
			if element.MinOccurs == "0" {
				goType = "*" + goType
			}

			// Handle multiple occurrences
			if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
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

			// Standard field generation for referenced elements
			xmlTag := buildXMLTag(xmlName, element.MinOccurs == "0", isAttribute)
			g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
			return true
		}
	}

	var fieldName string
	if fieldRegistry != nil {
		fieldName = fieldRegistry.generateUniqueFieldName(element.Name, false)
	} else {
		fieldName = toGoName(element.Name)
	}

	// Determine the Go type
	var goType string
	if element.Type != "" {
		rawType := mapXSDTypeToGoWithContext(element.Type, ctx)
		goType = convertToQualifiedType(rawType, g)
		// Handle complex type references - use the Go type name for complex types only
		if complexType := ctx.resolveComplexType(element.Type); complexType != nil {
			goType = toGoName(extractLocalName(element.Type))
		}
	} else if element.ComplexType != nil {
		// Inline complex type
		if ctx != nil && parentElementName != "" {
			// Check if this inline complex type should be RawXML (contains xs:any)
			if shouldUseRawXMLForComplexType(element.ComplexType) {
				goType = "RawXML"
			} else {
				// Use same naming convention as generateTypeName
				inlineTypeName := toGoName(parentElementName) + "_" + toGoName(element.Name)
				if ctx.anonymousTypes[inlineTypeName] {
					goType = inlineTypeName
				} else {
					goType = "RawXML"
				}
			}
		} else {
			// No parent context available, fallback to RawXML
			goType = "RawXML"
		}
	} else {
		goType = g.QualifiedGoIdent(codegen.StringIdent) // fallback
	}

	xmlName := element.Name

	// Handle optional elements
	if element.MinOccurs == "0" {
		if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") {
			goType = "*" + goType
		}
	}

	// Handle multiple occurrences
	if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
		// For []byte (raw XML capture), don't create [][]byte - keep as []byte to capture all XML content
		if goType != "[]byte" && goType != "*[]byte" {
			goType = "[]" + strings.TrimPrefix(goType, "*")
		}
	}

	// Generate XML tag
	// Otherwise use element names for proper XML structure parsing
	var xmlTag string
	if goType == "RawXML" && singleRawXMLCount == 1 {
		xmlTag = ",innerxml"
	} else {
		xmlTag = buildXMLTag(xmlName, element.MinOccurs == "0", isAttribute)
	}

	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
	return true
}

// generateAttributeFieldWithFieldRegistry generates a Go struct field from an XSD attribute with collision detection
func generateAttributeFieldWithFieldRegistry(g *codegen.File, attr *xsd.Attribute, ctx *SchemaContext, fieldRegistry *FieldRegistry) bool {
	if attr.Name == "" {
		return false
	}

	var fieldName string
	if fieldRegistry != nil {
		fieldName = fieldRegistry.generateUniqueFieldName(attr.Name, true)
	} else {
		fieldName = toGoName(attr.Name)
	}

	// Map XSD type to Go type
	goType := mapXSDTypeToGoWithContext(attr.Type, ctx)
	goType = convertToQualifiedType(goType, g)

	// Handle optional attributes
	if attr.Use != "required" {
		if !strings.HasPrefix(goType, "*") {
			goType = "*" + goType
		}
	}

	// Generate XML tag for attribute
	xmlTag := buildXMLTag(attr.Name, attr.Use != "required", true)
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlTag, "\"`")
	return true
}

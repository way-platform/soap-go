package soapgen

import (
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

// generateXMLNameField generates an XMLName field with appropriate namespace handling
func generateXMLNameField(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	elementName := element.Name

	// For operation elements (used in SOAP messages), include the target namespace
	// This ensures proper WSDL compliance for both requests and responses
	if ctx.generator != nil && ctx.generator.isOperationMessageElement(elementName) {
		if ctx.schema.TargetNamespace != "" {
			g.P("\tXMLName ", g.QualifiedGoIdent(codegen.XMLNameIdent), " `xml:\"", ctx.schema.TargetNamespace, " ", elementName, "\"`")
		} else {
			g.P("\tXMLName ", g.QualifiedGoIdent(codegen.XMLNameIdent), " `xml:\"", elementName, "\"`")
		}
	} else {
		// For non-operation elements, use flexible namespace handling
		g.P("\tXMLName ", g.QualifiedGoIdent(codegen.XMLNameIdent), " `xml:\"", elementName, "\"`")
	}
}

// convertToQualifiedType converts raw type strings to use QualifiedGoIdent for proper import management
func convertToQualifiedType(rawType string, g *codegen.File) string {
	switch rawType {
	case "time.Time":
		return g.QualifiedGoIdent(codegen.TimeIdent)
	case "time.Duration":
		return g.QualifiedGoIdent(codegen.GoIdent{GoImportPath: "time", GoName: "Duration"})
	case "string":
		return g.QualifiedGoIdent(codegen.StringIdent)
	case "bool":
		return g.QualifiedGoIdent(codegen.BoolIdent)
	case "int":
		return g.QualifiedGoIdent(codegen.IntIdent)
	case "[]byte":
		return "[]" + g.QualifiedGoIdent(codegen.ByteIdent)
	default:
		return rawType
	}
}

// shouldUseRawXMLForComplexType determines if a complex type should be represented as RawXML
// instead of generating a structured type. This is true for complex types that contain xs:any elements.
func shouldUseRawXMLForComplexType(complexType *xsd.ComplexType) bool {
	if complexType.Sequence != nil {
		// Check if the sequence contains xs:any elements
		if len(complexType.Sequence.Any) > 0 {
			return true
		}

		// Check if all elements are untyped (no type attribute and no inline complex type)
		hasTypedElements := false
		for _, elem := range complexType.Sequence.Elements {
			if elem.Type != "" || elem.ComplexType != nil {
				hasTypedElements = true
				break
			}
		}

		// If there are only untyped elements, use RawXML
		if !hasTypedElements && len(complexType.Sequence.Elements) > 0 {
			return true
		}
	}

	return false
}

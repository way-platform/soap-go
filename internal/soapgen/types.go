package soapgen

import (
	"strings"

	"github.com/way-platform/soap-go/xsd"
)

// toGoName converts an XML name to a Go identifier (PascalCase)
func toGoName(name string) string {
	if name == "" {
		return ""
	}

	// Trim spaces to avoid issues with malformed XML element names
	name = strings.TrimSpace(name)

	// Split on common separators and capitalize each part
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '_' || r == '-' || r == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	// Handle the case where name doesn't need splitting
	if len(parts) <= 1 {
		result.Reset()
		result.WriteString(strings.ToUpper(name[:1]))
		if len(name) > 1 {
			result.WriteString(name[1:])
		}
	}

	return result.String()
}

// mapXSDTypeToGoWithContext maps XSD types to Go types using schema context for better resolution
func mapXSDTypeToGoWithContext(xsdType string, ctx *SchemaContext) string {
	if xsdType == "" {
		return "[]byte" // fallback for empty types - capture raw XML
	}

	// First try to resolve as a simple type in the schema
	if simpleType := ctx.resolveSimpleType(xsdType); simpleType != nil {
		// For simple types with restrictions, check if it's an enumeration
		if simpleType.Restriction != nil && simpleType.Restriction.Base != "" {
			// If it has enumerations, keep it as a custom type (enum)
			if len(simpleType.Restriction.Enumerations) > 0 {
				return toGoName(extractLocalName(xsdType))
			}
			// Otherwise, resolve to the base type (for simple restrictions)
			return mapXSDTypeToGoWithContext(simpleType.Restriction.Base, ctx)
		}
		// If no restriction, treat as the simple type name
		return toGoName(extractLocalName(xsdType))
	}

	// Then try to resolve as a complex type in the schema
	if complexType := ctx.resolveComplexType(xsdType); complexType != nil {
		// For named complex types, generate a Go type name
		return toGoName(extractLocalName(xsdType))
	}

	// Check if this is a custom type (contains namespace prefix or ends with "Type")
	localName := extractLocalName(xsdType)
	if isCustomTypeName(localName) {
		// For custom types not defined in this schema, use a reasonable Go type
		return inferGoTypeFromCustomTypeName(localName)
	}

	// Try standard XSD type parsing
	parsedType := xsd.ParseType(xsdType)
	if !parsedType.IsCustomType() {
		return parsedType.ToGoType()
	}

	// For truly unknown/custom types, use []byte to capture raw XML
	return "[]byte"
}

// isCustomTypeName checks if a type name looks like a custom type
func isCustomTypeName(typeName string) bool {
	// Check if it ends with "Type" suffix (common pattern)
	if len(typeName) > 4 && strings.HasSuffix(typeName, "Type") {
		return true
	}

	// Check if it's not a standard XSD type
	parsedType := xsd.ParseType(typeName)
	return parsedType.IsCustomType()
}

// inferGoTypeFromCustomTypeName attempts to infer appropriate Go type from custom type name
func inferGoTypeFromCustomTypeName(typeName string) string {
	// Handle common patterns based on naming conventions
	name := strings.ToLower(typeName)

	// ID types are typically numeric
	if strings.Contains(name, "id") && strings.HasSuffix(name, "type") {
		return "int64"
	}

	// Timestamp types are typically strings (custom format)
	if strings.Contains(name, "timestamp") {
		return "string"
	}

	// Version types are typically numeric
	if strings.Contains(name, "version") {
		return "int64"
	}

	// Limit, offset, size types are typically numeric
	if strings.Contains(name, "limit") || strings.Contains(name, "offset") || strings.Contains(name, "size") {
		return "int64"
	}

	// Session types are typically strings
	if strings.Contains(name, "session") {
		return "string"
	}

	// For other custom types ending in "Type", assume string (safest default)
	if strings.HasSuffix(name, "type") {
		return "string"
	}

	// Default to generating a proper Go type name for complex types
	return xsd.ToGoTypeName(typeName)
}

// extractLocalName removes namespace prefix from a type name
func extractLocalName(typeName string) string {
	if colonIdx := strings.LastIndex(typeName, ":"); colonIdx != -1 {
		return typeName[colonIdx+1:]
	}
	return typeName
}

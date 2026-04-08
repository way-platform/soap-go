package soapgen

import (
	"strings"
	"unicode"

	"github.com/way-platform/soap-go/xsd"
)

// toGoName converts an XML name to a valid Go identifier (PascalCase).
//
// XSD names and enumeration values can contain characters that are not valid
// in Go identifiers (spaces, slashes, colons, plus signs, asterisks, etc.).
// This function splits on any non-letter, non-digit character, capitalizes
// the first letter of each part, and preserves the rest of each part's casing
// to maintain camelCase names like "GetWeather" or "DownloadRequest".
func toGoName(name string) string {
	if name == "" {
		return ""
	}

	name = strings.TrimSpace(name)

	// Split on any character that is not a letter or digit.
	// This handles spaces, slashes, colons, plus signs, asterisks, dots,
	// hyphens, underscores, and any other non-identifier characters.
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	if len(parts) == 0 {
		return "Value"
	}

	var result strings.Builder
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		// Capitalize first letter, preserve the rest of the casing.
		// This maintains camelCase names (e.g. "DownloadRequest" stays as-is)
		// while still capitalizing words from split values
		// (e.g. "EIR Sync Error" → "EIR" + "Sync" + "Error").
		result.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			result.WriteString(part[1:])
		}
	}

	// If the result starts with a digit, prefix to make it a valid Go identifier.
	s := result.String()
	if len(s) > 0 && unicode.IsDigit(rune(s[0])) {
		return "V" + s
	}

	return s
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
		return mapXSDTypeToGo(parsedType)
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
	return toGoTypeName(typeName)
}

// extractLocalName removes namespace prefix from a type name
func extractLocalName(typeName string) string {
	if colonIdx := strings.LastIndex(typeName, ":"); colonIdx != -1 {
		return typeName[colonIdx+1:]
	}
	return typeName
}

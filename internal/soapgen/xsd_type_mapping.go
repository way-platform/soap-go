package soapgen

import (
	"strings"

	"github.com/way-platform/soap-go/xsd"
)

// mapXSDTypeToGo maps XSD types to Go types.
func mapXSDTypeToGo(xsdType xsd.Type) string {
	switch xsdType {
	// String types
	case xsd.String, xsd.NormalizedString, xsd.Token, xsd.Language, xsd.NMTOKEN, xsd.Name, xsd.NCName,
		xsd.ID, xsd.IDREF, xsd.ENTITY, xsd.AnyURI:
		return "string"

	// String slice types
	case xsd.NMTOKENS, xsd.IDREFS, xsd.ENTITIES:
		return "[]string"

	// Boolean type
	case xsd.Boolean:
		return "bool"

	// Floating point types - all map to float64 for precision
	case xsd.Float, xsd.Double, xsd.Decimal:
		return "float64"

	// Signed integer types
	case xsd.Byte:
		return "int8"
	case xsd.Short:
		return "int16"
	case xsd.Int:
		return "int32"
	case xsd.Long, xsd.Integer, xsd.NonPositiveInteger, xsd.NegativeInteger:
		return "int64"

	// Unsigned integer types
	case xsd.UnsignedByte:
		return "uint8"
	case xsd.UnsignedShort:
		return "uint16"
	case xsd.UnsignedInt:
		return "uint32"
	case xsd.UnsignedLong, xsd.NonNegativeInteger, xsd.PositiveInteger:
		return "uint64"

	// Time types
	case xsd.DateTime, xsd.Time, xsd.Date:
		return "time.Time"
	case xsd.Duration:
		return "time.Duration"
	case xsd.GYearMonth, xsd.GYear, xsd.GMonthDay, xsd.GDay, xsd.GMonth:
		return "string" // These require custom parsing

	// Binary types
	case xsd.HexBinary, xsd.Base64Binary:
		return "[]byte"

	// Special types
	case xsd.QName:
		return "xml.Name"
	case xsd.NOTATION:
		return "string"

	default:
		// For custom/unknown types, convert to proper Go type name
		return toGoTypeName(string(xsdType))
	}
}

// getRequiredImports returns the import paths required for this XSD type.
func getRequiredImports(xsdType xsd.Type) []string {
	switch xsdType {
	case xsd.DateTime, xsd.Time, xsd.Date, xsd.Duration:
		return []string{"time"}
	case xsd.QName:
		return []string{"encoding/xml"}
	default:
		return nil
	}
}

// toGoTypeName converts a type name to a proper Go type name using PascalCase
func toGoTypeName(name string) string {
	if name == "" {
		return "string"
	}

	// Split on common separators and capitalize each part
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '_' || r == '-' || r == '.' || r == ':'
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

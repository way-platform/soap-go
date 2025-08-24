package xsd

// Type represents an XSD datatype as defined in XML Schema Part 2: Datatypes.
// This enum captures all primitive and built-in derived types from the XSD 1.0 specification.
type Type string

// Primitive XSD datatypes as defined in section 3.2 of the XSD 1.0 specification
const (
	// String represents the xs:string primitive datatype
	String Type = "string"

	// Boolean represents the xs:boolean primitive datatype
	Boolean Type = "boolean"

	// Decimal represents the xs:decimal primitive datatype
	Decimal Type = "decimal"

	// Float represents the xs:float primitive datatype
	Float Type = "float"

	// Double represents the xs:double primitive datatype
	Double Type = "double"

	// Duration represents the xs:duration primitive datatype
	Duration Type = "duration"

	// DateTime represents the xs:dateTime primitive datatype
	DateTime Type = "dateTime"

	// Time represents the xs:time primitive datatype
	Time Type = "time"

	// Date represents the xs:date primitive datatype
	Date Type = "date"

	// GYearMonth represents the xs:gYearMonth primitive datatype
	GYearMonth Type = "gYearMonth"

	// GYear represents the xs:gYear primitive datatype
	GYear Type = "gYear"

	// GMonthDay represents the xs:gMonthDay primitive datatype
	GMonthDay Type = "gMonthDay"

	// GDay represents the xs:gDay primitive datatype
	GDay Type = "gDay"

	// GMonth represents the xs:gMonth primitive datatype
	GMonth Type = "gMonth"

	// HexBinary represents the xs:hexBinary primitive datatype
	HexBinary Type = "hexBinary"

	// Base64Binary represents the xs:base64Binary primitive datatype
	Base64Binary Type = "base64Binary"

	// AnyURI represents the xs:anyURI primitive datatype
	AnyURI Type = "anyURI"

	// QName represents the xs:QName primitive datatype
	QName Type = "QName"

	// NOTATION represents the xs:NOTATION primitive datatype
	NOTATION Type = "NOTATION"
)

// Built-in derived XSD datatypes as defined in section 3.3 of the XSD 1.0 specification
const (
	// NormalizedString represents the xs:normalizedString derived datatype
	NormalizedString Type = "normalizedString"

	// Token represents the xs:token derived datatype
	Token Type = "token"

	// Language represents the xs:language derived datatype
	Language Type = "language"

	// NMTOKEN represents the xs:NMTOKEN derived datatype
	NMTOKEN Type = "NMTOKEN"

	// NMTOKENS represents the xs:NMTOKENS derived datatype
	NMTOKENS Type = "NMTOKENS"

	// Name represents the xs:Name derived datatype
	Name Type = "Name"

	// NCName represents the xs:NCName derived datatype
	NCName Type = "NCName"

	// ID represents the xs:ID derived datatype
	ID Type = "ID"

	// IDREF represents the xs:IDREF derived datatype
	IDREF Type = "IDREF"

	// IDREFS represents the xs:IDREFS derived datatype
	IDREFS Type = "IDREFS"

	// ENTITY represents the xs:ENTITY derived datatype
	ENTITY Type = "ENTITY"

	// ENTITIES represents the xs:ENTITIES derived datatype
	ENTITIES Type = "ENTITIES"

	// Integer represents the xs:integer derived datatype
	Integer Type = "integer"

	// NonPositiveInteger represents the xs:nonPositiveInteger derived datatype
	NonPositiveInteger Type = "nonPositiveInteger"

	// NegativeInteger represents the xs:negativeInteger derived datatype
	NegativeInteger Type = "negativeInteger"

	// Long represents the xs:long derived datatype
	Long Type = "long"

	// Int represents the xs:int derived datatype
	Int Type = "int"

	// Short represents the xs:short derived datatype
	Short Type = "short"

	// Byte represents the xs:byte derived datatype
	Byte Type = "byte"

	// NonNegativeInteger represents the xs:nonNegativeInteger derived datatype
	NonNegativeInteger Type = "nonNegativeInteger"

	// UnsignedLong represents the xs:unsignedLong derived datatype
	UnsignedLong Type = "unsignedLong"

	// UnsignedInt represents the xs:unsignedInt derived datatype
	UnsignedInt Type = "unsignedInt"

	// UnsignedShort represents the xs:unsignedShort derived datatype
	UnsignedShort Type = "unsignedShort"

	// UnsignedByte represents the xs:unsignedByte derived datatype
	UnsignedByte Type = "unsignedByte"

	// PositiveInteger represents the xs:positiveInteger derived datatype
	PositiveInteger Type = "positiveInteger"
)

// IsPrimitive returns true if the type is a primitive XSD datatype.
func (t Type) IsPrimitive() bool {
	switch t {
	case String, Boolean, Decimal, Float, Double, Duration,
		DateTime, Time, Date, GYearMonth, GYear, GMonthDay,
		GDay, GMonth, HexBinary, Base64Binary, AnyURI, QName, NOTATION:
		return true
	default:
		return false
	}
}

// IsDerived returns true if the type is a built-in derived XSD datatype.
func (t Type) IsDerived() bool {
	return !t.IsPrimitive() && t.IsBuiltIn()
}

// IsBuiltIn returns true if the type is a built-in XSD datatype (primitive or derived).
func (t Type) IsBuiltIn() bool {
	switch t {
	case String, Boolean, Decimal, Float, Double, Duration,
		DateTime, Time, Date, GYearMonth, GYear, GMonthDay,
		GDay, GMonth, HexBinary, Base64Binary, AnyURI, QName, NOTATION,
		NormalizedString, Token, Language, NMTOKEN, NMTOKENS,
		Name, NCName, ID, IDREF, IDREFS, ENTITY, ENTITIES,
		Integer, NonPositiveInteger, NegativeInteger, Long, Int, Short, Byte,
		NonNegativeInteger, UnsignedLong, UnsignedInt, UnsignedShort, UnsignedByte,
		PositiveInteger:
		return true
	default:
		return false
	}
}

// IsNumeric returns true if the type represents a numeric datatype.
func (t Type) IsNumeric() bool {
	switch t {
	case Decimal, Float, Double, Integer, NonPositiveInteger, NegativeInteger,
		Long, Int, Short, Byte, NonNegativeInteger, UnsignedLong, UnsignedInt,
		UnsignedShort, UnsignedByte, PositiveInteger:
		return true
	default:
		return false
	}
}

// IsString returns true if the type represents a string-based datatype.
func (t Type) IsString() bool {
	switch t {
	case String, NormalizedString, Token, Language, NMTOKEN, NMTOKENS,
		Name, NCName, ID, IDREF, IDREFS, ENTITY, ENTITIES:
		return true
	default:
		return false
	}
}

// IsTemporal returns true if the type represents a date/time datatype.
func (t Type) IsTemporal() bool {
	switch t {
	case Duration, DateTime, Time, Date, GYearMonth, GYear, GMonthDay, GDay, GMonth:
		return true
	default:
		return false
	}
}

// IsBinary returns true if the type represents binary data.
func (t Type) IsBinary() bool {
	switch t {
	case HexBinary, Base64Binary:
		return true
	default:
		return false
	}
}

// ToGoType maps the XSD type to the appropriate Go type.
func (t Type) ToGoType() string {
	switch t {
	// String types
	case String, NormalizedString, Token, Language, NMTOKEN, Name, NCName,
		ID, IDREF, ENTITY, AnyURI:
		return "string"

	// String slice types
	case NMTOKENS, IDREFS, ENTITIES:
		return "[]string"

	// Boolean type
	case Boolean:
		return "bool"

	// Floating point types
	case Float:
		return "float32"
	case Double, Decimal:
		return "float64"

	// Signed integer types
	case Byte:
		return "int8"
	case Short:
		return "int16"
	case Int:
		return "int32"
	case Long, Integer, NonPositiveInteger, NegativeInteger:
		return "int64"

	// Unsigned integer types
	case UnsignedByte:
		return "uint8"
	case UnsignedShort:
		return "uint16"
	case UnsignedInt:
		return "uint32"
	case UnsignedLong, NonNegativeInteger, PositiveInteger:
		return "uint64"

	// Time types
	case DateTime, Time, Date:
		return "time.Time"
	case Duration:
		return "time.Duration"
	case GYearMonth, GYear, GMonthDay, GDay, GMonth:
		return "string" // These require custom parsing

	// Binary types
	case HexBinary, Base64Binary:
		return "[]byte"

	// Special types
	case QName:
		return "xml.Name"
	case NOTATION:
		return "string"

	default:
		// For unknown types, default to string
		return "string"
	}
}

// RequiresImport returns the import paths required for this type.
func (t Type) RequiresImport() []string {
	switch t {
	case DateTime, Time, Date, Duration:
		return []string{"time"}
	case QName:
		return []string{"encoding/xml"}
	default:
		return nil
	}
}

// ParseType parses a string representation of an XSD type and returns the corresponding Type.
// It handles both local names (e.g., "string") and qualified names (e.g., "xs:string", "xsd:string").
func ParseType(typeStr string) Type {
	// Handle qualified names by extracting the local part
	if colonIdx := lastIndexOf(typeStr, ":"); colonIdx != -1 {
		typeStr = typeStr[colonIdx+1:]
	}

	return Type(typeStr)
}

// lastIndexOf returns the last index of substr in s, or -1 if not found.
func lastIndexOf(s, substr string) int {
	idx := -1
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			idx = i
		}
	}
	return idx
}

// String returns the string representation of the type.
func (t Type) String() string {
	return string(t)
}

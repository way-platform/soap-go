package kitchensink

import (
	"encoding/xml"
	"time"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// Enumeration types

// PriorityType represents an enumeration type
type PriorityType string

// PriorityType enumeration values
const (
	PriorityType1 PriorityType = "1"
	PriorityType2 PriorityType = "2"
	PriorityType3 PriorityType = "3"
)

// String returns the string representation of PriorityType
func (e PriorityType) String() string {
	return string(e)
}

// IsValid returns true if the PriorityType value is valid
func (e PriorityType) IsValid() bool {
	switch e {
	case PriorityType1, PriorityType2, PriorityType3:
		return true
	default:
		return false
	}
}

// StatusType represents an enumeration type
type StatusType string

// StatusType enumeration values
const (
	StatusTypeACTIVE   StatusType = "ACTIVE"
	StatusTypeINACTIVE StatusType = "INACTIVE"
	StatusTypePENDING  StatusType = "PENDING"
)

// String returns the string representation of StatusType
func (e StatusType) String() string {
	return string(e)
}

// IsValid returns true if the StatusType value is valid
func (e StatusType) IsValid() bool {
	switch e {
	case StatusTypeACTIVE, StatusTypeINACTIVE, StatusTypePENDING:
		return true
	default:
		return false
	}
}

// InlineTypesTest_Customer represents an inline complex type
type InlineTypesTest_Customer struct {
	Name    string                          `xml:"name"`
	Address InlinetypestestCustomer_Address `xml:"address"`
}

// InlinetypestestCustomer_Address represents an inline complex type
type InlinetypestestCustomer_Address struct {
	Street string `xml:"street"`
	City   string `xml:"city"`
}

// InlineTypesTest_Items represents an inline complex type
type InlineTypesTest_Items struct {
	Item []InlinetypestestItems_Item `xml:"item"`
}

// InlinetypestestItems_Item represents an inline complex type
type InlinetypestestItems_Item struct {
	Product  string `xml:"product"`
	Quantity int32  `xml:"quantity"`
}

// Inline complex types

// UntypedFieldsTest_ComplexData represents an inline complex type
type UntypedFieldsTest_ComplexData struct {
	InnerField string `xml:"innerField"`
}

// UntypedFieldsTest_MultipleComplexData represents an inline complex type
type UntypedFieldsTest_MultipleComplexData struct {
	InnerField int32 `xml:"innerField"`
}

// Complex types

// AddressType represents the AddressType complex type
type AddressType struct {
	Street   string `xml:"street"`
	City     string `xml:"city"`
	ZipCode  string `xml:"zipCode"`
	Country  string `xml:"country,attr"`
	Verified *bool  `xml:"verified,attr,omitempty"`
}

// UserInfoType represents the UserInfoType complex type
type UserInfoType struct {
	UserId int64      `xml:"userId"`
	Status StatusType `xml:"status"`
	Email  string     `xml:"email"`
}

// Tag represents the Tag element
type Tag struct {
	XMLName xml.Name `xml:"Tag"`
	Value   string   `xml:",chardata"`
}

// KitchenSinkRequestWrapper represents the KitchenSinkRequest element
type KitchenSinkRequestWrapper struct {
	XMLName                 xml.Name      `xml:"http://example.com/typetest KitchenSinkRequest"`
	StringField             string        `xml:"stringField"`
	BooleanField            bool          `xml:"booleanField"`
	IntField                int32         `xml:"intField"`
	LongField               int64         `xml:"longField"`
	ShortField              int16         `xml:"shortField"`
	ByteField               int8          `xml:"byteField"`
	FloatField              float32       `xml:"floatField"`
	DoubleField             float64       `xml:"doubleField"`
	DecimalField            float64       `xml:"decimalField"`
	DateTimeField           time.Time     `xml:"dateTimeField"`
	DateField               time.Time     `xml:"dateField"`
	TimeField               time.Time     `xml:"timeField"`
	DurationField           time.Duration `xml:"durationField"`
	UnsignedLongField       uint64        `xml:"unsignedLongField"`
	UnsignedIntField        uint32        `xml:"unsignedIntField"`
	UnsignedShortField      uint16        `xml:"unsignedShortField"`
	UnsignedByteField       uint8         `xml:"unsignedByteField"`
	IntegerField            int64         `xml:"integerField"`
	PositiveIntegerField    uint64        `xml:"positiveIntegerField"`
	NonNegativeIntegerField uint64        `xml:"nonNegativeIntegerField"`
	NegativeIntegerField    int64         `xml:"negativeIntegerField"`
	NonPositiveIntegerField int64         `xml:"nonPositiveIntegerField"`
	NormalizedStringField   string        `xml:"normalizedStringField"`
	TokenField              string        `xml:"tokenField"`
	LanguageField           string        `xml:"languageField"`
	NmtokenField            string        `xml:"nmtokenField"`
	NameField               string        `xml:"nameField"`
	NcnameField             string        `xml:"ncnameField"`
	IdField                 string        `xml:"idField"`
	IdrefField              string        `xml:"idrefField"`
	AnyUriField             string        `xml:"anyUriField"`
	QnameField              xml.Name      `xml:"qnameField"`
	HexBinaryField          []byte        `xml:"hexBinaryField"`
	Base64BinaryField       []byte        `xml:"base64BinaryField"`
	GYearField              string        `xml:"gYearField"`
	GMonthField             string        `xml:"gMonthField"`
	GDayField               string        `xml:"gDayField"`
	GYearMonthField         string        `xml:"gYearMonthField"`
	GMonthDayField          string        `xml:"gMonthDayField"`
	OptionalString          *string       `xml:"optionalString,omitempty"`
	OptionalInt             *int32        `xml:"optionalInt,omitempty"`
	Tags                    []string      `xml:"tags"`
	Numbers                 []int32       `xml:"numbers"`
	OptionalTags            []string      `xml:"optionalTags,omitempty"`
	Status                  StatusType    `xml:"status"`
	Priority                PriorityType  `xml:"priority"`
	OptionalStatus          *StatusType   `xml:"optionalStatus,omitempty"`
	Address                 AddressType   `xml:"address"`
	OptionalAddress         *AddressType  `xml:"optionalAddress,omitempty"`
	SimpleElement           string        `xml:"simpleElement"`
	Metadata                *AddressType  `xml:"metadata,omitempty"`
	Version                 string        `xml:"version,attr"`
	Debug                   *bool         `xml:"debug,attr,omitempty"`
	Timestamp               *time.Time    `xml:"timestamp,attr,omitempty"`
}

// KitchenSinkResponseWrapper represents the KitchenSinkResponse element
type KitchenSinkResponseWrapper struct {
	XMLName xml.Name `xml:"http://example.com/typetest KitchenSinkResponse"`
	Result  string   `xml:"result"`
}

// InlineTypesTest represents the InlineTypesTest element
type InlineTypesTest struct {
	XMLName  xml.Name                 `xml:"InlineTypesTest"`
	Customer InlineTypesTest_Customer `xml:"customer"`
	Items    InlineTypesTest_Items    `xml:"items"`
}

// PersonName represents the PersonName element
type PersonName struct {
	XMLName xml.Name `xml:"PersonName"`
	Value   string   `xml:",chardata"`
}

// PersonAge represents the PersonAge element
type PersonAge struct {
	XMLName xml.Name `xml:"PersonAge"`
	Value   int32    `xml:",chardata"`
}

// PersonInfo represents the PersonInfo element
type PersonInfo struct {
	XMLName    xml.Name   `xml:"PersonInfo"`
	PersonName PersonName `xml:"PersonName"`
	PersonAge  PersonAge  `xml:"PersonAge"`
	Tag        *Tag       `xml:"Tag,omitempty"`
}

// UntypedFieldsTest represents the UntypedFieldsTest element
type UntypedFieldsTest struct {
	XMLName             xml.Name                                `xml:"UntypedFieldsTest"`
	UnknownField        string                                  `xml:"unknownField"`
	UnknownArray        []string                                `xml:"unknownArray"`
	OptionalUnknown     *string                                 `xml:"optionalUnknown,omitempty"`
	ComplexData         UntypedFieldsTest_ComplexData           `xml:"complexData"`
	MultipleComplexData []UntypedFieldsTest_MultipleComplexData `xml:"multipleComplexData"`
}

// UserTest represents the UserTest element
type UserTest struct {
	XMLName xml.Name     `xml:"UserTest"`
	Value   UserInfoType `xml:",chardata"`
}

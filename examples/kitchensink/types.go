package kitchensink

import (
	"encoding/xml"
	"github.com/way-platform/soap-go"
	"time"
)

// Enumeration constants

// PriorityType enumeration values
const (
	PriorityType1 = "1"
	PriorityType2 = "2"
	PriorityType3 = "3"
)

// StatusType enumeration values
const (
	StatusTypeACTIVE   = "ACTIVE"
	StatusTypeINACTIVE = "INACTIVE"
	StatusTypePENDING  = "PENDING"
)

// Complex types

// AddressType represents the AddressType complex type
type AddressType struct {
	Street   string `xml:"street"`
	City     string `xml:"city"`
	ZipCode  string `xml:"zipCode"`
	Country  string `xml:"country,attr"`
	Verified *bool  `xml:"verified,attr"`
}

// UserInfoType represents the UserInfoType complex type
type UserInfoType struct {
	UserId int64  `xml:"userId"`
	Status string `xml:"status"`
	Email  string `xml:"email"`
}

// InlineTypesTest_Customer represents an inline complex type
type InlineTypesTest_Customer struct {
	Name    string      `xml:"name"`
	Address soap.RawXML `xml:"address"`
}

// InlinetypestestCustomer_Address represents an inline complex type
type InlinetypestestCustomer_Address struct {
	Street string `xml:"street"`
	City   string `xml:"city"`
}

// InlineTypesTest_Items represents an inline complex type
type InlineTypesTest_Items struct {
	Item []soap.RawXML `xml:"item"`
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

// Tag represents the Tag element
type Tag struct {
	Value string `xml:",chardata"`
}

// KitchenSinkRequest represents the KitchenSinkRequest element
type KitchenSinkRequest struct {
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
	OptionalString          *string       `xml:"optionalString"`
	OptionalInt             *int32        `xml:"optionalInt"`
	Tags                    []string      `xml:"tags"`
	Numbers                 []int32       `xml:"numbers"`
	OptionalTags            []string      `xml:"optionalTags"`
	Status                  string        `xml:"status"`
	Priority                int32         `xml:"priority"`
	OptionalStatus          *string       `xml:"optionalStatus"`
	Address                 AddressType   `xml:"address"`
	OptionalAddress         *AddressType  `xml:"optionalAddress"`
	SimpleElement           string        `xml:"simpleElement"`
	Metadata                *AddressType  `xml:"metadata"`
	Version                 string        `xml:"version,attr"`
	Debug                   *bool         `xml:"debug,attr"`
	Timestamp               *time.Time    `xml:"timestamp,attr"`
}

// KitchenSinkResponse represents the KitchenSinkResponse element
type KitchenSinkResponse struct {
	Result string `xml:"result"`
}

// InlineTypesTest represents the InlineTypesTest element
type InlineTypesTest struct {
	Customer soap.RawXML `xml:"customer"`
	Items    soap.RawXML `xml:"items"`
}

// PersonName represents the PersonName element
type PersonName struct {
	Value string `xml:",chardata"`
}

// PersonAge represents the PersonAge element
type PersonAge struct {
	Value int32 `xml:",chardata"`
}

// PersonInfo represents the PersonInfo element
type PersonInfo struct {
	PersonName PersonName `xml:"PersonName"`
	PersonAge  PersonAge  `xml:"PersonAge"`
	Tag        *Tag       `xml:"Tag"`
}

// UntypedFieldsTest represents the UntypedFieldsTest element
type UntypedFieldsTest struct {
	UnknownField        string        `xml:"unknownField"`
	UnknownArray        []string      `xml:"unknownArray"`
	OptionalUnknown     *string       `xml:"optionalUnknown"`
	ComplexData         soap.RawXML   `xml:",innerxml"`
	MultipleComplexData []soap.RawXML `xml:"multipleComplexData"`
}

// UserTest represents the UserTest element
type UserTest struct {
	Value UserInfoType `xml:",chardata"`
}

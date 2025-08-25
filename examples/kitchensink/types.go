package kitchensink

import (
	"encoding/xml"
	"time"
)

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
}

// KitchenSinkResponse represents the KitchenSinkResponse element
type KitchenSinkResponse struct {
	Result string `xml:"result"`
}

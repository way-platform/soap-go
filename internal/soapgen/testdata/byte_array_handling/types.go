package byte_array_handling

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// DataContainer_InlineData represents an inline complex type
type DataContainer_InlineData struct {
	InnerField string `xml:"innerField"`
}

// DataContainer_MultipleInlineData represents an inline complex type
type DataContainer_MultipleInlineData struct {
	InnerField int32 `xml:"innerField"`
}

// Inline complex types

// DataContainer represents the DataContainer element
type DataContainer struct {
	XMLName              xml.Name                           `xml:"http://example.com/test DataContainer"`
	SingleData           string                             `xml:"singleData"`
	MultipleData         []string                           `xml:"multipleData"`
	OptionalData         *string                            `xml:"optionalData,omitempty"`
	OptionalMultipleData []string                           `xml:"optionalMultipleData,omitempty"`
	KnownString          string                             `xml:"knownString"`
	KnownStringArray     []string                           `xml:"knownStringArray"`
	InlineData           DataContainer_InlineData           `xml:"inlineData"`
	MultipleInlineData   []DataContainer_MultipleInlineData `xml:"multipleInlineData"`
}

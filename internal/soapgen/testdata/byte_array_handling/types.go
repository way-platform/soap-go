package byte_array_handling

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
	SingleData           string   `xml:"singleData"`
	MultipleData         []string `xml:"multipleData"`
	OptionalData         *string  `xml:"optionalData"`
	OptionalMultipleData []string `xml:"optionalMultipleData"`
	KnownString          string   `xml:"knownString"`
	KnownStringArray     []string `xml:"knownStringArray"`
	InlineData           []byte   `xml:"inlineData"`
	MultipleInlineData   []byte   `xml:"multipleInlineData"`
}

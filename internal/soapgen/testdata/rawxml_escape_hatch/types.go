package rawxml_escape_hatch

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// DocumentWithRawContent_DynamicContent represents an inline complex type
type DocumentWithRawContent_DynamicContent struct {
	// No fields defined
}

// Inline complex types

// DocumentWithMultipleRawContent_Header represents an inline complex type
type DocumentWithMultipleRawContent_Header struct {
	// No fields defined
}

// DocumentWithMultipleRawContent_Body represents an inline complex type
type DocumentWithMultipleRawContent_Body struct {
	// No fields defined
}

// DocumentWithRawContent represents the DocumentWithRawContent element
type DocumentWithRawContent struct {
	XMLName        xml.Name                              `xml:"http://example.com/test DocumentWithRawContent"`
	Title          string                                `xml:"title"`
	Version        string                                `xml:"version"`
	DynamicContent DocumentWithRawContent_DynamicContent `xml:"dynamicContent"`
}

// DocumentWithMultipleRawContent represents the DocumentWithMultipleRawContent element
type DocumentWithMultipleRawContent struct {
	XMLName xml.Name                              `xml:"http://example.com/test DocumentWithMultipleRawContent"`
	Title   string                                `xml:"title"`
	Header  DocumentWithMultipleRawContent_Header `xml:"header"`
	Body    DocumentWithMultipleRawContent_Body   `xml:"body"`
}

// PureUntypedDocument represents the PureUntypedDocument element
type PureUntypedDocument struct {
	XMLName xml.Name `xml:"http://example.com/test PureUntypedDocument"`
	Content RawXML   `xml:",innerxml"`
}

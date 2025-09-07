package rawxml_escape_hatch

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// DocumentWithRawContent_DynamicContent represents an inline complex type
type DocumentWithRawContent_DynamicContent struct {
	Content RawXML `xml:",innerxml"`
}

// Inline complex types

// DocumentWithMultipleRawContent_Header represents an inline complex type
type DocumentWithMultipleRawContent_Header struct {
	Content RawXML `xml:",innerxml"`
}

// DocumentWithMultipleRawContent_Body represents an inline complex type
type DocumentWithMultipleRawContent_Body struct {
	Content RawXML `xml:",innerxml"`
}

// DocumentWithRawContent represents the DocumentWithRawContent element
type DocumentWithRawContent struct {
	XMLName        xml.Name `xml:"DocumentWithRawContent"`
	Title          string   `xml:"title"`
	Version        string   `xml:"version"`
	DynamicContent RawXML   `xml:",innerxml"`
}

// DocumentWithMultipleRawContent represents the DocumentWithMultipleRawContent element
type DocumentWithMultipleRawContent struct {
	XMLName xml.Name                              `xml:"DocumentWithMultipleRawContent"`
	Title   string                                `xml:"title"`
	Header  DocumentWithMultipleRawContent_Header `xml:"header"`
	Body    DocumentWithMultipleRawContent_Body   `xml:"body"`
}

// PureUntypedDocument represents the PureUntypedDocument element
type PureUntypedDocument struct {
	XMLName xml.Name `xml:"PureUntypedDocument"`
	Content RawXML   `xml:",innerxml"`
}

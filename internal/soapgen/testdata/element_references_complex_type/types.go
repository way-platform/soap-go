package element_references_complex_type

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// ResponseType_Items represents an inline complex type
type ResponseType_Items struct {
	ID    int64  `xml:"ID"`
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

// Inline complex types

// Complex types

// ResponseType represents the ResponseType complex type
type ResponseType struct {
	Items []ResponseType_Items `xml:"Items,omitempty"`
	Total *int32               `xml:"total,attr,omitempty"`
}

// Response represents the Response element
type Response struct {
	XMLName xml.Name             `xml:"Response"`
	Items   []ResponseType_Items `xml:"Items,omitempty"`
	Total   *int32               `xml:"total,attr,omitempty"`
}

// GetItemsWrapper represents the getItems element
type GetItemsWrapper struct {
	XMLName xml.Name `xml:"http://example.com/test getItems"`
	Request string   `xml:"request"`
}

// GetItemsResponseWrapper represents the getItemsResponse element
type GetItemsResponseWrapper struct {
	XMLName  xml.Name `xml:"http://example.com/test getItemsResponse"`
	Response Response `xml:"Response"`
}

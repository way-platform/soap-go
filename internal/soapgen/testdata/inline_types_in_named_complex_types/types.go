package inline_types_in_named_complex_types

import (
	"encoding/xml"
	"time"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// ResponseType_Data represents an inline complex type
type ResponseType_Data struct {
	Id       string                    `xml:"id"`
	Value    string                    `xml:"value"`
	Metadata ResponsetypeData_Metadata `xml:"metadata"`
}

// ResponsetypeData_Metadata represents an inline complex type
type ResponsetypeData_Metadata struct {
	Timestamp time.Time `xml:"timestamp"`
	Source    string    `xml:"source"`
}

// ResponseType_Items represents an inline complex type
type ResponseType_Items struct {
	ItemId   string `xml:"itemId"`
	ItemName string `xml:"itemName"`
}

// Inline complex types

// Complex types

// ResponseType represents the ResponseType complex type
type ResponseType struct {
	Status string               `xml:"status"`
	Data   ResponseType_Data    `xml:"data"`
	Items  []ResponseType_Items `xml:"items"`
}

// Response represents the Response element
type Response struct {
	XMLName xml.Name             `xml:"Response"`
	Status  string               `xml:"status"`
	Data    ResponseType_Data    `xml:"data"`
	Items   []ResponseType_Items `xml:"items"`
}

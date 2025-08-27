package inline_complex_types

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// Order_Customer represents an inline complex type
type Order_Customer struct {
	Name    string `xml:"name"`
	Address RawXML `xml:"address"`
}

// OrderCustomer_Address represents an inline complex type
type OrderCustomer_Address struct {
	Street string `xml:"street"`
	City   string `xml:"city"`
}

// Order_Items represents an inline complex type
type Order_Items struct {
	Item []RawXML `xml:"item"`
}

// OrderItems_Item represents an inline complex type
type OrderItems_Item struct {
	Product  string `xml:"product"`
	Quantity int32  `xml:"quantity"`
}

// Inline complex types

// Order represents the Order element
type Order struct {
	XMLName  xml.Name `xml:"http://example.com/test Order"`
	Customer RawXML   `xml:"customer"`
	Items    RawXML   `xml:"items"`
}

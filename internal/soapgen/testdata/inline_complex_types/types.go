package inline_complex_types

import (
	"github.com/way-platform/soap-go"
)

// Order_Customer represents an inline complex type
type Order_Customer struct {
	Name    string      `xml:"name"`
	Address soap.RawXML `xml:"address"`
}

// OrderCustomer_Address represents an inline complex type
type OrderCustomer_Address struct {
	Street string `xml:"street"`
	City   string `xml:"city"`
}

// Order_Items represents an inline complex type
type Order_Items struct {
	Item []soap.RawXML `xml:"item"`
}

// OrderItems_Item represents an inline complex type
type OrderItems_Item struct {
	Product  string `xml:"product"`
	Quantity int32  `xml:"quantity"`
}

// Inline complex types

// Order represents the Order element
type Order struct {
	Customer soap.RawXML `xml:"customer"`
	Items    soap.RawXML `xml:"items"`
}

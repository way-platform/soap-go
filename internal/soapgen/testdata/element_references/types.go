package element_references

import (
	"encoding/xml"
)

// PersonName represents the PersonName element
type PersonName struct {
	XMLName xml.Name `xml:"PersonName"`
	Value   string   `xml:",chardata"`
}

// PersonAge represents the PersonAge element
type PersonAge struct {
	XMLName xml.Name `xml:"PersonAge"`
	Value   int32    `xml:",chardata"`
}

// Address represents the Address element
type Address struct {
	XMLName xml.Name `xml:"Address"`
	Street  string   `xml:"street"`
	City    string   `xml:"city"`
	ZipCode string   `xml:"zipCode"`
}

// Person represents the Person element
type Person struct {
	XMLName        xml.Name    `xml:"Person"`
	PersonName     PersonName  `xml:"PersonName"`
	PersonAge      PersonAge   `xml:"PersonAge"`
	Address        Address     `xml:"Address"`
	PersonNameElem *PersonName `xml:"PersonName,omitempty"`
	AddressElem    []Address   `xml:"Address"`
}

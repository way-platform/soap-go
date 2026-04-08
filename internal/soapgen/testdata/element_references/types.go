package element_references

import (
	"encoding/xml"
)

// PersonNameWrapper represents the PersonName element
type PersonNameWrapper struct {
	XMLName xml.Name `xml:"PersonName"`
	Value   string   `xml:",chardata"`
}

// PersonAgeWrapper represents the PersonAge element
type PersonAgeWrapper struct {
	XMLName xml.Name `xml:"PersonAge"`
	Value   int32    `xml:",chardata"`
}

// AddressWrapper represents the Address element
type AddressWrapper struct {
	XMLName xml.Name `xml:"Address"`
	Street  string   `xml:"street"`
	City    string   `xml:"city"`
	ZipCode string   `xml:"zipCode"`
}

// PersonWrapper represents the Person element
type PersonWrapper struct {
	XMLName        xml.Name           `xml:"Person"`
	PersonName     PersonNameWrapper  `xml:"PersonName"`
	PersonAge      PersonAgeWrapper   `xml:"PersonAge"`
	Address        AddressWrapper     `xml:"Address"`
	PersonNameElem *PersonNameWrapper `xml:"PersonName,omitempty"`
	AddressElem    []AddressWrapper   `xml:"Address"`
}

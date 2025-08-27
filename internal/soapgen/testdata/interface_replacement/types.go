package interface_replacement

import (
	"encoding/xml"
)

// MixedContent represents the MixedContent element
type MixedContent struct {
	XMLName         xml.Name `xml:"http://example.com/test MixedContent"`
	KnownField      string   `xml:"knownField"`
	UnknownField    string   `xml:"unknownField"`
	UnknownArray    []string `xml:"unknownArray"`
	OptionalUnknown *string  `xml:"optionalUnknown,omitempty"`
}

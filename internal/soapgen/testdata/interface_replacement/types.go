package interface_replacement

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// MixedContentWrapper represents the MixedContent element
type MixedContentWrapper struct {
	XMLName         xml.Name `xml:"MixedContent"`
	KnownField      string   `xml:"knownField"`
	UnknownField    string   `xml:"unknownField"`
	UnknownArray    []string `xml:"unknownArray"`
	OptionalUnknown *string  `xml:"optionalUnknown,omitempty"`
}

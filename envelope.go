package soap

import (
	"encoding/xml"
)

// Namespace is the standard SOAP 1.1 envelope namespace
const Namespace = "http://schemas.xmlsoap.org/soap/envelope/"

// Envelope represents a SOAP 1.1 envelope according to the specification.
// It provides a complete implementation supporting headers, body, faults, and extensibility.
// Generates soap: prefixed XML for maximum compatibility with real-world SOAP services.
type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`

	// SOAP namespace declaration - this gets marshaled as xmlns:soap
	XMLNS string `xml:"xmlns:soap,attr"`

	// Optional encoding style as per SOAP 1.1 spec section 4.1.1
	EncodingStyle string `xml:"soap:encodingStyle,attr,omitempty"`

	// Optional header as per SOAP 1.1 spec section 4.2
	Header *Header `xml:"soap:Header,omitempty"`

	// Mandatory body as per SOAP 1.1 spec section 4.3
	Body Body `xml:"soap:Body"`

	// Additional attributes for extensibility as per SOAP 1.1 spec section 4.1
	Attrs []xml.Attr `xml:",any,attr"`
}

// Header represents a SOAP header containing header entries.
// Each header entry can have mustUnderstand and actor attributes as per SOAP 1.1 spec section 4.2.
type Header struct {
	// Header entries - flexible content allowing any XML
	Entries []HeaderEntry `xml:",any"`

	// Additional attributes for extensibility
	Attrs []xml.Attr `xml:",any,attr"`
}

// HeaderEntry represents a single header entry with SOAP-specific attributes.
// Implements the mustUnderstand and actor semantics from SOAP 1.1 spec sections 4.2.2 and 4.2.3.
type HeaderEntry struct {
	XMLName xml.Name

	// MustUnderstand attribute as per SOAP 1.1 spec section 4.2.3
	// Values: true (1) means mandatory, false (0) or nil means optional
	MustUnderstand *bool `xml:"mustUnderstand,attr,omitempty"`

	// Actor attribute as per SOAP 1.1 spec section 4.2.2
	// Specifies the intended recipient of this header entry
	Actor string `xml:"actor,attr,omitempty"`

	// Content as raw XML for maximum flexibility
	Content []byte `xml:",innerxml"`

	// Additional attributes for extensibility
	Attrs []xml.Attr `xml:",any,attr"`
}

// Body represents a SOAP body containing the main message payload.
// As per SOAP 1.1 spec section 4.3, it contains body entries.
type Body struct {
	// Content as raw XML for maximum flexibility
	// This allows both simple payloads and complex nested structures
	Content []byte `xml:",innerxml"`

	// Additional attributes for extensibility
	Attrs []xml.Attr `xml:",any,attr"`
}

// Fault represents a SOAP fault element as per SOAP 1.1 spec section 4.4.
// Used for error reporting within SOAP messages.
type Fault struct {
	XMLName xml.Name `xml:"soap:Fault"`

	// FaultCode is mandatory and provides algorithmic fault identification
	FaultCode string `xml:"faultcode"`

	// FaultString is mandatory and provides human-readable fault description
	FaultString string `xml:"faultstring"`

	// FaultActor is optional and identifies the fault source
	FaultActor string `xml:"faultactor,omitempty"`

	// Detail is optional and contains application-specific error information
	Detail *Detail `xml:"detail,omitempty"`
}

// Detail represents fault detail information.
// Contains application-specific error data as per SOAP 1.1 spec section 4.4.
type Detail struct {
	// Content as raw XML to accommodate any application-specific error data
	Content []byte `xml:",innerxml"`

	// Additional attributes for extensibility
	Attrs []xml.Attr `xml:",any,attr"`
}

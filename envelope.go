package soap

import (
	"encoding/xml"
)

// Namespace is the standard SOAP 1.1 envelope namespace
const Namespace = "http://schemas.xmlsoap.org/soap/envelope/"

// Envelope represents a SOAP envelope with flexible namespace support.
// It can handle any namespace prefix and URI, making it compatible with various SOAP implementations.
// The XMLName field determines the actual element name and namespace used in marshaling/unmarshaling.
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`

	// Optional encoding style as per SOAP 1.1 spec section 4.1.1
	EncodingStyle string `xml:"encodingStyle,attr,omitempty"`

	// Optional header as per SOAP 1.1 spec section 4.2
	Header *Header `xml:"Header,omitempty"`

	// Mandatory body as per SOAP 1.1 spec section 4.3
	Body Body `xml:"Body"`

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
	XMLName xml.Name `xml:"Fault"`

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

// NewEnvelope creates a new SOAP envelope with the standard SOAP 1.1 namespace.
// This is a convenience function for the most common use case.
func NewEnvelope() *Envelope {
	return &Envelope{
		XMLName: xml.Name{
			Space: Namespace,
			Local: "Envelope",
		},
	}
}

// NewEnvelopeWithNamespace creates a new SOAP envelope with a custom namespace URI and prefix.
// This allows for maximum flexibility when working with different SOAP implementations.
func NewEnvelopeWithNamespace(namespaceURI, prefix string) *Envelope {
	return &Envelope{
		XMLName: xml.Name{
			Space: namespaceURI,
			Local: "Envelope",
		},
		Attrs: []xml.Attr{
			{
				Name:  xml.Name{Local: "xmlns:" + prefix},
				Value: namespaceURI,
			},
		},
	}
}

// NewEnvelopeWithBody creates a new SOAP envelope with the specified body content.
func NewEnvelopeWithBody(bodyContent []byte) *Envelope {
	env := NewEnvelope()
	env.Body = Body{Content: bodyContent}
	return env
}

// NewEnvelopeWithBodyAndNamespace creates a new SOAP envelope with custom namespace and body content.
func NewEnvelopeWithBodyAndNamespace(namespaceURI, prefix string, bodyContent []byte) *Envelope {
	env := NewEnvelopeWithNamespace(namespaceURI, prefix)
	env.Body = Body{Content: bodyContent}
	return env
}

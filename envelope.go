package soap

import (
	"encoding/xml"
	"fmt"
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
// It implements the error interface to allow SOAP faults to be used as Go errors.
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

// Error implements the error interface for Fault.
// Returns the fault string as the error message.
func (f *Fault) Error() string {
	return fmt.Sprintf("SOAP fault %s: %s", f.FaultString, f.FaultCode)
}

// Detail represents fault detail information.
// Contains application-specific error data as per SOAP 1.1 spec section 4.4.
type Detail struct {
	// Content as raw XML to accommodate any application-specific error data
	Content []byte `xml:",innerxml"`

	// Additional attributes for extensibility
	Attrs []xml.Attr `xml:",any,attr"`
}

// EnvelopeOption is a function that configures an Envelope.
type EnvelopeOption func(*Envelope) error

// WithNamespace sets the namespace for the Envelope.
func WithNamespace(namespace string) EnvelopeOption {
	return func(env *Envelope) error {
		env.XMLName.Space = namespace
		env.XMLName.Local = "Envelope"
		env.Attrs = []xml.Attr{
			{
				Name:  xml.Name{Local: "xmlns:" + namespace},
				Value: namespace,
			},
		}
		return nil
	}
}

// WithBody sets the body for the Envelope.
func WithBody(body any) EnvelopeOption {
	return func(env *Envelope) error {
		if body == nil {
			return fmt.Errorf("body is nil")
		}
		switch body := body.(type) {
		case []byte:
			env.Body = Body{Content: body}
			return nil
		default:
			xmlData, err := xml.Marshal(body)
			if err != nil {
				return err
			}
			env.Body = Body{Content: xmlData}
			return nil
		}
	}
}

// NewEnvelope creates a new SOAP envelope with the specified options.
func NewEnvelope(opts ...EnvelopeOption) (*Envelope, error) {
	var result Envelope
	for _, opt := range opts {
		if err := opt(&result); err != nil {
			return nil, err
		}
	}
	return &result, nil
}

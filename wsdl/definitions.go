package wsdl

import (
	"encoding/xml"
	"os"

	"github.com/way-platform/soap-go/xsd"
)

// ParseFromFile reads a WSDL file from disk and unmarshals it.
func ParseFromFile(filename string) (*Definitions, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var defs Definitions
	if err := xml.Unmarshal(data, &defs); err != nil {
		return nil, err
	}
	return &defs, nil
}

// Definitions represents a WSDL 1.1 file, corresponding to the <definitions> element.
// It can handle both namespaced and non-namespaced WSDL documents.
type Definitions struct {
	XMLName         xml.Name `xml:"definitions"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Name            string   `xml:"name,attr"`

	Imports  []Import   `xml:"import"`
	Types    *Types     `xml:"types"`
	Messages []Message  `xml:"message"`
	PortType []PortType `xml:"portType"`
	Binding  []Binding  `xml:"binding"`
	Service  []Service  `xml:"service"`
}

// Import corresponds to the <import> element.
type Import struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

// Types corresponds to the <types> element.
type Types struct {
	Schemas []xsd.Schema `xml:"schema"`
}

// Message corresponds to the <message> element.
type Message struct {
	Name  string `xml:"name,attr"`
	Parts []Part `xml:"part"`
}

// Part corresponds to the <part> element within a <message>.
type Part struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
	Type    string `xml:"type,attr"`
}

// PortType corresponds to the <portType> element.
type PortType struct {
	Name       string      `xml:"name,attr"`
	Operations []Operation `xml:"operation"`
}

// Operation corresponds to the <operation> element within a <portType>.
type Operation struct {
	Name           string   `xml:"name,attr"`
	ParameterOrder string   `xml:"parameterOrder,attr"`
	Input          *Input   `xml:"input"`
	Output         *Output  `xml:"output"`
	Faults         []*Fault `xml:"fault"`
}

// Input corresponds to the <input> element of an <operation>.
type Input struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
}

// Output corresponds to the <output> element of an <operation>.
type Output struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
}

// Fault corresponds to the <fault> element of an <operation>.
type Fault struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
}

// Binding corresponds to the <binding> element.
type Binding struct {
	Name              string             `xml:"name,attr"`
	Type              string             `xml:"type,attr"`
	SOAP11Binding     *SOAPBinding       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
	SOAP12Binding     *SOAPBinding       `xml:"http://schemas.xmlsoap.org/wsdl/soap12/ binding"`
	HTTPBinding       *HTTPBinding       `xml:"http://schemas.xmlsoap.org/wsdl/http/ binding"`
	BindingOperations []BindingOperation `xml:"operation"`
}

// SOAPBinding corresponds to the <soap:binding> or <soap12:binding> element.
type SOAPBinding struct {
	Style     string `xml:"style,attr"`
	Transport string `xml:"transport,attr"`
}

// HTTPBinding corresponds to the <http:binding> element.
type HTTPBinding struct {
	Verb string `xml:"verb,attr"`
}

// BindingOperation corresponds to the <operation> element within a <binding>.
type BindingOperation struct {
	Name            string         `xml:"name,attr"`
	SOAP11Operation *SOAPOperation `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
	SOAP12Operation *SOAPOperation `xml:"http://schemas.xmlsoap.org/wsdl/soap12/ operation"`
	HTTPOperation   *HTTPOperation `xml:"http://schemas.xmlsoap.org/wsdl/http/ operation"`
	Input           *BindingBody   `xml:"input"`
	Output          *BindingBody   `xml:"output"`
	Faults          []BindingFault `xml:"fault"`
}

// SOAPOperation corresponds to the <soap:operation> or <soap12:operation> element.
type SOAPOperation struct {
	SOAPAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

// HTTPOperation corresponds to the <http:operation> element.
type HTTPOperation struct {
	Location string `xml:"location,attr"`
}

// BindingBody represents <input> or <output> elements inside a <binding><operation>.
type BindingBody struct {
	SOAP11Body           *SOAPBody             `xml:"http://schemas.xmlsoap.org/wsdl/soap/ body"`
	SOAP12Body           *SOAPBody             `xml:"http://schemas.xmlsoap.org/wsdl/soap12/ body"`
	URLReplacement       *URLReplacement       `xml:"http://schemas.xmlsoap.org/wsdl/http/ urlReplacement"`
	URLEncoded           *URLEncoded           `xml:"http://schemas.xmlsoap.org/wsdl/http/ urlEncoded"`
	MIMEContent          []*MIMEContent        `xml:"http://schemas.xmlsoap.org/wsdl/mime/ content"`
	MIMEMultipartRelated *MIMEMultipartRelated `xml:"http://schemas.xmlsoap.org/wsdl/mime/ multipartRelated"`
	MIMEXML              *MIMEXML              `xml:"http://schemas.xmlsoap.org/wsdl/mime/ mimeXml"`
}

// SOAPBody corresponds to the <soap:body> or <soap12:body> element.
type SOAPBody struct {
	Use           string `xml:"use,attr"`
	Namespace     string `xml:"namespace,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
	Parts         string `xml:"parts,attr"`
}

// URLReplacement corresponds to the <http:urlReplacement> element.
type URLReplacement struct{}

// URLEncoded corresponds to the <http:urlEncoded> element.
type URLEncoded struct{}

// MIMEContent corresponds to the <mime:content> element.
type MIMEContent struct {
	Part string `xml:"part,attr"`
	Type string `xml:"type,attr"`
}

// MIMEMultipartRelated corresponds to the <mime:multipartRelated> element.
type MIMEMultipartRelated struct {
	Parts []MIMEPart `xml:"part"`
}

// MIMEPart corresponds to the <mime:part> element.
type MIMEPart struct {
	SOAP11Body           *SOAPBody             `xml:"http://schemas.xmlsoap.org/wsdl/soap/ body"`
	MIMEContent          []*MIMEContent        `xml:"http://schemas.xmlsoap.org/wsdl/mime/ content"`
	MIMEMultipartRelated *MIMEMultipartRelated `xml:"http://schemas.xmlsoap.org/wsdl/mime/ multipartRelated"`
	MIMEXML              *MIMEXML              `xml:"http://schemas.xmlsoap.org/wsdl/mime/ mimeXml"`
}

// MIMEXML corresponds to the <mime:mimeXml> element.
type MIMEXML struct {
	Part string `xml:"part,attr"`
}

// BindingFault corresponds to the <fault> element within a <binding><operation>.
type BindingFault struct {
	Name        string     `xml:"name,attr"`
	SOAP11Fault *SOAPFault `xml:"http://schemas.xmlsoap.org/wsdl/soap/ fault"`
}

// SOAPFault corresponds to the <soap:fault> element.
type SOAPFault struct {
	Name          string `xml:"name,attr"`
	Use           string `xml:"use,attr"`
	Namespace     string `xml:"namespace,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
}

// Service corresponds to the <service> element.
type Service struct {
	Name  string `xml:"name,attr"`
	Ports []Port `xml:"port"`
}

// Port corresponds to the <port> element within a <service>.
type Port struct {
	Name          string       `xml:"name,attr"`
	Binding       string       `xml:"binding,attr"`
	SOAP11Address *SOAPAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
	SOAP12Address *SOAPAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap12/ address"`
	HTTPAddress   *HTTPAddress `xml:"http://schemas.xmlsoap.org/wsdl/http/ address"`
}

// SOAPAddress corresponds to the <soap:address> or <soap12:address> element.
type SOAPAddress struct {
	Location string `xml:"location,attr"`
}

// HTTPAddress corresponds to the <http:address> element.
type HTTPAddress struct {
	Location string `xml:"location,attr"`
}

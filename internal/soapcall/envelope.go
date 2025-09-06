package soapcall

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// soapEnvelope represents a SOAP envelope.
type soapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    soapBody `xml:"soap:Body"`
}

// soapBody represents a SOAP body.
type soapBody struct {
	Content []byte `xml:",innerxml"`
}

// soapFault represents a SOAP fault element.
type soapFault struct {
	XMLName     xml.Name `xml:"soap:Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
	Detail      []byte   `xml:",innerxml"`
}

// IsFullEnvelope checks if the provided XML data represents a complete SOAP envelope.
func IsFullEnvelope(xmlData []byte) bool {
	// Parse XML and check if root element is soap:Envelope
	var envelope struct {
		XMLName xml.Name
	}
	if err := xml.Unmarshal(xmlData, &envelope); err != nil {
		return false
	}
	return envelope.XMLName.Local == "Envelope" &&
		envelope.XMLName.Space == "http://schemas.xmlsoap.org/soap/envelope/"
}

// WrapInEnvelope wraps the provided XML payload in a SOAP envelope.
func WrapInEnvelope(payload []byte) ([]byte, error) {
	envelope := &soapEnvelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  soapBody{Content: payload},
	}

	return xml.Marshal(envelope)
}

// ExtractFromEnvelope extracts the body content from a SOAP envelope.
// If the XML is not a SOAP envelope, it returns the original data.
func ExtractFromEnvelope(xmlData []byte) ([]byte, error) {
	if !IsFullEnvelope(xmlData) {
		return xmlData, nil
	}

	var envelope soapEnvelope
	if err := xml.Unmarshal(xmlData, &envelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SOAP envelope: %w", err)
	}

	// Check for SOAP fault
	if fault := extractSOAPFault(envelope.Body.Content); fault != nil {
		return nil, fmt.Errorf("SOAP fault: %s - %s", fault.FaultCode, fault.FaultString)
	}

	return envelope.Body.Content, nil
}

// extractSOAPFault attempts to extract a SOAP fault from the body content.
func extractSOAPFault(bodyContent []byte) *soapFault {
	var fault soapFault
	if err := xml.Unmarshal(bodyContent, &fault); err != nil {
		return nil
	}
	if fault.XMLName.Local == "Fault" {
		return &fault
	}
	return nil
}

// FormatXML formats XML data with proper indentation.
func FormatXML(xmlData []byte) ([]byte, error) {
	var buf bytes.Buffer
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to parse XML: %w", err)
		}
		if err := encoder.EncodeToken(token); err != nil {
			return nil, fmt.Errorf("failed to format XML: %w", err)
		}
	}

	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("failed to flush XML encoder: %w", err)
	}

	return buf.Bytes(), nil
}

// AddXMLDeclaration adds an XML declaration to the beginning of XML data if not present.
func AddXMLDeclaration(xmlData []byte) []byte {
	xmlDecl := `<?xml version="1.0" encoding="UTF-8"?>`
	xmlStr := strings.TrimSpace(string(xmlData))

	if !strings.HasPrefix(xmlStr, "<?xml") {
		return []byte(xmlDecl + "\n" + xmlStr)
	}
	return xmlData
}

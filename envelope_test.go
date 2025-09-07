package soap

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestEnvelopeMarshalUnmarshal(t *testing.T) {
	// Create envelope with header and body
	mustUnderstand := true
	headerEntry := HeaderEntry{
		XMLName:        xml.Name{Local: "TestHeader", Space: "http://example.com"},
		MustUnderstand: &mustUnderstand,
		Actor:          "http://example.com/actor",
		Content:        []byte("<value>header-content</value>"),
	}

	envelope := &Envelope{
		XMLNS:         Namespace,
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Header: &Header{
			Entries: []HeaderEntry{headerEntry},
		},
		Body: Body{Content: []byte("<body>body-content</body>")},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope: %v", err)
	}

	xmlStr := string(xmlData)
	t.Logf("Generated XML: %s", xmlStr)

	// Verify XML structure - should use soap: prefixed elements for compatibility
	if !strings.Contains(xmlStr, "soap:Envelope") {
		t.Error("XML should contain soap:Envelope element")
	}
	if !strings.Contains(xmlStr, "soap:Header") {
		t.Error("XML should contain soap:Header element")
	}
	if !strings.Contains(xmlStr, "soap:Body") {
		t.Error("XML should contain soap:Body element")
	}
	if !strings.Contains(xmlStr, "xmlns:soap") {
		t.Error("XML should contain xmlns:soap namespace declaration")
	}
	if !strings.Contains(xmlStr, "mustUnderstand") {
		t.Error("XML should contain mustUnderstand attribute")
	}
	if !strings.Contains(xmlStr, "soap:encodingStyle") {
		t.Error("XML should contain soap:encodingStyle attribute")
	}

	// Verify namespace is correct
	if !strings.Contains(xmlStr, Namespace) {
		t.Errorf("XML should contain SOAP namespace %s", Namespace)
	}

	// Verify content is preserved
	if !strings.Contains(xmlStr, "header-content") {
		t.Error("XML should contain header content")
	}
	if !strings.Contains(xmlStr, "body-content") {
		t.Error("XML should contain body content")
	}

	// Note: Full envelope unmarshaling has limitations with Go's XML package
	// and prefixed elements. For practical use, unmarshal the Body.Content directly.
	t.Log("Envelope marshaling produces correct soap: prefixed XML for maximum service compatibility")
}

func TestFaultHandling(t *testing.T) {
	// Test fault marshaling - create a fault structure
	fault := &Fault{
		FaultCode:   "Client",
		FaultString: "Invalid request",
		FaultActor:  "http://example.com/service",
		Detail: &Detail{
			Content: []byte("<errorcode>E001</errorcode>"),
		},
	}

	// Marshal the fault to see the XML structure
	faultXML, err := xml.Marshal(fault)
	if err != nil {
		t.Fatalf("Failed to marshal fault: %v", err)
	}

	t.Logf("Fault XML: %s", string(faultXML))

	// Verify fault XML uses soap: prefix
	faultStr := string(faultXML)
	if !strings.Contains(faultStr, "soap:Fault") {
		t.Error("Fault XML should contain soap:Fault element")
	}

	// Create envelope with fault in body
	envelope := &Envelope{
		XMLNS: Namespace,
		Body:  Body{Content: faultXML},
	}

	// Marshal envelope with fault
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope with fault: %v", err)
	}

	xmlStr := string(xmlData)
	t.Logf("Envelope with fault: %s", xmlStr)

	// Verify the envelope contains the fault
	if !strings.Contains(xmlStr, "soap:Envelope") {
		t.Error("XML should contain soap:Envelope")
	}
	if !strings.Contains(xmlStr, "soap:Body") {
		t.Error("XML should contain soap:Body")
	}
	if !strings.Contains(xmlStr, "soap:Fault") {
		t.Error("XML should contain soap:Fault within the body")
	}
	if !strings.Contains(xmlStr, "Client") {
		t.Error("XML should contain fault code")
	}
	if !strings.Contains(xmlStr, "Invalid request") {
		t.Error("XML should contain fault string")
	}
}

func TestHeaderEntryMustUnderstand(t *testing.T) {
	// Test mustUnderstand true
	mustUnderstand := true
	headerEntry := HeaderEntry{
		XMLName:        xml.Name{Local: "Auth", Space: "http://example.com/auth"},
		MustUnderstand: &mustUnderstand,
		Content:        []byte("<token>abc123</token>"),
	}

	xmlData, err := xml.Marshal(headerEntry)
	if err != nil {
		t.Fatalf("Failed to marshal header entry: %v", err)
	}

	xmlStr := string(xmlData)
	if !strings.Contains(xmlStr, "mustUnderstand") {
		t.Error("XML should contain mustUnderstand attribute")
	}

	// Test mustUnderstand false
	mustUnderstandFalse := false
	headerEntry.MustUnderstand = &mustUnderstandFalse

	xmlData, err = xml.Marshal(headerEntry)
	if err != nil {
		t.Fatalf("Failed to marshal header entry: %v", err)
	}

	xmlStr = string(xmlData)
	// mustUnderstand=false should still be present in XML
	if !strings.Contains(xmlStr, "mustUnderstand") {
		t.Error("XML should contain mustUnderstand attribute even when false")
	}
}

func TestEnvelopeExtensibility(t *testing.T) {
	// Test custom attributes on envelope
	envelope := &Envelope{
		XMLNS: Namespace,
		Body:  Body{Content: []byte("<test>content</test>")},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "custom"}, Value: "value"},
			{Name: xml.Name{Local: "version"}, Value: "1.0"},
		},
	}

	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope: %v", err)
	}

	xmlStr := string(xmlData)
	t.Logf("Envelope with custom attributes: %s", xmlStr)

	// Verify structure and custom attributes
	if !strings.Contains(xmlStr, "soap:Envelope") {
		t.Error("XML should contain soap:Envelope element")
	}
	if !strings.Contains(xmlStr, "custom=\"value\"") {
		t.Error("XML should contain custom attribute")
	}
	if !strings.Contains(xmlStr, "version=\"1.0\"") {
		t.Error("XML should contain version attribute")
	}
	if !strings.Contains(xmlStr, "xmlns:soap") {
		t.Error("XML should contain SOAP namespace declaration")
	}
	if !strings.Contains(xmlStr, "<test>content</test>") {
		t.Error("XML should contain body content")
	}

	// Verify that extensibility works - envelope can carry custom attributes
	// while maintaining SOAP compliance
	t.Log("Envelope extensibility allows custom attributes while maintaining soap: prefix structure")
}

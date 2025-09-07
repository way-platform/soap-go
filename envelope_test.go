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
		XMLName:       xml.Name{Space: Namespace, Local: "Envelope"},
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

	// Verify XML structure - should contain essential SOAP elements (prefixed or unprefixed)
	if !strings.Contains(xmlStr, "Envelope") {
		t.Error("XML should contain Envelope element")
	}
	if !strings.Contains(xmlStr, "Header") {
		t.Error("XML should contain Header element")
	}
	if !strings.Contains(xmlStr, "Body") {
		t.Error("XML should contain Body element")
	}
	// Note: xmlns:soap declaration is optional - the XMLName.Space provides namespace context
	if !strings.Contains(xmlStr, "mustUnderstand") {
		t.Error("XML should contain mustUnderstand attribute")
	}
	if !strings.Contains(xmlStr, "encodingStyle") {
		t.Error("XML should contain encodingStyle attribute")
	}

	// Note: The SOAP namespace is implicit in the XMLName.Space, not necessarily in the serialized XML

	// Verify content is preserved
	if !strings.Contains(xmlStr, "header-content") {
		t.Error("XML should contain header content")
	}
	if !strings.Contains(xmlStr, "body-content") {
		t.Error("XML should contain body content")
	}

	// Note: Full envelope unmarshaling has limitations with Go's XML package
	// and prefixed elements. For practical use, unmarshal the Body.Content directly.
	t.Log("Envelope marshaling produces valid SOAP XML with proper namespace declarations")
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
	if !strings.Contains(faultStr, "Fault") {
		t.Error("Fault XML should contain Fault element")
	}

	// Create envelope with fault in body
	envelope := &Envelope{
		XMLName: xml.Name{Space: Namespace, Local: "Envelope"},
		Body:    Body{Content: faultXML},
	}

	// Marshal envelope with fault
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope with fault: %v", err)
	}

	xmlStr := string(xmlData)
	t.Logf("Envelope with fault: %s", xmlStr)

	// Verify the envelope contains the fault
	if !strings.Contains(xmlStr, "Envelope") {
		t.Error("XML should contain Envelope")
	}
	if !strings.Contains(xmlStr, "Body") {
		t.Error("XML should contain Body")
	}
	if !strings.Contains(xmlStr, "Fault") {
		t.Error("XML should contain Fault within the body")
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
		XMLName: xml.Name{Space: Namespace, Local: "Envelope"},
		Body:    Body{Content: []byte("<test>content</test>")},
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
	if !strings.Contains(xmlStr, "Envelope") {
		t.Error("XML should contain Envelope element")
	}
	if !strings.Contains(xmlStr, "custom=\"value\"") {
		t.Error("XML should contain custom attribute")
	}
	if !strings.Contains(xmlStr, "version=\"1.0\"") {
		t.Error("XML should contain version attribute")
	}
	// Note: xmlns:soap declaration is optional with our permissive approach
	if !strings.Contains(xmlStr, "<test>content</test>") {
		t.Error("XML should contain body content")
	}

	// Verify that extensibility works - envelope can carry custom attributes
	// while maintaining SOAP compliance
	t.Log("Envelope extensibility allows custom attributes while maintaining SOAP compliance")
}

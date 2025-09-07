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

	// Verify XML contains expected elements
	if !strings.Contains(xmlStr, "Envelope") {
		t.Error("XML should contain Envelope element")
	}
	if !strings.Contains(xmlStr, "Header") {
		t.Error("XML should contain Header element")
	}
	if !strings.Contains(xmlStr, "Body") {
		t.Error("XML should contain Body element")
	}
	if !strings.Contains(xmlStr, "mustUnderstand") {
		t.Error("XML should contain mustUnderstand attribute")
	}
	if !strings.Contains(xmlStr, "encodingStyle") {
		t.Error("XML should contain encodingStyle attribute")
	}

	// Unmarshal back
	var unmarshaled Envelope
	if err := xml.Unmarshal(xmlData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal envelope: %v", err)
	}

	// Verify unmarshaled data - XMLName should have the correct namespace
	if unmarshaled.XMLName.Space != Namespace {
		t.Errorf("Expected XMLName.Space %s, got %s", Namespace, unmarshaled.XMLName.Space)
	}

	if unmarshaled.XMLName.Local != "Envelope" {
		t.Errorf("Expected XMLName.Local 'Envelope', got %s", unmarshaled.XMLName.Local)
	}

	if unmarshaled.Header == nil {
		t.Error("Expected header to be present")
	}

	if len(unmarshaled.Header.Entries) != 1 {
		t.Errorf("Expected 1 header entry, got %d", len(unmarshaled.Header.Entries))
	}
}

func TestFaultHandling(t *testing.T) {
	// Create envelope with fault in body
	faultXML := `<Fault>
		<faultcode>Client</faultcode>
		<faultstring>Invalid request</faultstring>
		<faultactor>http://example.com/service</faultactor>
		<detail>
			<errorcode>E001</errorcode>
		</detail>
	</Fault>`

	envelope := &Envelope{
		XMLNS: Namespace,
		Body:  Body{Content: []byte(faultXML)},
	}

	// Marshal and unmarshal to ensure proper handling
	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope: %v", err)
	}

	var unmarshaled Envelope
	if err := xml.Unmarshal(xmlData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal envelope: %v", err)
	}

	// Verify fault can be extracted
	var fault Fault
	if err := xml.Unmarshal(unmarshaled.Body.Content, &fault); err != nil {
		t.Fatalf("Failed to unmarshal fault: %v", err)
	}

	if fault.FaultCode != "Client" {
		t.Errorf("Expected fault code 'Client', got %s", fault.FaultCode)
	}

	if fault.FaultString != "Invalid request" {
		t.Errorf("Expected fault string 'Invalid request', got %s", fault.FaultString)
	}

	if fault.FaultActor != "http://example.com/service" {
		t.Errorf("Expected fault actor 'http://example.com/service', got %s", fault.FaultActor)
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
		},
	}

	xmlData, err := xml.Marshal(envelope)
	if err != nil {
		t.Fatalf("Failed to marshal envelope: %v", err)
	}

	xmlStr := string(xmlData)
	if !strings.Contains(xmlStr, "custom=\"value\"") {
		t.Error("XML should contain custom attribute")
	}

	// Unmarshal back
	var unmarshaled Envelope
	if err := xml.Unmarshal(xmlData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal envelope: %v", err)
	}

	// Verify custom attribute is preserved
	found := false
	for _, attr := range unmarshaled.Attrs {
		if attr.Name.Local == "custom" && attr.Value == "value" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Custom attribute should be preserved after unmarshal")
	}
}

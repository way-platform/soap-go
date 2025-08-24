package wsdl11_test

import (
	"testing"

	"github.com/way-platform/soap-go/wsdl11"
)

func TestParseNumberConversion(t *testing.T) {
	// Test parsing NumberConversion.wsdl using the new ParseFromFile function
	defs, err := wsdl11.ParseFromFile("../testdata/NumberConversion.wsdl")
	if err != nil {
		t.Fatalf("failed to parse NumberConversion.wsdl: %v", err)
	}

	// Verify basic WSDL structure
	if defs.TargetNamespace != "http://www.dataaccess.com/webservicesserver/" {
		t.Errorf("expected targetNamespace to be 'http://www.dataaccess.com/webservicesserver/', got %q", defs.TargetNamespace)
	}

	if defs.Name != "NumberConversion" {
		t.Errorf("expected name to be 'NumberConversion', got %q", defs.Name)
	}

	// Verify that the XSD schema was parsed
	if defs.Types == nil {
		t.Fatal("expected Types to be non-nil")
	}

	if len(defs.Types.Schemas) != 1 {
		t.Fatalf("expected 1 schema, got %d", len(defs.Types.Schemas))
	}

	schema := defs.Types.Schemas[0]
	if schema.TargetNamespace != "http://www.dataaccess.com/webservicesserver/" {
		t.Errorf("expected schema targetNamespace to be 'http://www.dataaccess.com/webservicesserver/', got %q", schema.TargetNamespace)
	}

	if schema.ElementFormDefault != "qualified" {
		t.Errorf("expected schema elementFormDefault to be 'qualified', got %q", schema.ElementFormDefault)
	}

	// Verify that XSD elements are parsed correctly
	expectedElements := []string{"NumberToWords", "NumberToWordsResponse", "NumberToDollars", "NumberToDollarsResponse"}
	if len(schema.Elements) != len(expectedElements) {
		t.Errorf("expected %d schema elements, got %d", len(expectedElements), len(schema.Elements))
	}

	elementNames := make(map[string]bool)
	for _, el := range schema.Elements {
		elementNames[el.Name] = true
	}

	for _, expected := range expectedElements {
		if !elementNames[expected] {
			t.Errorf("expected schema element %q not found", expected)
		}
	}

	// Verify complex types in elements
	for _, el := range schema.Elements {
		if el.ComplexType == nil {
			t.Errorf("expected element %q to have a complex type", el.Name)
			continue
		}
		if el.ComplexType.Sequence == nil {
			t.Errorf("expected element %q complex type to have a sequence", el.Name)
			continue
		}
		if len(el.ComplexType.Sequence.Elements) == 0 {
			t.Errorf("expected element %q sequence to have elements", el.Name)
		}
	}

	// Verify WSDL messages
	expectedMessages := []string{
		"NumberToWordsSoapRequest", "NumberToWordsSoapResponse",
		"NumberToDollarsSoapRequest", "NumberToDollarsSoapResponse",
	}
	if len(defs.Messages) != len(expectedMessages) {
		t.Errorf("expected %d messages, got %d", len(expectedMessages), len(defs.Messages))
	}

	messageNames := make(map[string]bool)
	for _, msg := range defs.Messages {
		messageNames[msg.Name] = true
	}

	for _, expected := range expectedMessages {
		if !messageNames[expected] {
			t.Errorf("expected message %q not found", expected)
		}
	}

	// Verify port types
	if len(defs.PortType) != 1 {
		t.Errorf("expected 1 port type, got %d", len(defs.PortType))
	}

	if defs.PortType[0].Name != "NumberConversionSoapType" {
		t.Errorf("expected port type name to be 'NumberConversionSoapType', got %q", defs.PortType[0].Name)
	}

	expectedOperations := []string{"NumberToWords", "NumberToDollars"}
	if len(defs.PortType[0].Operations) != len(expectedOperations) {
		t.Errorf("expected %d operations, got %d", len(expectedOperations), len(defs.PortType[0].Operations))
	}

	// Verify bindings
	expectedBindings := []string{"NumberConversionSoapBinding", "NumberConversionSoapBinding12"}
	if len(defs.Binding) != len(expectedBindings) {
		t.Errorf("expected %d bindings, got %d", len(expectedBindings), len(defs.Binding))
	}

	// Verify services
	if len(defs.Service) != 1 {
		t.Errorf("expected 1 service, got %d", len(defs.Service))
	}

	if defs.Service[0].Name != "NumberConversion" {
		t.Errorf("expected service name to be 'NumberConversion', got %q", defs.Service[0].Name)
	}

	expectedPorts := []string{"NumberConversionSoap", "NumberConversionSoap12"}
	if len(defs.Service[0].Ports) != len(expectedPorts) {
		t.Errorf("expected %d ports, got %d", len(expectedPorts), len(defs.Service[0].Ports))
	}
}

func TestParseGlobalWeatherWithNewAPI(t *testing.T) {
	// Test parsing GlobalWeather.wsdl using the new ParseFromFile function
	defs, err := wsdl11.ParseFromFile("../testdata/GlobalWeather.wsdl")
	if err != nil {
		t.Fatalf("failed to parse GlobalWeather.wsdl: %v", err)
	}

	// Basic validation
	if defs.TargetNamespace != "http://www.webserviceX.NET" {
		t.Errorf("expected targetNamespace to be 'http://www.webserviceX.NET', got %q", defs.TargetNamespace)
	}

	// Verify that the XSD schema was parsed
	if defs.Types == nil || len(defs.Types.Schemas) == 0 {
		t.Fatal("expected to parse schema content, but got none")
	}

	schema := defs.Types.Schemas[0]
	if schema.TargetNamespace != "http://www.webserviceX.NET" {
		t.Errorf("expected schema targetNamespace to be 'http://www.webserviceX.NET', got %q", schema.TargetNamespace)
	}

	if schema.ElementFormDefault != "qualified" {
		t.Errorf("expected schema elementFormDefault to be 'qualified', got %q", schema.ElementFormDefault)
	}

	// Should have parsed some schema elements
	if len(schema.Elements) == 0 {
		t.Error("expected to parse some schema elements, got none")
	}

	// Verify messages were parsed
	if len(defs.Messages) == 0 {
		t.Error("expected to parse some messages, got none")
	}

	// Verify port types were parsed
	if len(defs.PortType) == 0 {
		t.Error("expected to parse some port types, got none")
	}
}

package xsd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/way-platform/soap-go/xsd"
)

func TestParseSoapEnvelope(t *testing.T) {
	// Read the SOAP envelope XSD file
	data, err := os.ReadFile("testdata/soap-envelope.xsd")
	if err != nil {
		t.Fatalf("failed to read soap-envelope.xsd: %v", err)
	}

	// Parse the XSD
	schema, err := xsd.Parse(strings.NewReader(string(data)))
	if err != nil {
		t.Fatalf("failed to parse soap-envelope.xsd: %v", err)
	}

	// Validate the parsed schema
	if schema.TargetNamespace != "http://schemas.xmlsoap.org/soap/envelope/" {
		t.Errorf("expected targetNamespace to be 'http://schemas.xmlsoap.org/soap/envelope/', got %q", schema.TargetNamespace)
	}

	// Check that we have the expected elements
	expectedElements := []string{"Envelope", "Header", "Body", "Fault"}
	if len(schema.Elements) != len(expectedElements) {
		t.Errorf("expected %d elements, got %d", len(expectedElements), len(schema.Elements))
	}

	elementNames := make(map[string]bool)
	for _, el := range schema.Elements {
		elementNames[el.Name] = true
	}

	for _, expected := range expectedElements {
		if !elementNames[expected] {
			t.Errorf("expected element %q not found", expected)
		}
	}

	// Check complex types
	expectedComplexTypes := []string{"Envelope", "Header", "Body", "Fault", "detail"}
	if len(schema.ComplexTypes) != len(expectedComplexTypes) {
		t.Errorf("expected %d complex types, got %d", len(expectedComplexTypes), len(schema.ComplexTypes))
	}

	complexTypeNames := make(map[string]bool)
	for _, ct := range schema.ComplexTypes {
		complexTypeNames[ct.Name] = true
	}

	for _, expected := range expectedComplexTypes {
		if !complexTypeNames[expected] {
			t.Errorf("expected complex type %q not found", expected)
		}
	}

	// Check attributes
	expectedAttributes := []string{"mustUnderstand", "actor", "encodingStyle"}
	if len(schema.Attributes) != len(expectedAttributes) {
		t.Errorf("expected %d attributes, got %d", len(expectedAttributes), len(schema.Attributes))
	}

	// Check that Envelope complex type has a sequence
	var envelopeType *xsd.ComplexType
	for _, ct := range schema.ComplexTypes {
		if ct.Name == "Envelope" {
			envelopeType = &ct
			break
		}
	}

	if envelopeType == nil {
		t.Fatal("Envelope complex type not found")
	}

	if envelopeType.Sequence == nil {
		t.Error("Envelope complex type should have a sequence")
	}

	// Check sequence elements
	if len(envelopeType.Sequence.Elements) != 2 {
		t.Errorf("expected 2 elements in Envelope sequence, got %d", len(envelopeType.Sequence.Elements))
	}

	// Check that sequence has Any elements
	if len(envelopeType.Sequence.Any) != 1 {
		t.Errorf("expected 1 Any element in Envelope sequence, got %d", len(envelopeType.Sequence.Any))
	}

	// Check anyAttribute
	if envelopeType.AnyAttribute == nil {
		t.Error("Envelope complex type should have anyAttribute")
	}

	// Check simple types
	expectedSimpleTypes := []string{"encodingStyle"}
	if len(schema.SimpleTypes) != len(expectedSimpleTypes) {
		t.Errorf("expected %d simple types, got %d", len(expectedSimpleTypes), len(schema.SimpleTypes))
	}

	// Check attribute groups
	expectedAttributeGroups := []string{"encodingStyle"}
	if len(schema.AttributeGroups) != len(expectedAttributeGroups) {
		t.Errorf("expected %d attribute groups, got %d", len(expectedAttributeGroups), len(schema.AttributeGroups))
	}
}

func TestParseIBM(t *testing.T) {
	// Read the IBM XSD file
	data, err := os.ReadFile("testdata/ibm.xsd")
	if err != nil {
		t.Fatalf("failed to read ibm.xsd: %v", err)
	}

	// Parse the XSD
	schema, err := xsd.Parse(strings.NewReader(string(data)))
	if err != nil {
		t.Fatalf("failed to parse ibm.xsd: %v", err)
	}

	// Validate the parsed schema
	if schema.TargetNamespace != "http://com.ibm.wbit.comptest.controller" {
		t.Errorf("expected targetNamespace to be 'http://com.ibm.wbit.comptest.controller', got %q", schema.TargetNamespace)
	}

	// Check complex types
	expectedComplexTypes := []string{"TestResults", "TestSuiteRun", "TestCaseRun", "VariationRun", "TestRun"}
	if len(schema.ComplexTypes) != len(expectedComplexTypes) {
		t.Errorf("expected %d complex types, got %d", len(expectedComplexTypes), len(schema.ComplexTypes))
	}

	complexTypeNames := make(map[string]bool)
	for _, ct := range schema.ComplexTypes {
		complexTypeNames[ct.Name] = true
	}

	for _, expected := range expectedComplexTypes {
		if !complexTypeNames[expected] {
			t.Errorf("expected complex type %q not found", expected)
		}
	}

	// Check simple types
	expectedSimpleTypes := []string{"Severity"}
	if len(schema.SimpleTypes) != len(expectedSimpleTypes) {
		t.Errorf("expected %d simple types, got %d", len(expectedSimpleTypes), len(schema.SimpleTypes))
	}

	// Check that Severity simple type has restriction with enumerations
	var severityType *xsd.SimpleType
	for _, st := range schema.SimpleTypes {
		if st.Name == "Severity" {
			severityType = &st
			break
		}
	}

	if severityType == nil {
		t.Fatal("Severity simple type not found")
	}

	if severityType.Restriction == nil {
		t.Error("Severity simple type should have a restriction")
	}

	if severityType.Restriction.Base != "xsd:string" {
		t.Errorf("expected Severity restriction base to be 'xsd:string', got %q", severityType.Restriction.Base)
	}

	expectedEnums := []string{"pass", "fail", "error"}
	if len(severityType.Restriction.Enumerations) != len(expectedEnums) {
		t.Errorf("expected %d enumerations, got %d", len(expectedEnums), len(severityType.Restriction.Enumerations))
	}

	for i, enum := range severityType.Restriction.Enumerations {
		if i < len(expectedEnums) && enum.Value != expectedEnums[i] {
			t.Errorf("expected enumeration value %q, got %q", expectedEnums[i], enum.Value)
		}
	}

	// Check complex content and extension
	var testSuiteRunType *xsd.ComplexType
	for _, ct := range schema.ComplexTypes {
		if ct.Name == "TestSuiteRun" {
			testSuiteRunType = &ct
			break
		}
	}

	if testSuiteRunType == nil {
		t.Fatal("TestSuiteRun complex type not found")
	}

	if testSuiteRunType.ComplexContent == nil {
		t.Error("TestSuiteRun should have complex content")
	}

	if testSuiteRunType.ComplexContent.Extension == nil {
		t.Error("TestSuiteRun complex content should have extension")
	}

	if testSuiteRunType.ComplexContent.Extension.Base != "Q1:TestRun" {
		t.Errorf("expected extension base to be 'Q1:TestRun', got %q", testSuiteRunType.ComplexContent.Extension.Base)
	}

	// Check that extension has sequence and attributes
	if testSuiteRunType.ComplexContent.Extension.Sequence == nil {
		t.Error("TestSuiteRun extension should have sequence")
	}

	if len(testSuiteRunType.ComplexContent.Extension.Attributes) != 1 {
		t.Errorf("expected 1 attribute in TestSuiteRun extension, got %d", len(testSuiteRunType.ComplexContent.Extension.Attributes))
	}
}

func TestParseInvalidXSD(t *testing.T) {
	invalidXSD := `<?xml version="1.0"?>
<invalid>not an XSD</invalid>`

	_, err := xsd.Parse(strings.NewReader(invalidXSD))
	if err == nil {
		t.Error("expected error parsing invalid XSD, got nil")
	}
}

func TestParseEmptySchema(t *testing.T) {
	emptySchema := `<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
</xs:schema>`

	schema, err := xsd.Parse(strings.NewReader(emptySchema))
	if err != nil {
		t.Fatalf("failed to parse empty schema: %v", err)
	}

	if len(schema.Elements) != 0 {
		t.Errorf("expected 0 elements in empty schema, got %d", len(schema.Elements))
	}

	if len(schema.ComplexTypes) != 0 {
		t.Errorf("expected 0 complex types in empty schema, got %d", len(schema.ComplexTypes))
	}

	if len(schema.SimpleTypes) != 0 {
		t.Errorf("expected 0 simple types in empty schema, got %d", len(schema.SimpleTypes))
	}
}

func TestParseSchemaWithAnnotations(t *testing.T) {
	schemaWithAnnotations := `<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
	<xs:element name="test">
		<xs:annotation>
			<xs:documentation>This is a test element</xs:documentation>
		</xs:annotation>
		<xs:complexType>
			<xs:sequence>
				<xs:element name="value" type="xs:string"/>
			</xs:sequence>
		</xs:complexType>
	</xs:element>
</xs:schema>`

	schema, err := xsd.Parse(strings.NewReader(schemaWithAnnotations))
	if err != nil {
		t.Fatalf("failed to parse schema with annotations: %v", err)
	}

	if len(schema.Elements) != 1 {
		t.Fatalf("expected 1 element, got %d", len(schema.Elements))
	}

	element := schema.Elements[0]
	if element.Name != "test" {
		t.Errorf("expected element name 'test', got %q", element.Name)
	}

	if element.Annotation == nil {
		t.Error("expected element to have annotation")
	}

	if len(element.Annotation.Documentation) != 1 {
		t.Errorf("expected 1 documentation, got %d", len(element.Annotation.Documentation))
	}

	if element.Annotation.Documentation[0].Content != "This is a test element" {
		t.Errorf("expected documentation content 'This is a test element', got %q", element.Annotation.Documentation[0].Content)
	}
}

package soap_test

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/way-platform/soap-go"
)

// ExampleEnvelope_basic demonstrates creating a basic SOAP envelope with just a body.
func ExampleEnvelope_basic() {
	// Create a simple SOAP envelope using the new API
	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<GetWeather><city>London</city></GetWeather>`)))
	if err != nil {
		log.Fatal(err)
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Body><GetWeather><city>London</city></GetWeather></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleEnvelope_withHeader demonstrates creating a SOAP envelope with headers.
func ExampleEnvelope_withHeader() {
	// Create header entries
	mustUnderstand := true
	authHeader := soap.HeaderEntry{
		XMLName:        xml.Name{Local: "Authentication", Space: "http://example.com/auth"},
		MustUnderstand: &mustUnderstand,
		Actor:          "http://example.com/gateway",
		Content:        []byte(`<token>abc123xyz</token><user>john.doe</user>`),
	}
	transactionHeader := soap.HeaderEntry{
		XMLName: xml.Name{Local: "Transaction", Space: "http://example.com/tx"},
		Content: []byte(`<id>tx-456</id>`),
	}
	// Create envelope with headers using the new API
	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<GetUserProfile><userId>12345</userId></GetUserProfile>`)))
	if err != nil {
		log.Fatal(err)
	}

	// Add headers manually for advanced use cases
	envelope.Header = &soap.Header{
		XMLName: xml.Name{Local: "soapenv:Header"},
		Entries: []soap.HeaderEntry{authHeader, transactionHeader},
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Header>
	//     <Authentication xmlns="http://example.com/auth" mustUnderstand="true" actor="http://example.com/gateway"><token>abc123xyz</token><user>john.doe</user></Authentication>
	//     <Transaction xmlns="http://example.com/tx"><id>tx-456</id></Transaction>
	//   </soapenv:Header>
	//   <soapenv:Body><GetUserProfile><userId>12345</userId></GetUserProfile></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleEnvelope_withEncodingStyle demonstrates setting encoding style.
func ExampleEnvelope_withEncodingStyle() {
	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<GetStockPrice><symbol>AAPL</symbol></GetStockPrice>`)))
	if err != nil {
		log.Fatal(err)
	}

	// Set encoding style for advanced use cases
	envelope.EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Body><GetStockPrice><symbol>AAPL</symbol></GetStockPrice></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleEnvelope_realWorld demonstrates creating a SOAP envelope for real-world use.
func ExampleEnvelope_realWorld() {
	// Create a SOAP envelope for a weather service request using the new API
	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<GetTemperature xmlns="http://weather.example.com/"><city>Paris</city></GetTemperature>`)))
	if err != nil {
		log.Fatal(err)
	}

	// Marshal to XML (this is what gets sent to SOAP services)
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Body><GetTemperature xmlns="http://weather.example.com/"><city>Paris</city></GetTemperature></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleFault demonstrates creating and handling SOAP faults.
func ExampleFault() {
	// Create a fault
	fault := soap.Fault{
		FaultCode:   "Client",
		FaultString: "Invalid authentication credentials",
		FaultActor:  "http://example.com/auth-service",
		Detail: &soap.Detail{
			Content: []byte(`<error><code>AUTH001</code><message>Token expired</message></error>`),
		},
	}
	faultXMLData, _ := xml.Marshal(fault)
	envelope, err := soap.NewEnvelope(soap.WithBody(faultXMLData))
	if err != nil {
		log.Fatal(err)
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Body><Fault><faultcode>Client</faultcode><faultstring>Invalid authentication credentials</faultstring><faultactor>http://example.com/auth-service</faultactor><detail><error><code>AUTH001</code><message>Token expired</message></error></detail></Fault></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleEnvelope_extensibility demonstrates using custom attributes for extensibility.
func ExampleEnvelope_extensibility() {
	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<ProcessOrder><orderId>12345</orderId></ProcessOrder>`)))
	if err != nil {
		log.Fatal(err)
	}

	// Add custom attributes to the body for extensibility
	envelope.Body.Attrs = []xml.Attr{
		{Name: xml.Name{Local: "priority"}, Value: "high"},
	}

	// Add custom attributes to the envelope for extensibility
	envelope.Attrs = append(envelope.Attrs, []xml.Attr{
		{Name: xml.Name{Local: "version"}, Value: "1.2"},
		{Name: xml.Name{Local: "trace", Space: "http://example.com/trace"}, Value: "enabled"},
	}...)

	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" version="1.2" xmlns:trace="http://example.com/trace" trace:trace="enabled">
	//   <soapenv:Body priority="high"><ProcessOrder><orderId>12345</orderId></ProcessOrder></soapenv:Body>
	// </soapenv:Envelope>
}

// ExampleHeaderEntry_mustUnderstand demonstrates the mustUnderstand attribute usage.
func ExampleHeaderEntry_mustUnderstand() {
	// Header that MUST be understood by the receiver
	mustUnderstand := true
	criticalHeader := soap.HeaderEntry{
		XMLName:        xml.Name{Local: "Security", Space: "http://example.com/security"},
		MustUnderstand: &mustUnderstand,
		Content:        []byte(`<signature>digital_signature_here</signature>`),
	}

	// Header that's optional
	optionalHeader := soap.HeaderEntry{
		XMLName: xml.Name{Local: "Metadata", Space: "http://example.com/meta"},
		Content: []byte(`<version>2.0</version>`),
		// MustUnderstand is nil, so it's optional
	}

	envelope, err := soap.NewEnvelope(soap.WithBody([]byte(`<SecureOperation><data>sensitive</data></SecureOperation>`)))
	if err != nil {
		log.Fatal(err)
	}

	// Add headers manually for advanced use cases
	envelope.Header = &soap.Header{
		XMLName: xml.Name{Local: "soapenv:Header"},
		Entries: []soap.HeaderEntry{criticalHeader, optionalHeader},
	}

	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	//   <soapenv:Header>
	//     <Security xmlns="http://example.com/security" mustUnderstand="true"><signature>digital_signature_here</signature></Security>
	//     <Metadata xmlns="http://example.com/meta"><version>2.0</version></Metadata>
	//   </soapenv:Header>
	//   <soapenv:Body><SecureOperation><data>sensitive</data></SecureOperation></soapenv:Body>
	// </soapenv:Envelope>
}

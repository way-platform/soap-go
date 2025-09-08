package soap_test

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/way-platform/soap-go"
)

// ExampleEnvelope_basic demonstrates creating a basic SOAP envelope with just a body.
func ExampleEnvelope_basic() {
	// Create a simple SOAP envelope
	envelope := &soap.Envelope{
		XMLName: xml.Name{Space: soap.Namespace, Local: "Envelope"},
		Body: soap.Body{
			Content: []byte(`<GetWeather><city>London</city></GetWeather>`),
		},
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <Envelope>
	//   <Body><GetWeather><city>London</city></GetWeather></Body>
	// </Envelope>
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
	// Create envelope with headers
	envelope := &soap.Envelope{
		XMLName: xml.Name{Space: soap.Namespace, Local: "Envelope"},
		Header: &soap.Header{
			Entries: []soap.HeaderEntry{authHeader, transactionHeader},
		},
		Body: soap.Body{
			Content: []byte(`<GetUserProfile><userId>12345</userId></GetUserProfile>`),
		},
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <Envelope>
	//   <Header>
	//     <Authentication xmlns="http://example.com/auth" mustUnderstand="true" actor="http://example.com/gateway"><token>abc123xyz</token><user>john.doe</user></Authentication>
	//     <Transaction xmlns="http://example.com/tx"><id>tx-456</id></Transaction>
	//   </Header>
	//   <Body><GetUserProfile><userId>12345</userId></GetUserProfile></Body>
	// </Envelope>
}

// ExampleEnvelope_withEncodingStyle demonstrates setting encoding style.
func ExampleEnvelope_withEncodingStyle() {
	envelope := &soap.Envelope{
		XMLName:       xml.Name{Space: soap.Namespace, Local: "Envelope"},
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: soap.Body{
			Content: []byte(`<GetStockPrice><symbol>AAPL</symbol></GetStockPrice>`),
		},
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <Envelope encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	//   <Body><GetStockPrice><symbol>AAPL</symbol></GetStockPrice></Body>
	// </Envelope>
}

// ExampleEnvelope_realWorld demonstrates creating a SOAP envelope for real-world use.
func ExampleEnvelope_realWorld() {
	// Create a SOAP envelope for a weather service request
	envelope := &soap.Envelope{
		XMLName: xml.Name{Space: soap.Namespace, Local: "Envelope"},
		Body: soap.Body{
			Content: []byte(`<GetTemperature xmlns="http://weather.example.com/"><city>Paris</city></GetTemperature>`),
		},
	}

	// Marshal to XML (this is what gets sent to SOAP services)
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <Envelope>
	//   <Body><GetTemperature xmlns="http://weather.example.com/"><city>Paris</city></GetTemperature></Body>
	// </Envelope>
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
	envelope := &soap.Envelope{
		XMLName: xml.Name{Space: soap.Namespace, Local: "Envelope"},
		Body: soap.Body{
			Content: faultXMLData,
		},
	}
	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	// Output: <Envelope>
	//   <Body><Fault><faultcode>Client</faultcode><faultstring>Invalid authentication credentials</faultstring><faultactor>http://example.com/auth-service</faultactor><detail><error><code>AUTH001</code><message>Token expired</message></error></detail></Fault></Body>
	// </Envelope>
}

// ExampleEnvelope_extensibility demonstrates using custom attributes for extensibility.
func ExampleEnvelope_extensibility() {
	envelope := &soap.Envelope{
		XMLName: xml.Name{Space: soap.Namespace, Local: "Envelope"},
		Body: soap.Body{
			Content: []byte(`<ProcessOrder><orderId>12345</orderId></ProcessOrder>`),
			Attrs: []xml.Attr{
				{Name: xml.Name{Local: "priority"}, Value: "high"},
			},
		},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "version"}, Value: "1.2"},
			{Name: xml.Name{Local: "trace", Space: "http://example.com/trace"}, Value: "enabled"},
		},
	}

	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <Envelope version="1.2" xmlns:trace="http://example.com/trace" trace:trace="enabled">
	//   <Body priority="high"><ProcessOrder><orderId>12345</orderId></ProcessOrder></Body>
	// </Envelope>
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

	envelope := &soap.Envelope{
		XMLName: xml.Name{Local: "soapenv:Envelope"},
		Header: &soap.Header{
			XMLName: xml.Name{Local: "soapenv:Header"},
			Entries: []soap.HeaderEntry{criticalHeader, optionalHeader},
		},
		Body: soap.Body{
			XMLName: xml.Name{Local: "soapenv:Body"},
			Content: []byte(`<SecureOperation><data>sensitive</data></SecureOperation>`),
		},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "xmlns:soapenv"}, Value: soap.Namespace},
		},
	}

	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(xmlData))
	// Output: <Envelope>
	//   <Header>
	//     <Security xmlns="http://example.com/security" mustUnderstand="true"><signature>digital_signature_here</signature></Security>
	//     <Metadata xmlns="http://example.com/meta"><version>2.0</version></Metadata>
	//   </Header>
	//   <Body><SecureOperation><data>sensitive</data></SecureOperation></Body>
	// </Envelope>
}

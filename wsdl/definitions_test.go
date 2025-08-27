package wsdl_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

func TestUnmarshalGlobalWeather(t *testing.T) {
	// Read the WSDL file
	wsdlBytes, err := os.ReadFile("../examples/GlobalWeather.wsdl")
	if err != nil {
		t.Fatalf("reading GlobalWeather.wsdl should not fail: %v", err)
	}

	// Unmarshal the WSDL content into our Go struct
	var defs wsdl.Definitions
	err = xml.Unmarshal(wsdlBytes, &defs)
	if err != nil {
		t.Fatalf("unmarshalling WSDL should not fail: %v", err)
	}

	// Construct the expected Go struct representing the GlobalWeather.wsdl content
	expected := wsdl.Definitions{
		XMLName: xml.Name{
			Space: "http://schemas.xmlsoap.org/wsdl/",
			Local: "definitions",
		},
		TargetNamespace: "http://www.webserviceX.NET",
		Types: &wsdl.Types{
			Schemas: []xsd.Schema{
				{},
			},
		},
		Messages: []wsdl.Message{
			{Name: "GetWeatherSoapIn", Parts: []wsdl.Part{{Name: "parameters", Element: "tns:GetWeather"}}},
			{Name: "GetWeatherSoapOut", Parts: []wsdl.Part{{Name: "parameters", Element: "tns:GetWeatherResponse"}}},
			{Name: "GetCitiesByCountrySoapIn", Parts: []wsdl.Part{{Name: "parameters", Element: "tns:GetCitiesByCountry"}}},
			{Name: "GetCitiesByCountrySoapOut", Parts: []wsdl.Part{{Name: "parameters", Element: "tns:GetCitiesByCountryResponse"}}},
			{Name: "GetWeatherHttpGetIn", Parts: []wsdl.Part{{Name: "CityName", Type: "s:string"}, {Name: "CountryName", Type: "s:string"}}},
			{Name: "GetWeatherHttpGetOut", Parts: []wsdl.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetCitiesByCountryHttpGetIn", Parts: []wsdl.Part{{Name: "CountryName", Type: "s:string"}}},
			{Name: "GetCitiesByCountryHttpGetOut", Parts: []wsdl.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetWeatherHttpPostIn", Parts: []wsdl.Part{{Name: "CityName", Type: "s:string"}, {Name: "CountryName", Type: "s:string"}}},
			{Name: "GetWeatherHttpPostOut", Parts: []wsdl.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetCitiesByCountryHttpPostIn", Parts: []wsdl.Part{{Name: "CountryName", Type: "s:string"}}},
			{Name: "GetCitiesByCountryHttpPostOut", Parts: []wsdl.Part{{Name: "Body", Element: "tns:string"}}},
		},
		PortType: []wsdl.PortType{
			{
				Name: "GlobalWeatherSoap",
				Operations: []wsdl.Operation{
					{Name: "GetWeather", Documentation: "\n                Get weather report for all major cities around the world.\n            ", Input: &wsdl.Input{Message: "tns:GetWeatherSoapIn"}, Output: &wsdl.Output{Message: "tns:GetWeatherSoapOut"}},
					{Name: "GetCitiesByCountry", Documentation: "Get all major\n                cities by country name(full / part).", Input: &wsdl.Input{Message: "tns:GetCitiesByCountrySoapIn"}, Output: &wsdl.Output{Message: "tns:GetCitiesByCountrySoapOut"}},
				},
			},
			{
				Name: "GlobalWeatherHttpGet",
				Operations: []wsdl.Operation{
					{Name: "GetWeather", Documentation: "\n                Get weather report for all major cities around the world.\n            ", Input: &wsdl.Input{Message: "tns:GetWeatherHttpGetIn"}, Output: &wsdl.Output{Message: "tns:GetWeatherHttpGetOut"}},
					{Name: "GetCitiesByCountry", Documentation: "Get all major\n                cities by country name(full / part).", Input: &wsdl.Input{Message: "tns:GetCitiesByCountryHttpGetIn"}, Output: &wsdl.Output{Message: "tns:GetCitiesByCountryHttpGetOut"}},
				},
			},
			{
				Name: "GlobalWeatherHttpPost",
				Operations: []wsdl.Operation{
					{Name: "GetWeather", Documentation: "\n                Get weather report for all major cities around the world.\n            ", Input: &wsdl.Input{Message: "tns:GetWeatherHttpPostIn"}, Output: &wsdl.Output{Message: "tns:GetWeatherHttpPostOut"}},
					{Name: "GetCitiesByCountry", Documentation: "Get all major\n                cities by country name(full / part).", Input: &wsdl.Input{Message: "tns:GetCitiesByCountryHttpPostIn"}, Output: &wsdl.Output{Message: "tns:GetCitiesByCountryHttpPostOut"}},
				},
			},
		},
		Binding: []wsdl.Binding{
			{
				Name:          "GlobalWeatherSoap",
				Type:          "tns:GlobalWeatherSoap",
				SOAP11Binding: &wsdl.SOAPBinding{Transport: "http://schemas.xmlsoap.org/soap/http"},
				BindingOperations: []wsdl.BindingOperation{
					{
						Name:            "GetWeather",
						SOAP11Operation: &wsdl.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetWeather", Style: "document"},
						Input:           &wsdl.BindingBody{SOAP11Body: &wsdl.SOAPBody{Use: "literal"}},
						Output:          &wsdl.BindingBody{SOAP11Body: &wsdl.SOAPBody{Use: "literal"}},
					},
					{
						Name:            "GetCitiesByCountry",
						SOAP11Operation: &wsdl.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetCitiesByCountry", Style: "document"},
						Input:           &wsdl.BindingBody{SOAP11Body: &wsdl.SOAPBody{Use: "literal"}},
						Output:          &wsdl.BindingBody{SOAP11Body: &wsdl.SOAPBody{Use: "literal"}},
					},
				},
			},
			{
				Name:          "GlobalWeatherSoap12",
				Type:          "tns:GlobalWeatherSoap",
				SOAP12Binding: &wsdl.SOAPBinding{Transport: "http://schemas.xmlsoap.org/soap/http"},
				BindingOperations: []wsdl.BindingOperation{
					{
						Name:            "GetWeather",
						SOAP12Operation: &wsdl.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetWeather", Style: "document"},
						Input:           &wsdl.BindingBody{SOAP12Body: &wsdl.SOAPBody{Use: "literal"}},
						Output:          &wsdl.BindingBody{SOAP12Body: &wsdl.SOAPBody{Use: "literal"}},
					},
					{
						Name:            "GetCitiesByCountry",
						SOAP12Operation: &wsdl.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetCitiesByCountry", Style: "document"},
						Input:           &wsdl.BindingBody{SOAP12Body: &wsdl.SOAPBody{Use: "literal"}},
						Output:          &wsdl.BindingBody{SOAP12Body: &wsdl.SOAPBody{Use: "literal"}},
					},
				},
			},
			{
				Name:        "GlobalWeatherHttpGet",
				Type:        "tns:GlobalWeatherHttpGet",
				HTTPBinding: &wsdl.HTTPBinding{Verb: "GET"},
				BindingOperations: []wsdl.BindingOperation{
					{
						Name:          "GetWeather",
						HTTPOperation: &wsdl.HTTPOperation{Location: "/GetWeather"},
						Input:         &wsdl.BindingBody{URLEncoded: &wsdl.URLEncoded{}},
						Output:        &wsdl.BindingBody{MIMEXML: &wsdl.MIMEXML{Part: "Body"}},
					},
					{
						Name:          "GetCitiesByCountry",
						HTTPOperation: &wsdl.HTTPOperation{Location: "/GetCitiesByCountry"},
						Input:         &wsdl.BindingBody{URLEncoded: &wsdl.URLEncoded{}},
						Output:        &wsdl.BindingBody{MIMEXML: &wsdl.MIMEXML{Part: "Body"}},
					},
				},
			},
			{
				Name:        "GlobalWeatherHttpPost",
				Type:        "tns:GlobalWeatherHttpPost",
				HTTPBinding: &wsdl.HTTPBinding{Verb: "POST"},
				BindingOperations: []wsdl.BindingOperation{
					{
						Name:          "GetWeather",
						HTTPOperation: &wsdl.HTTPOperation{Location: "/GetWeather"},
						Input:         &wsdl.BindingBody{MIMEContent: []*wsdl.MIMEContent{{Type: "application/x-www-form-urlencoded"}}},
						Output:        &wsdl.BindingBody{MIMEXML: &wsdl.MIMEXML{Part: "Body"}},
					},
					{
						Name:          "GetCitiesByCountry",
						HTTPOperation: &wsdl.HTTPOperation{Location: "/GetCitiesByCountry"},
						Input:         &wsdl.BindingBody{MIMEContent: []*wsdl.MIMEContent{{Type: "application/x-www-form-urlencoded"}}},
						Output:        &wsdl.BindingBody{MIMEXML: &wsdl.MIMEXML{Part: "Body"}},
					},
				},
			},
		},
		Service: []wsdl.Service{
			{
				Name: "GlobalWeather",
				Ports: []wsdl.Port{
					{Name: "GlobalWeatherSoap", Binding: "tns:GlobalWeatherSoap", SOAP11Address: &wsdl.SOAPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherSoap12", Binding: "tns:GlobalWeatherSoap12", SOAP12Address: &wsdl.SOAPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherHttpGet", Binding: "tns:GlobalWeatherHttpGet", HTTPAddress: &wsdl.HTTPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherHttpPost", Binding: "tns:GlobalWeatherHttpPost", HTTPAddress: &wsdl.HTTPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
				},
			},
		},
	}

	// Verify that the schema was parsed correctly, but don't check exact content
	// as it's complex and not the focus of this test.
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

	if len(schema.Elements) != 5 {
		t.Errorf("expected 5 schema elements, got %d", len(schema.Elements))
	}

	// Replace with a simple empty schema for the comparison test
	defs.Types.Schemas[0] = xsd.Schema{}

	if diff := cmp.Diff(expected, defs); diff != "" {
		t.Errorf("Definitions mismatch (-want +got):\n%s", diff)
	}
}

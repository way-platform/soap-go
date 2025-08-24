package wsdl11_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/way-platform/soap-go/wsdl11"
	"github.com/way-platform/soap-go/xsd10"
)

func TestUnmarshalGlobalWeather(t *testing.T) {
	// Read the WSDL file
	wsdlBytes, err := os.ReadFile("../testdata/GlobalWeather.wsdl")
	if err != nil {
		t.Fatalf("reading GlobalWeather.wsdl should not fail: %v", err)
	}

	// Unmarshal the WSDL content into our Go struct
	var defs wsdl11.Definitions
	err = xml.Unmarshal(wsdlBytes, &defs)
	if err != nil {
		t.Fatalf("unmarshalling WSDL should not fail: %v", err)
	}

	// Construct the expected Go struct representing the GlobalWeather.wsdl content
	expected := wsdl11.Definitions{
		XMLName: xml.Name{
			Space: "http://schemas.xmlsoap.org/wsdl/",
			Local: "definitions",
		},
		TargetNamespace: "http://www.webserviceX.NET",
		Types: &wsdl11.Types{
			Schemas: []xsd10.Schema{
				{},
			},
		},
		Messages: []wsdl11.Message{
			{Name: "GetWeatherSoapIn", Parts: []wsdl11.Part{{Name: "parameters", Element: "tns:GetWeather"}}},
			{Name: "GetWeatherSoapOut", Parts: []wsdl11.Part{{Name: "parameters", Element: "tns:GetWeatherResponse"}}},
			{Name: "GetCitiesByCountrySoapIn", Parts: []wsdl11.Part{{Name: "parameters", Element: "tns:GetCitiesByCountry"}}},
			{Name: "GetCitiesByCountrySoapOut", Parts: []wsdl11.Part{{Name: "parameters", Element: "tns:GetCitiesByCountryResponse"}}},
			{Name: "GetWeatherHttpGetIn", Parts: []wsdl11.Part{{Name: "CityName", Type: "s:string"}, {Name: "CountryName", Type: "s:string"}}},
			{Name: "GetWeatherHttpGetOut", Parts: []wsdl11.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetCitiesByCountryHttpGetIn", Parts: []wsdl11.Part{{Name: "CountryName", Type: "s:string"}}},
			{Name: "GetCitiesByCountryHttpGetOut", Parts: []wsdl11.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetWeatherHttpPostIn", Parts: []wsdl11.Part{{Name: "CityName", Type: "s:string"}, {Name: "CountryName", Type: "s:string"}}},
			{Name: "GetWeatherHttpPostOut", Parts: []wsdl11.Part{{Name: "Body", Element: "tns:string"}}},
			{Name: "GetCitiesByCountryHttpPostIn", Parts: []wsdl11.Part{{Name: "CountryName", Type: "s:string"}}},
			{Name: "GetCitiesByCountryHttpPostOut", Parts: []wsdl11.Part{{Name: "Body", Element: "tns:string"}}},
		},
		PortType: []wsdl11.PortType{
			{
				Name: "GlobalWeatherSoap",
				Operations: []wsdl11.Operation{
					{Name: "GetWeather", Input: &wsdl11.Input{Message: "tns:GetWeatherSoapIn"}, Output: &wsdl11.Output{Message: "tns:GetWeatherSoapOut"}},
					{Name: "GetCitiesByCountry", Input: &wsdl11.Input{Message: "tns:GetCitiesByCountrySoapIn"}, Output: &wsdl11.Output{Message: "tns:GetCitiesByCountrySoapOut"}},
				},
			},
			{
				Name: "GlobalWeatherHttpGet",
				Operations: []wsdl11.Operation{
					{Name: "GetWeather", Input: &wsdl11.Input{Message: "tns:GetWeatherHttpGetIn"}, Output: &wsdl11.Output{Message: "tns:GetWeatherHttpGetOut"}},
					{Name: "GetCitiesByCountry", Input: &wsdl11.Input{Message: "tns:GetCitiesByCountryHttpGetIn"}, Output: &wsdl11.Output{Message: "tns:GetCitiesByCountryHttpGetOut"}},
				},
			},
			{
				Name: "GlobalWeatherHttpPost",
				Operations: []wsdl11.Operation{
					{Name: "GetWeather", Input: &wsdl11.Input{Message: "tns:GetWeatherHttpPostIn"}, Output: &wsdl11.Output{Message: "tns:GetWeatherHttpPostOut"}},
					{Name: "GetCitiesByCountry", Input: &wsdl11.Input{Message: "tns:GetCitiesByCountryHttpPostIn"}, Output: &wsdl11.Output{Message: "tns:GetCitiesByCountryHttpPostOut"}},
				},
			},
		},
		Binding: []wsdl11.Binding{
			{
				Name:          "GlobalWeatherSoap",
				Type:          "tns:GlobalWeatherSoap",
				SOAP11Binding: &wsdl11.SOAPBinding{Transport: "http://schemas.xmlsoap.org/soap/http"},
				BindingOperations: []wsdl11.BindingOperation{
					{
						Name:            "GetWeather",
						SOAP11Operation: &wsdl11.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetWeather", Style: "document"},
						Input:           &wsdl11.BindingBody{SOAP11Body: &wsdl11.SOAPBody{Use: "literal"}},
						Output:          &wsdl11.BindingBody{SOAP11Body: &wsdl11.SOAPBody{Use: "literal"}},
					},
					{
						Name:            "GetCitiesByCountry",
						SOAP11Operation: &wsdl11.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetCitiesByCountry", Style: "document"},
						Input:           &wsdl11.BindingBody{SOAP11Body: &wsdl11.SOAPBody{Use: "literal"}},
						Output:          &wsdl11.BindingBody{SOAP11Body: &wsdl11.SOAPBody{Use: "literal"}},
					},
				},
			},
			{
				Name:          "GlobalWeatherSoap12",
				Type:          "tns:GlobalWeatherSoap",
				SOAP12Binding: &wsdl11.SOAPBinding{Transport: "http://schemas.xmlsoap.org/soap/http"},
				BindingOperations: []wsdl11.BindingOperation{
					{
						Name:            "GetWeather",
						SOAP12Operation: &wsdl11.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetWeather", Style: "document"},
						Input:           &wsdl11.BindingBody{SOAP12Body: &wsdl11.SOAPBody{Use: "literal"}},
						Output:          &wsdl11.BindingBody{SOAP12Body: &wsdl11.SOAPBody{Use: "literal"}},
					},
					{
						Name:            "GetCitiesByCountry",
						SOAP12Operation: &wsdl11.SOAPOperation{SOAPAction: "http://www.webserviceX.NET/GetCitiesByCountry", Style: "document"},
						Input:           &wsdl11.BindingBody{SOAP12Body: &wsdl11.SOAPBody{Use: "literal"}},
						Output:          &wsdl11.BindingBody{SOAP12Body: &wsdl11.SOAPBody{Use: "literal"}},
					},
				},
			},
			{
				Name:        "GlobalWeatherHttpGet",
				Type:        "tns:GlobalWeatherHttpGet",
				HTTPBinding: &wsdl11.HTTPBinding{Verb: "GET"},
				BindingOperations: []wsdl11.BindingOperation{
					{
						Name:          "GetWeather",
						HTTPOperation: &wsdl11.HTTPOperation{Location: "/GetWeather"},
						Input:         &wsdl11.BindingBody{URLEncoded: &wsdl11.URLEncoded{}},
						Output:        &wsdl11.BindingBody{MIMEXML: &wsdl11.MIMEXML{Part: "Body"}},
					},
					{
						Name:          "GetCitiesByCountry",
						HTTPOperation: &wsdl11.HTTPOperation{Location: "/GetCitiesByCountry"},
						Input:         &wsdl11.BindingBody{URLEncoded: &wsdl11.URLEncoded{}},
						Output:        &wsdl11.BindingBody{MIMEXML: &wsdl11.MIMEXML{Part: "Body"}},
					},
				},
			},
			{
				Name:        "GlobalWeatherHttpPost",
				Type:        "tns:GlobalWeatherHttpPost",
				HTTPBinding: &wsdl11.HTTPBinding{Verb: "POST"},
				BindingOperations: []wsdl11.BindingOperation{
					{
						Name:          "GetWeather",
						HTTPOperation: &wsdl11.HTTPOperation{Location: "/GetWeather"},
						Input:         &wsdl11.BindingBody{MIMEContent: []*wsdl11.MIMEContent{{Type: "application/x-www-form-urlencoded"}}},
						Output:        &wsdl11.BindingBody{MIMEXML: &wsdl11.MIMEXML{Part: "Body"}},
					},
					{
						Name:          "GetCitiesByCountry",
						HTTPOperation: &wsdl11.HTTPOperation{Location: "/GetCitiesByCountry"},
						Input:         &wsdl11.BindingBody{MIMEContent: []*wsdl11.MIMEContent{{Type: "application/x-www-form-urlencoded"}}},
						Output:        &wsdl11.BindingBody{MIMEXML: &wsdl11.MIMEXML{Part: "Body"}},
					},
				},
			},
		},
		Service: []wsdl11.Service{
			{
				Name: "GlobalWeather",
				Ports: []wsdl11.Port{
					{Name: "GlobalWeatherSoap", Binding: "tns:GlobalWeatherSoap", SOAP11Address: &wsdl11.SOAPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherSoap12", Binding: "tns:GlobalWeatherSoap12", SOAP12Address: &wsdl11.SOAPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherHttpGet", Binding: "tns:GlobalWeatherHttpGet", HTTPAddress: &wsdl11.HTTPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
					{Name: "GlobalWeatherHttpPost", Binding: "tns:GlobalWeatherHttpPost", HTTPAddress: &wsdl11.HTTPAddress{Location: "http://www.webservicex.com/globalweather.asmx"}},
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
	defs.Types.Schemas[0] = xsd10.Schema{}

	if diff := cmp.Diff(expected, defs); diff != "" {
		t.Errorf("Definitions mismatch (-want +got):\n%s", diff)
	}
}

package globalweather

import (
	"encoding/xml"
)

// GetWeatherWrapper represents the GetWeather element
type GetWeatherWrapper struct {
	XMLName     xml.Name `xml:"http://www.webserviceX.NET GetWeather"`
	CityName    *string  `xml:"CityName,omitempty"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetWeatherResponseWrapper represents the GetWeatherResponse element
type GetWeatherResponseWrapper struct {
	XMLName          xml.Name `xml:"http://www.webserviceX.NET GetWeatherResponse"`
	GetWeatherResult *string  `xml:"GetWeatherResult,omitempty"`
}

// GetCitiesByCountryWrapper represents the GetCitiesByCountry element
type GetCitiesByCountryWrapper struct {
	XMLName     xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountry"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetCitiesByCountryResponseWrapper represents the GetCitiesByCountryResponse element
type GetCitiesByCountryResponseWrapper struct {
	XMLName                  xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountryResponse"`
	GetCitiesByCountryResult *string  `xml:"GetCitiesByCountryResult,omitempty"`
}

// StringWrapper represents the string element
type StringWrapper struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET string"`
	Value   string   `xml:",chardata"`
}

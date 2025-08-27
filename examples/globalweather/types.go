package globalweather

import (
	"encoding/xml"
)

// GetWeather represents the GetWeather element
type GetWeather struct {
	XMLName     xml.Name `xml:"http://www.webserviceX.NET GetWeather"`
	CityName    *string  `xml:"CityName,omitempty"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetWeatherResponse represents the GetWeatherResponse element
type GetWeatherResponse struct {
	XMLName          xml.Name `xml:"http://www.webserviceX.NET GetWeatherResponse"`
	GetWeatherResult *string  `xml:"GetWeatherResult,omitempty"`
}

// GetCitiesByCountry represents the GetCitiesByCountry element
type GetCitiesByCountry struct {
	XMLName     xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountry"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetCitiesByCountryResponse represents the GetCitiesByCountryResponse element
type GetCitiesByCountryResponse struct {
	XMLName                  xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountryResponse"`
	GetCitiesByCountryResult *string  `xml:"GetCitiesByCountryResult,omitempty"`
}

// String represents the string element
type String struct {
	XMLName xml.Name `xml:"http://www.webserviceX.NET string"`
	Value   string   `xml:",chardata"`
}

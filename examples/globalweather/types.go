package globalweather

import (
	"encoding/xml"
)

// GetWeather represents the GetWeather element
type GetWeather struct {
	XMLName     xml.Name `xml:"GetWeather"`
	CityName    *string  `xml:"CityName,omitempty"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetWeatherResponse represents the GetWeatherResponse element
type GetWeatherResponse struct {
	XMLName          xml.Name `xml:"GetWeatherResponse"`
	GetWeatherResult *string  `xml:"GetWeatherResult,omitempty"`
}

// GetCitiesByCountry represents the GetCitiesByCountry element
type GetCitiesByCountry struct {
	XMLName     xml.Name `xml:"GetCitiesByCountry"`
	CountryName *string  `xml:"CountryName,omitempty"`
}

// GetCitiesByCountryResponse represents the GetCitiesByCountryResponse element
type GetCitiesByCountryResponse struct {
	XMLName                  xml.Name `xml:"GetCitiesByCountryResponse"`
	GetCitiesByCountryResult *string  `xml:"GetCitiesByCountryResult,omitempty"`
}

// String represents the string element
type String struct {
	XMLName xml.Name `xml:"string"`
	Value   string   `xml:",chardata"`
}

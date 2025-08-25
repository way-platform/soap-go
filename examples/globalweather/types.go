package globalweather

// GetWeather represents the GetWeather element
type GetWeather struct {
	CityName    *string `xml:"CityName"`
	CountryName *string `xml:"CountryName"`
}

// GetWeatherResponse represents the GetWeatherResponse element
type GetWeatherResponse struct {
	GetWeatherResult *string `xml:"GetWeatherResult"`
}

// GetCitiesByCountry represents the GetCitiesByCountry element
type GetCitiesByCountry struct {
	CountryName *string `xml:"CountryName"`
}

// GetCitiesByCountryResponse represents the GetCitiesByCountryResponse element
type GetCitiesByCountryResponse struct {
	GetCitiesByCountryResult *string `xml:"GetCitiesByCountryResult"`
}

// String represents the string element
type String struct {
	Value string `xml:",chardata"`
}

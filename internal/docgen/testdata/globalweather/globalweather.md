# 

**Namespace:** `http://www.webserviceX.NET`

## GlobalWeather

### GetWeather

Get weather report for all major cities around the world.

#### Request

- **GetWeather** (complex)
  - **CityName** (s:string) (0..1)
  - **CountryName** (s:string) (0..1)

#### Response

- **GetWeatherResponse** (complex)
  - **GetWeatherResult** (s:string) (0..1)


### GetCitiesByCountry

Get all major cities by country name(full / part).

#### Request

- **GetCitiesByCountry** (complex)
  - **CountryName** (s:string) (0..1)

#### Response

- **GetCitiesByCountryResponse** (complex)
  - **GetCitiesByCountryResult** (s:string) (0..1)



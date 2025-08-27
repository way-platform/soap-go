# GlobalWeather API Documentation

**Namespace:** `http://www.webserviceX.NET`
**Endpoint:** `http://www.webservicex.com/globalweather.asmx`

## Available Operations

- **[GetWeather](#getweather)** - Get weather report for all major cities around the world.
- **[GetCitiesByCountry](#getcitiesbycountry)** - Get all major cities by country name(full / part).

## Operations

### GetWeather {#getweather}

> Get weather report for all major cities around the world.

**SOAP Action:** `http://www.webserviceX.NET/GetWeather`

#### Request

**Message:** `GetWeatherSoapIn`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetWeather | object | Yes | - |
| GetWeather.CityName | s:string | No | - |
| GetWeather.CountryName | s:string | No | - |

#### Response

**Message:** `GetWeatherSoapOut`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetWeatherResponse | object | Yes | - |
| GetWeatherResponse.GetWeatherResult | s:string | No | - |


### GetCitiesByCountry {#getcitiesbycountry}

> Get all major cities by country name(full / part).

**SOAP Action:** `http://www.webserviceX.NET/GetCitiesByCountry`

#### Request

**Message:** `GetCitiesByCountrySoapIn`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetCitiesByCountry | object | Yes | - |
| GetCitiesByCountry.CountryName | s:string | No | - |

#### Response

**Message:** `GetCitiesByCountrySoapOut`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetCitiesByCountryResponse | object | Yes | - |
| GetCitiesByCountryResponse.GetCitiesByCountryResult | s:string | No | - |



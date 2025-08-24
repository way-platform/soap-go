# Examples

This directory contains examples of using the `soap-gen-go` tool to generate Go types from WSDL files.

## Generated Packages

### NumberConversion

Generated from `../testdata/NumberConversion.wsdl`, this package provides Go types for the NumberConversion web service:

- `NumberToWords` - Request struct for converting numbers to words
- `NumberToWordsResponse` - Response struct containing the word representation
- `NumberToDollars` - Request struct for converting numbers to dollar amounts
- `NumberToDollarsResponse` - Response struct containing the dollar representation

### GlobalWeather

Generated from `../testdata/GlobalWeather.wsdl`, this package provides Go types for the GlobalWeather web service:

- `GetWeather` - Request struct for getting weather information
- `GetWeatherResponse` - Response struct containing weather data
- `GetCitiesByCountry` - Request struct for getting cities by country
- `GetCitiesByCountryResponse` - Response struct containing city list

## Usage

### Running Code Generation

To regenerate the packages, run:

```bash
go generate
```

This will execute the `//go:generate` directives in `generate.go` and recreate the packages.

### Using the Generated Types

Each package includes example functions and tests showing how to:

1. **Create request/response structs**
2. **Marshal to XML** for sending SOAP requests
3. **Unmarshal from XML** for processing SOAP responses

Example:

```go
// Create a request
request := numberconversion.NumberToWords{
    UbiNum: 42,
}

// Marshal to XML
xmlData, err := xml.Marshal(&request)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(xmlData))
// Output: <NumberToWords><ubiNum>42</ubiNum></NumberToWords>
```

### Running Tests

Each package includes comprehensive tests:

```bash
cd numberconversion && go test -v
cd globalweather && go test -v
```

## XML Tags

All generated structs include proper XML tags for marshaling/unmarshaling:

- Field names are mapped to the correct XML element names from the WSDL
- Types are correctly mapped from XSD types to Go types:
  - `xs:string` → `string`
  - `xs:unsignedLong` → `uint64`
  - `xs:decimal` → `float64`

## Next Steps

These generated types can be used to build complete SOAP clients by:

1. Creating SOAP envelopes with the generated request types
2. Sending HTTP requests to the web service endpoints
3. Parsing responses using the generated response types

See the individual package tests and examples for more detailed usage patterns.

# AGENTS.md

## Specs

When developing this SDK, use the WSDL and SOAP specifications in [docs](./docs):

- [SOAP 1.1](./docs/soap-1.1.md)
- [WSDL 1.1](./docs/wsdl-1.1.md)
- [XSD 1.0 (Structures)](./docs/xsd-1.0-structures.md)
- [XSD 1.0 (Data Types)](./docs/xsd-1.0-datatypes.md)

## Structure

- The project uses a [tools](./tools/) directory with a separate Go module containing tools for building, linting and generating code.

- The project uses [Mage](https://magefile.org) with build tasks declared in [magefile.go](./tools/magefile.go).

- The CLI in [cmd/soap](./cmd/soap) depends on an already released version of the SDK, to make installation as easy as possible.

## Developing

- Run tests with `./tools/mage test`

- Lint with `./tools/mage lint`

- Re-generate code with `./tools/mage generate`

- Leave all version control and git to the user/developer. If you see a build error related to having a git diff, this is normal.

## Testing

- When fixing issues in the code generator, add tests in [internal/soapgen/golden_test.go](./internal/soapgen/golden_test.go).

- When fixing issues in how the generated code parses raw XML, add tests in [internal/soapgen/comprehensive_parsing_test.go](./internal/soapgen/comprehensive_parsing_test.go) or [internal/soapgen/integration_test.go](./internal/soapgen/integration_test.go).

- Golden files are updated by running `go test` with the `-update` flag.

- Golden file tests prefixed with `proprietary_` may contain proprietary data and are not checked into git. When an issue is detected in proprietary test data, create test cases using non-proprietary data to isolate the issue.

## Principles

- The generated code should be as idiomatic as possible and leverage the latest available Go features.

- Work within the constraints of Go's XML parser. Find simple and effective solutions that align with how "encoding/xml" works.

- Avoid custom UnmarshalXML methods when possible. Prefer modelling the generated types to capture all the XML data through standard Unmarshalling.

- Keep the `wsdl` and `xsd` packages focused on describing the WSDL and XSD file formats. Keep them clean of business logic.

- For extreme cases of complex data structures that cannot be modelled through standard Unmarshalling, use the `RawXML` type.

- Avoid configuration options. Pick sane defaults that "just work". The SDK, generated code and tools should be as simple to use as possible.

- Public API surfaces should be small, simple and orthogonal. Less is more.

- Be permissive with XML namespaces. The generated code should be able to handle any namespace prefix and URI, even when the data returned from the server differs from the WSDL file.

- When in doubt about how Go handles XML - refer to [docs/go-encoding-xml.txt](./docs/go-encoding-xml.txt).

# AGENTS.md

## Specs

When developing this SDK, use the WSDL and SOAP specifications in [docs](./docs):

- [SOAP 1.1](./docs/soap-1.1.md)
- [WSDL 1.1](./docs/wsdl-1.1.md)
- [XSD 1.0 (Structures)](./docs/xsd-1.0-structures.md)
- [XSD 1.0 (Data Types)](./docs/xsd-1.0-datatypes.md)
- [Go encoding/xml reference](./docs/go-encoding-xml.txt)

## Architecture

### Module structure

Multi-module repo with three `go.mod` files:

- `/go.mod` — SDK root (minimal deps: `go-cmp`, `x/text`)
- `/cmd/soap/go.mod` — CLI (depends on _released_ SDK version, not local)
- `/tools/go.mod` — build tools (Mage, golangci-lint)

### Package dependency graph

```
cmd/soap/internal/{gen,doc,call}  (CLI commands)
         │
         ▼
    soap (root pkg)          ← public SDK: Client, Envelope, Error
         │
    ┌────┼────┐
    ▼    ▼    ▼
  wsdl  xsd  internal/codegen   ← parsing + codegen utilities
              │
              ▼
         internal/soapgen       ← WSDL→Go code generator
         internal/docgen        ← WSDL→Markdown doc generator
```

### Key packages

| Package                  | Role                                                                  |
| ------------------------ | --------------------------------------------------------------------- |
| `soap` (root)            | SOAP 1.1 client, envelope, fault, error types                         |
| `wsdl`                   | WSDL 1.1 XML unmarshaling — pure data, no logic                       |
| `xsd`                    | XSD 1.0 schema unmarshaling + type constants — pure data, no logic    |
| `internal/codegen`       | Generic Go file builder with import management (AST-based formatting) |
| `internal/soapgen`       | WSDL→Go code generator (~5k lines, golden-file tested)                |
| `internal/docgen`        | WSDL→Markdown documentation generator                                 |
| `cmd/soap/internal/gen`  | `soap gen` CLI command                                                |
| `cmd/soap/internal/doc`  | `soap doc` CLI command (with interactive TUI pager)                   |
| `cmd/soap/internal/call` | `soap call` CLI command                                               |

## Developing

- Build (full CI): `./tools/mage build`
- Test: `./tools/mage test`
- Lint: `./tools/mage lint`
- Generate: `./tools/mage generate`
- Integration tests: `./tools/mage integrationtest`
- Leave all version control and git to the user/developer. If you see a build error related to having a git diff, this is normal.

## Testing

- When fixing issues in the code generator, add tests in [internal/soapgen/golden_test.go](./internal/soapgen/golden_test.go).
- When fixing issues in how the generated code parses raw XML, add tests in [internal/soapgen/comprehensive_parsing_test.go](./internal/soapgen/comprehensive_parsing_test.go) or [internal/soapgen/integration_test.go](./internal/soapgen/integration_test.go).
- Golden files are updated by running `go test` with the `-update` flag.
- Golden file tests prefixed with `proprietary_` may contain proprietary data and are not checked into git. When an issue is detected in proprietary test data, create test cases using non-proprietary data to isolate the issue.

### Golden test structure

Each test case lives in `internal/soapgen/testdata/{name}/`:

- `definitions.wsdl` — input
- `types.go` — expected generated types (golden)
- `client.go` — expected generated client (golden, optional)
- `error.txt` — expected error message (for failure cases)

50+ test cases covering enums, complex types, inline types, element refs, collisions, content extensions, and edge cases.

## Design Patterns

### Functional options

Both `Client` and `Envelope` use functional options for configuration:

```go
// Client options apply at creation AND per-call (per-call overrides client defaults)
client, _ := soap.NewClient(soap.WithEndpoint(url), soap.WithMaxRetries(5))
env, _ := client.Call(ctx, action, envelope, soap.WithTimeout(30*time.Second))
```

### Raw XML for flexibility

`Body.Content`, `HeaderEntry.Content`, and `Detail.Content` use `[]byte` with `xml:",innerxml"` — enables a generic client that works with any SOAP payload without schema knowledge.

### Two-pass code generation

1. **Registration pass** — discover and register all types (inline enums, anonymous types, named types) without emitting code
2. **Generation pass** — emit Go code with full type reference information available

Registries prevent duplicates and resolve collisions:

- `TypeRegistry` — tracks generated Go type names
- `FieldRegistry` — detects field name collisions within structs
- `InlineEnumInfo` — deduplicates inline enums via hash signatures

### Unified error type

`soap.Error` combines HTTP status, raw response body, parsed SOAP envelope, and SOAP fault into one type. Use `errors.As(err, &soapErr)` to inspect.

### Retry with exponential backoff

Built into `Client.Call` — retries 5xx, 429, 420, and network errors with exponential backoff + jitter. Respects `Retry-After` headers. Context-aware.

## Principles

- The generated code should be as idiomatic as possible and leverage the latest available Go features.
- Work within the constraints of Go's XML parser. Find simple and effective solutions that align with how `encoding/xml` works.
- Avoid custom `UnmarshalXML` methods when possible. Prefer modelling the generated types to capture all the XML data through standard unmarshalling.
- Keep the `wsdl` and `xsd` packages focused on describing the WSDL and XSD file formats. Keep them clean of business logic.
- For extreme cases of complex data structures that cannot be modelled through standard unmarshalling, use the `RawXML` type.
- Avoid configuration options. Pick sane defaults that "just work". The SDK, generated code and tools should be as simple to use as possible.
- Public API surfaces should be small, simple and orthogonal. Less is more.
- Be permissive with XML namespaces. The generated code should be able to handle any namespace prefix and URI, even when the data returned from the server differs from the WSDL file.
- When in doubt about how Go handles XML — refer to [docs/go-encoding-xml.txt](./docs/go-encoding-xml.txt).

## Code Generation Conventions

### XSD → Go type mapping

| XSD                                         | Go                  | Notes                                        |
| ------------------------------------------- | ------------------- | -------------------------------------------- |
| `string`, `token`, `normalizedString`, etc. | `string`            |                                              |
| `boolean`                                   | `bool`              |                                              |
| `float`, `double`, `decimal`                | `float64`           | Unified for simplicity                       |
| `int`, `integer` / `long`                   | `int32` / `int64`   | Sized per spec                               |
| `unsignedInt` / `unsignedLong`              | `uint32` / `uint64` |                                              |
| `dateTime`, `date`, `time`                  | `time.Time`         |                                              |
| `duration`                                  | `string`            | ISO 8601, requires custom parsing            |
| `hexBinary`, `base64Binary`                 | `[]byte`            |                                              |
| `QName`                                     | `xml.Name`          |                                              |
| Unknown / `xs:any`                          | `[]byte` (RawXML)   | Fallback                                     |

### Generated struct conventions

- `XMLName xml.Name` with namespace-qualified tag: `xml:"http://ns Element"`
- Optional fields: `*T` pointer + `omitempty`
- Required fields: value type, no `omitempty`
- Attribute fields: `xml:"name,attr"`
- Attribute-element collision: element keeps name, attribute gets `Attr` suffix
- Enums: `type FooType string` + constants `FooTypeVALUE` + `IsValid()` method

### Generated client conventions

- Wraps `soap.Client` via embedding
- `NewClient` sets default endpoint from WSDL, accepts `...ClientOption` overrides
- One method per SOAP operation, accepting context + typed request + per-call options
- Methods handle envelope wrapping/unwrapping and XML marshal/unmarshal

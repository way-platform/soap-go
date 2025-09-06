# SOAP Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/way-platform/soap-go)](https://pkg.go.dev/github.com/way-platform/soap-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/way-platform/soap-go)](https://goreportcard.com/report/github.com/way-platform/soap-go)
[![CI](https://github.com/way-platform/soap-go/actions/workflows/release.yaml/badge.svg)](https://github.com/way-platform/soap-go/actions/workflows/release.yaml)

A Go SDK and CLI tool for SOAP web services.

## Features

- Support for SOAP 1.1, WSDL 1.1, and XSD 1.0
- Code generation from WSDL files
- Documentation generation

## Developing

See [AGENTS.md](./AGENTS.md).

## CLI tool

The `soap` CLI tool can generate code, documentation, and call SOAP APIs on the fly.

```bash
$ soap

  Multi-tool for SOAP APIs

  USAGE

    soap [command] [--flags]  

  CODE GENERATION

    gen [--flags]         Generate code for a SOAP API

  DOCUMENTATION

    doc [--flags]         Display documentation for a SOAP API

  NETWORK OPERATIONS

    call [--flags]        Call a SOAP action

  UTILS

    completion [command]  Generate the autocompletion script for the specified shell
    help [command]        Help about any command

  FLAGS

    -h --help             Help for soap
    -v --version          Version for soap
```

### Installing

```bash
go install github.com/way-platform/soap-go/cmd/soap@latest
```

Prebuilt binaries for Linux, Windows, and Mac are available from the [Releases](https://github.com/way-platform/soap-go/releases).

## License

This SDK is published under the [MIT License](./LICENSE).

## Security

Security researchers, see the [Security Policy](https://github.com/way-platform/soap-go?tab=security-ov-file#readme).

## Code of Conduct

Be nice. For more info, see the [Code of Conduct](https://github.com/way-platform/soap-go?tab=coc-ov-file#readme).

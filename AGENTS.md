# AGENTS.md

## Specs

When developing this SDK, use the WSDL and SOAP specifications in [docs](./docs):

- [SOAP 1.1](./docs/soap-1.1.md)
- [WSDL 1.1](./docs/wsdl-1.1.md)

## Structure

- The project uses a [tools](./tools/) directory with a separate Go module containing tools for building, linting and generating code.

- The project uses [Mage](https://magefile.org) with build tasks declared in [magefile.go](./tools/magefile.go).

## Developing

- Run tests with `./tools/mage test`

- Lint with `./tools/mage lint`

- Leave all version control and git to the user/developer. If you see a build error related to having a git diff, this is normal.

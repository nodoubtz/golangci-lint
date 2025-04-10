# GolangCI-Lint

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

GolangCI-Lint is a fast and flexible Go linters runner. Originally forked from the [official GolangCI-Lint](https://github.com/golangci/golangci-lint), this project aims to provide seamless linting capabilities for Go projects.

## Features

- **Fast Execution**: Optimized for speed to lint large codebases.
- **Multiple Linters**: Supports a wide variety of Go linters.
- **Highly Configurable**: Customize linting rules and configurations.
- **Extensible**: Add custom linters to fit your specific project needs.

## Getting Started

Follow the instructions below to integrate GolangCI-Lint into your workflow.

### Prerequisites

- Go version 1.20+ installed.
- Ensure your `GOPATH` is properly configured.

### Installation

You can install GolangCI-Lint using the following command:

```bash
go install github.com/nodoubtz/golangci-lint/cmd/golangci-lint@latest
```

### Running GolangCI-Lint

To lint your Go project, run:

```bash
golangci-lint run
```

For detailed output with issues grouped by files, use:

```bash
golangci-lint run --verbose
```

### Configuration

You can configure GolangCI-Lint using a `.golangci.yml` file. Here's an example configuration:

```yaml
linters:
  enable:
    - govet
    - errcheck
run:
  timeout: 5m
```

See the [official configuration guide](https://golangci-lint.run/usage/configuration/) for more details.

## Documentation

For full documentation, visit the [GolangCI-Lint website](https://golangci-lint.run).

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork this repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear description of your changes.

## License

This project is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0).

## Acknowledgments

This repository is a fork of the official [GolangCI-Lint](https://github.com/golangci/golangci-lint). We are grateful for the original maintainers' work and contributions.

## Contact

For questions or feedback, feel free to open an issue or reach out via [my GitHub profile](https://github.com/nodoubtz).
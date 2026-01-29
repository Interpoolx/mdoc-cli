# mdoc-cli

A high-performance Markdown CLI tool built with Go, distributed via NPM.

## Features

- **Blazing Fast**: Sub-50ms startup time.
- **Multi-platform**: Native binaries for Windows, Linux, and macOS.
- **Conversion**: Convert Markdown to HTML, JSON, and more.
- **Batch Processing**: Process entire directories concurrently.
- **Zero Dependencies**: Single binary execution.

### Release a new version

```bash
npm run publish-version 1.0.1
```

This automates versioning, cross-platform Go builds, and NPM publication.

## Installation

```bash
npx mdoc-cli --help
```

Or install globally:

```bash
npm install -g mdoc-cli@latest
```

## Usage

### Convert Markdown to HTML

```bash
# Using npx (Recommended)
npx mdoc-cli convert README.md -o output/README.html

# If installed globally
mdoc-cli convert README.md -o output/README.html
```

### Batch Conversion

```bash
npx mdoc-cli convert --recursive examples/ -o output/
```

### Validate Markdown

```bash
npx mdoc-cli validate README.md
```

## Configuration

`mdoc-cli` can be configured via flags, environment variables, or a `mdoc-cli.yaml` file.

```yaml
verbose: true
theme: dark
```

## Platform Support
`mdoc-cli` handles platform-specific binaries automatically via the NPM wrapper:
- **Windows**: Executes `md-cli-windows-amd64.exe`
- **Linux**: Executes `md-cli-linux-amd64` (x64)
- **macOS**: Executes `md-cli-darwin-arm64` (Apple Silicon)

## License

MIT © [Interpoolx](https://github.com/interpoolx)

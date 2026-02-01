# CLI Overview

The `aha` command-line tool provides access to the Aha! API from your terminal.

## Installation

```bash
go install github.com/grokify/go-aha/v3/cmd/aha@latest
```

## Global Flags

These flags are available for all commands:

| Flag | Short | Environment Variable | Description |
|------|-------|---------------------|-------------|
| `--domain` | `-d` | `AHA_DOMAIN` | Aha! subdomain |
| `--api-key` | `-k` | `AHA_API_KEY` | API key |
| `--help` | `-h` | - | Show help |

## Commands

| Command | Description |
|---------|-------------|
| [ideas](ideas.md) | Manage ideas |
| [features](features.md) | Manage features |
| [products](products.md) | Manage products |
| [releases](releases.md) | Manage releases |
| [report](report.md) | Generate reports |

## Usage Pattern

```bash
aha [global-flags] <command> <subcommand> [flags] [args]
```

## Examples

```bash
# List ideas
aha ideas list

# Get a specific feature
aha features get FEAT-123

# Generate a report
aha report idea-feature --all -f xlsx

# With explicit credentials
aha -d mycompany -k abc123 ideas list
```

## Output

By default, commands output JSON to stdout. Use flags to control output:

- `--format` / `-f` - Output format (json, markdown, xlsx)
- `--output` / `-o` - Write to file instead of stdout

## Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Error (credentials, API error, etc.) |

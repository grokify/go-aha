# Installation

## CLI Installation

Install the `aha` command-line tool:

```bash
go install github.com/grokify/go-aha/v3/cmd/aha@latest
```

Verify the installation:

```bash
aha --help
```

## SDK Installation

Add the SDK to your Go project:

```bash
go get github.com/grokify/go-aha/v3
```

## Requirements

- Go 1.21 or later
- An Aha! account with API access
- An API key from your Aha! account

## Getting an API Key

1. Log in to your Aha! account
2. Go to **Settings** > **Account** > **Personal**
3. Scroll to **Developer** section
4. Click **Generate API key**
5. Copy and save your API key securely

!!! warning "Keep your API key secure"
    Never commit your API key to version control. Use environment variables or a secure secrets manager.

## Next Steps

- [Configuration](configuration.md) - Set up your credentials
- [Quick Start](quickstart.md) - Run your first commands

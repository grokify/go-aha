# Configuration

## Authentication

The Aha! API requires two pieces of information:

| Parameter | Description | Example |
|-----------|-------------|---------|
| Domain | Your Aha! subdomain | `mycompany` (from `mycompany.aha.io`) |
| API Key | Your personal API key | `abc123...` |

## CLI Configuration

### Environment Variables (Recommended)

Set environment variables in your shell profile:

```bash
export AHA_DOMAIN="mycompany"
export AHA_API_KEY="your-api-key-here"
```

Then use the CLI without flags:

```bash
aha ideas list
```

### Command-Line Flags

Pass credentials directly (useful for scripts or one-off commands):

```bash
aha --domain mycompany --api-key your-api-key ideas list
```

Or use short flags:

```bash
aha -d mycompany -k your-api-key ideas list
```

### Priority Order

The CLI uses credentials in this order:

1. Command-line flags (`--domain`, `--api-key`)
2. Environment variables (`AHA_DOMAIN`, `AHA_API_KEY`)

## SDK Configuration

### Using Environment Variables

```go
import (
    "os"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
)

func main() {
    domain := os.Getenv("AHA_DOMAIN")
    apiKey := os.Getenv("AHA_API_KEY")

    cfg, err := client.NewConfiguration(domain, apiKey)
    if err != nil {
        panic(err)
    }

    apiClient := aha.NewAPIClient(cfg)
    // Use apiClient...
}
```

### Direct Configuration

```go
cfg, err := client.NewConfiguration("mycompany", "your-api-key")
if err != nil {
    panic(err)
}

apiClient := aha.NewAPIClient(cfg)
```

## Security Best Practices

!!! tip "Environment Variables"
    Always use environment variables in production. Never hardcode API keys.

!!! warning "Git Ignore"
    Add `.env` and similar files to `.gitignore` to prevent accidental commits.

Example `.env` file (do not commit):

```bash
AHA_DOMAIN=mycompany
AHA_API_KEY=your-api-key-here
```

Load with a tool like [direnv](https://direnv.net/) or source manually:

```bash
source .env
```

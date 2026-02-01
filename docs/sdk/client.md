# API Client

## Creating a Client

### Using the Helper

The `client` package provides a helper for creating configurations:

```go
import (
    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
)

// Create configuration
cfg, err := client.NewConfiguration("mycompany", "api-key-here")
if err != nil {
    panic(err)
}

// Create API client
apiClient := aha.NewAPIClient(cfg)
```

### Direct Configuration

You can also configure the client directly:

```go
import "github.com/grokify/go-aha/v3/oag7/aha"

cfg := aha.NewConfiguration()
cfg.Servers = aha.ServerConfigurations{
    {
        URL: "https://mycompany.aha.io/api/v1",
    },
}
cfg.AddDefaultHeader("Authorization", "Bearer your-api-key")

apiClient := aha.NewAPIClient(cfg)
```

## Configuration Options

### Timeout

```go
cfg, _ := client.NewConfiguration(domain, apiKey)

// Access underlying HTTP client
cfg.HTTPClient.Timeout = 30 * time.Second
```

### Custom HTTP Client

```go
import "net/http"

customClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns: 10,
    },
}

cfg, _ := client.NewConfiguration(domain, apiKey)
cfg.HTTPClient = customClient
```

### Debug Mode

```go
cfg, _ := client.NewConfiguration(domain, apiKey)
cfg.Debug = true  // Logs HTTP requests/responses
```

## Simple Client

For direct HTTP requests without the generated client:

```go
import "github.com/grokify/go-aha/v3/oag7/client"

sc, err := client.NewSimpleClient("mycompany", "api-key")
if err != nil {
    panic(err)
}

// Make direct requests
resp, err := sc.Get("/ideas")
```

## Context Usage

All API calls require a context:

```go
import "context"

// Basic context
ctx := context.Background()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
// Call cancel() when done
```

## Error Handling

```go
ideas, resp, err := apiClient.IdeasAPI.ListIdeas(ctx).Execute()

if err != nil {
    // Check for specific error types
    if apiErr, ok := err.(*aha.GenericOpenAPIError); ok {
        fmt.Printf("API Error: %s\n", apiErr.Body())
    }
    return err
}

// Always check HTTP status
if resp.StatusCode == 401 {
    return fmt.Errorf("unauthorized: check API key")
}
if resp.StatusCode == 404 {
    return fmt.Errorf("not found")
}
if resp.StatusCode >= 400 {
    return fmt.Errorf("API error: %s", resp.Status)
}
```

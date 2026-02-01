# SDK Overview

The go-aha SDK provides programmatic access to the Aha! API in Go applications.

## Packages

| Package | Description |
|---------|-------------|
| `oag7/aha` | Generated API client (OpenAPI Generator v7) |
| `oag7/client` | Client configuration helpers |
| `oag7/ideas` | High-level idea helpers and reports |
| `oag7/features` | Feature helper utilities |

## Basic Usage

```go
import (
    "context"
    "os"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
)

func main() {
    // Create configuration
    cfg, err := client.NewConfiguration(
        os.Getenv("AHA_DOMAIN"),
        os.Getenv("AHA_API_KEY"),
    )
    if err != nil {
        panic(err)
    }

    // Create API client
    apiClient := aha.NewAPIClient(cfg)

    // Use the client
    ctx := context.Background()
    ideas, _, err := apiClient.IdeasAPI.ListIdeas(ctx).Execute()
}
```

## Available APIs

The client provides access to these API services:

| Service | Description |
|---------|-------------|
| `IdeasAPI` | Create, read, update ideas |
| `FeaturesAPI` | Manage features |
| `ProductsAPI` | List and get products |
| `ReleasesAPI` | Manage releases |

## Request Pattern

All API methods follow a builder pattern:

```go
// Build the request with optional parameters
request := apiClient.IdeasAPI.ListIdeas(ctx).
    Q("search term").           // Query string
    WorkflowStatus("Planned").  // Filter
    Page(1).                    // Pagination
    PerPage(50)

// Execute the request
response, httpResponse, err := request.Execute()
```

## Response Handling

```go
ideas, resp, err := apiClient.IdeasAPI.ListIdeas(ctx).Execute()

// Check for errors
if err != nil {
    log.Fatalf("API error: %v", err)
}

// Check HTTP status
if resp.StatusCode >= 400 {
    log.Fatalf("HTTP error: %s", resp.Status)
}

// Access data
for _, idea := range ideas.Ideas {
    fmt.Printf("%s: %s\n", idea.ReferenceNum, idea.Name)
}

// Access pagination info
if ideas.Pagination != nil {
    fmt.Printf("Page %d of %d\n",
        ideas.Pagination.GetCurrentPage(),
        ideas.Pagination.GetTotalPages())
}
```

## Next Steps

- [API Client](client.md) - Detailed client configuration
- [Ideas API](ideas.md) - Working with ideas
- [Features API](features.md) - Working with features
- [Reports](reports.md) - Generating reports

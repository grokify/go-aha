# Quick Start

This guide walks you through common tasks with both the CLI and SDK.

## Prerequisites

1. [Install](installation.md) the CLI or SDK
2. [Configure](configuration.md) your credentials

## CLI Quick Start

### List Ideas

```bash
# List first 10 ideas
aha ideas list --per-page 10

# Search for ideas by name
aha ideas list --query "authentication"

# Filter by workflow status
aha ideas list --status "Under consideration"
```

### Get a Specific Idea

```bash
aha ideas get IDEA-123
```

### List Products

```bash
aha products list
```

### Generate a Report

```bash
# Generate JSON report of ideas with features
aha report idea-feature --per-page 50

# Export all ideas to Excel
aha report idea-feature --all -f xlsx -o ideas.xlsx

# Export to Markdown
aha report idea-feature --all -f markdown -o ideas.md
```

## SDK Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
)

func main() {
    // Create client
    cfg, err := client.NewConfiguration(
        os.Getenv("AHA_DOMAIN"),
        os.Getenv("AHA_API_KEY"),
    )
    if err != nil {
        panic(err)
    }

    apiClient := aha.NewAPIClient(cfg)
    ctx := context.Background()

    // List ideas
    ideas, _, err := apiClient.IdeasAPI.ListIdeas(ctx).
        PerPage(10).
        Execute()
    if err != nil {
        panic(err)
    }

    for _, idea := range ideas.Ideas {
        fmt.Printf("%s: %s\n", idea.ReferenceNum, idea.Name)
    }
}
```

### Generate a Report

```go
package main

import (
    "context"
    "os"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
    "github.com/grokify/go-aha/v3/oag7/ideas"
)

func main() {
    cfg, _ := client.NewConfiguration(
        os.Getenv("AHA_DOMAIN"),
        os.Getenv("AHA_API_KEY"),
    )
    apiClient := aha.NewAPIClient(cfg)

    // Fetch all ideas with automatic pagination
    req := ideas.ListIdeasRequest{
        FetchAll: true,
        PerPage:  100,
    }

    reportSet, err := ideas.GetIdeaFeatureReports(
        context.Background(),
        apiClient,
        req,
    )
    if err != nil {
        panic(err)
    }

    // Filter to only ideas with features
    withFeatures := reportSet.FilterByHasFeature(true)

    // Export to Excel
    withFeatures.Table().WriteXLSX("promoted_ideas.xlsx", "Ideas")
}
```

## Next Steps

- [CLI Reference](../cli/overview.md) - Full command documentation
- [SDK Guide](../sdk/overview.md) - Detailed SDK usage
- [Examples](../examples/cli.md) - More usage patterns

# go-aha

A Go SDK and CLI for the [Aha! Roadmap API](https://www.aha.io/api).

## Features

- **Full API Coverage** - Generated client from OpenAPI specification
- **CLI Tool** - Command-line interface for common operations
- **Report Generation** - Create idea-feature-release reports with XLSX/Markdown export
- **Helper Libraries** - High-level functions for common workflows

## Quick Example

### CLI

```bash
# List ideas
aha ideas list --per-page 10

# Generate a report of all ideas with their features and releases
aha report idea-feature --all -f xlsx -o ideas_report.xlsx
```

### SDK

```go
package main

import (
    "context"
    "fmt"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
    "github.com/grokify/go-aha/v3/oag7/ideas"
)

func main() {
    cfg, _ := client.NewConfiguration("your-domain", "your-api-key")
    apiClient := aha.NewAPIClient(cfg)

    // Generate a comprehensive report
    req := ideas.ListIdeasRequest{FetchAll: true}
    reportSet, _ := ideas.GetIdeaFeatureReports(context.Background(), apiClient, req)

    // Export to XLSX
    reportSet.Table().WriteXLSX("report.xlsx", "Ideas")
}
```

## Installation

```bash
go install github.com/grokify/go-aha/v3/cmd/aha@latest
```

## Documentation

- [Getting Started](getting-started/installation.md) - Installation and configuration
- [CLI Reference](cli/overview.md) - Command-line interface documentation
- [SDK Guide](sdk/overview.md) - Using the Go SDK in your applications
- [Examples](examples/cli.md) - Common usage patterns

## License

MIT License - see [LICENSE](https://github.com/grokify/go-aha/blob/master/LICENSE) for details.

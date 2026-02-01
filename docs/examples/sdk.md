# SDK Examples

Common usage patterns for the go-aha SDK.

## Setup

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
)

func getClient() *aha.APIClient {
    cfg, err := client.NewConfiguration(
        os.Getenv("AHA_DOMAIN"),
        os.Getenv("AHA_API_KEY"),
    )
    if err != nil {
        panic(err)
    }
    return aha.NewAPIClient(cfg)
}
```

## Listing Ideas

### Basic List

```go
func listIdeas() {
    client := getClient()
    ctx := context.Background()

    ideas, _, err := client.IdeasAPI.ListIdeas(ctx).
        PerPage(20).
        Execute()
    if err != nil {
        panic(err)
    }

    for _, idea := range ideas.Ideas {
        fmt.Printf("%s: %s (%d votes)\n",
            idea.ReferenceNum,
            idea.Name,
            idea.GetVotes())
    }
}
```

### Paginate All Ideas

```go
func getAllIdeas() []aha.Idea {
    client := getClient()
    ctx := context.Background()

    var allIdeas []aha.Idea
    page := int32(1)

    for {
        ideas, _, err := client.IdeasAPI.ListIdeas(ctx).
            Page(page).
            PerPage(100).
            Execute()
        if err != nil {
            panic(err)
        }

        allIdeas = append(allIdeas, ideas.Ideas...)

        if ideas.Pagination == nil ||
           int64(page) >= ideas.Pagination.GetTotalPages() {
            break
        }
        page++
    }

    return allIdeas
}
```

### Search and Filter

```go
func searchIdeas(query, status string) {
    client := getClient()
    ctx := context.Background()

    req := client.IdeasAPI.ListIdeas(ctx)

    if query != "" {
        req = req.Q(query)
    }
    if status != "" {
        req = req.WorkflowStatus(status)
    }

    ideas, _, err := req.PerPage(50).Execute()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Found %d ideas\n", len(ideas.Ideas))
}
```

## Working with Features

### Get Feature with Release

```go
func getFeatureDetails(refNum string) {
    client := getClient()
    ctx := context.Background()

    resp, _, err := client.FeaturesAPI.GetFeature(ctx, refNum).Execute()
    if err != nil {
        panic(err)
    }

    f := resp.Feature
    fmt.Printf("Feature: %s\n", f.Name)

    if f.WorkflowStatus != nil {
        fmt.Printf("Status: %s\n", f.WorkflowStatus.GetName())
    }

    if f.Release != nil {
        fmt.Printf("Release: %s (due %s)\n",
            f.Release.GetName(),
            f.Release.GetReleaseDate())
    }
}
```

## Generating Reports

### Full Report to Excel

```go
import "github.com/grokify/go-aha/v3/oag7/ideas"

func exportToExcel() {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{
        FetchAll: true,
        PerPage:  100,
    }

    reportSet, err := ideas.GetIdeaFeatureReports(ctx, client, req)
    if err != nil {
        panic(err)
    }

    err = reportSet.Table().WriteXLSX("ideas_report.xlsx", "Ideas")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Exported %d ideas to ideas_report.xlsx\n", reportSet.Len())
}
```

### Filtered Report

```go
func promotedIdeasReport() {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{
        FetchAll:       true,
        WorkflowStatus: "Planned",
    }

    reportSet, err := ideas.GetIdeaFeatureReports(ctx, client, req)
    if err != nil {
        panic(err)
    }

    // Filter to only promoted ideas with releases
    promoted := reportSet.
        FilterByHasFeature(true).
        FilterByHasRelease(true)

    // Sort by votes
    promoted.SortByVotes()

    // Export
    promoted.Table().WriteXLSX("promoted_ideas.xlsx", "Promoted")
}
```

### Markdown Report

```go
func markdownReport() {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{
        FetchAll: true,
    }

    reportSet, err := ideas.GetIdeaFeatureReports(ctx, client, req)
    if err != nil {
        panic(err)
    }

    // Generate markdown with links
    table := reportSet.TableWithLinks(
        "https://portal.aha.io",
        "https://company.aha.io",
    )

    md := table.Markdown("\n", true)
    fmt.Println(md)
}
```

## Analysis Examples

### Count by Status

```go
func countByStatus() {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{FetchAll: true}
    reportSet, _ := ideas.GetIdeaFeatureReports(ctx, client, req)

    counts := make(map[string]int)
    for _, r := range reportSet.Reports {
        counts[r.IdeaStatus]++
    }

    for status, count := range counts {
        fmt.Printf("%s: %d\n", status, count)
    }
}
```

### Top Voted Ideas

```go
func topVotedIdeas(n int) {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{FetchAll: true}
    reportSet, _ := ideas.GetIdeaFeatureReports(ctx, client, req)

    reportSet.SortByVotes()

    fmt.Printf("Top %d ideas by votes:\n", n)
    for i, r := range reportSet.Reports {
        if i >= n {
            break
        }
        fmt.Printf("%d. %s (%d votes) - %s\n",
            i+1, r.IdeaRefNum, r.IdeaVotes, r.IdeaName)
    }
}
```

### Ideas Without Features

```go
func backlogIdeas() {
    client := getClient()
    ctx := context.Background()

    req := ideas.ListIdeasRequest{FetchAll: true}
    reportSet, _ := ideas.GetIdeaFeatureReports(ctx, client, req)

    backlog := reportSet.FilterByHasFeature(false)
    backlog.SortByVotes()

    fmt.Printf("Backlog: %d ideas not yet promoted\n", backlog.Len())

    // Top 10 backlog items
    for i, r := range backlog.Reports {
        if i >= 10 {
            break
        }
        fmt.Printf("- %s: %s (%d votes)\n",
            r.IdeaRefNum, r.IdeaName, r.IdeaVotes)
    }
}
```

## Error Handling

```go
func robustAPICall() {
    client := getClient()
    ctx := context.Background()

    ideas, resp, err := client.IdeasAPI.ListIdeas(ctx).Execute()

    // Handle errors
    if err != nil {
        if apiErr, ok := err.(*aha.GenericOpenAPIError); ok {
            fmt.Printf("API Error Body: %s\n", apiErr.Body())
        }
        panic(err)
    }

    // Check HTTP status
    switch resp.StatusCode {
    case 401:
        panic("Unauthorized: check API key")
    case 403:
        panic("Forbidden: insufficient permissions")
    case 404:
        panic("Not found")
    case 429:
        panic("Rate limited: slow down requests")
    }

    if resp.StatusCode >= 400 {
        panic(fmt.Sprintf("HTTP error: %s", resp.Status))
    }

    // Process ideas...
    _ = ideas
}
```

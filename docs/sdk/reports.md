# Reports

The `oag7/ideas` package provides high-level helpers for generating comprehensive reports.

## Idea-Feature-Release Report

This report combines data from ideas, their promoted features, and assigned releases into a single view.

### Basic Usage

```go
import (
    "context"

    "github.com/grokify/go-aha/v3/oag7/aha"
    "github.com/grokify/go-aha/v3/oag7/client"
    "github.com/grokify/go-aha/v3/oag7/ideas"
)

func main() {
    cfg, _ := client.NewConfiguration(domain, apiKey)
    apiClient := aha.NewAPIClient(cfg)

    req := ideas.ListIdeasRequest{
        FetchAll: true,  // Paginate through all ideas
        PerPage:  100,   // Larger page size for efficiency
    }

    reportSet, err := ideas.GetIdeaFeatureReports(
        context.Background(),
        apiClient,
        req,
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("Total ideas: %d\n", reportSet.Len())
}
```

### Request Options

```go
req := ideas.ListIdeasRequest{
    // Pagination
    FetchAll: true,           // Fetch all pages automatically
    Page:     1,              // Starting page (if not FetchAll)
    PerPage:  100,            // Results per page

    // Filters
    Query:          "auth",   // Search query
    WorkflowStatus: "Planned", // Status filter
    Tag:            "mobile", // Tag filter
    UserID:         "user-id", // Creator filter

    // Date filters
    CreatedSince: &startDate,  // Created after
    UpdatedSince: &updateDate, // Updated after

    // Other
    Spam: false,              // Include spam
}
```

### Report Data

Each `IdeaFeatureReport` contains:

```go
type IdeaFeatureReport struct {
    // Idea fields
    IdeaID         string
    IdeaRefNum     string
    IdeaName       string
    IdeaStatus     string
    IdeaVotes      int64
    IdeaCategories []string
    IdeaCreatedAt  time.Time
    IdeaUpdatedAt  time.Time

    // Feature fields
    HasFeature       bool
    FeatureID        string
    FeatureRefNum    string
    FeatureName      string
    FeatureStatus    string
    FeatureStartDate string
    FeatureDueDate   string
    FeatureJiraKey   string
    FeatureJiraURL   string

    // Release fields
    HasRelease      bool
    ReleaseID       string
    ReleaseName     string
    ReleaseDate     string
    ReleaseReleased bool
}
```

## Filtering Reports

### By Feature Status

```go
// Only ideas promoted to features
withFeatures := reportSet.FilterByHasFeature(true)

// Ideas NOT yet promoted
withoutFeatures := reportSet.FilterByHasFeature(false)
```

### By Release Status

```go
// Features assigned to releases
withReleases := reportSet.FilterByHasRelease(true)

// Features not yet scheduled
unscheduled := reportSet.FilterByHasRelease(false)
```

### By Workflow Status

```go
// Filter by idea status
planned := reportSet.FilterByIdeaStatus("Planned")

// Filter by feature status
inDev := reportSet.FilterByFeatureStatus("In development")
```

### Chaining Filters

```go
// Ideas promoted to features, assigned to releases
result := reportSet.
    FilterByHasFeature(true).
    FilterByHasRelease(true)
```

## Sorting

```go
// Sort by vote count (highest first)
reportSet.SortByVotes()

// Sort by creation date (newest first)
reportSet.SortByCreatedAt()
```

## Exporting

### To Excel (XLSX)

```go
// Full table
err := reportSet.Table().WriteXLSX("report.xlsx", "Ideas")

// Compact table
err := reportSet.TableCompact().WriteXLSX("report.xlsx", "Ideas")
```

### To Markdown

```go
// Full table
md := reportSet.Table().Markdown("\n", true)
fmt.Println(md)

// Compact table
md := reportSet.TableCompact().Markdown("\n", true)

// With links
md := reportSet.TableWithLinks(
    "https://portal.aha.io",    // Idea portal URL
    "https://company.aha.io",   // Feature base URL
).Markdown("\n", true)
```

### To JSON

```go
import "encoding/json"

data, err := json.MarshalIndent(reportSet.Reports, "", "  ")
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

## Table Formats

### Full Table

Includes all columns:

```go
tbl := reportSet.Table()
```

Columns: Idea Ref, Name, Status, Categories, Votes, Has Feature, Created, Feature Ref, Feature Name, Feature Status, Start, Due, Has Release, Release Name, Release Date, Released, Jira Key

### Compact Table

Essential columns only:

```go
tbl := reportSet.TableCompact()
```

Columns: Idea Ref, Name, Status, Votes, Feature Ref, Feature Status, Release Name, Release Date

### Table with Links

Adds markdown links for idea and feature names:

```go
tbl := reportSet.TableWithLinks(
    "https://portal.aha.io",   // Idea links
    "https://company.aha.io",  // Feature links
)
```

## Single Idea Report

Fetch a report for a single idea:

```go
report, err := ideas.GetIdeaFeatureReport(ctx, apiClient, "IDEA-123")
if err != nil {
    panic(err)
}

fmt.Printf("Idea: %s\n", report.IdeaName)
if report.HasFeature {
    fmt.Printf("Feature: %s\n", report.FeatureName)
}
if report.HasRelease {
    fmt.Printf("Release: %s (%s)\n", report.ReleaseName, report.ReleaseDate)
}
```

## Accessing by Idea Reference

```go
// Get specific idea from set
if report, ok := reportSet.ByIdea["IDEA-123"]; ok {
    fmt.Printf("Found: %s\n", report.IdeaName)
}
```

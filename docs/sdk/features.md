# Features API

## Listing Features

### Basic List

```go
ctx := context.Background()
features, _, err := apiClient.FeaturesAPI.GetFeatures(ctx).Execute()
if err != nil {
    return err
}

for _, feature := range features.Features {
    fmt.Printf("%s: %s\n", feature.ReferenceNum, feature.Name)
}
```

### With Filters

```go
features, _, err := apiClient.FeaturesAPI.GetFeatures(ctx).
    Q("payment").                    // Search query
    Tag("Q1-2024").                  // Tag filter
    AssignedToUser("user@email.com"). // Assignee filter
    UpdatedSince(time.Now().AddDate(0, -1, 0)). // Last month
    Page(1).
    PerPage(50).
    Execute()
```

### Available Filters

| Method | Description |
|--------|-------------|
| `Q(string)` | Search query |
| `Tag(string)` | Filter by tag |
| `AssignedToUser(string)` | Filter by assignee email or ID |
| `UpdatedSince(time.Time)` | Updated after date |
| `Page(int32)` | Page number |
| `PerPage(int32)` | Results per page |

## Getting a Single Feature

```go
feature, _, err := apiClient.FeaturesAPI.GetFeature(ctx, "FEAT-123").Execute()
if err != nil {
    return err
}

f := feature.Feature
fmt.Printf("Feature: %s\n", f.Name)
fmt.Printf("Reference: %s\n", f.ReferenceNum)

// Workflow status
if f.WorkflowStatus != nil {
    fmt.Printf("Status: %s\n", f.WorkflowStatus.GetName())
}

// Release info
if f.Release != nil {
    fmt.Printf("Release: %s\n", f.Release.GetName())
    fmt.Printf("Release Date: %s\n", f.Release.GetReleaseDate())
}

// Dates
fmt.Printf("Start: %s\n", f.GetStartDate())
fmt.Printf("Due: %s\n", f.GetDueDate())
```

## Feature Model

Key fields in the `aha.Feature` struct:

| Field | Type | Description |
|-------|------|-------------|
| `Id` | `string` | Unique identifier |
| `ReferenceNum` | `string` | Human-readable reference (e.g., FEAT-123) |
| `Name` | `string` | Feature title |
| `WorkflowStatus` | `*FeatureWorkflowStatus` | Current status |
| `Release` | `*Release` | Assigned release |
| `StartDate` | `*string` | Start date (YYYY-MM-DD) |
| `DueDate` | `*string` | Due date (YYYY-MM-DD) |
| `Tags` | `[]string` | Assigned tags |
| `IntegrationFields` | `[]IntegrationField` | External integrations (Jira, etc.) |
| `CreatedAt` | `time.Time` | Creation timestamp |
| `Url` | `*string` | Web URL |

## Integration Fields (Jira, etc.)

Features can have integration fields linking to external systems:

```go
import "github.com/grokify/go-aha/v3/oag7/features"

// Use the features helper
f := features.Feature(*featureResp.Feature)

// Get Jira key
jiraKey := f.JiraKey()
jiraURL := f.JiraURL()

if jiraKey != "" {
    fmt.Printf("Jira: %s (%s)\n", jiraKey, jiraURL)
}
```

### Manual Access

```go
for _, field := range feature.IntegrationFields {
    if field.ServiceName != nil && *field.ServiceName == "Jira" {
        if field.Name != nil && *field.Name == "key" {
            fmt.Printf("Jira Key: %s\n", field.GetValue())
        }
    }
}
```

## Release Features

Get features for a specific release:

```go
features, _, err := apiClient.FeaturesAPI.
    GetReleaseFeatures(ctx, "release-id").
    Execute()
```

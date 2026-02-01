# Ideas API

## Listing Ideas

### Basic List

```go
ctx := context.Background()
ideas, _, err := apiClient.IdeasAPI.ListIdeas(ctx).Execute()
if err != nil {
    return err
}

for _, idea := range ideas.Ideas {
    fmt.Printf("%s: %s (%d votes)\n",
        idea.ReferenceNum,
        idea.Name,
        idea.GetVotes())
}
```

### With Filters

```go
ideas, _, err := apiClient.IdeasAPI.ListIdeas(ctx).
    Q("authentication").           // Search query
    WorkflowStatus("Planned").     // Status filter
    Tag("security").               // Tag filter
    Page(1).                       // Page number
    PerPage(50).                   // Results per page
    Execute()
```

### Available Filters

| Method | Description |
|--------|-------------|
| `Q(string)` | Search query |
| `WorkflowStatus(string)` | Filter by status name or ID |
| `Tag(string)` | Filter by tag |
| `UserId(string)` | Filter by creator |
| `CreatedSince(time.Time)` | Created after date |
| `UpdatedSince(time.Time)` | Updated after date |
| `Sort(string)` | Sort: recent, trending, popular |
| `Spam(bool)` | Include spam ideas |
| `Page(int32)` | Page number |
| `PerPage(int32)` | Results per page |

### Pagination

```go
page := int32(1)
perPage := int32(100)

for {
    ideas, _, err := apiClient.IdeasAPI.ListIdeas(ctx).
        Page(page).
        PerPage(perPage).
        Execute()
    if err != nil {
        return err
    }

    for _, idea := range ideas.Ideas {
        processIdea(idea)
    }

    // Check if there are more pages
    if ideas.Pagination == nil {
        break
    }
    if int64(page) >= ideas.Pagination.GetTotalPages() {
        break
    }
    page++
}
```

## Getting a Single Idea

```go
idea, _, err := apiClient.IdeasAPI.GetIdea(ctx, "IDEA-123").Execute()
if err != nil {
    return err
}

fmt.Printf("Idea: %s\n", idea.Idea.Name)
fmt.Printf("Status: %s\n", idea.Idea.WorkflowStatus.GetName())
fmt.Printf("Votes: %d\n", idea.Idea.GetVotes())

// Check if promoted to feature
if idea.Idea.Feature != nil {
    fmt.Printf("Feature: %s\n", idea.Idea.Feature.GetReferenceNum())
}
```

## Working with Idea Fields

### Optional Fields

Many fields are optional and use pointer types:

```go
idea := ideas.Ideas[0]

// Use getter methods for safe access
votes := idea.GetVotes()           // Returns 0 if nil
createdAt := idea.GetCreatedAt()   // Returns zero time if nil

// Check if field exists
if idea.HasVotes() {
    fmt.Printf("Votes: %d\n", idea.GetVotes())
}

// Direct pointer access (check nil first)
if idea.WorkflowStatus != nil && idea.WorkflowStatus.Name != nil {
    fmt.Printf("Status: %s\n", *idea.WorkflowStatus.Name)
}
```

### Categories

```go
for _, category := range idea.Categories {
    fmt.Printf("Category: %s\n", category.Name)
}
```

### Feature Reference

```go
if idea.Feature != nil {
    fmt.Printf("Promoted to: %s\n", idea.Feature.GetReferenceNum())
    fmt.Printf("Feature name: %s\n", idea.Feature.GetName())
}
```

## Idea Model

Key fields in the `aha.Idea` struct:

| Field | Type | Description |
|-------|------|-------------|
| `Id` | `string` | Unique identifier |
| `ReferenceNum` | `string` | Human-readable reference (e.g., IDEA-123) |
| `Name` | `string` | Idea title |
| `WorkflowStatus` | `*FeatureWorkflowStatus` | Current status |
| `Votes` | `*int32` | Vote count |
| `Categories` | `[]Category` | Assigned categories |
| `Feature` | `*IdeaFeature` | Promoted feature (if any) |
| `CreatedAt` | `*time.Time` | Creation timestamp |
| `UpdatedAt` | `*time.Time` | Last update timestamp |

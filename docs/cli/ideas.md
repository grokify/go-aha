# aha ideas

Manage Aha! ideas.

## Commands

| Command | Description |
|---------|-------------|
| `list` | List ideas with optional filtering |
| `get` | Get a specific idea by ID or reference |

---

## aha ideas list

List ideas with optional filtering.

### Usage

```bash
aha ideas list [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--query` | `-q` | | Search term to match against idea name |
| `--status` | `-s` | | Filter by workflow status ID or name |
| `--sort` | | | Sort by: recent, trending, or popular |
| `--tag` | `-t` | | Filter by tag |
| `--user-id` | | | Filter by creator user ID |
| `--created-since` | | | Only ideas created after timestamp (RFC3339) |
| `--updated-since` | | | Only ideas updated after timestamp (RFC3339) |
| `--page` | `-p` | 1 | Page number |
| `--per-page` | `-n` | 30 | Results per page |
| `--spam` | | false | Include spam ideas |

### Examples

```bash
# List first 10 ideas
aha ideas list --per-page 10

# Search by name
aha ideas list -q "login feature"

# Filter by status
aha ideas list --status "Under consideration"

# Filter by tag
aha ideas list --tag "mobile"

# Sort by popularity
aha ideas list --sort popular

# Get ideas updated in the last week
aha ideas list --updated-since "2024-01-01T00:00:00Z"

# Combine filters
aha ideas list -q "auth" --status "Planned" --per-page 50
```

### Output

JSON array of ideas:

```json
{
  "ideas": [
    {
      "id": "123456",
      "reference_num": "IDEA-42",
      "name": "Add SSO support",
      "workflow_status": {
        "name": "Under consideration"
      },
      "votes": 15,
      "categories": [
        {"name": "Authentication"}
      ],
      "created_at": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "total_records": 150,
    "total_pages": 5,
    "current_page": 1
  }
}
```

---

## aha ideas get

Get details for a specific idea.

### Usage

```bash
aha ideas get <idea-id>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `idea-id` | Idea ID or reference number (e.g., `IDEA-123`) |

### Examples

```bash
# Get by reference number
aha ideas get IDEA-123

# Get by ID
aha ideas get 6789abcd-1234-5678-9abc-def012345678
```

### Output

JSON object with full idea details:

```json
{
  "idea": {
    "id": "6789abcd-1234-5678-9abc-def012345678",
    "reference_num": "IDEA-123",
    "name": "Add dark mode support",
    "description": {
      "body": "Users have requested..."
    },
    "workflow_status": {
      "id": "status-id",
      "name": "Planned"
    },
    "votes": 42,
    "categories": [
      {"id": "cat-1", "name": "UI/UX"}
    ],
    "feature": {
      "id": "feature-id",
      "reference_num": "FEAT-456",
      "name": "Dark Mode"
    },
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-02-01T14:22:00Z"
  }
}
```

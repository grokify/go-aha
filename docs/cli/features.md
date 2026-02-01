# aha features

Manage Aha! features.

## Commands

| Command | Description |
|---------|-------------|
| `list` | List features with optional filtering |
| `get` | Get a specific feature by ID or reference |

---

## aha features list

List features with optional filtering.

### Usage

```bash
aha features list [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--query` | `-q` | | Search term to match against feature name or ID |
| `--tag` | `-t` | | Filter by tag |
| `--assigned-to` | | | Filter by assigned user ID or email |
| `--updated-since` | | | Only features updated after timestamp (RFC3339) |
| `--page` | `-p` | 1 | Page number |
| `--per-page` | `-n` | 30 | Results per page |

### Examples

```bash
# List first 20 features
aha features list --per-page 20

# Search by name
aha features list -q "payment"

# Filter by tag
aha features list --tag "Q1-2024"

# Filter by assignee
aha features list --assigned-to "john@example.com"

# Get recently updated features
aha features list --updated-since "2024-01-01T00:00:00Z"
```

### Output

JSON array of features:

```json
{
  "features": [
    {
      "id": "abc123",
      "reference_num": "FEAT-100",
      "name": "User Authentication",
      "workflow_status": {
        "name": "In development"
      },
      "release": {
        "name": "Q1 2024",
        "release_date": "2024-03-31"
      },
      "created_at": "2024-01-10T09:00:00Z"
    }
  ],
  "pagination": {
    "total_records": 75,
    "total_pages": 3,
    "current_page": 1
  }
}
```

---

## aha features get

Get details for a specific feature.

### Usage

```bash
aha features get <feature-id>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `feature-id` | Feature ID or reference number (e.g., `FEAT-123`) |

### Examples

```bash
# Get by reference number
aha features get FEAT-123

# Get by ID
aha features get abc123def456
```

### Output

JSON object with full feature details:

```json
{
  "feature": {
    "id": "abc123def456",
    "reference_num": "FEAT-123",
    "name": "OAuth 2.0 Integration",
    "workflow_status": {
      "id": "status-id",
      "name": "In development"
    },
    "release": {
      "id": "release-id",
      "name": "Q1 2024",
      "release_date": "2024-03-31"
    },
    "start_date": "2024-01-15",
    "due_date": "2024-02-28",
    "tags": ["security", "auth"],
    "integration_fields": [
      {
        "name": "Jira",
        "value": "JIRA-456"
      }
    ],
    "created_at": "2024-01-10T09:00:00Z",
    "url": "https://company.aha.io/features/FEAT-123"
  }
}
```

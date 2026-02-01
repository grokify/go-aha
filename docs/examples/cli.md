# CLI Examples

Common usage patterns for the `aha` command-line tool.

## Setup

```bash
# Set credentials (add to ~/.bashrc or ~/.zshrc)
export AHA_DOMAIN="mycompany"
export AHA_API_KEY="your-api-key"
```

## Ideas

### List and Search

```bash
# List all ideas (first page)
aha ideas list

# Search by keyword
aha ideas list -q "mobile app"

# Filter by status
aha ideas list --status "Under consideration"

# Filter by tag
aha ideas list --tag "customer-request"

# Combine filters
aha ideas list -q "login" --status "Planned" --tag "security"

# Sort by popularity
aha ideas list --sort popular --per-page 20

# Get ideas updated this month
aha ideas list --updated-since "2024-01-01T00:00:00Z"
```

### Get Details

```bash
# Get by reference number
aha ideas get IDEA-123

# Get by ID
aha ideas get abc123-def456
```

## Features

### List and Search

```bash
# List features
aha features list

# Search
aha features list -q "authentication"

# Filter by assignee
aha features list --assigned-to "john@example.com"

# Filter by tag
aha features list --tag "Q1-2024"
```

### Get Details

```bash
aha features get FEAT-456
```

## Products and Releases

```bash
# List products
aha products list

# Get product details
aha products get MYAPP

# List releases for a product
aha releases list --product MYAPP

# Get release details
aha releases get MYAPP-R-1
```

## Reports

### Basic Reports

```bash
# JSON output (default)
aha report idea-feature

# First 50 ideas
aha report idea-feature --per-page 50

# All ideas (automatic pagination)
aha report idea-feature --all
```

### Export Formats

```bash
# Excel spreadsheet
aha report idea-feature --all -f xlsx -o ideas_report.xlsx

# Markdown table
aha report idea-feature --all -f markdown -o ideas_report.md

# Compact format
aha report idea-feature --all -f markdown --compact
```

### Filtered Reports

```bash
# Only ideas with features
aha report idea-feature --all --has-feature yes

# Ideas promoted to features AND scheduled
aha report idea-feature --all --has-feature yes --has-release yes

# Unpromoted ideas (backlog)
aha report idea-feature --all --has-feature no

# By workflow status
aha report idea-feature --all --status "Planned"

# Search + filter
aha report idea-feature --all -q "auth" --has-feature yes
```

### With Links

```bash
# Markdown with clickable links
aha report idea-feature --all -f markdown \
  --idea-portal-url "https://portal.aha.io" \
  --feature-base-url "https://mycompany.aha.io" \
  -o report.md
```

## Scripting Examples

### Export All Ideas Weekly

```bash
#!/bin/bash
DATE=$(date +%Y%m%d)
aha report idea-feature --all -f xlsx -o "ideas_backup_${DATE}.xlsx"
```

### Filter and Process with jq

```bash
# Get idea names only
aha ideas list --per-page 100 | jq -r '.ideas[].name'

# Count by status
aha report idea-feature --all | jq 'group_by(.IdeaStatus) | map({status: .[0].IdeaStatus, count: length})'

# Ideas with most votes
aha report idea-feature --all | jq 'sort_by(-.IdeaVotes) | .[0:10] | .[] | "\(.IdeaVotes) votes: \(.IdeaName)"'
```

### Pipeline to Other Tools

```bash
# Convert to CSV (using jq)
aha report idea-feature --all | jq -r '
  ["Idea","Status","Votes","Feature","Release"],
  (.[] | [.IdeaRefNum, .IdeaStatus, .IdeaVotes, .FeatureRefNum, .ReleaseName])
  | @csv'

# Send to Slack (example)
COUNT=$(aha report idea-feature --all --has-feature no | jq length)
curl -X POST -H 'Content-type: application/json' \
  --data "{\"text\":\"Backlog: ${COUNT} ideas not yet promoted\"}" \
  $SLACK_WEBHOOK_URL
```

## Tips

!!! tip "Large Exports"
    For large datasets, use `--per-page 100` with `--all` to reduce API calls.

!!! tip "Debugging"
    Check your credentials:
    ```bash
    echo "Domain: $AHA_DOMAIN"
    aha products list  # Quick test
    ```

!!! warning "Rate Limits"
    The Aha! API has rate limits. If you hit them, add delays between requests in scripts.

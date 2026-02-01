# aha report

Generate reports from Aha! data.

## Commands

| Command | Description |
|---------|-------------|
| `idea-feature` | Generate idea-feature-release report |

---

## aha report idea-feature

Generate a comprehensive report of ideas with their associated features and releases.

This command:

1. Fetches ideas (with optional filtering)
2. For each idea with a promoted feature, fetches full feature details
3. Extracts release information from features
4. Outputs a combined report

### Usage

```bash
aha report idea-feature [flags]
```

### Flags

#### Pagination & Filtering

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--all` | `-a` | false | Fetch all pages automatically |
| `--query` | `-q` | | Search term to match against idea name |
| `--status` | `-s` | | Filter by workflow status |
| `--tag` | `-t` | | Filter by tag |
| `--page` | `-p` | 1 | Page number (ignored if `--all`) |
| `--per-page` | `-n` | 30 | Results per page |

#### Post-fetch Filters

| Flag | Default | Description |
|------|---------|-------------|
| `--has-feature` | | Filter by has feature (yes/no) |
| `--has-release` | | Filter by has release (yes/no) |

#### Output Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | json | Output format: json, markdown, xlsx |
| `--output` | `-o` | | Output file path |
| `--compact` | | false | Use compact table format |
| `--idea-portal-url` | | | Base URL for idea portal links |
| `--feature-base-url` | | | Base URL for feature links |

### Examples

#### Basic Usage

```bash
# Get first page of ideas as JSON
aha report idea-feature

# Get all ideas
aha report idea-feature --all

# Search and filter
aha report idea-feature --all -q "authentication" --status "Planned"
```

#### Export Formats

```bash
# Export to Excel
aha report idea-feature --all -f xlsx -o ideas_report.xlsx

# Export to Markdown
aha report idea-feature --all -f markdown -o ideas_report.md

# Compact table format
aha report idea-feature --all -f markdown --compact
```

#### Filtering Results

```bash
# Only ideas that have been promoted to features
aha report idea-feature --all --has-feature yes

# Ideas with features assigned to releases
aha report idea-feature --all --has-feature yes --has-release yes

# Ideas NOT yet promoted
aha report idea-feature --all --has-feature no
```

#### With Links

```bash
# Add clickable links in Markdown output
aha report idea-feature --all -f markdown \
  --idea-portal-url "https://portal.aha.io" \
  --feature-base-url "https://company.aha.io"
```

### Output Formats

#### JSON

Default format. Full data for programmatic processing:

```json
[
  {
    "IdeaID": "idea-123",
    "IdeaRefNum": "IDEA-42",
    "IdeaName": "Add SSO support",
    "IdeaStatus": "Planned",
    "IdeaVotes": 15,
    "IdeaCategories": ["Security", "Authentication"],
    "IdeaCreatedAt": "2024-01-15T10:30:00Z",
    "HasFeature": true,
    "FeatureRefNum": "FEAT-100",
    "FeatureName": "SSO Integration",
    "FeatureStatus": "In development",
    "HasRelease": true,
    "ReleaseName": "Q2 2024",
    "ReleaseDate": "2024-06-30",
    "FeatureJiraKey": "JIRA-456"
  }
]
```

#### Markdown

Table format for documentation or sharing:

```markdown
| Idea Ref | Idea Name | Idea Status | Idea Votes | Feature Ref | Feature Status | Release Name | Release Date |
|----------|-----------|-------------|------------|-------------|----------------|--------------|--------------|
| IDEA-42 | Add SSO support | Planned | 15 | FEAT-100 | In development | Q2 2024 | 2024-06-30 |
```

#### XLSX (Excel)

Spreadsheet with formatted columns. Automatically generates a filename if `--output` is not specified.

### Report Columns

#### Full Table

| Column | Description |
|--------|-------------|
| Idea Ref | Idea reference number |
| Idea Name | Idea title |
| Idea Status | Workflow status |
| Idea Categories | Categories (semicolon-separated) |
| Idea Votes | Vote count |
| Has Feature | Yes/No |
| Idea Created | Creation date |
| Feature Ref | Feature reference number |
| Feature Name | Feature title |
| Feature Status | Feature workflow status |
| Feature Start | Start date |
| Feature Due | Due date |
| Has Release | Yes/No |
| Release Name | Release name |
| Release Date | Target release date |
| Released | Yes/No |
| Jira Key | Linked Jira issue |

#### Compact Table (`--compact`)

| Column | Description |
|--------|-------------|
| Idea Ref | Idea reference number |
| Idea Name | Idea title |
| Idea Status | Workflow status |
| Idea Votes | Vote count |
| Feature Ref | Feature reference number |
| Feature Status | Feature workflow status |
| Release Name | Release name |
| Release Date | Target release date |

### Performance Notes

!!! tip "Large datasets"
    When fetching all ideas with `--all`, consider:

    - Use `--per-page 100` for fewer API calls
    - Each idea with a feature requires an additional API call
    - Progress is not shown; be patient for large datasets

!!! info "API rate limits"
    The Aha! API has rate limits. If you hit them, wait and retry.

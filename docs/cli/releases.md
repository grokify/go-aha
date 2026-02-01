# aha releases

Manage Aha! releases.

## Commands

| Command | Description |
|---------|-------------|
| `list` | List releases for a product |
| `get` | Get a specific release by ID or key |

---

## aha releases list

List releases for a specific product.

### Usage

```bash
aha releases list --product <product-id> [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--product` | | **required** | Product ID |
| `--page` | `-p` | 1 | Page number |
| `--per-page` | `-n` | 30 | Results per page |

### Examples

```bash
# List releases for a product
aha releases list --product MYAPP

# With pagination
aha releases list --product MYAPP --per-page 50
```

### Output

JSON array of releases:

```json
{
  "releases": [
    {
      "id": "release-123",
      "reference_num": "MYAPP-R-1",
      "name": "Q1 2024",
      "start_date": "2024-01-01",
      "release_date": "2024-03-31",
      "released": false,
      "parking_lot": false
    }
  ],
  "pagination": {
    "total_records": 8,
    "total_pages": 1,
    "current_page": 1
  }
}
```

---

## aha releases get

Get details for a specific release.

### Usage

```bash
aha releases get <release-id>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `release-id` | Release ID or key |

### Examples

```bash
# Get by ID
aha releases get release-123

# Get by reference
aha releases get MYAPP-R-1
```

### Output

JSON object with full release details:

```json
{
  "release": {
    "id": "release-123",
    "reference_num": "MYAPP-R-1",
    "name": "Q1 2024",
    "start_date": "2024-01-01",
    "release_date": "2024-03-31",
    "external_release_date": "2024-04-01",
    "released": false,
    "parking_lot": false,
    "url": "https://company.aha.io/releases/MYAPP-R-1"
  }
}
```

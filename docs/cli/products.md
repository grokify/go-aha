# aha products

Manage Aha! products.

## Commands

| Command | Description |
|---------|-------------|
| `list` | List all products |
| `get` | Get a specific product by ID or key |

---

## aha products list

List all products.

### Usage

```bash
aha products list [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--page` | `-p` | 1 | Page number |
| `--per-page` | `-n` | 30 | Results per page |

### Examples

```bash
# List all products
aha products list

# With pagination
aha products list --page 2 --per-page 10
```

### Output

JSON array of products:

```json
{
  "products": [
    {
      "id": "prod-123",
      "reference_prefix": "MYAPP",
      "name": "My Application",
      "product_line": false,
      "created_at": "2023-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "total_records": 5,
    "total_pages": 1,
    "current_page": 1
  }
}
```

---

## aha products get

Get details for a specific product.

### Usage

```bash
aha products get <product-id>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `product-id` | Product ID or key |

### Examples

```bash
# Get by ID
aha products get prod-123

# Get by key
aha products get MYAPP
```

### Output

JSON object with full product details:

```json
{
  "product": {
    "id": "prod-123",
    "reference_prefix": "MYAPP",
    "name": "My Application",
    "product_line": false,
    "created_at": "2023-01-01T00:00:00Z",
    "url": "https://company.aha.io/products/MYAPP"
  }
}
```

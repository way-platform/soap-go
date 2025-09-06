# Kitchensink API Documentation

| | |
|---|---|
| **Namespace** | `http://example.com/typetest` |

## Operations

## Custom Types

This section documents the custom data types defined in the schema.

### Simple Types

#### `StatusType`

**Base Type:** `string`

**Allowed Values:**
- `ACTIVE`
- `INACTIVE`
- `PENDING`


#### `PriorityType`

**Base Type:** `int`

**Allowed Values:**
- `1`
- `2`
- `3`


#### `UserIdType`

**Base Type:** `long`

**Pattern:**
- `[0-9]+`



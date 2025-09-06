# FilteringService API Documentation

| | |
|---|---|
| **Namespace** | `http://example.com/complexfiltering` |
| **Endpoint** | `http://example.com/complexfiltering` |

## Overview

Service for testing complex type filtering

## Available Operations

- **[TestFiltering](#testfiltering)** - Tests complex type filtering based on documentation

## Operations

### TestFiltering

> Tests complex type filtering based on documentation

**SOAP Action:** `http://example.com/complexfiltering/TestFiltering`

#### Request

**Message:** `TestRequest`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<TestRequest>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<DocumentedField>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<field1>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<UndocumentedField>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<field1>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<SimpleField>` | [tns:DocumentedSimpleType](#documentedsimpletype) | Yes |  |

#### Response

**Message:** `TestResponse`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<TestResponse>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Success>` | xs:boolean | Yes |  |


## Custom Types

This section documents the custom data types defined in the schema.

### Simple Types

#### `UndocumentedSimpleType`

**Base Type:** `string`

**Allowed Values:**
- `VALUE1`
- `VALUE2`


#### `DocumentedSimpleType`

**Base Type:** `string`

This simple type has documentation

**Allowed Values:**
- `OPTION1`
- `OPTION2`


### Complex Types

#### `DocumentedType`

This type has documentation and should appear in Custom Types section



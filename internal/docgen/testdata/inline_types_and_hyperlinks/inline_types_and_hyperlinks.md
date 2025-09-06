# InlineTypesService API Documentation

| | |
|---|---|
| **Namespace** | `http://example.com/inlinetypes` |
| **Endpoint** | `http://example.com/inlinetypes` |

## Overview

Service for testing inline simple types and hyperlinks

## Available Operations

- **[TestInlineTypes](#testinlinetypes)** - Tests inline simple types and hyperlink generation

## Operations

### TestInlineTypes

> Tests inline simple types and hyperlink generation

**SOAP Action:** `http://example.com/inlinetypes/TestInlineTypes`

#### Request

**Message:** `TestInlineTypesRequest`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<TestInlineTypesRequest>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;Version |  (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;RequestStatus | [tns:NamedStatusType](#namedstatustype) (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Priority>` | [Priority (inline)](#priority-inline) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Category>` | [Category (inline)](#category-inline) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Code>` | [Code (inline)](#code-inline) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<ShortText>` | [ShortText (inline)](#shorttext-inline) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<SecondPriority>` | [SecondPriority (inline)](#secondpriority-inline) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Status>` | [tns:NamedStatusType](#namedstatustype) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Details>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<field1>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Description>` | xs:string | Yes |  |

#### Response

**Message:** `TestInlineTypesResponse`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<TestInlineTypesResponse>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Success>` | xs:boolean | Yes |  |


## Custom Types

This section documents the custom data types defined in the schema.

### Simple Types

#### `NamedStatusType`

**Base Type:** `string`

A named status type for testing hyperlinks

**Allowed Values:**
- `ACTIVE`
- `INACTIVE`


#### `Category (inline)`

**Base Type:** `string`

**Allowed Values:**
- `URGENT`
- `NORMAL`


#### `Code (inline)`

**Base Type:** `string`

**Pattern:**
- `[A-Z]{2}-[0-9]{4}`


#### `Priority (inline)`

**Base Type:** `string`

Priority level for the request

**Allowed Values:**
- `HIGH`
- `MEDIUM`
- `LOW`


#### `SecondPriority (inline)`

**Base Type:** `string`

**Allowed Values:**
- `HIGH`
- `MEDIUM`
- `LOW`


#### `ShortText (inline)`

**Base Type:** `string`

**Length Constraints:**
- minLength: 1
- maxLength: 50


### Complex Types

#### `DocumentedComplexType`

This complex type has documentation and should appear in the Custom Types section



# TestService API Documentation

| | |
|---|---|
| **Namespace** | `http://example.com/custom` |
| **Endpoint** | `http://example.com/test` |

## Overview

Test service with custom namespace prefix binding

## Available Operations

- **[testOperation](#testoperation)** - Test operation with custom namespace prefix.

## Operations

### testOperation

> Test operation with custom namespace prefix.

**SOAP Action:** `testOperation`

#### Request

**Message:** `TestRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<parameters>` | element: custom:TestRequest | Unknown |  |

#### Response

**Message:** `TestResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<parameters>` | element: custom:TestResponse | Unknown |  |



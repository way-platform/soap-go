# TestService API Documentation

**Namespace:** `http://example.com/test`
**Endpoint:** `http://example.com/test`

## Overview

Test service with namespace prefixed binding reference

## Available Operations

- **[testOperation](#testoperation)** - Test operation for binding resolution.

## Operations

### testOperation {#testoperation}

> Test operation for binding resolution.

**SOAP Action:** `testOperation`

#### Request

**Message:** `TestRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| parameters | element: tns:TestRequest | Unknown | - |

#### Response

**Message:** `TestResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| parameters | element: tns:TestResponse | Unknown | - |



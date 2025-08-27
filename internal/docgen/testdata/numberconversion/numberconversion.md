# NumberConversion API Documentation

**Namespace:** `http://www.dataaccess.com/webservicesserver/`

## Overview

The Number Conversion Web Service, implemented with Visual DataFlex, provides functions that convert numbers into words or dollar amounts.

## Available Operations

- **[NumberToWords](#numbertowords)** - Returns the word corresponding to the positive number passed as parameter. Limited to quadrillions.
- **[NumberToDollars](#numbertodollars)** - Returns the non-zero dollar amount of the passed number.

## Operations

### NumberToWords {#numbertowords}

> Returns the word corresponding to the positive number passed as parameter. Limited to quadrillions.

#### Request

**Message:** `NumberToWordsSoapRequest`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| NumberToWords | object | Yes | - |
| NumberToWords.ubiNum | xs:unsignedLong | Yes | - |

#### Response

**Message:** `NumberToWordsSoapResponse`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| NumberToWordsResponse | object | Yes | - |
| NumberToWordsResponse.NumberToWordsResult | xs:string | Yes | - |


### NumberToDollars {#numbertodollars}

> Returns the non-zero dollar amount of the passed number.

#### Request

**Message:** `NumberToDollarsSoapRequest`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| NumberToDollars | object | Yes | - |
| NumberToDollars.dNum | xs:decimal | Yes | - |

#### Response

**Message:** `NumberToDollarsSoapResponse`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| NumberToDollarsResponse | object | Yes | - |
| NumberToDollarsResponse.NumberToDollarsResult | xs:string | Yes | - |



# PersonService API Documentation

| | |
|---|---|
| **Namespace** | `http://example.com/elementrefs` |
| **Endpoint** | `http://example.com/personservice` |

## Overview

Service for managing person information with element references

## Available Operations

- **[processPerson](#processperson)** - Process a person with addresses and contact information using element references
- **[getPersonInfo](#getpersoninfo)** - Get person information by ID with optional address details

## Operations

### processPerson

> Process a person with addresses and contact information using element references

**SOAP Action:** `processPerson`

#### Request

**Message:** `ProcessPersonRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<ProcessPersonRequest>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;requestId | xs:string (attribute) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Person>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;id | xs:long (attribute) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<firstName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<lastName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<age>` | xs:int | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Address>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;verified | xs:boolean (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<street>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<city>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<zipCode>` | xs:string | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<country>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<ContactInfo>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<email>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<phone>` | xs:string | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<preferences>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<newsletter>` | xs:boolean | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<notifications>` | xs:boolean | Yes |  |

#### Response

**Message:** `ProcessPersonResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<ProcessPersonResponse>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;responseId | xs:string (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<success>` | xs:boolean | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<message>` | xs:string | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<updatedPerson>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<Person>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;id | xs:long (attribute) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<firstName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<lastName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<age>` | xs:int | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<lastUpdated>` | xs:dateTime | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<validatedAddresses>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<Address>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;verified | xs:boolean (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<street>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<city>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<zipCode>` | xs:string | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<country>` | xs:string | Yes |  |


### getPersonInfo

> Get person information by ID with optional address details

**SOAP Action:** `getPersonInfo`

#### Request

**Message:** `GetPersonInfoRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<GetPersonInfoRequest>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<personId>` | xs:long | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<includeAddresses>` | xs:boolean | Yes |  |

#### Response

**Message:** `GetPersonInfoResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `<GetPersonInfoResponse>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<Person>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;id | xs:long (attribute) | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<firstName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<lastName>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<age>` | xs:int | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<addresses>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<Address>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;verified | xs:boolean (attribute) | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<street>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<city>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<zipCode>` | xs:string | No |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<country>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;`<ContactInfo>` | object | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<email>` | xs:string | Yes |  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`<phone>` | xs:string | No |  |



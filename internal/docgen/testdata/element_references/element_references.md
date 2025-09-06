# PersonService API Documentation

**Namespace:** `http://example.com/elementrefs`
**Endpoint:** `http://example.com/personservice`

## Overview

Service for managing person information with element references

## Available Operations

- **[processPerson](#processperson)** - Process a person with addresses and contact information using element references
- **[getPersonInfo](#getpersoninfo)** - Get person information by ID with optional address details

## Operations

### processPerson {#processperson}

> Process a person with addresses and contact information using element references

**SOAP Action:** `processPerson`

#### Request

**Message:** `ProcessPersonRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ProcessPersonRequest | object | Yes | - |
| ProcessPersonRequest.Person | object | Yes | - |
| ProcessPersonRequest.Person.firstName | xs:string | Yes | - |
| ProcessPersonRequest.Person.lastName | xs:string | Yes | - |
| ProcessPersonRequest.Person.age | xs:int | No | - |
| ProcessPersonRequest.Person.@id | xs:long (attribute) | Yes | - |
| ProcessPersonRequest.Address | object | Yes | - |
| ProcessPersonRequest.Address.street | xs:string | Yes | - |
| ProcessPersonRequest.Address.city | xs:string | Yes | - |
| ProcessPersonRequest.Address.zipCode | xs:string | No | - |
| ProcessPersonRequest.Address.country | xs:string | Yes | - |
| ProcessPersonRequest.Address.@verified | xs:boolean (attribute) | No | - |
| ProcessPersonRequest.ContactInfo | object | Yes | - |
| ProcessPersonRequest.ContactInfo.email | xs:string | Yes | - |
| ProcessPersonRequest.ContactInfo.phone | xs:string | No | - |
| ProcessPersonRequest.preferences | object | Yes | - |
| ProcessPersonRequest.preferences.newsletter | xs:boolean | Yes | - |
| ProcessPersonRequest.preferences.notifications | xs:boolean | Yes | - |
| ProcessPersonRequest.@requestId | xs:string (attribute) | Yes | - |

#### Response

**Message:** `ProcessPersonResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ProcessPersonResponse | object | Yes | - |
| ProcessPersonResponse.success | xs:boolean | Yes | - |
| ProcessPersonResponse.message | xs:string | No | - |
| ProcessPersonResponse.updatedPerson | object | Yes | - |
| ProcessPersonResponse.updatedPerson.Person | object | Yes | - |
| ProcessPersonResponse.updatedPerson.Person.firstName | xs:string | Yes | - |
| ProcessPersonResponse.updatedPerson.Person.lastName | xs:string | Yes | - |
| ProcessPersonResponse.updatedPerson.Person.age | xs:int | No | - |
| ProcessPersonResponse.updatedPerson.Person.@id | xs:long (attribute) | Yes | - |
| ProcessPersonResponse.updatedPerson.lastUpdated | xs:dateTime | Yes | - |
| ProcessPersonResponse.validatedAddresses | object | Yes | - |
| ProcessPersonResponse.validatedAddresses.Address | object | Yes | - |
| ProcessPersonResponse.validatedAddresses.Address.street | xs:string | Yes | - |
| ProcessPersonResponse.validatedAddresses.Address.city | xs:string | Yes | - |
| ProcessPersonResponse.validatedAddresses.Address.zipCode | xs:string | No | - |
| ProcessPersonResponse.validatedAddresses.Address.country | xs:string | Yes | - |
| ProcessPersonResponse.validatedAddresses.Address.@verified | xs:boolean (attribute) | No | - |
| ProcessPersonResponse.@responseId | xs:string (attribute) | No | - |


### getPersonInfo {#getpersoninfo}

> Get person information by ID with optional address details

**SOAP Action:** `getPersonInfo`

#### Request

**Message:** `GetPersonInfoRequestMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetPersonInfoRequest | object | Yes | - |
| GetPersonInfoRequest.personId | xs:long | Yes | - |
| GetPersonInfoRequest.includeAddresses | xs:boolean | Yes | - |

#### Response

**Message:** `GetPersonInfoResponseMessage`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| GetPersonInfoResponse | object | Yes | - |
| GetPersonInfoResponse.Person | object | Yes | - |
| GetPersonInfoResponse.Person.firstName | xs:string | Yes | - |
| GetPersonInfoResponse.Person.lastName | xs:string | Yes | - |
| GetPersonInfoResponse.Person.age | xs:int | No | - |
| GetPersonInfoResponse.Person.@id | xs:long (attribute) | Yes | - |
| GetPersonInfoResponse.addresses | object | Yes | - |
| GetPersonInfoResponse.addresses.Address | object | Yes | - |
| GetPersonInfoResponse.addresses.Address.street | xs:string | Yes | - |
| GetPersonInfoResponse.addresses.Address.city | xs:string | Yes | - |
| GetPersonInfoResponse.addresses.Address.zipCode | xs:string | No | - |
| GetPersonInfoResponse.addresses.Address.country | xs:string | Yes | - |
| GetPersonInfoResponse.addresses.Address.@verified | xs:boolean (attribute) | No | - |
| GetPersonInfoResponse.ContactInfo | object | Yes | - |
| GetPersonInfoResponse.ContactInfo.email | xs:string | Yes | - |
| GetPersonInfoResponse.ContactInfo.phone | xs:string | No | - |



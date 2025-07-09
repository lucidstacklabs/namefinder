# DNS Record Management API

## Overview

The DNS Record Management API provides endpoints for managing DNS records within namespaces. The API supports two authentication methods:
- **Admin Bearer Token**: Full administrative access via `/admin/api/v1/` endpoints
- **API Key**: Programmatic access via `/api/v1/` endpoints

## Authentication

### Admin Bearer Token
```
Authorization: Bearer <token>
```

### API Key
```
Authorization: ApiKey <api_key>
```

## Base URLs

- **Admin API**: `http://localhost:5301/admin/api/v1/`
- **Public API**: `http://localhost:5301/api/v1/`

## Endpoints

### Create DNS Record

Creates a new DNS record in the specified namespace.

#### Admin Endpoint
```http
POST /admin/api/v1/namespaces/{namespace_id}/records
```

#### Public Endpoint
```http
POST /api/v1/namespaces/{namespace_id}/records
```

**Headers:**
- `Authorization: Bearer <token>` (Admin) or `Authorization: ApiKey <api_key>` (Public)
- `Content-Type: application/json`

**Path Parameters:**
- `namespace_id` (string): The unique identifier of the namespace

**Request Body:**
```json
{
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN"
}
```

**Response:**
```json
{
  "id": "686e90ba259b44824fc012dc",
  "namespace_id": "686e814c7a17b87d6c8f5c1a",
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN",
  "creator_type": "admin",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T21:24:34.42219904+05:30",
  "updated_at": "2025-07-09T21:24:34.42219918+05:30"
}
```

### List DNS Records

Retrieves all DNS records in the specified namespace.

#### Admin Endpoint
```http
GET /admin/api/v1/namespaces/{namespace_id}/records
```

#### Public Endpoint
```http
GET /api/v1/namespaces/{namespace_id}/records
```

**Headers:**
- `Authorization: Bearer <token>` (Admin) or `Authorization: ApiKey <api_key>` (Public)

**Path Parameters:**
- `namespace_id` (string): The unique identifier of the namespace

**Response:**
```json
[
  {
    "id": "686e918f0c8222466821c565",
    "namespace_id": "686e814c7a17b87d6c8f5c1a",
    "name": "www.example.com",
    "type": "CNAME",
    "value": "example.github.io",
    "ttl": 60,
    "class": "IN",
    "creator_type": "admin",
    "creator_id": "686de8b94f3ea24b4a887a67",
    "created_at": "2025-07-09T15:58:07.394Z",
    "updated_at": "2025-07-09T15:58:07.394Z"
  }
]
```

### Get DNS Record

Retrieves a specific DNS record by its ID.

#### Admin Endpoint
```http
GET /admin/api/v1/namespaces/{namespace_id}/records/{record_id}
```

#### Public Endpoint
```http
GET /api/v1/namespaces/{namespace_id}/records/{record_id}
```

**Headers:**
- `Authorization: Bearer <token>` (Admin) or `Authorization: ApiKey <api_key>` (Public)

**Path Parameters:**
- `namespace_id` (string): The unique identifier of the namespace
- `record_id` (string): The unique identifier of the DNS record

**Response:**
```json
{
  "id": "686e918f0c8222466821c565",
  "namespace_id": "686e814c7a17b87d6c8f5c1a",
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN",
  "creator_type": "admin",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T15:58:07.394Z",
  "updated_at": "2025-07-09T15:58:07.394Z"
}
```

### Update DNS Record

Updates an existing DNS record.

#### Admin Endpoint
```http
PUT /admin/api/v1/namespaces/{namespace_id}/records/{record_id}
```

#### Public Endpoint
```http
PUT /api/v1/namespaces/{namespace_id}/records/{record_id}
```

**Headers:**
- `Authorization: Bearer <token>` (Admin) or `Authorization: ApiKey <api_key>` (Public)
- `Content-Type: application/json`

**Path Parameters:**
- `namespace_id` (string): The unique identifier of the namespace
- `record_id` (string): The unique identifier of the DNS record

**Request Body:**
```json
{
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN"
}
```

**Response:**
```json
{
  "id": "686e918f0c8222466821c565",
  "namespace_id": "686e814c7a17b87d6c8f5c1a",
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN",
  "creator_type": "admin",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T15:58:07.394Z",
  "updated_at": "2025-07-09T15:59:13.748Z"
}
```

### Delete DNS Record

Deletes a DNS record from the namespace.

#### Admin Endpoint
```http
DELETE /admin/api/v1/namespaces/{namespace_id}/records/{record_id}
```

#### Public Endpoint
```http
DELETE /api/v1/namespaces/{namespace_id}/records/{record_id}
```

**Headers:**
- `Authorization: Bearer <token>` (Admin) or `Authorization: ApiKey <api_key>` (Public)

**Path Parameters:**
- `namespace_id` (string): The unique identifier of the namespace
- `record_id` (string): The unique identifier of the DNS record

**Response:**
```json
{
  "id": "686e918f0c8222466821c565",
  "namespace_id": "686e814c7a17b87d6c8f5c1a",
  "name": "www.example.com",
  "type": "CNAME",
  "value": "example.github.io",
  "ttl": 60,
  "class": "IN",
  "creator_type": "admin",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T15:58:07.394Z",
  "updated_at": "2025-07-09T15:59:13.748Z"
}
```

## Data Models

### DNS Record Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier for the DNS record |
| `namespace_id` | string | ID of the namespace containing this record |
| `name` | string | Domain name for the DNS record |
| `type` | string | DNS record type (A, AAAA, CNAME, MX, TXT, etc.) |
| `value` | string | DNS record value |
| `ttl` | integer | Time-to-live in seconds |
| `class` | string | DNS class (typically "IN" for Internet) |
| `creator_type` | string | Type of creator ("admin" or "apikey") |
| `creator_id` | string | ID of the creator |
| `created_at` | string | ISO 8601 timestamp of creation |
| `updated_at` | string | ISO 8601 timestamp of last update |

### Request Body (Create/Update)

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Domain name for the DNS record |
| `type` | string | Yes | DNS record type |
| `value` | string | Yes | DNS record value |
| `ttl` | integer | Yes | Time-to-live in seconds |
| `class` | string | Yes | DNS class |

## Error Responses

The API returns appropriate HTTP status codes:
- `200 OK` - Successful operation
- `201 Created` - Record created successfully
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Authentication required
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Example Usage

### Using Admin Bearer Token

```bash
# Create a CNAME record
curl -X POST \
  'http://localhost:5301/admin/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/records' \
  -H 'Authorization: Bearer $TOKEN' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "www.example.com",
    "type": "CNAME",
    "value": "example.github.io",
    "ttl": 60,
    "class": "IN"
  }'
```

### Using API Key

```bash
# List all records
curl -X GET \
  'http://localhost:5301/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/records' \
  -H 'Authorization: ApiKey $API_KEY'
```
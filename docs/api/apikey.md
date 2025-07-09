# API Key Management

## Overview

The API Key management endpoints allow admin users to create, manage, and revoke API keys for accessing the DNS service. All API key management operations require admin authentication.

## Base URL
```
http://localhost:5301/api/v1
```

## Authentication

All API key management endpoints require admin authentication using a Bearer token in the `Authorization` header:
```
Authorization: Bearer <admin_token>
```

## API Key Operations

### Create API Key

Create a new API key for accessing the DNS service.

**Endpoint:** `POST /api-keys`

**Description:** Creates a new API key with the specified name.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Body:**
```json
{
  "name": "Test API Key"
}
```

**Parameters:**
- `name` (string, required): A descriptive name for the API key

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de95d4f3ea24b4a887a68",
  "name": "Test API Key",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T09:30:29.488204709+05:30",
  "updated_at": "2025-07-09T09:30:29.488204839+05:30"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the API key
- `name` (string): The descriptive name of the API key
- `creator_id` (string): ID of the admin user who created this API key
- `created_at` (string): ISO 8601 timestamp of when the API key was created
- `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/api-keys \
  -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Test API Key"}'
```

### List All API Keys

Retrieve a list of all API keys in the system.

**Endpoint:** `GET /api-keys`

**Description:** Returns a list of all API keys with their metadata (excluding secrets).

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Parameters:** None

#### Response

**Status Code:** `200 OK`

**Body:**
```json
[
  {
    "id": "686de95d4f3ea24b4a887a68",
    "name": "Test API Key",
    "creator_id": "686de8b94f3ea24b4a887a67",
    "created_at": "2025-07-09T04:00:29.488Z",
    "updated_at": "2025-07-09T04:00:29.488Z"
  }
]
```

**Response Fields:**
- Array of API key objects, each containing:
    - `id` (string): Unique identifier for the API key
    - `name` (string): The descriptive name of the API key
    - `creator_id` (string): ID of the admin user who created this API key
    - `created_at` (string): ISO 8601 timestamp of when the API key was created
    - `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/api-keys \
  -H "Authorization: Bearer $TOKEN"
```

### Get API Key by ID

Retrieve information about a specific API key by its ID.

**Endpoint:** `GET /api-keys/{id}`

**Description:** Returns the metadata of a specific API key identified by its ID (excluding the secret).

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the API key

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de95d4f3ea24b4a887a68",
  "name": "Test API Key",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T04:00:29.488Z",
  "updated_at": "2025-07-09T04:00:29.488Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the API key
- `name` (string): The descriptive name of the API key
- `creator_id` (string): ID of the admin user who created this API key
- `created_at` (string): ISO 8601 timestamp of when the API key was created
- `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/api-keys/686de95d4f3ea24b4a887a68 \
  -H "Authorization: Bearer $TOKEN"
```

### Update API Key

Update the name of an existing API key.

**Endpoint:** `PUT /api-keys/{id}`

**Description:** Updates the name of a specific API key.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the API key

**Body:**
```json
{
  "name": "Test API Key"
}
```

**Parameters:**
- `name` (string, required): The new descriptive name for the API key

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de95d4f3ea24b4a887a68",
  "name": "Test API Key",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T04:00:29.488Z",
  "updated_at": "2025-07-09T04:01:37.84Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the API key
- `name` (string): The updated descriptive name of the API key
- `creator_id` (string): ID of the admin user who created this API key
- `created_at` (string): ISO 8601 timestamp of when the API key was created
- `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/api-keys/686de95d4f3ea24b4a887a68 \
  -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Test API Key"}'
```

### Get API Key Secret

Retrieve the secret value of an API key.

**Endpoint:** `GET /api-keys/{id}/secret`

**Description:** Returns the secret value of a specific API key. This is the actual key value used for authentication.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the API key

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "secret": "puast73Zc1OAqBv5vzPjDYTN3AT2C4HaBTolsJBWu6iS2zZjhIfNpbNGWPmm87mmlRYNFd5bW5OQZH8okIIxsUjYy5RInLeSqaI7lz0nMgYqCa4TtO8CDlhzlii9m1ng"
}
```

**Response Fields:**
- `secret` (string): The secret value of the API key used for authentication

#### Example

```bash
curl localhost:5301/api/v1/api-keys/686de95d4f3ea24b4a887a68/secret \
  -H "Authorization: Bearer $TOKEN"
```

### Regenerate API Key Secret

Generate a new secret value for an existing API key.

**Endpoint:** `PUT /api-keys/{id}/secret`

**Description:** Regenerates the secret value for a specific API key. This invalidates the old secret and creates a new one.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the API key

**Body:** None

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "secret": "psb0wuPnqQ0lUUSB5PHnPun3LwFbNITlHuVWEBAF9wpqB4SmKVzsfI0bMScO51t9OoMt3lYDpXCYhvsXjrc6E3ZvvoaWTAc99Zp1afRSKihRNAS5VLoSoJssgzZOnZN4"
}
```

**Response Fields:**
- `secret` (string): The newly generated secret value of the API key

#### Example

```bash
curl localhost:5301/api/v1/api-keys/686de95d4f3ea24b4a887a68/secret \
  -X PUT \
  -H "Authorization: Bearer $TOKEN"
```

### Delete API Key

Delete an API key from the system.

**Endpoint:** `DELETE /api-keys/{id}`

**Description:** Deletes a specific API key from the system. This permanently revokes the API key and makes it unusable.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the API key to delete

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de95d4f3ea24b4a887a68",
  "name": "Test API Key",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T04:00:29.488Z",
  "updated_at": "2025-07-09T04:02:12.254Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the deleted API key
- `name` (string): The descriptive name of the deleted API key
- `creator_id` (string): ID of the admin user who created this API key
- `created_at` (string): ISO 8601 timestamp of when the API key was created
- `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/api-keys/686de95d4f3ea24b4a887a68 \
  -X DELETE \
  -H "Authorization: Bearer $TOKEN"
```

## Security Notes

- API key secrets are long, randomly generated strings that provide authentication to the DNS service
- Once an API key is deleted, it cannot be recovered and becomes permanently unusable
- Regenerating an API key secret invalidates the old secret immediately
- Store API key secrets securely and never expose them in logs or client-side code
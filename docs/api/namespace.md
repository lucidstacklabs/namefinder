# Namespace Management API

## Overview

Namespaces provide logical separation and organization of DNS resources within the authoritative DNS service. Each namespace can have specific API key access permissions, allowing fine-grained control over who can perform operations within each namespace.

## Base URL
```
http://localhost:5301/api/v1
```

## Authentication

All namespace management endpoints require admin authentication using a Bearer token in the `Authorization` header:
```
Authorization: Bearer <admin_token>
```

## Namespace Operations

### Create Namespace

Create a new namespace for organizing DNS resources.

**Endpoint:** `POST /namespaces`

**Description:** Creates a new namespace with the specified name.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Body:**
```json
{
  "name": "namespace1"
}
```

**Parameters:**
- `name` (string, required): A unique name for the namespace

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686e7fff7a17b87d6c8f5c18",
  "name": "namespace1",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T20:13:11.916534926+05:30",
  "updated_at": "2025-07-09T20:13:11.916535057+05:30"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the namespace
- `name` (string): The name of the namespace
- `creator_id` (string): ID of the admin user who created this namespace
- `created_at` (string): ISO 8601 timestamp of when the namespace was created
- `updated_at` (string): ISO 8601 timestamp of when the namespace was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces \
  -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "namespace1"}'
```

### List All Namespaces

Retrieve a list of all namespaces in the system.

**Endpoint:** `GET /namespaces`

**Description:** Returns a list of all namespaces with their metadata.

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
    "id": "686e7fff7a17b87d6c8f5c18",
    "name": "namespace1",
    "creator_id": "686de8b94f3ea24b4a887a67",
    "created_at": "2025-07-09T14:43:11.916Z",
    "updated_at": "2025-07-09T14:43:11.916Z"
  }
]
```

**Response Fields:**
- Array of namespace objects, each containing:
    - `id` (string): Unique identifier for the namespace
    - `name` (string): The name of the namespace
    - `creator_id` (string): ID of the admin user who created this namespace
    - `created_at` (string): ISO 8601 timestamp of when the namespace was created
    - `updated_at` (string): ISO 8601 timestamp of when the namespace was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces \
  -H "Authorization: Bearer $TOKEN"
```

### Get Namespace by ID

Retrieve information about a specific namespace by its ID.

**Endpoint:** `GET /namespaces/{id}`

**Description:** Returns the metadata of a specific namespace identified by its ID.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686e7fff7a17b87d6c8f5c18",
  "name": "namespace1",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T14:43:11.916Z",
  "updated_at": "2025-07-09T14:43:11.916Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the namespace
- `name` (string): The name of the namespace
- `creator_id` (string): ID of the admin user who created this namespace
- `created_at` (string): ISO 8601 timestamp of when the namespace was created
- `updated_at` (string): ISO 8601 timestamp of when the namespace was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e7fff7a17b87d6c8f5c18 \
  -H "Authorization: Bearer $TOKEN"
```

### Update Namespace

Update the name of an existing namespace.

**Endpoint:** `PUT /namespaces/{id}`

**Description:** Updates the name of a specific namespace.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

**Body:**
```json
{
  "name": "namespace1"
}
```

**Parameters:**
- `name` (string, required): The new name for the namespace

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686e7fff7a17b87d6c8f5c18",
  "name": "namespace1",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T14:43:11.916Z",
  "updated_at": "2025-07-09T14:43:11.916Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the namespace
- `name` (string): The updated name of the namespace
- `creator_id` (string): ID of the admin user who created this namespace
- `created_at` (string): ISO 8601 timestamp of when the namespace was created
- `updated_at` (string): ISO 8601 timestamp of when the namespace was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e7fff7a17b87d6c8f5c18 \
  -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "namespace1"}'
```

### Delete Namespace

Delete a namespace from the system.

**Endpoint:** `DELETE /namespaces/{id}`

**Description:** Deletes a specific namespace from the system. This permanently removes the namespace and all its associated resources.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace to delete

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686e7fff7a17b87d6c8f5c18",
  "name": "namespace1",
  "creator_id": "686de8b94f3ea24b4a887a67",
  "created_at": "2025-07-09T14:43:11.916Z",
  "updated_at": "2025-07-09T14:44:31.828Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the deleted namespace
- `name` (string): The name of the deleted namespace
- `creator_id` (string): ID of the admin user who created this namespace
- `created_at` (string): ISO 8601 timestamp of when the namespace was created
- `updated_at` (string): ISO 8601 timestamp of when the namespace was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e7fff7a17b87d6c8f5c18 \
  -X DELETE \
  -H "Authorization: Bearer $TOKEN"
```

## API Key Access Management

### Grant API Key Access to Namespace

Grant specific permissions to an API key for a namespace.

**Endpoint:** `POST /namespaces/{id}/api-keys`

**Description:** Grants an API key specific permissions (actions) within a namespace.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

**Body:**
```json
{
  "api_key_id": "686dea5d4f3ea24b4a887a69",
  "actions": ["create", "read", "update", "delete"]
}
```

**Parameters:**
- `api_key_id` (string, required): The ID of the API key to grant access to
- `actions` (array of strings, required): List of actions the API key can perform. Valid actions: `create`, `read`, `update`, `delete`

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "message": "api key access added successfully to namespace",
  "success": true
}
```

**Response Fields:**
- `message` (string): Success message
- `success` (boolean): Indicates if the operation was successful

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/api-keys \
  -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"api_key_id": "686dea5d4f3ea24b4a887a69", "actions": ["create", "read", "update", "delete"]}'
```

### List API Key Access for Namespace

Retrieve all API key permissions for a specific namespace.

**Endpoint:** `GET /namespaces/{id}/api-keys`

**Description:** Returns a list of all API keys that have access to the namespace and their permissions.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

#### Response

**Status Code:** `200 OK`

**Body:**
```json
[
  {
    "access": {
      "id": "686e816e1f2721837b763ce3",
      "namespace_id": "686e814c7a17b87d6c8f5c1a",
      "api_key_id": "686dea5d4f3ea24b4a887a69",
      "actions": ["create", "read", "update", "delete"],
      "creator_id": "686de8b94f3ea24b4a887a67",
      "created_at": "2025-07-09T14:49:18.929Z",
      "updated_at": "2025-07-09T14:49:18.929Z"
    },
    "api_key": {
      "id": "686dea5d4f3ea24b4a887a69",
      "name": "Test API Key",
      "creator_id": "686de8b94f3ea24b4a887a67",
      "created_at": "2025-07-09T04:04:45.433Z",
      "updated_at": "2025-07-09T04:04:45.433Z"
    }
  }
]
```

**Response Fields:**
- Array of objects, each containing:
    - `access` (object): Access control information
        - `id` (string): Unique identifier for the access record
        - `namespace_id` (string): ID of the namespace
        - `api_key_id` (string): ID of the API key
        - `actions` (array of strings): List of permitted actions
        - `creator_id` (string): ID of the admin who granted this access
        - `created_at` (string): ISO 8601 timestamp of when the access was granted
        - `updated_at` (string): ISO 8601 timestamp of when the access was last updated
    - `api_key` (object): API key information
        - `id` (string): Unique identifier for the API key
        - `name` (string): Name of the API key
        - `creator_id` (string): ID of the admin who created the API key
        - `created_at` (string): ISO 8601 timestamp of when the API key was created
        - `updated_at` (string): ISO 8601 timestamp of when the API key was last updated

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/api-keys \
  -H "Authorization: Bearer $TOKEN"
```

### Remove Specific API Key Permissions

Remove specific permissions from an API key for a namespace.

**Endpoint:** `DELETE /namespaces/{id}/api-keys`

**Description:** Removes specific actions/permissions from an API key's access to a namespace.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

**Body:**
```json
{
  "api_key_id": "686dea5d4f3ea24b4a887a69",
  "actions": ["delete"]
}
```

**Parameters:**
- `api_key_id` (string, required): The ID of the API key to remove permissions from
- `actions` (array of strings, required): List of actions to remove from the API key

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "message": "api key access deleted successfully from namespace",
  "success": true
}
```

**Response Fields:**
- `message` (string): Success message
- `success` (boolean): Indicates if the operation was successful

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/api-keys \
  -X DELETE \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"api_key_id": "686dea5d4f3ea24b4a887a69", "actions": ["delete"]}'
```

### Revoke All API Key Access

Completely revoke an API key's access to a namespace.

**Endpoint:** `POST /namespaces/{id}/api-keys/destroy`

**Description:** Completely removes all access permissions for an API key from a namespace.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the namespace

**Body:**
```json
{
  "api_key_id": "686dea5d4f3ea24b4a887a69"
}
```

**Parameters:**
- `api_key_id` (string, required): The ID of the API key to completely revoke access for

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "message": "api key access destroyed successfully from namespace",
  "success": true
}
```

**Response Fields:**
- `message` (string): Success message
- `success` (boolean): Indicates if the operation was successful

#### Example

```bash
curl localhost:5301/api/v1/namespaces/686e814c7a17b87d6c8f5c1a/api-keys/destroy \
  -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"api_key_id": "686dea5d4f3ea24b4a887a69"}'
```

## Permission Actions

The following actions can be granted to API keys for namespace access:

- `create`: Allows creating new DNS records and resources within the namespace
- `read`: Allows reading/viewing DNS records and resources within the namespace
- `update`: Allows modifying existing DNS records and resources within the namespace
- `delete`: Allows deleting DNS records and resources within the namespace

## Security Notes

- Namespaces provide logical isolation of DNS resources
- API key permissions are namespace-specific and can be fine-tuned per namespace
- Deleting a namespace removes all associated resources and API key permissions
- Use the principle of least privilege when granting API key permissions
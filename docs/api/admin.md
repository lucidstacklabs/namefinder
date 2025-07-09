# Admin Management API

### Base URL
```
http://localhost:5301/api/v1
```

### Initialize Admin User

Create the initial admin user for the DNS service.

**Endpoint:** `POST /admins/init`

**Description:** Creates the first admin user in the system. This endpoint is typically used during initial setup.

#### Request

**Headers:**
```
Content-Type: application/json
```

**Body:**
```json
{
  "username": "admin",
  "password": "123"
}
```

**Parameters:**
- `username` (string, required): The admin username
- `password` (string, required): The admin password

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de4174f3ea24b4a887a65",
  "username": "admin",
  "creator_id": "",
  "created_at": "2025-07-09T09:07:59.158383808+05:30",
  "updated_at": "2025-07-09T09:07:59.158383858+05:30"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the admin user
- `username` (string): The admin username
- `creator_id` (string): ID of the user who created this admin (empty for initial admin)
- `created_at` (string): ISO 8601 timestamp of when the admin was created
- `updated_at` (string): ISO 8601 timestamp of when the admin was last updated

#### Example

```bash
curl localhost:5301/api/v1/admins/init \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "123"}'
```

### Get Admin Token

Authenticate an admin user and retrieve a JWT token for API access.

**Endpoint:** `POST /admins/token`

**Description:** Authenticates an admin user and returns a JWT token for subsequent API requests.

#### Request

**Headers:**
```
Content-Type: application/json
```

**Body:**
```json
{
  "username": "admin",
  "password": "123"
}
```

**Parameters:**
- `username` (string, required): The admin username
- `password` (string, required): The admin password

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJuYW1lZmluZGVyIiwiZXhwIjoxNzUyMTE5Mzc1LCJpYXQiOjE3NTIwMzI5NzUsImlzcyI6Im5hbWVmaW5kZXIiLCJzdWIiOiI2ODZkZTQxNzRmM2VhMjRiNGE4ODdhNjUiLCJ0eXBlIjoiYWRtaW4ifQ.cWy6JG55C8yCg9YNVQAwd2ejrlYXjoTMr3cX-gymSls"
}
```

**Response Fields:**
- `token` (string): JWT token for authentication. Use this token in the `Authorization` header for subsequent API requests as `Bearer <token>`

#### Example

```bash
curl localhost:5301/api/v1/admins/token \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "123"}'
```

### Get Current Admin

Retrieve information about the currently authenticated admin user.

**Endpoint:** `GET /admins/current`

**Description:** Returns the profile information of the currently authenticated admin user.

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
{
  "id": "686de4174f3ea24b4a887a65",
  "username": "admin",
  "creator_id": "",
  "created_at": "2025-07-09T03:37:59.158Z",
  "updated_at": "2025-07-09T03:37:59.158Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the admin user
- `username` (string): The admin username
- `creator_id` (string): ID of the user who created this admin (empty for initial admin)
- `created_at` (string): ISO 8601 timestamp of when the admin was created
- `updated_at` (string): ISO 8601 timestamp of when the admin was last updated

#### Example

```bash
curl localhost:5301/api/v1/admins/current \
  -H "Authorization: Bearer $TOKEN"
```

### Update Current Admin Password

Update the password for the currently authenticated admin user.

**Endpoint:** `PUT /admins/current/password`

**Description:** Updates the password for the currently authenticated admin user.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Body:**
```json
{
  "password": "123"
}
```

**Parameters:**
- `password` (string, required): The new password for the admin user

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de4174f3ea24b4a887a65",
  "username": "admin",
  "creator_id": "",
  "created_at": "2025-07-09T03:37:59.158Z",
  "updated_at": "2025-07-09T03:52:20.899Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the admin user
- `username` (string): The admin username
- `creator_id` (string): ID of the user who created this admin (empty for initial admin)
- `created_at` (string): ISO 8601 timestamp of when the admin was created
- `updated_at` (string): ISO 8601 timestamp of when the admin was last updated (reflects the password update)

#### Example

```bash
curl localhost:5301/api/v1/admins/current/password \
  -X PUT \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"password": "123"}'
```

### Create Admin User

Create a new admin user in the system.

**Endpoint:** `POST /admins`

**Description:** Creates a new admin user. The system will generate a random password for the new admin.

#### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Body:**
```json
{
  "username": "admin1"
}
```

**Parameters:**
- `username` (string, required): The username for the new admin user

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "admin": {
    "id": "686de7b84f3ea24b4a887a66",
    "username": "admin1",
    "creator_id": "686de4174f3ea24b4a887a65",
    "created_at": "2025-07-09T09:23:28.474235137+05:30",
    "updated_at": "2025-07-09T09:23:28.474235187+05:30"
  },
  "password": "dWg59uGruHyl5tBY"
}
```

**Response Fields:**
- `admin` (object): The created admin user object
    - `id` (string): Unique identifier for the new admin user
    - `username` (string): The admin username
    - `creator_id` (string): ID of the admin user who created this admin
    - `created_at` (string): ISO 8601 timestamp of when the admin was created
    - `updated_at` (string): ISO 8601 timestamp of when the admin was last updated
- `password` (string): The generated password for the new admin user

#### Example

```bash
curl localhost:5301/api/v1/admins \
  -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"username": "admin1"}'
```

### List All Admin Users

Retrieve a list of all admin users in the system.

**Endpoint:** `GET /admins`

**Description:** Returns a list of all admin users in the system.

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
    "id": "686de4174f3ea24b4a887a65",
    "username": "admin",
    "creator_id": "",
    "created_at": "2025-07-09T03:37:59.158Z",
    "updated_at": "2025-07-09T03:52:20.899Z"
  },
  {
    "id": "686de7b84f3ea24b4a887a66",
    "username": "admin1",
    "creator_id": "686de4174f3ea24b4a887a65",
    "created_at": "2025-07-09T03:53:28.474Z",
    "updated_at": "2025-07-09T03:53:28.474Z"
  }
]
```

**Response Fields:**
- Array of admin objects, each containing:
    - `id` (string): Unique identifier for the admin user
    - `username` (string): The admin username
    - `creator_id` (string): ID of the user who created this admin (empty for initial admin)
    - `created_at` (string): ISO 8601 timestamp of when the admin was created
    - `updated_at` (string): ISO 8601 timestamp of when the admin was last updated

#### Example

```bash
curl localhost:5301/api/v1/admins \
  -H "Authorization: Bearer $TOKEN"
```

### Get Admin User by ID

Retrieve information about a specific admin user by their ID.

**Endpoint:** `GET /admins/{id}`

**Description:** Returns the profile information of a specific admin user identified by their ID.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the admin user

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de4174f3ea24b4a887a65",
  "username": "admin",
  "creator_id": "",
  "created_at": "2025-07-09T03:37:59.158Z",
  "updated_at": "2025-07-09T03:52:20.899Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the admin user
- `username` (string): The admin username
- `creator_id` (string): ID of the user who created this admin (empty for initial admin)
- `created_at` (string): ISO 8601 timestamp of when the admin was created
- `updated_at` (string): ISO 8601 timestamp of when the admin was last updated

#### Example

```bash
curl localhost:5301/api/v1/admins/686de4174f3ea24b4a887a65 \
  -H "Authorization: Bearer $TOKEN"
```

### Reset Admin User Password

Reset the password for a specific admin user by their ID.

**Endpoint:** `PUT /admins/{id}/password`

**Description:** Resets the password for a specific admin user. The system will generate a new random password for the user.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the admin user

**Body:** None

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "admin": {
    "id": "686de4174f3ea24b4a887a65",
    "username": "admin",
    "creator_id": "",
    "created_at": "2025-07-09T03:37:59.158Z",
    "updated_at": "2025-07-09T03:54:55.247Z"
  },
  "password": "cFOKrg64NnmtNdKw"
}
```

**Response Fields:**
- `admin` (object): The updated admin user object
    - `id` (string): Unique identifier for the admin user
    - `username` (string): The admin username
    - `creator_id` (string): ID of the user who created this admin
    - `created_at` (string): ISO 8601 timestamp of when the admin was created
    - `updated_at` (string): ISO 8601 timestamp of when the admin was last updated (reflects the password reset)
- `password` (string): The newly generated password for the admin user

#### Example

```bash
curl localhost:5301/api/v1/admins/686de4174f3ea24b4a887a65/password \
  -X PUT \
  -H "Authorization: Bearer $TOKEN"
```

### Delete Admin User

Delete a specific admin user by their ID.

**Endpoint:** `DELETE /admins/{id}`

**Description:** Deletes a specific admin user from the system. Returns the deleted admin user information.

#### Request

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `id` (string, required): The unique identifier of the admin user to delete

#### Response

**Status Code:** `200 OK`

**Body:**
```json
{
  "id": "686de4174f3ea24b4a887a65",
  "username": "admin",
  "creator_id": "",
  "created_at": "2025-07-09T03:37:59.158Z",
  "updated_at": "2025-07-09T03:54:55.247Z"
}
```

**Response Fields:**
- `id` (string): Unique identifier for the deleted admin user
- `username` (string): The admin username
- `creator_id` (string): ID of the user who created this admin
- `created_at` (string): ISO 8601 timestamp of when the admin was created
- `updated_at` (string): ISO 8601 timestamp of when the admin was last updated

#### Example

```bash
curl localhost:5301/api/v1/admins/686de4174f3ea24b4a887a65 \
  -X DELETE \
  -H "Authorization: Bearer $TOKEN"
```
# Visit Sidayu Back End API

REST API documentation for the **Visit Sidayu Back End**.

This API serves as the communication bridge between the Visit Sidayu Front End and Back End applications.

---

# API Version

> [!IMPORTANT]
> The current API version is **v1**.

All endpoints are prefixed with:

```text
/api/v1
```

For example:

```text
GET /api/v1/blogs
```

---

# API Explorer

> [!TIP]
> You can retrieve all available endpoints programmatically.
> You can use this static page to explore the API [here](./playground/api-playground.html)

```http
GET /api/v1/expose
```

---

# Authentication

Most write operations require JWT authentication.

Protected endpoints require the following header:

```http
Authorization: Bearer <your-jwt-token>
```

---

# Common Query Parameters

Most collection endpoints support pagination and searching.

| Parameter | Type    | Description              |
| --------- | ------- | ------------------------ |
| search    | string  | Search keyword           |
| page      | integer | Current page             |
| limit     | integer | Number of items per page |

Example:

```http
GET /api/v1/blogs?search=history&page=1&limit=10
```

---

# Standard Response Format

## Success Response

```json
{
    "success": true,
    "status_code": 200,
    "message": "...",
    "data": {},
    "jwt_token": "",
    "meta": {}
}
```

---

## Error Response

```json
{
    "success": false,
    "status_code": 400,
    "message": "...",
    "error": "..."
}
```

---

# Endpoints

---

# Health Check

## GET `/ping`

Check whether the API is running.

### Request

```http
GET /api/v1/ping
```

### Request Example

```json
{}
```

### Response Example

```json
{
    "message": "pong",
    "status": 200
}
```

---

# Authentication

## POST `/superadmins/auth/login`

Login as Superadmin.

### Content-Type

```text
application/json
```

### Request Body

```json
{
    "identity": "fkhrl",
    "password": "test-pw-123"
}
```

### Response Example

```json
{
    "succes": true,
    "status_code": 200,
    "message": "Login success!",
    "data": {
        "ID": "cd35b8a9-df03-4010-bb57-e5c420a3b809",
        "CreatedAt": "2026-08-04T12:15:21.447328+07:00",
        "UpdatedAt": "2026-08-04T12:15:21.447328+07:00",
        "DeletedAt": null,
        "username": "fkhrl",
        "password": "$2a$10$...",
        "email": "fkhrl@email.com"
    },
    "jwt_token": "eyJh..."
}
```

---

## POST `/superadmins/auth/register`

Register a Superadmin.

### Content-Type

```text
application/json
```

### Request Body

```json
{
    "username": "fkhrl",
    "password": "test-pw-123",
    "email": "fkhrl@email.com"
}
```

### Response Example

```json
{
    "succes": true,
    "status_code": 201,
    "message": "User registered!",
    "data": {
        "ID": "cd35b8a9-df03-4010-bb57-e5c420a3b809",
        "CreatedAt": "2026-08-04T12:15:21.4473286+07:00",
        "UpdatedAt": "2026-08-04T12:15:21.4473286+07:00",
        "DeletedAt": null,
        "username": "fkhrl",
        "email": "fkhrl@email.com"
    }
}
```

---

# Blogs

## GET `/blogs`

Retrieve all blogs.

### Query Parameters

| Name   | Required |
| ------ | -------- |
| search | No       |
| page   | No       |
| limit  | No       |

### Request

```http
GET /api/v1/blogs?search=&page=&limit=
```

### Response Example

```json
{
    "succes": true,
    "status_code": 200,
    "message": "Blogs data fetched successfully",
    "data": [
        {
            "ID": "bd39a45e-a08c-492d-9da5-e1d5a7c0570b",
            "CreatedAt": "2026-08-04T08:29:50.682287+07:00",
            "UpdatedAt": "2026-08-04T08:29:50.682287+07:00",
            "DeletedAt": null,
            "title": "Test Blog",
            "slug": "test-blog-08-29-50-04-08-2026",
            "description": "This is test blog",
            "tags": ["Tag1", "Tag2", "tag3"],
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ",
            "author": "John Doe",
            "blog_written_datetime": "2026-07-20T16:41:48+07:00",
            "estimated_minutes_read": 60,
            "thumbnail_id": "af4bbe1e-4583-414b-a6a7-2c3ea9ff9462",
            "thumbnail": {
                "ID": "af4bbe1e-4583-414b-a6a7-2c3ea9ff9462",
                "CreatedAt": "2026-08-04T08:29:47.970294+07:00",
                "UpdatedAt": "2026-08-04T08:29:47.970294+07:00",
                "DeletedAt": null,
                "image_url": "https://example.com/sidayu.jpg",
                "file_id": "",
                "name": "Pemandangan Sidayu",
                "custom_name": "",
                "description": "Foto indah"
            },
            "location": "Sidayu",
            "timeline": {
                "ID": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95",
                "CreatedAt": "2026-08-04T08:29:51.228923+07:00",
                "UpdatedAt": "2026-08-04T09:32:22.802003+07:00",
                "DeletedAt": null,
                "name": "Timeline test blog",
                "description": "timeline untuk test blog",
                "blog_id": "bd39a45e-a08c-492d-9da5-e1d5a7c0570b",
                "timeline_data": [
                    {
                        "ID": "a0ac2c30-610d-4783-bb64-2af312775798",
                        "CreatedAt": "2026-08-04T08:29:51.77303+07:00",
                        "UpdatedAt": "2026-08-04T09:32:23.87478+07:00",
                        "DeletedAt": null,
                        "name": "Important event 2 Updated lagi",
                        "timeline_datetime": "2026-07-20T16:41:48+07:00",
                        "timelines_id": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95"
                    },
                    {
                        "ID": "7de1e0de-d850-4963-8ba0-88180b7061f3",
                        "CreatedAt": "2026-08-04T09:21:22.596248+07:00",
                        "UpdatedAt": "2026-08-04T09:32:23.344397+07:00",
                        "DeletedAt": null,
                        "name": "Important event 1 updated lagi",
                        "timeline_datetime": "0001-01-01T07:00:00+07:00",
                        "description": "Keterangan event 1 lagi",
                        "timelines_id": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95"
                    }
                ]
            }
        }
    ],
    "meta": {
        "page": 1,
        "limit": 20,
        "total_rows": 1,
        "total_pages": 1
    }
}
```

---

## GET `/blogs/id/:id`

### Request

```http
GET /api/v1/blogs/id/{uuid}
```

### Response Example

```json
{
    "succes": true,
    "status_code": 200,
    "message": "Blogs data fetched successfully",
    "data": {
        "ID": "bd39a45e-a08c-492d-9da5-e1d5a7c0570b",
        "CreatedAt": "2026-08-04T08:29:50.682287+07:00",
        "UpdatedAt": "2026-08-04T08:29:50.682287+07:00",
        "DeletedAt": null,
        "title": "Test Blog",
        "slug": "test-blog-08-29-50-04-08-2026",
        "description": "This is test blog",
        "tags": ["Tag1", "Tag2", "tag3"],
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum ",
        "author": "John Doe",
        "blog_written_datetime": "2026-07-20T16:41:48+07:00",
        "estimated_minutes_read": 60,
        "thumbnail_id": "af4bbe1e-4583-414b-a6a7-2c3ea9ff9462",
        "thumbnail": {
            "ID": "af4bbe1e-4583-414b-a6a7-2c3ea9ff9462",
            "CreatedAt": "2026-08-04T08:29:47.970294+07:00",
            "UpdatedAt": "2026-08-04T08:29:47.970294+07:00",
            "DeletedAt": null,
            "image_url": "https://example.com/sidayu.jpg",
            "file_id": "",
            "name": "Pemandangan Sidayu",
            "custom_name": "",
            "description": "Foto indah"
        },
        "location": "Sidayu",
        "timeline": {
            "ID": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95",
            "CreatedAt": "2026-08-04T08:29:51.228923+07:00",
            "UpdatedAt": "2026-08-04T09:32:22.802003+07:00",
            "DeletedAt": null,
            "name": "Timeline test blog",
            "description": "timeline untuk test blog",
            "blog_id": "bd39a45e-a08c-492d-9da5-e1d5a7c0570b",
            "timeline_data": [
                {
                    "ID": "a0ac2c30-610d-4783-bb64-2af312775798",
                    "CreatedAt": "2026-08-04T08:29:51.77303+07:00",
                    "UpdatedAt": "2026-08-04T09:32:23.87478+07:00",
                    "DeletedAt": null,
                    "name": "Important event 2 Updated lagi",
                    "timeline_datetime": "2026-07-20T16:41:48+07:00",
                    "timelines_id": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95"
                },
                {
                    "ID": "7de1e0de-d850-4963-8ba0-88180b7061f3",
                    "CreatedAt": "2026-08-04T09:21:22.596248+07:00",
                    "UpdatedAt": "2026-08-04T09:32:23.344397+07:00",
                    "DeletedAt": null,
                    "name": "Important event 1 updated lagi",
                    "timeline_datetime": "0001-01-01T07:00:00+07:00",
                    "description": "Keterangan event 1 lagi",
                    "timelines_id": "f10a6ca1-0e75-42fe-b1ed-a43071eb9d95"
                }
            ]
        }
    }
}
```

---

## GET `/blogs/slug/:slug`

### Request

```http
GET /api/v1/blogs/slug/{slug}
```

### Response Example

```json
{
    // TODO
}
```

---

## POST `/blogs`

> [!IMPORTANT]
> JWT Required

### Content-Type

```text
application/json
```

### Request Body

```json
{
    // TODO
}
```

### Response Example

```json
{
    // TODO
}
```

---

## PATCH `/blogs/id/:id`

> [!IMPORTANT]
> JWT Required

### Request Body

```json
{
    // TODO
}
```

### Response Example

```json
{
    // TODO
}
```

---

## DELETE `/blogs/id/:id`

> [!IMPORTANT]
> JWT Required

### Response Example

```json
{
    // TODO
}
```

---

# Other Resources

The following resources follow the same documentation format:

- Images
- Culture Blogs
- Demographies
- Galleries
- Geographies
- Industries Blogs
- Officials
- Shops & UMKMs
- Timelines
- Roles

Each resource contains:

- GET Collection
- GET By ID
- GET By Slug (if applicable)
- POST
- PATCH
- DELETE

---

# Status Codes

| Code | Meaning               |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 409  | Conflict              |
| 422  | Validation Error      |
| 500  | Internal Server Error |

---

# Content Types

| Endpoint | Content-Type        |
| -------- | ------------------- |
| Images   | multipart/form-data |
| Others   | application/json    |

---

# Rate Limit

The API uses rate limiting.

If the limit is exceeded:

```http
429 Too Many Requests
```

---

# Notes

- All IDs are UUID v4.
- Dates use ISO 8601 format.
- All timestamps use UTC unless otherwise stated.
- Pagination starts from page **1**.

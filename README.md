# TrustWell/FoodLogiQ Event Management API

This project implements a RESTful API for managing events related to shipping and receiving using Go (Gin framework) and MongoDB. The API supports CRUD operations and utilizes rudimentary authentication.

### Features

1. Create an event
2. Retrieve a specific event
3. List all events for a user
4. Delete (soft delete) an event

### Prerequisites

- Docker
- Docker Compose
- Go (if running locally)

### Setup

1. Clone the Repository

```bash
git clone https://github.com/bridgekeeper/trustwell-api.git

cd trustwell-api
```

2. Run Docker Containers:

```bash
docker-compose up --build
```

3. Check if Containers are Running:

```bash
docker-compose ps
```

#### Accessing the API

The API will be accessible at http://localhost:8080.

#### API Endpoints

##### Authentication

- Use the following Bearer tokens to authenticate your requests:

```
74edf612f393b4eb01fbc2c29dd96671 for user 12345 (Acme)
d88b4b1e77c70ba780b56032db1c259b for user 98765 (Ajax)
```

##### Endpoints

- Create an Event: `POST /events`

Headers: Authorization: Bearer <token>

Body:

```json
{
  "type": "shipping",
  "contents": [
    {
      "gtin": "12345678901234",
      "lot": "abc123",
      "bestByDate": "2024-12-31",
      "expirationDate": "2025-01-15"
    }
  ]
}
```

Response:

```json
{
"id": "generated_event_id",
"createdAt": "timestamp",
"createdBy": "user_id",
"isDeleted": false,
"type": "shipping",
"contents": [...]
}
```

- Retrieve an Event: `GET /events/:id`

Response:

```json
{
"id": "event_id",
"createdAt": "timestamp",
"createdBy": "user_id",
"isDeleted": false,
"type": "shipping",
"contents": [...]
}
```

- List All Events `GET /events`

Headers: Authorization: Bearer <token>

Response:

```json

[
{
"id": "event_id",
"createdAt": "timestamp",
"createdBy": "user_id",
"isDeleted": false,
"type": "shipping",
"contents": [...]
},
...
]
```

- Delete an Event `DELETE /events/:id`

Headers: Authorization: Bearer <token>

Response:

```json
{
  "message": "Event deleted successfully"
}
```

.

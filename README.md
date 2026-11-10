# Go Notes CRUD API

A hands-on REST API project built while following **Sangam's GoLang Full Course 2026** on YouTube.

The project focuses on learning backend development with Go, Gin and MongoDB by building a Notes CRUD application
progressively.

## Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Gin
- **Database:** MongoDB Atlas
- **Database Driver:** MongoDB Go Driver
- **Configuration:** Environment variables (`.env`)
- **Live Reloading:** Air
- **API Testing:** Postman
- **Version Control:** Git

## Project Structure

```text
04-notes-crud/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── db/
│   │   └── mongo.go
│   │
│   ├── notes/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── handler.go
│   │   └── routes.go
│   │
│   └── server/
│       └── router.go
│
├── tmp/
├── .air.toml
├── .env
├── .gitignore
├── go.mod
└── README.md
```

## Recommended Implementation Sequence

The following sequence is our preferred approach for building this Gin-based backend project. Each layer builds upon the
previous one.

| Order | File                           | Responsibility                                         |
|------:|--------------------------------|--------------------------------------------------------|
|     1 | `internal/config/config.go`    | Load environment configuration                         |
|     2 | `internal/db/mongo.go`         | Establish and manage MongoDB connections               |
|     3 | `internal/notes/model.go`      | Define data structures and request DTOs                |
|     4 | `internal/notes/repository.go` | Implement database operations                          |
|     5 | `internal/notes/handler.go`    | Handle HTTP requests and application logic             |
|     6 | `internal/notes/routes.go`     | Register Notes API endpoints                           |
|     7 | `internal/server/router.go`    | Initialise Gin and combine application routes          |
|     8 | `cmd/api/main.go`              | Initialise dependencies and manage application startup |

*Note: This is our learning and implementation sequence, not a mandatory Go convention.*

## Application Architecture

The project follows a simple layered architecture with separation of concerns.

```text
                HTTP Client
               (Postman)
                    |
                    v
              Gin Router
            (server/router.go)
                    |
                    v
             Notes Routes
            (notes/routes.go)
                    |
                    v
             Notes Handler
            (notes/handler.go)
                    |
                    v
            Notes Repository
           (notes/repository.go)
                    |
                    v
              MongoDB Atlas
```

### Layer Responsibilities

**1. Configuration (`config/`)**

- Load environment variables.
- Manage application configuration such as MongoDB URI and server port.

**2. Database (`db/`)**

- Establish MongoDB connections.
- Verify database connectivity using Ping.
- Manage database disconnection and cleanup.

**3. Models (`notes/model.go`)**

- Define the Note struct.
- Define request DTOs such as CreateNoteRequest.
- Use JSON and BSON struct tags for API and database field mapping.

**4. Repository (`notes/repository.go`)**

- Encapsulate MongoDB database operations.
- Implement CRUD functionality.
- Keep database queries separate from HTTP handling.

**5. Handlers (`notes/handler.go`)**

- Receive HTTP requests.
- Bind and validate JSON payloads.
- Call repository methods.
- Return appropriate HTTP responses.

**6. Routes (`notes/routes.go`)**

- Create repository and handler dependencies.
- Register Notes endpoints using Gin.
- Group related endpoints under `/notes`.

**7. Server (`server/router.go`)**

- Initialise the Gin engine.
- Register application routes.
- Return the configured router.

**8. Main (`cmd/api/main.go`)**

- Load configuration.
- Establish database connectivity.
- Initialise the router.
- Start the HTTP server.
- Handle database cleanup during shutdown.

## Current API Endpoints

| HTTP Method | Endpoint     | Description           | Status      |
|-------------|--------------|-----------------------|-------------|
| GET         | `/ping`      | Health check          | Implemented |
| POST        | `/notes`     | Create a note         | In progress |
| GET         | `/notes`     | Retrieve all notes    | Planned     |
| GET         | `/notes/:id` | Retrieve a note by ID | Planned     |
| PUT         | `/notes/:id` | Update a note         | Planned     |
| DELETE      | `/notes/:id` | Delete a note         | Planned     |

## Current Progress

- [x] Project structure established.
- [x] Go modules initialised.
- [x] Environment configuration implemented.
- [x] MongoDB Atlas connection established.
- [x] MongoDB connectivity verified using Ping.
- [x] Air installed for live reloading.
- [x] Gin server initialised.
- [x] `/ping` endpoint responds with `pong`.
- [x] Note and CreateNoteRequest structs defined.
- [x] Notes route registration started.
- [ ] Complete repository implementation.
- [ ] Complete Notes HTTP handlers.
- [ ] Complete CRUD endpoints.
- [ ] Test endpoints using Postman.
- [ ] Initialise and maintain Git version control.

## Key Go Concepts Being Practised

- Structs and custom types
- Pointers and receivers
- Interfaces and dependency injection
- Error handling and error wrapping
- Context and timeout management
- Defer and resource cleanup
- Struct tags and reflection
- JSON encoding and decoding
- MongoDB ObjectID and BSON
- Gin routing and request binding
- Layered architecture and separation of concerns

## Running the Project

Start the application normally:

```bash
go run ./cmd/api
```

Or use Air for automatic reloading during development:

```bash
air
```

The server currently runs on the configured port (default configuration to be verified).

Test the health-check endpoint:

```bash
curl http://localhost:8080/ping
```

Expected response:

```text
pong
```

## Environment Configuration

Create a `.env` file in the project root.

Example:

```env
MONGO_URI=your_mongodb_connection_string
MONGO_DB=notesdb
SERVER_PORT=8080
```

Never commit real credentials or connection strings to Git. Keep `.env` in `.gitignore` and provide a sanitised
`.env.example` for other developers.

## Learning Approach

This project is being developed progressively alongside Sangam's GoLang Full Course 2026.

The focus is on understanding:

1. What each file and function does.
2. Why a particular Go construct or design pattern is used.
3. How the different application layers communicate.
4. How to implement backend functionality rather than simply copy code.
5. How to test and debug the application.

**Documentation will be updated as the tutorial progresses**, including additional endpoints, implementation steps, Go
concepts, architectural decisions and lessons learned.

---

*Learning Go by building practical backend applications.*
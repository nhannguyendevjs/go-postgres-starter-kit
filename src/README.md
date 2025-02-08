# go-postgres-starter-kit

Go Postgres starter kit


## Folder Structure

```bash
/go-postgres-starter-kit
│── cmd/                # Entry points for different services/applications
│   ├── app/            # Main application entry point
│   │   ├── main.go
│   ├── worker/         # If you have a worker process
│   │   ├── main.go
│
│── pkg/                # Reusable libraries used across the project (optional)
│   ├── logger/         # Logging utilities
│   ├── middleware/     # Middleware functions
│   ├── utils/          # General utilities
│
│── internal/           # Private application code
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and migrations
│   ├── models/         # Data models
│   ├── repository/     # Database operations (repository pattern)
│   ├── services/       # Business logic layer
│   ├── transport/      # Handlers for HTTP / gRPC
│   ├── worker/         # Background jobs
│
│── api/                # API specifications (OpenAPI, GraphQL, etc.)
│
│── web/                # Frontend files (if applicable)
│
│── migrations/         # Database migrations
│
│── scripts/            # Deployment and management scripts
│
│── test/               # Additional test helpers and integration tests
│
│── docs/               # Documentation files
│
│── .env                # Environment variables
│── .gitignore          # Git ignore file
│── go.mod              # Go modules file
│── go.sum              # Go modules dependencies
│── Makefile            # Build automation
│── Dockerfile          # Docker containerization
│── README.md           # Project documentation
```

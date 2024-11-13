# PhotoFlock

## Project Structure
```
.
├── api/                    # API layer: handlers, middleware, routes
│   ├── handlers/          # Request handlers
│   ├── middleware/        # Custom middleware
│   └── routes.go          # Route definitions
├── cmd/                   # Main applications for this project
│   └── api/              # API server entry point
├── internal/              # Private application and library code
│   ├── config/           # Configuration
│   ├── models/           # Domain models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic layer
├── pkg/                  # Library code that could be used by external applications
│   └── utils/           # Shared utilities
└── scripts/             # Scripts for development, CI/CD, etc.
```

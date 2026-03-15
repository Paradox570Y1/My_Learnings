
# Production Go Backend File Structure

```cmd
my-go-service/
│
├── cmd/                         # Entry points of the application
│   └── server/
│       └── main.go
│
├── internal/                    # Private application code
│   │
│   ├── config/                  # Environment & configuration loading
│   │   └── config.go
│   │
│   ├── handlers/                # HTTP handlers (controllers)
│   │   └── user_handler.go
│   │
│   ├── services/                # Business logic layer
│   │   └── user_service.go
│   │
│   ├── repository/              # Database interaction layer
│   │   └── user_repository.go
│   │
│   ├── models/                  # Database models
│   │   └── user.go
│   │
│   ├── dto/                     # Request/Response objects
│   │   └── user_dto.go
│   │
│   ├── middleware/              # HTTP middleware
│   │   └── auth_middleware.go
│   │
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   │
│   ├── database/                # DB connection setup
│   │   └── mysql.go
│   │
│   ├── utils/                   # Helper utilities
│   │   └── response.go
│   │
│   └── constants/               # App constants
│       └── constants.go
│
├── pkg/                         # Public reusable packages
│   └── logger/
│       └── logger.go
│
├── migrations/                  # Database migrations
│   └── 001_create_users.sql
│
├── scripts/                     # Dev / deployment scripts
│
├── tests/                       # Integration tests
│
├── .env
├── go.mod
├── go.sum
└── README.md
```


# Folder Responsibilities (Production Explanation)

- [[cmd]]
- [[internal]]
- [[pkg]]
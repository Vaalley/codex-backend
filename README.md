# Codex Backend

work needs to be done...

## Todo

- Add User Registering
- Add User Logging

## File Structure

.
├── main.go           # Entry point
├── config/
│   └── config.go     # Configuration management
│
├── api/             # API-related code
│   ├── routes.go    # Route definitions
│   └── middleware/  # Added when needed
│       ├── auth.go
│       └── logging.go
│
├── handlers/        # One file per feature
│   ├── games.go
│   ├── platforms.go
│   ├── users.go
│   ├── stats.go
│   └── social.go
│
├── models/          # Split by domain
│   ├── game.go
│   ├── platform.go
│   ├── user.go
│   └── stats.go
│
├── db/             # Database operations
│   ├── mongo.go    # DB connection/setup
│   ├── games.go
│   ├── platforms.go
│   └── users.go
│
├── services/       # Added for complex business logic
│   ├── auth.go     # Authentication logic
│   ├── search.go   # Search functionality
│   └── stats.go    # Statistics calculation
│
├── utils/          # Shared utilities
│   ├── errors.go
│   └── validation.go
│
├── tests/          # Tests directory
│   ├── handlers_test.go
│   └── services_test.go
│
├── .env
├── .env.example
├── go.mod
├── go.sum
└── README.md       # <-- You're here!

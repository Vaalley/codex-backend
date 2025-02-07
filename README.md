# Codex Backend

work needs to be done...

## Todo

- Add User Registering
- Add User Logging

## File Structure

.
├── main.go           # Thin entry point (keeps initialization clean)
├── config/
│   ├── config.go     # Load env vars, DB connection, etc.
│
├── api/
│   ├── routes.go     # Route definitions
│   └── handlers/     # Group handlers by domain
│       ├── users.go  # Handlers with direct DB access (for now)
│       └── games.go
│
├── models/           # Database models (structs)
│   ├── user.go
│   └── game.go
│
├── db/               # DB connection + direct queries
│   ├── mongo.go      # Initialize connection
│   ├── users.go      # Raw queries (e.g., GetUserByID)
│   └── games.go
│
├── utils/            # Helpers (e.g., JSON responses, validation)
│   └── http.go
│
├── .env
├── .env.example
├── go.mod
└── README.md

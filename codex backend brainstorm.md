# Codex Backend brainstorm

## Database Schema

### Users Table

| Column Name     | Data Type | Description                                        |
| --------------- | --------- | -------------------------------------------------- |
| `id`            | UUID      | Unique ID for the user.                            |
| `username`      | VARCHAR   | Username chosen by the user.                       |
| `email`         | VARCHAR   | User's email address.                              |
| `password_hash` | VARCHAR   | Hashed password for authentication.                |
| `created_at`    | TIMESTAMP | Timestamp of when the user account was created.    |
| `last_login`    | TIMESTAMP | Timestamp of the last user login.                  |
| `is_admin`      | BOOLEAN   | Indicates if the user is an admin.                 |
| `avatar_url`    | VARCHAR   | URL to the user's avatar image.                    |
| `bio`           | TEXT      | A short biography or description of the user.      |
| `timezone`      | VARCHAR   | User's preferred timezone.                         |
| `is_trusted`    | BOOLEAN   | Indicates if the user is trusted (e.g., verified). |
| `is_private`    | BOOLEAN   | Indicates if the user's profile is private.        |

### Platforms Table

| Column Name    | Data Type | Description                   |
| -------------- | --------- | ----------------------------- |
| `id`           | UUID      | Unique ID for the platform.   |
| `name`         | VARCHAR   | Name of the platform.         |
| `manufacturer` | VARCHAR   | Manufacturer of the platform. |

### Games Table

| Column Name        | Data Type      | Description                                             |
| ------------------ | -------------- | ------------------------------------------------------- |
| `id`               | UUID           | Unique ID for the game.                                 |
| `title`            | VARCHAR        | Title of the game.                                      |
| `developer`        | VARCHAR        | Developer of the game.                                  |
| `publisher`        | VARCHAR        | Publisher of the game.                                  |
| `release_date`     | DATE           | Release date of the game.                               |
| `genres`           | ARRAY of UUIDs | Array of genre IDs associated with the game.            |
| `platforms`        | ARRAY of UUIDs | Array of platform IDs the game is available on.         |
| `cover_image_url`  | VARCHAR        | URL to the cover image of the game.                     |
| `description`      | TEXT           | Brief description or synopsis of the game.              |
| `metacritic_score` | INT            | Aggregate review score (e.g., Metacritic).              |
| `user_scores`      | JSON           | JSON object storing user ratings and the overall score. |

### Genres Table

| Column Name   | Data Type | Description                     |
| ------------- | --------- | ------------------------------- |
| `id`          | UUID      | Unique ID for the genre.        |
| `name`        | VARCHAR   | Name of the genre.              |
| `description` | TEXT      | Brief description of the genre. |

### Game Progress Table

| Column Name      | Data Type | Description                                                                  |
| ---------------- | --------- | ---------------------------------------------------------------------------- |
| `user_id`        | UUID      | Reference to the user.                                                       |
| `game_id`        | UUID      | Reference to the game.                                                       |
| `status`         | VARCHAR   | Status of the game (`not started`, `in progress`, `completed`, `abandoned`). |
| `rating`         | INT       | User's rating for the game (optional).                                       |
| `review`         | TEXT      | User's written review of the game (optional).                                |
| `hours_played`   | INT       | Number of hours the user has played the game.                                |
| `date_completed` | TIMESTAMP | Date when the user marked the game as completed (optional).                  |

### Friendship Table

| Column Name | Data Type | Description                                           |
| ----------- | --------- | ----------------------------------------------------- |
| `user_id`   | UUID      | Reference to the user.                                |
| `friend_id` | UUID      | Reference to the friend (another user).               |
| `status`    | VARCHAR   | Friendship status (`pending`, `accepted`, `blocked`). |

## API Routes

### User Management

- `POST /create-user` - Register a new user.
- `POST /login` - User login.
- `POST /logout` - User logout.
- `POST /update-user` - Update user information (e.g., avatar, bio).
- `DELETE /delete-user` - Delete a user account.
- `GET /get-user?id=123` - Gets basic non-sensitive user info.

### Game Management

- `GET /get-game?id=123` - Get detailed information about a game.
- `POST /add-game` - Add a new game (admin route).
- `POST /update-game?id=123` - Update game information (admin route).
- `DELETE /delete-game?id=123` - Delete a game (admin route).

### Platform Management

- `GET /get-platform?id=123` - Get detailed information about a platform.
- `GET /get-platforms` - Get a list of all platforms.
- `POST /add-platform` - Add a new platform (admin route).
- `POST /update-platform?id=123` - Update platform information (admin route).
- `DELETE /delete-platform?id=123` - Delete a platform (admin route).

### Tracking Progress

- `POST /add-progress` - Add or update progress on a game for a user.
- `GET /get-progress?user_id=123&game_id=456` - Get progress of a specific game for a user.
- `GET /get-user-progress?user_id=123` - Get all progress for a user.

### Social Features

- `POST /add-friend?user_id=123&friend_id=456` - Send a friend request.
- `POST /accept-friend?user_id=123&friend_id=456` - Accept a friend request.
- `POST /remove-friend?user_id=123&friend_id=456` - Remove a friend.

### User Scores

- `POST /rate-game` - Rate a game and update the overall score.
- `GET /get-game-score?id=123` - Get the overall user score for a game.


## File Structure

```text
codex-backend/
├── cmd/
│   └── main.go                  # Entry point for the application, initializes app and routes
├── config/
│   └── config.go                # Loads and stores app configuration (e.g., database URI, JWT secret)
├── controllers/                 # Handles HTTP requests and responses
│   ├── user_controller.go       # User-related operations
│   ├── game_controller.go       # Game-related operations
│   ├── platform_controller.go   # Platform-related operations
│   └── progress_controller.go   # Game progress operations
├── models/                      # Data models and MongoDB schemas
│   ├── user.go                  # User model
│   ├── game.go                  # Game model
│   ├── platform.go              # Platform model
│   └── progress.go              # Game progress model
├── routes/
│   └── routes.go                # Organizes and registers all routes with Fiber
├── services/                    # Contains business logic
│   ├── auth_service.go          # Authentication (e.g., JWT generation)
│   └── user_service.go          # User-related operations (e.g., password hashing)
├── db/                          # Database setup
│   └── mongo.go                 # MongoDB connection logic
├── middleware/
│   └── auth_middleware.go       # Middleware for authentication (e.g., JWT validation)
├── utils/                       # Utility functions
│   └── hash.go                  # Password hashing utilities
├── .env                         # Environment variables (excluded from version control)
├── go.mod                       # Go module file
└── go.sum                       # Dependency checksum file
```
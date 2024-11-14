# Codex - Video Game Completion Tracker

A Go-based backend service for tracking your video game completion progress across different platforms. Built with clean architecture principles, Codex helps you manage your gaming journey and track your achievements.

## 🎮 Features

Current:
- Platform management (PlayStation, Xbox, Nintendo, etc.)
- Clean architecture implementation
- MongoDB integration
- Input validation
- Error handling
- Dependency injection

Coming Soon:
- Game management and tracking
- Completion status tracking
- Achievement tracking
- Progress statistics
- User profiles and authentication
- Game library management
- Custom completion criteria
- Progress sharing

## 📋 Prerequisites

- Go 1.21 or higher
- MongoDB
- Air (for live reloading during development)

## 🛠️ Tech Stack

- [Go Fiber](https://gofiber.io/) - Web framework
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - MongoDB driver
- [Validator](https://github.com/go-playground/validator) - Struct and field validation
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps

## 🏗️ Project Structure

```
.
├── cmd/                # Application entry points
├── config/            # Configuration
├── controllers/       # HTTP request handlers
├── db/               # Database connection and setup
├── docs/             # Documentation
├── middleware/       # HTTP middleware
├── models/           # Data models
├── repositories/     # Data access layer
├── routes/           # Route definitions
├── services/         # Business logic
└── utils/            # Utility functions
```

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd codex-backend
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   # Development (with live reload)
   air

   # Production
   go run cmd/main.go
   ```

## 📡 API Endpoints

### Currently Implemented

#### Platforms
- `GET /api/platforms` - List all gaming platforms
- `GET /api/platforms/:id` - Get platform details
- `POST /api/platforms` - Add new gaming platform
- `PUT /api/platforms/:id` - Update platform information
- `DELETE /api/platforms/:id` - Remove platform

### Coming Soon

#### Games
- `GET /api/games` - List all games
- `GET /api/games/:id` - Get game details
- `POST /api/games` - Add new game
- `PUT /api/games/:id` - Update game information
- `DELETE /api/games/:id` - Remove game

#### User Progress
- `GET /api/progress` - Get user's completion progress
- `POST /api/progress` - Update completion status
- `GET /api/statistics` - Get completion statistics

For detailed API documentation, see [request_flow.md](docs/request_flow.md).

## 🧪 Testing

Run tests:
```bash
go test ./...
```

## 🏛️ Architecture

This project follows clean architecture principles with distinct layers:

1. **Controllers**: HTTP request handling
2. **Services**: Business logic
3. **Repositories**: Data access
4. **Models**: Data structures

The clean architecture allows us to:
- Easily add new features and entities (games, achievements, etc.)
- Maintain separation of concerns
- Ensure testability
- Scale the application as it grows

For detailed architecture documentation, see [request_flow.md](docs/request_flow.md).

## 📝 Environment Variables

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=codex
PORT=3000
```

## 🎯 Roadmap

1. **Phase 1** 
   - Platform management
   - Basic API structure
   - Clean architecture implementation

2. **Phase 2** 
   - Game management
   - Basic completion tracking
   - User authentication

3. **Phase 3** 
   - Achievement tracking
   - Progress statistics
   - Social features

4. **Phase 4** 
   - Custom completion criteria
   - Advanced statistics
   - API documentation

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
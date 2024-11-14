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
- Progress statistics
- User profiles and authentication
- Game library management
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

## 📡 API Endpoints

### Currently Implemented

#### Platforms
- `GET /api/platforms` - List all gaming platforms
- `GET /api/platforms/:id` - Get platform details
- `POST /api/platforms` - Add new gaming platform
- `PUT /api/platforms/:id` - Update platform information
- `DELETE /api/platforms/:id` - Remove platform


For detailed API documentation, see [request_flow.md](docs/request_flow.md).

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

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
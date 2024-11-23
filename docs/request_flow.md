# Request Flow in Clean Architecture

## Overview
This document describes how requests flow through our clean architecture implementation.

## Layer Structure
```
HTTP Request
    ↓
Middleware (API Key validation, request validation)
    ↓
Routes (routing and DI)
    ↓
Controller (HTTP handling)
    ↓
Service (business logic)
    ↓
Repository (data access)
    ↓
Database
```

## Detailed Flow Example: Creating a Platform

### 1. HTTP Request
```http
POST /api/platforms
Content-Type: application/json
X-API-Key: your-api-key

{
    "name": "PlayStation 5",
    "manufacturer": "Sony"
}
```

### 2. Middleware Layer
```go
// API Key validation
api.Use(middleware.ValidateAPIKey)

// Request body validation
api.Post("/platforms", middleware.ValidateRequestBody(&models.Platform{}), ...)
```
Responsibilities:
- Validates API key for all requests
- Validates request body structure
- Handles validation errors

### 3. Routes Layer (`routes/routes.go`)
```go
platformRepo := repositories.NewMongoPlatformRepository(db.GetCollection("platforms"))
platformService := services.NewPlatformService(platformRepo)
platformController := controllers.NewPlatformController(platformService)

api.Post("/platforms", middleware.ValidateRequestBody(&models.Platform{}), platformController.CreatePlatform)
```
Responsibilities:
- Matches request to handler
- Sets up dependency injection chain
- Groups related routes
- Applies middleware

### 4. Controller Layer (`controllers/platform_controller.go`)
```go
func (pc *PlatformController) CreatePlatform(c fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var platform models.Platform
    if err := json.Unmarshal(c.Body(), &platform); err != nil {
        return c.Status(400).JSON(models.ErrorResponse{...})
    }

    if err := utils.ValidateStruct(platform); err != nil {
        return middleware.ErrorHandlerMiddleware(c, err)
    }

    if err := pc.service.CreatePlatform(ctx, &platform); err != nil {
        return c.Status(500).JSON(models.ErrorResponse{...})
    }

    return c.Status(201).JSON(platform)
}
```
Responsibilities:
- Sets up request context with timeout
- Extracts and validates request data
- Converts to domain model
- Calls appropriate service method
- Handles HTTP status codes
- Formats response

### 5. Service Layer (`services/platform_service.go`)
```go
func (s *platformService) CreatePlatform(ctx context.Context, platform *models.Platform) error {
    platform.ID = primitive.NewObjectID()
    platform.CreatedAt = time.Now()
    platform.UpdatedAt = time.Now()
    return s.repo.Create(ctx, platform)
}
```
Responsibilities:
- Implements business logic
- Manages entity lifecycle (IDs, timestamps)
- Validates business rules
- Coordinates with repository
- Handles domain-specific errors
- Transforms data if needed

### 6. Repository Layer (`repositories/platform_repository.go`)
```go
func (r *mongoPlatformRepository) Create(ctx context.Context, platform *models.Platform) error {
    _, err := r.collection.InsertOne(ctx, platform)
    return err
}
```
Responsibilities:
- Executes database operations
- Handles database-specific logic
- Maps between domain and database models
- Manages transactions

## Error Handling Flow

Errors flow up through the layers, with each layer adding appropriate context:

```
Repository Error (e.g., "duplicate key error")
    ↓
Service adds domain context (e.g., "platform already exists")
    ↓
Controller translates to HTTP status (e.g., 409 Conflict)
    ↓
HTTP Response with error message
```

Common error scenarios:
1. Validation Errors (400 Bad Request)
   - Invalid request body
   - Missing required fields
   - Invalid field formats

2. Authentication Errors (401 Unauthorized)
   - Missing API key
   - Invalid API key

3. Not Found Errors (404 Not Found)
   - Resource doesn't exist
   - Invalid ID format

4. Server Errors (500 Internal Server Error)
   - Database connection issues
   - Unexpected errors

## Available Routes

The API provides the following endpoints:

### Platforms
- `GET /api/platforms` - List all platforms
- `GET /api/platforms/:id` - Get platform by ID
- `POST /api/platforms` - Create new platform
- `PUT /api/platforms/:id` - Update platform
- `DELETE /api/platforms/:id` - Delete platform

### Games
- `GET /api/games` - List all games
- `GET /api/games/:id` - Get game by ID
- `POST /api/games` - Create new game
- `PUT /api/games/:id` - Update game
- `DELETE /api/games/:id` - Delete game

### Genres
- `GET /api/genres` - List all genres
- `GET /api/genres/:id` - Get genre by ID
- `POST /api/genres` - Create new genre
- `PUT /api/genres/:id` - Update genre
- `DELETE /api/genres/:id` - Delete genre

### Developers
- `GET /api/developers` - List all developers
- `GET /api/developers/:id` - Get developer by ID
- `POST /api/developers` - Create new developer
- `PUT /api/developers/:id` - Update developer
- `DELETE /api/developers/:id` - Delete developer

### Publishers
- `GET /api/publishers` - List all publishers
- `GET /api/publishers/:id` - Get publisher by ID
- `POST /api/publishers` - Create new publisher
- `PUT /api/publishers/:id` - Update publisher
- `DELETE /api/publishers/:id` - Delete publisher

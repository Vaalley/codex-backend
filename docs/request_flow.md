# Request Flow in Clean Architecture

## Overview
This document describes how requests flow through our clean architecture implementation.

## Layer Structure
```
HTTP Request
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

{
    "name": "PlayStation 5",
    "manufacturer": "Sony"
}
```

### 2. Routes Layer (`routes/routes.go`)
```go
api.Post("/platforms", middleware.ValidateRequestBody(&models.Platform{}), platformController.CreatePlatform)
```
- Matches request to handler
- Applies middleware (validation)
- Injects dependencies

### 3. Controller Layer (`controllers/platform_controller.go`)
```go
func (pc *PlatformController) CreatePlatform(c fiber.Ctx) error {
    // Parse request body
    // Call service
    // Handle response/errors
}
```
Responsibilities:
- Extracts data from HTTP request
- Converts to domain model
- Calls appropriate service method
- Handles HTTP status codes
- Formats response

### 4. Service Layer (`services/platform_service.go`)
```go
func (s *platformService) CreatePlatform(ctx context.Context, platform *models.Platform) error {
    // Apply business rules
    // Call repository
    // Handle domain errors
}
```
Responsibilities:
- Implements business logic
- Validates business rules
- Coordinates with repository
- Handles domain-specific errors
- Transforms data if needed

### 5. Repository Layer (`repositories/platform_repository.go`)
```go
func (r *mongoPlatformRepository) Create(ctx context.Context, platform *models.Platform) error {
    // Execute database operations
    // Handle database errors
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
Repository Error
    ↓
Service adds domain context
    ↓
Controller translates to HTTP status
    ↓
HTTP Response with error
```

Example:
1. Repository: "duplicate key error"
2. Service: "platform already exists"
3. Controller: HTTP 409 Conflict

## Validation Flow

Validation happens at multiple layers:

1. **HTTP/Request Validation** (Middleware)
   - JSON syntax
   - Required fields
   - Field formats

2. **Domain Validation** (Service)
   - Business rules
   - Complex validations
   - State checks

3. **Data Validation** (Repository)
   - Database constraints
   - Referential integrity

## Benefits of This Flow

1. **Separation of Concerns**
   - Each layer has a specific responsibility
   - Changes in one layer don't affect others

2. **Testability**
   - Can test each layer independently
   - Easy to mock dependencies
   - Clear boundaries for unit tests

3. **Maintainability**
   - Clear error handling
   - Consistent patterns
   - Easy to add new features

4. **Flexibility**
   - Can swap implementations
   - Easy to add middleware
   - Clear extension points

## Common Operations Flow

### GET Request
```
Request → Routes → Controller → Service → Repository → Database
Response ← Controller ← Service ← Repository ← Database
```

### POST/PUT Request
```
Request + Body → Routes → Validation → Controller → Service → Repository → Database
Response ← Controller ← Service ← Repository ← Database
```

### DELETE Request
```
Request → Routes → Controller → Service → Repository → Database
Response ← Controller ← Service ← Repository ← Database
```

## Cross-Cutting Concerns

These aspects affect multiple layers:

1. **Context & Timeouts**
   - Passed through all layers
   - Can cancel operations
   - Manages deadlines

2. **Logging**
   - Each layer can add context
   - Helps with debugging
   - Tracks request flow

3. **Error Handling**
   - Each layer handles appropriate errors
   - Adds context as needed
   - Maintains error chain

4. **Authentication/Authorization**
   - Verified at routes/middleware
   - Can be checked in service layer
   - Affects all operations

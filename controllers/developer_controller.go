package controllers

import (
	"codex-backend/middleware"
	"codex-backend/models"
	"codex-backend/services"
	"codex-backend/utils"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

type DeveloperController struct {
	service services.DeveloperService
}

func NewDeveloperController(service services.DeveloperService) *DeveloperController {
	return &DeveloperController{
		service: service,
	}
}

// GetDevelopers returns a list of all developers or filters by name if a query parameter is provided
func (dc *DeveloperController) GetDevelopers(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("📋 Fetching developers")
	name := c.Query("name")
	if name != "" {
		log.Printf("🔍 Searching for developers with name containing: %s", name)
	}

	developers, err := dc.service.GetDevelopers(ctx, name)
	if err != nil {
		log.Printf("❌ Error fetching developers: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch developers",
		})
	}

	log.Printf("✅ Successfully retrieved %d developers", len(developers))
	return c.JSON(developers)
}

// GetDeveloperByID returns a specific developer by its ID
func (dc *DeveloperController) GetDeveloperByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Fetching developer by ID")
	id := c.Params("id")
	log.Printf("🔍 Fetching developer with ID: %s", id)

	developer, err := dc.service.GetDeveloperByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			log.Printf("⚠️  Invalid developer ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			log.Printf("❌ Developer not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		default:
			log.Printf("❌ Error fetching developer: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch developer",
			})
		}
	}

	log.Printf("✅ Successfully retrieved developer: %s", developer.Name)
	return c.JSON(developer)
}

// CreateDeveloper adds a new developer
func (dc *DeveloperController) CreateDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("📝 Creating new developer")
	var developer models.Developer
	if err := json.Unmarshal(c.Body(), &developer); err != nil {
		log.Printf("❌ Error parsing developer data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(developer); err != nil {
		log.Printf("❌ Error validating developer data: %v", err)
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	log.Printf("📝 Creating new developer: %s", developer.Name)
	if err := dc.service.CreateDeveloper(ctx, &developer); err != nil {
		log.Printf("❌ Error creating developer: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create developer",
		})
	}

	log.Printf("✅ Successfully created developer: %s", developer.Name)
	return c.Status(201).JSON(developer)
}

// UpdateDeveloper updates an existing developer
func (dc *DeveloperController) UpdateDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Updating developer")
	id := c.Params("id")
	log.Printf("🔍 Updating developer with ID: %s", id)

	var update models.DeveloperUpdate
	if err := json.Unmarshal(c.Body(), &update); err != nil {
		log.Printf("❌ Error parsing update data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(update); err != nil {
		log.Printf("❌ Error validating update data: %v", err)
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	developer, err := dc.service.UpdateDeveloper(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			log.Printf("⚠️  Invalid developer ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			log.Printf("❌ Developer not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		case services.ErrNoDeveloperUpdateData:
			log.Printf("❌ No update data provided for developer with ID: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field must be provided for update",
			})
		default:
			log.Printf("❌ Error updating developer: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update developer",
			})
		}
	}

	log.Printf("✅ Successfully updated developer: %s", developer.Name)
	return c.JSON(developer)
}

// DeleteDeveloper removes a developer by its ID
func (dc *DeveloperController) DeleteDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Deleting developer")
	id := c.Params("id")
	log.Printf("🔍 Deleting developer with ID: %s", id)

	developer, err := dc.service.DeleteDeveloper(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			log.Printf("⚠️  Invalid developer ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			log.Printf("❌ Developer not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		default:
			log.Printf("❌ Error deleting developer: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete developer",
			})
		}
	}

	log.Printf("✅ Successfully deleted developer: %s", developer.Name)
	return c.JSON(developer)
}

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

type PlatformController struct {
	service services.PlatformService
}

func NewPlatformController(service services.PlatformService) *PlatformController {
	return &PlatformController{
		service: service,
	}
}

// GetPlatforms returns a list of all platforms or filters by name if a query parameter is provided
func (pc *PlatformController) GetPlatforms(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("📋 Fetching all platforms")
	name := c.Query("name")
	if name != "" {
		log.Printf("🔍 Searching for platforms with name containing: %s", name)
	}

	platforms, err := pc.service.GetPlatforms(ctx, name)
	if err != nil {
		log.Printf("❌ Error fetching platforms: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch platforms",
		})
	}

	log.Printf("✅ Successfully retrieved %d platforms", len(platforms))
	return c.JSON(platforms)
}

// GetPlatformByID returns a specific platform by its ID
func (pc *PlatformController) GetPlatformByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Fetching platform with ID: %s", id)

	platform, err := pc.service.GetPlatformByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			log.Printf("⚠️  Invalid platform ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrPlatformNotFound:
			log.Printf("❌ Platform not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			log.Printf("❌ Error fetching platform: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch platform",
			})
		}
	}

	log.Printf("✅ Successfully retrieved platform: %s", platform.Name)
	return c.JSON(platform)
}

// CreatePlatform adds a new platform
func (pc *PlatformController) CreatePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platform models.Platform
	if err := json.Unmarshal(c.Body(), &platform); err != nil {
		log.Printf("❌ Error parsing platform data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(platform); err != nil {
		log.Printf("❌ Error validating platform data: %v", err)
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	log.Printf("📝 Creating new platform: %s", platform.Name)
	if err := pc.service.CreatePlatform(ctx, &platform); err != nil {
		log.Printf("❌ Error creating platform: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create platform",
		})
	}

	log.Printf("✅ Successfully created platform: %s", platform.Name)
	return c.Status(201).JSON(platform)
}

// UpdatePlatform updates an existing platform
func (pc *PlatformController) UpdatePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updateData models.PlatformUpdate
	if err := json.Unmarshal(c.Body(), &updateData); err != nil {
		log.Printf("❌ Error parsing update data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(updateData); err != nil {
		log.Printf("❌ Error validating update data: %v", err)
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	id := c.Params("id")
	log.Printf("🔍 Updating platform with ID: %s", id)

	platform, err := pc.service.UpdatePlatform(ctx, id, &updateData)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			log.Printf("⚠️  Invalid platform ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrNoUpdateData:
			log.Printf("❌ No update data provided for platform with ID: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field (name or manufacturer) must be provided for update",
			})
		case services.ErrPlatformNotFound:
			log.Printf("❌ Platform not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			log.Printf("❌ Error updating platform: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update platform",
			})
		}
	}

	log.Printf("✅ Successfully updated platform: %s", platform.Name)
	return c.JSON(platform)
}

// DeletePlatform removes a platform by its ID
func (pc *PlatformController) DeletePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Deleting platform with ID: %s", id)

	platform, err := pc.service.DeletePlatform(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			log.Printf("⚠️  Invalid platform ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrPlatformNotFound:
			log.Printf("❌ Platform not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			log.Printf("❌ Error deleting platform: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete platform",
			})
		}
	}

	log.Printf("✅ Successfully deleted platform: %s", platform.Name)
	return c.Status(200).JSON(fiber.Map{
		"message": "Platform deleted successfully",
		"deleted": platform,
	})
}

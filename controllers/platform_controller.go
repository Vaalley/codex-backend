package controllers

import (
	"codex-backend/middleware"
	"codex-backend/models"
	"codex-backend/services"
	"codex-backend/utils"
	"context"
	"encoding/json"
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

	name := c.Query("name")
	platforms, err := pc.service.GetPlatforms(ctx, name)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch platforms",
		})
	}

	return c.JSON(platforms)
}

// GetPlatformByID returns a specific platform by its ID
func (pc *PlatformController) GetPlatformByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	platform, err := pc.service.GetPlatformByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrPlatformNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch platform",
			})
		}
	}

	return c.JSON(platform)
}

// CreatePlatform adds a new platform
func (pc *PlatformController) CreatePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platform models.Platform
	if err := json.Unmarshal(c.Body(), &platform); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(platform); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	if err := pc.service.CreatePlatform(ctx, &platform); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create platform",
		})
	}

	return c.Status(201).JSON(platform)
}

// UpdatePlatform updates an existing platform
func (pc *PlatformController) UpdatePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updateData models.PlatformUpdate
	if err := json.Unmarshal(c.Body(), &updateData); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(updateData); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	id := c.Params("id")
	platform, err := pc.service.UpdatePlatform(ctx, id, &updateData)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrNoUpdateData:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field (name or manufacturer) must be provided for update",
			})
		case services.ErrPlatformNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update platform",
			})
		}
	}

	return c.JSON(platform)
}

// DeletePlatform removes a platform by its ID
func (pc *PlatformController) DeletePlatform(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	platform, err := pc.service.DeletePlatform(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid platform ID",
			})
		case services.ErrPlatformNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Platform not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete platform",
			})
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Platform deleted successfully",
		"deleted": platform,
	})
}

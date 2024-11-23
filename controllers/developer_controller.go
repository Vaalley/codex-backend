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

	name := c.Query("name")
	developers, err := dc.service.GetDevelopers(ctx, name)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch developers",
		})
	}

	return c.JSON(developers)
}

// GetDeveloperByID returns a specific developer by its ID
func (dc *DeveloperController) GetDeveloperByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	developer, err := dc.service.GetDeveloperByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch developer",
			})
		}
	}

	return c.JSON(developer)
}

// CreateDeveloper adds a new developer
func (dc *DeveloperController) CreateDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var developer models.Developer
	if err := json.Unmarshal(c.Body(), &developer); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(developer); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	if err := dc.service.CreateDeveloper(ctx, &developer); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create developer",
		})
	}

	return c.Status(201).JSON(developer)
}

// UpdateDeveloper updates an existing developer
func (dc *DeveloperController) UpdateDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	var update models.DeveloperUpdate
	if err := json.Unmarshal(c.Body(), &update); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(update); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	developer, err := dc.service.UpdateDeveloper(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		case services.ErrNoDeveloperUpdateData:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field must be provided for update",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update developer",
			})
		}
	}

	return c.JSON(developer)
}

// DeleteDeveloper removes a developer by its ID
func (dc *DeveloperController) DeleteDeveloper(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	developer, err := dc.service.DeleteDeveloper(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidDeveloperID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid developer ID",
			})
		case services.ErrDeveloperNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Developer not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete developer",
			})
		}
	}

	return c.JSON(developer)
}

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

type PublisherController struct {
	service services.PublisherService
}

func NewPublisherController(service services.PublisherService) *PublisherController {
	return &PublisherController{
		service: service,
	}
}

// GetPublishers returns a list of all publishers or filters by name if a query parameter is provided
func (pc *PublisherController) GetPublishers(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := c.Query("name")
	publishers, err := pc.service.GetPublishers(ctx, name)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch publishers",
		})
	}

	return c.JSON(publishers)
}

// GetPublisherByID returns a specific publisher by its ID
func (pc *PublisherController) GetPublisherByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	publisher, err := pc.service.GetPublisherByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch publisher",
			})
		}
	}

	return c.JSON(publisher)
}

// CreatePublisher adds a new publisher
func (pc *PublisherController) CreatePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var publisher models.Publisher
	if err := json.Unmarshal(c.Body(), &publisher); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(publisher); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	if err := pc.service.CreatePublisher(ctx, &publisher); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create publisher",
		})
	}

	return c.Status(201).JSON(publisher)
}

// UpdatePublisher updates an existing publisher
func (pc *PublisherController) UpdatePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	var update models.PublisherUpdate
	if err := json.Unmarshal(c.Body(), &update); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(update); err != nil {
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	publisher, err := pc.service.UpdatePublisher(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		case services.ErrNoPublisherUpdateData:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field must be provided for update",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update publisher",
			})
		}
	}

	return c.JSON(publisher)
}

// DeletePublisher removes a publisher by its ID
func (pc *PublisherController) DeletePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	publisher, err := pc.service.DeletePublisher(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete publisher",
			})
		}
	}

	return c.JSON(publisher)
}

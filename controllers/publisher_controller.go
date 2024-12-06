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
	if name != "" {
		log.Printf("🔍 Searching for publishers with name containing: %s", name)
	} else {
		log.Println("📋 Fetching all publishers")
	}

	publishers, err := pc.service.GetPublishers(ctx, name)
	if err != nil {
		log.Printf("❌ Error fetching publishers: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch publishers",
		})
	}

	log.Printf("✅ Successfully retrieved %d publishers", len(publishers))
	return c.JSON(publishers)
}

// GetPublisherByID returns a specific publisher by its ID
func (pc *PublisherController) GetPublisherByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Fetching publisher with ID: %s", id)

	publisher, err := pc.service.GetPublisherByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			log.Printf("⚠️  Invalid publisher ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			log.Printf("❌ Publisher not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		default:
			log.Printf("❌ Error fetching publisher: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch publisher",
			})
		}
	}

	log.Printf("✅ Successfully retrieved publisher: %s", publisher.Name)
	return c.JSON(publisher)
}

// CreatePublisher adds a new publisher
func (pc *PublisherController) CreatePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var publisher models.Publisher
	if err := json.Unmarshal(c.Body(), &publisher); err != nil {
		log.Printf("❌ Error parsing publisher data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(publisher); err != nil {
		log.Printf("❌ Error validating publisher data: %v", err)
		return middleware.ErrorHandlerMiddleware(c, err)
	}

	log.Printf("📝 Creating new publisher: %s", publisher.Name)
	if err := pc.service.CreatePublisher(ctx, &publisher); err != nil {
		log.Printf("❌ Error creating publisher: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create publisher",
		})
	}

	log.Printf("✅ Successfully created publisher: %s", publisher.Name)
	return c.Status(201).JSON(publisher)
}

// UpdatePublisher updates an existing publisher
func (pc *PublisherController) UpdatePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Updating publisher with ID: %s", id)

	var update models.PublisherUpdate
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

	publisher, err := pc.service.UpdatePublisher(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			log.Printf("⚠️  Invalid publisher ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			log.Printf("❌ Publisher not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		case services.ErrNoPublisherUpdateData:
			log.Printf("❌ No update data provided for publisher with ID: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "At least one field must be provided for update",
			})
		default:
			log.Printf("❌ Error updating publisher: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update publisher",
			})
		}
	}

	log.Printf("✅ Successfully updated publisher: %s", publisher.Name)
	return c.JSON(publisher)
}

// DeletePublisher removes a publisher by its ID
func (pc *PublisherController) DeletePublisher(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Deleting publisher with ID: %s", id)

	publisher, err := pc.service.DeletePublisher(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidPublisherID:
			log.Printf("⚠️  Invalid publisher ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid publisher ID",
			})
		case services.ErrPublisherNotFound:
			log.Printf("❌ Publisher not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Publisher not found",
			})
		default:
			log.Printf("❌ Error deleting publisher: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete publisher",
			})
		}
	}

	log.Printf("✅ Successfully deleted publisher: %s", publisher.Name)
	return c.JSON(publisher)
}

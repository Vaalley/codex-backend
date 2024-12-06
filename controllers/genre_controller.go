package controllers

import (
	"codex-backend/models"
	"codex-backend/services"
	"codex-backend/utils"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

type GenreController struct {
	service services.GenreService
}

func NewGenreController(service services.GenreService) *GenreController {
	return &GenreController{
		service: service,
	}
}

// GetGenres returns a list of all genres or filters by name if a query parameter is provided
func (gc *GenreController) GetGenres(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("📋 Fetching genres")
	name := c.Query("name")
	if name != "" {
		log.Printf("🔍 Searching for genres with name containing: %s", name)
	}

	genres, err := gc.service.GetGenres(ctx, name)
	if err != nil {
		log.Printf("❌ Error fetching genres: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch genres",
		})
	}

	log.Printf("✅ Successfully retrieved %d genres", len(genres))
	return c.JSON(genres)
}

// GetGenreByID returns a specific genre by its ID
func (gc *GenreController) GetGenreByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Fetching genre")
	id := c.Params("id")
	log.Printf("🔍 Fetching genre with ID: %s", id)

	genre, err := gc.service.GetGenreByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			log.Printf("⚠️  Invalid genre ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			log.Printf("❌ Genre not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		default:
			log.Printf("❌ Error fetching genre: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch genre",
			})
		}
	}

	log.Printf("✅ Successfully retrieved genre: %s", genre.Name)
	return c.JSON(genre)
}

// CreateGenre adds a new genre
func (gc *GenreController) CreateGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("📝 Creating new genre")
	var genre models.Genre
	if err := json.Unmarshal(c.Body(), &genre); err != nil {
		log.Printf("❌ Error parsing genre data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(genre); err != nil {
		log.Printf("❌ Error validating genre data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	log.Printf("📝 Creating new genre: %s", genre.Name)
	if err := gc.service.CreateGenre(ctx, &genre); err != nil {
		log.Printf("❌ Error creating genre: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create genre",
		})
	}

	log.Printf("✅ Successfully created genre: %s", genre.Name)
	return c.Status(201).JSON(genre)
}

// UpdateGenre updates an existing genre
func (gc *GenreController) UpdateGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Updating genre")
	id := c.Params("id")
	log.Printf("🔍 Updating genre with ID: %s", id)

	var update models.GenreUpdate
	if err := json.Unmarshal(c.Body(), &update); err != nil {
		log.Printf("❌ Error parsing update data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(update); err != nil {
		log.Printf("❌ Error validating update data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	genre, err := gc.service.UpdateGenre(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			log.Printf("⚠️  Invalid genre ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			log.Printf("❌ Genre not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		case services.ErrNoGenreUpdateData:
			log.Printf("❌ No update data provided for genre with ID: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "No valid update data provided",
			})
		default:
			log.Printf("❌ Error updating genre: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update genre",
			})
		}
	}

	log.Printf("✅ Successfully updated genre: %s", genre.Name)
	return c.JSON(genre)
}

// DeleteGenre removes a genre by its ID
func (gc *GenreController) DeleteGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("🔍 Deleting genre")
	id := c.Params("id")
	log.Printf("🔍 Deleting genre with ID: %s", id)

	genre, err := gc.service.DeleteGenre(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			log.Printf("⚠️  Invalid genre ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			log.Printf("❌ Genre not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		default:
			log.Printf("❌ Error deleting genre: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete genre",
			})
		}
	}

	log.Printf("✅ Successfully deleted genre: %s", genre.Name)
	return c.JSON(genre)
}

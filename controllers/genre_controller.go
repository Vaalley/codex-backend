package controllers

import (
	"codex-backend/models"
	"codex-backend/services"
	"codex-backend/utils"
	"context"
	"encoding/json"
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

	name := c.Query("name")
	genres, err := gc.service.GetGenres(ctx, name)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch genres",
		})
	}

	return c.JSON(genres)
}

// GetGenreByID returns a specific genre by its ID
func (gc *GenreController) GetGenreByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	genre, err := gc.service.GetGenreByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch genre",
			})
		}
	}

	return c.JSON(genre)
}

// CreateGenre adds a new genre
func (gc *GenreController) CreateGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var genre models.Genre
	if err := json.Unmarshal(c.Body(), &genre); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(genre); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	if err := gc.service.CreateGenre(ctx, &genre); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create genre",
		})
	}

	return c.Status(201).JSON(genre)
}

// UpdateGenre updates an existing genre
func (gc *GenreController) UpdateGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	var update models.GenreUpdate
	if err := json.Unmarshal(c.Body(), &update); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(update); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	genre, err := gc.service.UpdateGenre(ctx, id, &update)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		case services.ErrNoGenreUpdateData:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "No valid update data provided",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update genre",
			})
		}
	}

	return c.JSON(genre)
}

// DeleteGenre removes a genre by its ID
func (gc *GenreController) DeleteGenre(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	genre, err := gc.service.DeleteGenre(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGenreID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid genre ID format",
			})
		case services.ErrGenreNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Genre not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete genre",
			})
		}
	}

	return c.JSON(genre)
}

package controllers

import (
	"codex-backend/models"
	"codex-backend/services"
	"codex-backend/utils"
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type GameController struct {
	service services.GameService
}

func NewGameController(service services.GameService) *GameController {
	return &GameController{
		service: service,
	}
}

// GetGames returns a list of all games or filters by title if a query parameter is provided
func (gc *GameController) GetGames(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	title := c.Query("title")
	games, err := gc.service.GetGames(ctx, title)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch games",
		})
	}

	return c.JSON(games)
}

// GetGameByID returns a specific game by its ID
func (gc *GameController) GetGameByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	game, err := gc.service.GetGameByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGameID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch game",
			})
		}
	}

	return c.JSON(game)
}

// CreateGame adds a new game
func (gc *GameController) CreateGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var game models.Game
	if err := json.Unmarshal(c.Body(), &game); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(game); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	if err := gc.service.CreateGame(ctx, &game); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create game",
		})
	}

	return c.Status(201).JSON(game)
}

// UpdateGame updates an existing game
func (gc *GameController) UpdateGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	var updateData map[string]interface{}
	if err := json.Unmarshal(c.Body(), &updateData); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	update := bson.M{}
	for key, value := range updateData {
		update[key] = value
	}

	game, err := gc.service.UpdateGame(ctx, id, update)
	if err != nil {
		switch err {
		case services.ErrInvalidGameID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		case services.ErrNoUpdateData:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "No update data provided",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update game",
			})
		}
	}

	return c.JSON(game)
}

// DeleteGame removes a game by its ID
func (gc *GameController) DeleteGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	game, err := gc.service.DeleteGame(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGameID:
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		default:
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete game",
			})
		}
	}

	return c.Status(200).JSON(game)
}

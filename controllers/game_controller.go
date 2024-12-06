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
	if title != "" {
		log.Printf("🔍 Searching for games with title containing: %s", title)
	} else {
		log.Println("📋 Fetching all games")
	}

	games, err := gc.service.GetGames(ctx, title)
	if err != nil {
		log.Printf("❌ Error fetching games: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to fetch games",
		})
	}

	log.Printf("✅ Successfully retrieved %d games", len(games))
	return c.JSON(games)
}

// GetGameByID returns a specific game by its ID
func (gc *GameController) GetGameByID(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Fetching game with ID: %s", id)

	game, err := gc.service.GetGameByID(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGameID:
			log.Printf("⚠️  Invalid game ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			log.Printf("❌ Game not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		default:
			log.Printf("❌ Error fetching game: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to fetch game",
			})
		}
	}

	log.Printf("✅ Successfully retrieved game: %s", game.Title)
	return c.JSON(game)
}

// CreateGame adds a new game
func (gc *GameController) CreateGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var game models.Game
	if err := json.Unmarshal(c.Body(), &game); err != nil {
		log.Printf("❌ Error parsing game data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
		})
	}

	if err := utils.ValidateStruct(game); err != nil {
		log.Printf("❌ Error validating game data: %v", err)
		return c.Status(400).JSON(models.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	log.Printf("📝 Creating new game: %s", game.Title)
	if err := gc.service.CreateGame(ctx, &game); err != nil {
		log.Printf("❌ Error creating game: %v", err)
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  500,
			Message: "Failed to create game",
		})
	}

	log.Printf("✅ Successfully created game: %s", game.Title)
	return c.Status(201).JSON(game)
}

// UpdateGame updates an existing game
func (gc *GameController) UpdateGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Updating game with ID: %s", id)

	var updateData map[string]interface{}
	if err := json.Unmarshal(c.Body(), &updateData); err != nil {
		log.Printf("❌ Error parsing update data: %v", err)
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
			log.Printf("⚠️  Invalid game ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			log.Printf("❌ Game not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		case services.ErrNoUpdateData:
			log.Printf("❌ No update data provided for game with ID: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "No update data provided",
			})
		default:
			log.Printf("❌ Error updating game: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to update game",
			})
		}
	}

	log.Printf("✅ Successfully updated game: %s", game.Title)
	return c.JSON(game)
}

// DeleteGame removes a game by its ID
func (gc *GameController) DeleteGame(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("🔍 Deleting game with ID: %s", id)

	game, err := gc.service.DeleteGame(ctx, id)
	if err != nil {
		switch err {
		case services.ErrInvalidGameID:
			log.Printf("⚠️  Invalid game ID format: %s", id)
			return c.Status(400).JSON(models.ErrorResponse{
				Status:  400,
				Message: "Invalid game ID format",
			})
		case services.ErrGameNotFound:
			log.Printf("❌ Game not found with ID: %s", id)
			return c.Status(404).JSON(models.ErrorResponse{
				Status:  404,
				Message: "Game not found",
			})
		default:
			log.Printf("❌ Error deleting game: %v", err)
			return c.Status(500).JSON(models.ErrorResponse{
				Status:  500,
				Message: "Failed to delete game",
			})
		}
	}

	log.Printf("✅ Successfully deleted game: %s", game.Title)
	return c.Status(200).JSON(game)
}

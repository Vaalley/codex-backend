package services

import (
	"codex-backend/models"
	"codex-backend/repositories"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrGameNotFound = errors.New("game not found")
	ErrInvalidGameID = errors.New("invalid game ID")
)

type GameService interface {
	GetGames(ctx context.Context, title string) ([]models.Game, error)
	GetGameByID(ctx context.Context, id string) (*models.Game, error)
	CreateGame(ctx context.Context, game *models.Game) error
	UpdateGame(ctx context.Context, id string, update bson.M) (*models.Game, error)
	DeleteGame(ctx context.Context, id string) (*models.Game, error)
}

type gameService struct {
	repo repositories.GameRepository
}

func NewGameService(repo repositories.GameRepository) GameService {
	return &gameService{
		repo: repo,
	}
}

func (s *gameService) GetGames(ctx context.Context, title string) ([]models.Game, error) {
	filter := bson.M{}
	if title != "" {
		filter = bson.M{"title": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: title, Options: "i"}}}}
	}

	return s.repo.FindAll(ctx, filter)
}

func (s *gameService) GetGameByID(ctx context.Context, id string) (*models.Game, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGameID
	}

	game, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}

func (s *gameService) CreateGame(ctx context.Context, game *models.Game) error {
	return s.repo.Create(ctx, game)
}

func (s *gameService) UpdateGame(ctx context.Context, id string, update bson.M) (*models.Game, error) {
	if len(update) == 0 {
		return nil, ErrNoUpdateData
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGameID
	}

	game, err := s.repo.Update(ctx, objID, update)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}

func (s *gameService) DeleteGame(ctx context.Context, id string) (*models.Game, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGameID
	}

	game, err := s.repo.Delete(ctx, objID)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}

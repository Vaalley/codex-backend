package services

import (
	"codex-backend/models"
	"codex-backend/repositories"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrGenreNotFound = errors.New("genre not found")
	ErrInvalidGenreID = errors.New("invalid genre ID")
	ErrNoGenreUpdateData = errors.New("at least one field must be provided for update")
)

type GenreService interface {
	GetGenres(ctx context.Context, name string) ([]models.Genre, error)
	GetGenreByID(ctx context.Context, id string) (*models.Genre, error)
	CreateGenre(ctx context.Context, genre *models.Genre) error
	UpdateGenre(ctx context.Context, id string, update *models.GenreUpdate) (*models.Genre, error)
	DeleteGenre(ctx context.Context, id string) (*models.Genre, error)
}

type genreService struct {
	repo repositories.GenreRepository
}

func NewGenreService(repo repositories.GenreRepository) GenreService {
	return &genreService{
		repo: repo,
	}
}

func (s *genreService) GetGenres(ctx context.Context, name string) ([]models.Genre, error) {
	filter := bson.M{}
	if name != "" {
		filter = bson.M{"name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: name, Options: "i"}}}}
	}

	return s.repo.FindAll(ctx, filter)
}

func (s *genreService) GetGenreByID(ctx context.Context, id string) (*models.Genre, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGenreID
	}

	genre, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if genre == nil {
		return nil, ErrGenreNotFound
	}

	return genre, nil
}

func (s *genreService) CreateGenre(ctx context.Context, genre *models.Genre) error {
	genre.CreatedAt = time.Now()
	genre.UpdatedAt = time.Now()
	return s.repo.Create(ctx, genre)
}

func (s *genreService) UpdateGenre(ctx context.Context, id string, update *models.GenreUpdate) (*models.Genre, error) {
	if update == nil || (update.Name == nil) {
		return nil, ErrNoGenreUpdateData
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGenreID
	}

	updateData := bson.M{
		"updatedAt": time.Now(),
	}
	if update.Name != nil {
		updateData["name"] = *update.Name
	}

	genre, err := s.repo.Update(ctx, objID, updateData)
	if err != nil {
		return nil, err
	}
	if genre == nil {
		return nil, ErrGenreNotFound
	}

	return genre, nil
}

func (s *genreService) DeleteGenre(ctx context.Context, id string) (*models.Genre, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidGenreID
	}

	genre, err := s.repo.Delete(ctx, objID)
	if err != nil {
		return nil, err
	}
	if genre == nil {
		return nil, ErrGenreNotFound
	}

	return genre, nil
}

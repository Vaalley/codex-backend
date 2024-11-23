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
	ErrDeveloperNotFound     = errors.New("developer not found")
	ErrInvalidDeveloperID    = errors.New("invalid developer ID")
	ErrNoDeveloperUpdateData = errors.New("at least one field must be provided for update")
)

type DeveloperService interface {
	GetDevelopers(ctx context.Context, name string) ([]models.Developer, error)
	GetDeveloperByID(ctx context.Context, id string) (*models.Developer, error)
	CreateDeveloper(ctx context.Context, developer *models.Developer) error
	UpdateDeveloper(ctx context.Context, id string, update *models.DeveloperUpdate) (*models.Developer, error)
	DeleteDeveloper(ctx context.Context, id string) (*models.Developer, error)
}

type developerService struct {
	repo repositories.DeveloperRepository
}

func NewDeveloperService(repo repositories.DeveloperRepository) DeveloperService {
	return &developerService{
		repo: repo,
	}
}

func (s *developerService) GetDevelopers(ctx context.Context, name string) ([]models.Developer, error) {
	filter := bson.M{}
	if name != "" {
		filter = bson.M{"name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: name, Options: "i"}}}}
	}

	return s.repo.FindAll(ctx, filter)
}

func (s *developerService) GetDeveloperByID(ctx context.Context, id string) (*models.Developer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidDeveloperID
	}

	developer, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if developer == nil {
		return nil, ErrDeveloperNotFound
	}

	return developer, nil
}

func (s *developerService) CreateDeveloper(ctx context.Context, developer *models.Developer) error {
	developer.ID = primitive.NewObjectID()
	developer.CreatedAt = time.Now()
	developer.UpdatedAt = time.Now()
	return s.repo.Create(ctx, developer)
}

func (s *developerService) UpdateDeveloper(ctx context.Context, id string, update *models.DeveloperUpdate) (*models.Developer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidDeveloperID
	}

	if update.Name == nil {
		return nil, ErrNoDeveloperUpdateData
	}

	updateDoc := bson.M{"$set": bson.M{
		"updatedAt": time.Now(),
	}}
	if update.Name != nil {
		updateDoc["$set"].(bson.M)["name"] = *update.Name
	}

	developer, err := s.repo.Update(ctx, objID, updateDoc)
	if err != nil {
		return nil, err
	}
	if developer == nil {
		return nil, ErrDeveloperNotFound
	}

	return developer, nil
}

func (s *developerService) DeleteDeveloper(ctx context.Context, id string) (*models.Developer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidDeveloperID
	}

	developer, err := s.repo.Delete(ctx, objID)
	if err != nil {
		return nil, err
	}
	if developer == nil {
		return nil, ErrDeveloperNotFound
	}

	return developer, nil
}

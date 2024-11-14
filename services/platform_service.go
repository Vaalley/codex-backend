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
	ErrPlatformNotFound = errors.New("platform not found")
	ErrInvalidID        = errors.New("invalid platform ID")
	ErrNoUpdateData     = errors.New("at least one field must be provided for update")
)

type PlatformService interface {
	GetPlatforms(ctx context.Context, name string) ([]models.Platform, error)
	GetPlatformByID(ctx context.Context, id string) (*models.Platform, error)
	CreatePlatform(ctx context.Context, platform *models.Platform) error
	UpdatePlatform(ctx context.Context, id string, update *models.PlatformUpdate) (*models.Platform, error)
	DeletePlatform(ctx context.Context, id string) (*models.Platform, error)
}

type platformService struct {
	repo repositories.PlatformRepository
}

func NewPlatformService(repo repositories.PlatformRepository) PlatformService {
	return &platformService{
		repo: repo,
	}
}

func (s *platformService) GetPlatforms(ctx context.Context, name string) ([]models.Platform, error) {
	filter := bson.M{}
	if name != "" {
		filter = bson.M{"name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: name, Options: "i"}}}}
	}

	return s.repo.FindAll(ctx, filter)
}

func (s *platformService) GetPlatformByID(ctx context.Context, id string) (*models.Platform, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	platform, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if platform == nil {
		return nil, ErrPlatformNotFound
	}

	return platform, nil
}

func (s *platformService) CreatePlatform(ctx context.Context, platform *models.Platform) error {
	platform.ID = primitive.NewObjectID()
	return s.repo.Create(ctx, platform)
}

func (s *platformService) UpdatePlatform(ctx context.Context, id string, update *models.PlatformUpdate) (*models.Platform, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	if update.Name == nil && update.Manufacturer == nil {
		return nil, ErrNoUpdateData
	}

	updateDoc := bson.M{"$set": bson.M{}}
	if update.Name != nil {
		updateDoc["$set"].(bson.M)["name"] = *update.Name
	}
	if update.Manufacturer != nil {
		updateDoc["$set"].(bson.M)["manufacturer"] = *update.Manufacturer
	}

	platform, err := s.repo.Update(ctx, objID, updateDoc)
	if err != nil {
		return nil, err
	}
	if platform == nil {
		return nil, ErrPlatformNotFound
	}

	return platform, nil
}

func (s *platformService) DeletePlatform(ctx context.Context, id string) (*models.Platform, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	platform, err := s.repo.Delete(ctx, objID)
	if err != nil {
		return nil, err
	}
	if platform == nil {
		return nil, ErrPlatformNotFound
	}

	return platform, nil
}

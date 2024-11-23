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
	ErrPublisherNotFound     = errors.New("publisher not found")
	ErrInvalidPublisherID    = errors.New("invalid publisher ID")
	ErrNoPublisherUpdateData = errors.New("at least one field must be provided for update")
)

type PublisherService interface {
	GetPublishers(ctx context.Context, name string) ([]models.Publisher, error)
	GetPublisherByID(ctx context.Context, id string) (*models.Publisher, error)
	CreatePublisher(ctx context.Context, publisher *models.Publisher) error
	UpdatePublisher(ctx context.Context, id string, update *models.PublisherUpdate) (*models.Publisher, error)
	DeletePublisher(ctx context.Context, id string) (*models.Publisher, error)
}

type publisherService struct {
	repo repositories.PublisherRepository
}

func NewPublisherService(repo repositories.PublisherRepository) PublisherService {
	return &publisherService{
		repo: repo,
	}
}

func (s *publisherService) GetPublishers(ctx context.Context, name string) ([]models.Publisher, error) {
	filter := bson.M{}
	if name != "" {
		filter = bson.M{"name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: name, Options: "i"}}}}
	}

	return s.repo.FindAll(ctx, filter)
}

func (s *publisherService) GetPublisherByID(ctx context.Context, id string) (*models.Publisher, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidPublisherID
	}

	publisher, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if publisher == nil {
		return nil, ErrPublisherNotFound
	}

	return publisher, nil
}

func (s *publisherService) CreatePublisher(ctx context.Context, publisher *models.Publisher) error {
	publisher.ID = primitive.NewObjectID()
	publisher.CreatedAt = time.Now()
	publisher.UpdatedAt = time.Now()
	return s.repo.Create(ctx, publisher)
}

func (s *publisherService) UpdatePublisher(ctx context.Context, id string, update *models.PublisherUpdate) (*models.Publisher, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidPublisherID
	}

	if update.Name == nil {
		return nil, ErrNoPublisherUpdateData
	}

	updateDoc := bson.M{"$set": bson.M{
		"updatedAt": time.Now(),
	}}
	if update.Name != nil {
		updateDoc["$set"].(bson.M)["name"] = *update.Name
	}

	publisher, err := s.repo.Update(ctx, objID, updateDoc)
	if err != nil {
		return nil, err
	}
	if publisher == nil {
		return nil, ErrPublisherNotFound
	}

	return publisher, nil
}

func (s *publisherService) DeletePublisher(ctx context.Context, id string) (*models.Publisher, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidPublisherID
	}

	publisher, err := s.repo.Delete(ctx, objID)
	if err != nil {
		return nil, err
	}
	if publisher == nil {
		return nil, ErrPublisherNotFound
	}

	return publisher, nil
}

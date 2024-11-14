package repositories

import (
	"codex-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlatformRepository interface {
	FindAll(ctx context.Context, filter bson.M) ([]models.Platform, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Platform, error)
	Create(ctx context.Context, platform *models.Platform) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Platform, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*models.Platform, error)
}

type mongoPlatformRepository struct {
	collection *mongo.Collection
}

func NewMongoPlatformRepository(collection *mongo.Collection) PlatformRepository {
	return &mongoPlatformRepository{
		collection: collection,
	}
}

func (r *mongoPlatformRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Platform, error) {
	var platforms []models.Platform
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &platforms); err != nil {
		return nil, err
	}
	return platforms, nil
}

func (r *mongoPlatformRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Platform, error) {
	var platform models.Platform
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&platform)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &platform, nil
}

func (r *mongoPlatformRepository) Create(ctx context.Context, platform *models.Platform) error {
	_, err := r.collection.InsertOne(ctx, platform)
	return err
}

func (r *mongoPlatformRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Platform, error) {
	var platform models.Platform
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		update,
		opts,
	).Decode(&platform)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &platform, nil
}

func (r *mongoPlatformRepository) Delete(ctx context.Context, id primitive.ObjectID) (*models.Platform, error) {
	var platform models.Platform
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&platform)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &platform, nil
}

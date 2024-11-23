package repositories

import (
	"codex-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeveloperRepository interface {
	FindAll(ctx context.Context, filter bson.M) ([]models.Developer, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Developer, error)
	Create(ctx context.Context, developer *models.Developer) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Developer, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*models.Developer, error)
}

type mongoDeveloperRepository struct {
	collection *mongo.Collection
}

func NewMongoDeveloperRepository(collection *mongo.Collection) DeveloperRepository {
	return &mongoDeveloperRepository{
		collection: collection,
	}
}

func (r *mongoDeveloperRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Developer, error) {
	var developers []models.Developer
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &developers); err != nil {
		return nil, err
	}
	return developers, nil
}

func (r *mongoDeveloperRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Developer, error) {
	var developer models.Developer
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&developer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &developer, nil
}

func (r *mongoDeveloperRepository) Create(ctx context.Context, developer *models.Developer) error {
	_, err := r.collection.InsertOne(ctx, developer)
	return err
}

func (r *mongoDeveloperRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Developer, error) {
	var developer models.Developer
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		update,
		opts,
	).Decode(&developer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &developer, nil
}

func (r *mongoDeveloperRepository) Delete(ctx context.Context, id primitive.ObjectID) (*models.Developer, error) {
	var developer models.Developer
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&developer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &developer, nil
}

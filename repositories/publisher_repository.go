package repositories

import (
	"codex-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PublisherRepository interface {
	FindAll(ctx context.Context, filter bson.M) ([]models.Publisher, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Publisher, error)
	Create(ctx context.Context, publisher *models.Publisher) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Publisher, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*models.Publisher, error)
}

type mongoPublisherRepository struct {
	collection *mongo.Collection
}

func NewMongoPublisherRepository(collection *mongo.Collection) PublisherRepository {
	return &mongoPublisherRepository{
		collection: collection,
	}
}

func (r *mongoPublisherRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Publisher, error) {
	var publishers []models.Publisher
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &publishers); err != nil {
		return nil, err
	}
	return publishers, nil
}

func (r *mongoPublisherRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Publisher, error) {
	var publisher models.Publisher
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&publisher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &publisher, nil
}

func (r *mongoPublisherRepository) Create(ctx context.Context, publisher *models.Publisher) error {
	_, err := r.collection.InsertOne(ctx, publisher)
	return err
}

func (r *mongoPublisherRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Publisher, error) {
	var publisher models.Publisher
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		update,
		opts,
	).Decode(&publisher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &publisher, nil
}

func (r *mongoPublisherRepository) Delete(ctx context.Context, id primitive.ObjectID) (*models.Publisher, error) {
	var publisher models.Publisher
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&publisher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &publisher, nil
}

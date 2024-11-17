package repositories

import (
	"codex-backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameRepository interface {
	FindAll(ctx context.Context, filter bson.M) ([]models.Game, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Game, error)
	Create(ctx context.Context, game *models.Game) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Game, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*models.Game, error)
}

type mongoGameRepository struct {
	collection *mongo.Collection
}

func NewMongoGameRepository(collection *mongo.Collection) GameRepository {
	return &mongoGameRepository{
		collection: collection,
	}
}

func (r *mongoGameRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Game, error) {
	var games []models.Game
	opts := options.Find().SetSort(bson.D{{Key: "title", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (r *mongoGameRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Game, error) {
	var game models.Game
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&game)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &game, nil
}

func (r *mongoGameRepository) Create(ctx context.Context, game *models.Game) error {
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, game)
	if err != nil {
		return err
	}
	game.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *mongoGameRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Game, error) {
	update["updatedAt"] = time.Now()
	after := options.After
	opts := options.FindOneAndUpdate().SetReturnDocument(after)
	
	result := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
		opts,
	)

	var updatedGame models.Game
	if err := result.Decode(&updatedGame); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &updatedGame, nil
}

func (r *mongoGameRepository) Delete(ctx context.Context, id primitive.ObjectID) (*models.Game, error) {
	var deletedGame models.Game
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&deletedGame)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &deletedGame, nil
}

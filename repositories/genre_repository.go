package repositories

import (
	"codex-backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenreRepository interface {
	FindAll(ctx context.Context, filter bson.M) ([]models.Genre, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Genre, error)
	Create(ctx context.Context, genre *models.Genre) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Genre, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*models.Genre, error)
}

type mongoGenreRepository struct {
	collection *mongo.Collection
}

func NewMongoGenreRepository(collection *mongo.Collection) GenreRepository {
	return &mongoGenreRepository{
		collection: collection,
	}
}

func (r *mongoGenreRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Genre, error) {
	var genres []models.Genre
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &genres); err != nil {
		return nil, err
	}
	return genres, nil
}

func (r *mongoGenreRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Genre, error) {
	var genre models.Genre
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&genre)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &genre, nil
}

func (r *mongoGenreRepository) Create(ctx context.Context, genre *models.Genre) error {
	if genre.ID.IsZero() {
		genre.ID = primitive.NewObjectID()
	}
	_, err := r.collection.InsertOne(ctx, genre)
	return err
}

func (r *mongoGenreRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*models.Genre, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedGenre models.Genre
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
		opts,
	).Decode(&updatedGenre)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &updatedGenre, nil
}

func (r *mongoGenreRepository) Delete(ctx context.Context, id primitive.ObjectID) (*models.Genre, error) {
	var deletedGenre models.Genre
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&deletedGenre)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &deletedGenre, nil
}

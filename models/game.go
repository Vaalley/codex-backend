package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game represents a video game
type Game struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title" bson:"title" validate:"required,min=1,max=100"`
	Description string               `json:"description" bson:"description" validate:"required,min=1"`
	Developers  []primitive.ObjectID `json:"developers" bson:"developers" validate:"required,min=1"`
	Publishers  []primitive.ObjectID `json:"publishers" bson:"publishers" validate:"required,min=1"`
	ReleaseDate time.Time            `json:"releaseDate" bson:"releaseDate" validate:"required"`
	Platforms   []primitive.ObjectID `json:"platforms" bson:"platforms" validate:"required,min=1"`
	Genres      []primitive.ObjectID `json:"genres" bson:"genres" validate:"required,min=1"`
	CoverImage  string               `json:"coverImage" bson:"coverImage" validate:"required"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt" bson:"updatedAt"`
}

// GameUpdate represents the fields that can be updated for a game
type GameUpdate struct {
	Title       *string              `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,min=1,max=100"`
	Description *string              `json:"description,omitempty" bson:"description,omitempty" validate:"omitempty,min=1"`
	Developers  []primitive.ObjectID `json:"developers,omitempty" bson:"developers,omitempty" validate:"omitempty,min=1"`
	Publishers  []primitive.ObjectID `json:"publishers,omitempty" bson:"publishers,omitempty" validate:"omitempty,min=1"`
	ReleaseDate *time.Time           `json:"releaseDate,omitempty" bson:"releaseDate,omitempty"`
	Platforms   []primitive.ObjectID `json:"platforms,omitempty" bson:"platforms,omitempty" validate:"omitempty,min=1"`
	Genres      []primitive.ObjectID `json:"genres,omitempty" bson:"genres,omitempty" validate:"omitempty,min=1"`
}

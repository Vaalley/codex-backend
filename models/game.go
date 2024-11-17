package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game represents a video game
type Game struct {
	ID              primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Title           string               `json:"title" bson:"title" validate:"required,min=1,max=100"`
	Description     string               `json:"description" bson:"description" validate:"required,min=1"`
	Developer       string               `json:"developer" bson:"developer" validate:"required,min=2,max=50"`
	Publisher       string               `json:"publisher" bson:"publisher" validate:"required,min=2,max=50"`
	ReleaseDate     time.Time            `json:"releaseDate" bson:"releaseDate" validate:"required"`
	Platforms       []primitive.ObjectID `json:"platforms" bson:"platforms" validate:"required,min=1"`
	Genre           []string             `json:"genre" bson:"genre" validate:"required,min=1"`
	CoverImage      string               `json:"coverImage" bson:"coverImage" validate:"required"`
	MetacriticScore float64              `json:"metacriticScore" bson:"metacriticScore" validate:"min=0,max=100"`
	UsersScore      float64              `json:"usersScore" bson:"usersScore" validate:"min=0,max=100"`
	CreatedAt       time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time            `json:"updatedAt" bson:"updatedAt"`
}

// GameUpdate represents the fields that can be updated for a game
type GameUpdate struct {
	Title           *string              `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,min=1,max=100"`
	Description     *string              `json:"description,omitempty" bson:"description,omitempty" validate:"omitempty,min=1"`
	Developer       *string              `json:"developer,omitempty" bson:"developer,omitempty" validate:"omitempty,min=2,max=50"`
	Publisher       *string              `json:"publisher,omitempty" bson:"publisher,omitempty" validate:"omitempty,min=2,max=50"`
	ReleaseDate     *time.Time           `json:"releaseDate,omitempty" bson:"releaseDate,omitempty"`
	Platforms       []primitive.ObjectID `json:"platforms,omitempty" bson:"platforms,omitempty" validate:"omitempty,min=1"`
	Genre           []string             `json:"genre,omitempty" bson:"genre,omitempty" validate:"omitempty,min=1"`
	MetacriticScore *float64             `json:"metacriticScore,omitempty" bson:"metacriticScore,omitempty" validate:"omitempty,min=0,max=100"`
	UsersScore      *float64             `json:"usersScore,omitempty" bson:"usersScore,omitempty" validate:"omitempty,min=0,max=100"`
}

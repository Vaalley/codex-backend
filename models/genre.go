package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Genre represents a genre (e.g., Action, Adventure)
type Genre struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,min=2,max=50"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// GenreUpdate represents the fields that can be updated for a genre
type GenreUpdate struct {
	Name *string `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,min=2,max=50"`
}

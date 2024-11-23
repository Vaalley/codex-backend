package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Developer represents a game development company
type Developer struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeveloperUpdate represents the fields that can be updated for a developer
type DeveloperUpdate struct {
	Name *string `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,min=2,max=100"`
}

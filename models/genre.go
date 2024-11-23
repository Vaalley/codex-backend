package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Genre represents a gaming genre (e.g., PlayStation, Xbox)
type Genre struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" validate:"required,min=2,max=50"`
}

// GenreUpdate represents the fields that can be updated for a genre
type GenreUpdate struct {
	Name *string `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,min=2,max=50"`
}

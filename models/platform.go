package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Platform represents a gaming platform (e.g., PlayStation, Xbox)
type Platform struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" validate:"required"`
	Manufacturer string             `json:"manufacturer" bson:"manufacturer" validate:"required"`
	Type         string             `json:"type" bson:"type"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Platform represents a gaming platform (e.g., PlayStation, Xbox)
type Platform struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" validate:"required,min=2,max=50"`
	Manufacturer string             `json:"manufacturer" bson:"manufacturer" validate:"required,min=2,max=50"`
}

// PlatformUpdate represents the fields that can be updated for a platform
type PlatformUpdate struct {
	Name         *string `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Manufacturer *string `json:"manufacturer,omitempty" bson:"manufacturer,omitempty" validate:"omitempty,min=2,max=50"`
}

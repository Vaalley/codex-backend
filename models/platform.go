package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Platform represents a gaming platform (e.g., PlayStation, Xbox)
type Platform struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" validate:"required,min=2,max=50"`
	Manufacturer string             `json:"manufacturer" bson:"manufacturer" validate:"required,min=2,max=50"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// PlatformUpdate represents the fields that can be updated for a platform
type PlatformUpdate struct {
	Name         *string `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Manufacturer *string `json:"manufacturer,omitempty" bson:"manufacturer,omitempty" validate:"omitempty,min=2,max=50"`
}

package models

import "time"

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Username  string    `bson:"username" validate:"required,min=3,max=32"`
	Email     string    `bson:"email" validate:"required,email"`
	Password  string    `bson:"password" validate:"required,min=8"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

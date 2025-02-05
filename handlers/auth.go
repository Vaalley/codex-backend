package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/vaalley/codex-backend/config"
	"github.com/vaalley/codex-backend/db"
	"github.com/vaalley/codex-backend/models"
	"github.com/vaalley/codex-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register handles the registration of a new user
func Register(c fiber.Ctx) error {
	// Parse request
	var req models.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if user exists
	collection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check email and username uniqueness
	var existing models.User
	err := collection.FindOne(ctx, bson.M{"$or": []bson.M{
		{"email": req.Email},
		{"username": req.Username},
	}}).Decode(&existing)

	if err == nil {
		return c.Status(409).JSON(fiber.Map{"error": "Email or username already exists"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Create user
	newUser := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Database error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":       result.InsertedID,
		"username": newUser.Username,
		"email":    newUser.Email,
	})
}

// Login handles the login of a user
func Login(c fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Find user by email
	collection := db.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT with user claims
	token, err := utils.GenerateJWT(user.ID, []string{"user"})
	if err != nil {
		log.Printf("JWT generation error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Set HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   config.IsProduction(),
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Logout handles the logout of a user
func Logout(c fiber.Ctx) error {
	// Clear session cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   config.IsProduction(),
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{"message": "Successfully logged out"})
}

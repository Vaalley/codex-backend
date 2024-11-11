package controllers

import (
	"codex-backend/db"
	"codex-backend/models"
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const platformCollectionName = "platforms"

// returns a list of all platforms
func GetPlatforms(c fiber.Ctx) error {
	collection := db.GetCollection(platformCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platforms []models.Platform
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := cursor.All(ctx, &platforms); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(platforms)
}

// returns a specific platform by its ID
func GetPlatformByID(c fiber.Ctx) error {
	collection := db.GetCollection(platformCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platform models.Platform
	id := c.Query("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid platform ID")
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&platform)
	if err != nil {
		return c.Status(404).SendString("Platform not found")
	}
	return c.JSON(platform)
}

// returns a specific platform by its name
func GetPlatformByName(c fiber.Ctx) error {
	collection := db.GetCollection(platformCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platform models.Platform
	name := c.Query("name")
	err := collection.FindOne(ctx, bson.M{"name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: name, Options: "i"}}}}).Decode(&platform)
	if err != nil {
		return c.Status(404).SendString("Platform not found")
	}
	return c.JSON(platform)
}

// adds a new platform
func CreatePlatform(c fiber.Ctx) error {
	collection := db.GetCollection(platformCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var platform models.Platform
	if err := json.Unmarshal(c.Body(), &platform); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	platform.ID = primitive.NewObjectID()
	platform.Type = "platform"
	_, err := collection.InsertOne(ctx, platform)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(platform)
}

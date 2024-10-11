package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// Platform represents a gaming platform
type Platform struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" validate:"required"`
	Name         string             `json:"name" validate:"required,min=3,max=50"`
	Manufacturer string             `json:"manufacturer" validate:"required,min=3,max=50"`
	Type         string             `json:"type" bson:"type" validate:"required"`
}

func main() {
	validate := validator.New()
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up MongoDB connection
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.bqe43.mongodb.net/?retryWrites=true&writeConcern=majority",
		os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"))
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Create a new Fiber app
	app := fiber.New()

	// API endpoint to make sure the api is working and running
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Hello World"})
	})

	// API endpoint to get all platforms
	app.Get("/api/get-platforms", func(c fiber.Ctx) error {
		platforms, err := getPlatforms(client)
		if err != nil {
			log.Printf("Error retrieving platforms: %v", err)
			return c.Status(500).JSON(map[string]string{"message": "Failed to retrieve platforms"})
		}
		return c.JSON(platforms)
	})

	// API endpoint to get a single platform by name
	app.Post("/api/get-platform-by-name", func(c fiber.Ctx) error {
		var request struct {
			Name string `json:"name"`
		}

		if err := json.Unmarshal(c.Body(), &request); err != nil {
			log.Printf("Error parsing request body: %v", err)
			return c.Status(400).JSON(map[string]string{"message": "Invalid request body"})
		}

		if err := validate.Struct(request); err != nil {
			log.Printf("Validation error: %v", err)
			return c.Status(400).JSON(map[string]string{"message": "Invalid request body"})
		}

		platform, err := getPlatformByName(client, request.Name)
		if err != nil {
			log.Printf("Error retrieving platform by name: %v", err)
			return c.Status(500).JSON(map[string]string{"message": "Failed to retrieve platform"})
		}

		return c.JSON(platform)
	})

	// API endpoint to get a single platform by ID
	app.Post("/api/get-platform-by-id", func(c fiber.Ctx) error {
		var request struct {
			ID string `json:"id"`
		}

		if err := json.Unmarshal(c.Body(), &request); err != nil {
			log.Printf("Error parsing request body: %v", err)
			return c.Status(400).JSON(map[string]string{"message": "Invalid request body"})
		}

		platform, err := getPlatformById(client, request.ID)
		if err != nil {
			log.Printf("Error retrieving platform by ID: %v", err)
			return c.Status(500).JSON(map[string]string{"message": "Failed to retrieve platform"})
		}

		return c.JSON(platform)
	})

	// API endpoint to add a new platform
	app.Post("/api/add-platform", func(c fiber.Ctx) error {
		var platform Platform
		if err := json.Unmarshal(c.Body(), &platform); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		platform.ID = primitive.NewObjectID()
		platform.Type = "platform"

		if err := validate.Struct(platform); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			var validationError string
			for _, e := range validationErrors {
				validationError += fmt.Sprintf(
					"%s should match %s %s\n",
					e.Field(),
					e.Tag(),
					e.Param(),
				)
			}
			return c.Status(400).JSON(fiber.Map{"error": validationError})
		}

		err := addPlatform(client, platform)

		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return c.Status(400).JSON(fiber.Map{"error": "Platform already exists"})
			}
			return c.Status(500).JSON(fiber.Map{"error": "Failed to add platform"})
		}
		return c.JSON(platform)
	})

	// API endpoint to update a platform
	app.Post("/api/update-platform", func(c fiber.Ctx) error {
		var request struct {
			ID           string `json:"id"`
			Name         string `json:"name,omitempty"`
			Manufacturer string `json:"manufacturer,omitempty"`
		}

		if err := json.Unmarshal(c.Body(), &request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Validate ID
		objectID, err := primitive.ObjectIDFromHex(request.ID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		// Create update document
		update := bson.M{}
		if request.Name != "" {
			update["name"] = request.Name
		}
		if request.Manufacturer != "" {
			update["manufacturer"] = request.Manufacturer
		}

		platform := Platform{
			ID:           objectID,
			Name:         request.Name,
			Manufacturer: request.Manufacturer,
			Type:         "platform",
		}

		err = updatePlatform(client, platform)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Platform with the same name already exists"})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update platform"})
		}

		return c.JSON(platform)
	})

	// API endpoint to delete a platform
	app.Post("/api/delete-platform", func(c fiber.Ctx) error {
		var request struct {
			ID string `json:"id"`
		}

		if err := json.Unmarshal(c.Body(), &request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Delete platform
		err = deletePlatform(client, request.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete platform"})
		}

		return c.JSON(fiber.Map{"message": "Platform deleted successfully"})
	})
	// Start the Fiber app
	log.Fatal(app.Listen("localhost:3000"))
}

// getPlatforms retrieves all platforms from the database
func getPlatforms(client *mongo.Client) ([]Platform, error) {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Define the filter to fetch only platforms
	filter := bson.D{{Key: "type", Value: "platform"}}

	// Find all platforms
	var platforms []Platform
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	// Iterate over the results
	for cur.Next(context.TODO()) {
		var platform Platform
		if err := cur.Decode(&platform); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}

	// Check for any errors that occurred during the iteration
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return platforms, nil
}

// getPlatformByName retrieves a single platform by name from the database
func getPlatformByName(client *mongo.Client, name string) (Platform, error) {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Define the filter to fetch a platform by name
	filter := bson.D{{Key: "name", Value: name}, {Key: "type", Value: "platform"}}

	// Find the platform
	var platform Platform
	err := collection.FindOne(context.TODO(), filter).Decode(&platform)
	if err != nil {
		return platform, err
	}

	return platform, nil
}

// getPlatformById retrieves a single platform by ID from the database
func getPlatformById(client *mongo.Client, id string) (Platform, error) {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Convert the string ID to an ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Platform{}, err
	}

	// Define the filter to fetch a platform by ID
	filter := bson.D{{Key: "_id", Value: objID}}

	// Find the platform
	var platform Platform
	err = collection.FindOne(context.TODO(), filter).Decode(&platform)
	if err != nil {
		return platform, err
	}

	return platform, nil
}

func addPlatform(client *mongo.Client, platform Platform) error {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Insert the platform
	_, err := collection.InsertOne(context.TODO(), platform)
	if err != nil {
		return err
	}

	return nil
}

func updatePlatform(client *mongo.Client, platform Platform) error {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Update the platform
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": platform.ID}, bson.M{"$set": platform})
	if err != nil {
		return err
	}

	return nil
}

func deletePlatform(client *mongo.Client, id string) error {
	// Get a collection handle
	collection := client.Database("codex-db").Collection("Codex")

	// Convert the string ID to an ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Delete the platform
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

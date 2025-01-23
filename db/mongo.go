package db

import (
	"context"
	"log"
	"time"

	"github.com/vaalley/codex-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// Sets up MongoDB connection
func ConnectMongo() {
	uri := config.GetEnv("MONGODB_URI")
	log.Printf("🔄 Attempting MongoDB connection to %s... 🗃️", MaskURI(uri))

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ MongoDB connection failed: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping failed: %v", err)
	}

	MongoClient = client
	log.Println("✅ Successfully connected to MongoDB!")
}

// Gets a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	dbName := config.GetEnv("DB_NAME")
	log.Printf("🔄 Accessing collection: %s in database: %s 📑", collectionName, dbName)

	collection := MongoClient.Database(dbName).Collection(collectionName)
	if collection == nil {
		log.Fatalf("❌ Failed to get collection: %s", collectionName)
	}
	return collection
}

// Masks sensitive information in MongoDB URI
func MaskURI(uri string) string {
	if uri == "" {
		return ""
	}

	const mask = "mongodb://*****"

	if len(uri) > 10 {
		return mask
	}
	return uri
}

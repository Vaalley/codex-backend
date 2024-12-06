package db

import (
	"codex-backend/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongo sets up MongoDB connection
func ConnectMongo() {
	uri := config.GetEnv("MONGODB_URI")
	log.Printf("🔄 Attempting MongoDB connection to %s...", maskURI(uri))
	
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

// GetCollection gets a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	dbName := config.GetEnv("DB_NAME")
	log.Printf("📑 Accessing collection: %s in database: %s", collectionName, dbName)
	return MongoClient.Database(dbName).Collection(collectionName)
}

// maskURI masks sensitive information in MongoDB URI
func maskURI(uri string) string {
	if uri == "" {
		return ""
	}
	// Simple masking that keeps protocol and host but masks everything else
	const mask = "mongodb://*****"
	if len(uri) > 10 {
		return mask
	}
	return uri
}

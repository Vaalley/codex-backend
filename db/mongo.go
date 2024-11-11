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
	clientOptions := options.Client().ApplyURI(config.GetEnv("MONGODB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB!")
}

// GetCollection gets a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database(config.GetEnv("DB_NAME")).Collection(collectionName)
}

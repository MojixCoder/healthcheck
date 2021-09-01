package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// MongoConnResult is the MongoDB connection result structure
type MongoConnResult struct {
	Client *mongo.Client
	Error  error
}

// GetMongoClient gets the MongoDB client
func GetMongoClient(mongoURI string, ch chan MongoConnResult) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ch <- MongoConnResult{
		Client: client,
		Error:  err,
	}
}

// ConnectToMongo initializes the MongoDB client
func ConnectToMongo(client *mongo.Client, ch chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	ch <- err
}

// MongoPing checks if a MongoDB server has been found
func MongoPing(client *mongo.Client, ch chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())
	ch <- err
}

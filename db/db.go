package db

import (
	"context"
	"github.com/MojixCoder/healthcheck/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// client is DB Client
var client *mongo.Client

// connectToDB connects to MongoDB
func connectToDB() *mongo.Client {
	mongoURI := config.GetAppConfig().MongoURI

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Init connects to MongoDB and set DB client
func Init() {
	client = connectToDB()
}

// GetCollection returns a Mongo collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DBName := config.GetAppConfig().DBName
	collection := client.Database(DBName).Collection(collectionName)
	return collection
}

// GetDBClient returns DB client
func GetDBClient() *mongo.Client {
	return client
}

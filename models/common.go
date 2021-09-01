package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOneResult is the structure of inserting an object to MongoDB
type InsertOneResult struct {
	InsertedID *mongo.InsertOneResult
	Error      error
}

// InsertOne inserts an object to MongoDB
func InsertOne(collection *mongo.Collection, obj interface{}, ch chan InsertOneResult) {
	insertedID, err := collection.InsertOne(context.TODO(), obj)
	result := InsertOneResult{insertedID, err}
	ch <- result
}

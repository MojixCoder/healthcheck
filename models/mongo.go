package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	MongoURI    string             `bson:"mongoURI" json:"mongoURI"`
	Status      string             `bson:"status" json:"status"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type MongoDBForm struct {
	MongoURI string `json:"mongoURI" validate:"required"`
}

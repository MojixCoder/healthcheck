package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Site is MongoDB collection structure
type Site struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	URL        string             `bson:"url" json:"url"`
	StatusCode int                `bson:"statusCode" json:"statusCode"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}

// SiteForm validates JSON body for checking website
type SiteForm struct {
	URL string `validate:"required,url"`
}

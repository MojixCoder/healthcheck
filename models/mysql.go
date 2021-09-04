package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MySQL is MongoDB collection structure for MySQL reports
type MySQL struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	ConnectionString string             `bson:"connectionString" json:"connectionString"`
	Status           string             `bson:"status" json:"status"`
	Description      string             `bson:"description" json:"description"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
}

// MySQLForm validates JSON body for checking MySQL
type MySQLForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	Host     string `json:"host" validate:"required,hostname|hostname_rfc1123"`
	Port     string `json:"port" validate:"required,numeric"`
	DBName   string `json:"DBName" validate:"required"`
}

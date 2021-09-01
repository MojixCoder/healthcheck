package controllers

import (
	"net/http"
	"time"

	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/helpers"
	"github.com/MojixCoder/healthcheck/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoHealthCheck is the MongoDB health checker
func MongoHealthCheck(c *gin.Context) {
	var MongoForm models.MongoDBForm

	if err := c.ShouldBindJSON(&MongoForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&MongoForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	// mongoConnChannel is MongoDB connection channel
	mongoConnChannel := make(chan helpers.MongoConnResult)

	// mongoErrChannel is the Channel for errors
	mongoErrChannel := make(chan error)

	// Getting MongoDB client
	go helpers.GetMongoClient(MongoForm.MongoURI, mongoConnChannel)
	mongoConn := <-mongoConnChannel
	if mongoConn.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": mongoConn.Error.Error(),
		})
		return
	}

	// Initializing client
	go helpers.ConnectToMongo(mongoConn.Client, mongoErrChannel)
	err := <-mongoErrChannel
	DTNow, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	mongoCollection := db.GetCollection(db.GetDBClient(), "mongodb")
	mongoChannel := make(chan models.InsertOneResult)
	if err != nil {
		connErr := models.MongoDB{
			ID:          primitive.NewObjectID(),
			MongoURI:    MongoForm.MongoURI,
			Status:      "Failed",
			Description: err.Error(),
			CreatedAt:   DTNow,
		}
		go models.InsertOne(mongoCollection, connErr, mongoChannel)
		_ = <-mongoChannel
		c.JSON(http.StatusOK, gin.H{
			"result": connErr,
		})
		return
	}

	// Checking connection
	go helpers.MongoPing(mongoConn.Client, mongoErrChannel)
	err = <-mongoErrChannel
	if err != nil {
		connErr := models.MongoDB{
			ID:          primitive.NewObjectID(),
			MongoURI:    MongoForm.MongoURI,
			Status:      "Failed",
			Description: err.Error(),
			CreatedAt:   DTNow,
		}
		go models.InsertOne(mongoCollection, connErr, mongoChannel)
		_ = <-mongoChannel
		c.JSON(http.StatusOK, gin.H{
			"result": connErr,
		})
		return
	}

	// Connection was OK!
	result := models.MongoDB{
		ID:          primitive.NewObjectID(),
		MongoURI:    MongoForm.MongoURI,
		Status:      "Success",
		Description: "Ok",
		CreatedAt:   DTNow,
	}

	go models.InsertOne(mongoCollection, result, mongoChannel)
	_ = <-mongoChannel

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

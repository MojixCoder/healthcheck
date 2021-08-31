package controllers

import (
	"context"
	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/helpers"
	"github.com/MojixCoder/healthcheck/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func SiteHealthCheck(c *gin.Context) {
	var siteForm models.SiteForm

	if err := c.ShouldBindJSON(&siteForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&siteForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	res, err := helpers.HeadRequest(siteForm.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	DTNow, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	var siteReport = models.Site{
		ID: primitive.NewObjectID(),
		URL: siteForm.URL,
		Status: res.Status,
		StatusCode: res.StatusCode,
		CreatedAt: DTNow,
	}

	siteCollection := db.GetCollection(db.GetDBClient(), "site")
	_, err = siteCollection.InsertOne(context.TODO(), siteReport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "unable to insert object",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": siteReport,
	})
}

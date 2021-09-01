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

// SiteHealthCheck is website health checker
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

	httpReqChannel := make(chan helpers.HttpRequestResult)

	// Performs a head request to URL
	go helpers.HeadRequest(siteForm.URL, httpReqChannel)
	reqResult := <-httpReqChannel
	if reqResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": reqResult.Error.Error(),
		})
		return
	}
	defer reqResult.Response.Body.Close()

	// Channel
	ch := make(chan models.InsertOneResult)
	// Site collection
	siteCollection := db.GetCollection(db.GetDBClient(), "site")

	// siteReport is the report of the head request to URL
	DTNow, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	var siteReport = models.Site{
		ID:         primitive.NewObjectID(),
		URL:        siteForm.URL,
		Status:     reqResult.Response.Status,
		StatusCode: reqResult.Response.StatusCode,
		CreatedAt:  DTNow,
	}

	// Insert site report to site collection
	go models.InsertOne(siteCollection, siteReport, ch)
	result := <-ch
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "unable to insert object",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": siteReport,
	})
}

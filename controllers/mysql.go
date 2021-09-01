package controllers

import (
	"net/http"
	"time"

	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/helpers"
	"github.com/MojixCoder/healthcheck/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MySQLHealthCheck is the MySQL health checker
func MySQLHealthCheck(c *gin.Context) {
	var mySQLForm models.MySQLForm

	if err := c.ShouldBindJSON(&mySQLForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&mySQLForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": err.Error(),
		})
		return
	}

	// mySQLChannel is MySQL channel
	mySQLChannel := make(chan helpers.MySQLConnResult)

	go helpers.ConnectToMySQL(
		mySQLForm.Username, mySQLForm.Password, mySQLForm.Host, mySQLForm.Port, mySQLForm.DBName,
		mySQLChannel,
	)
	// connResult is MySQL connection result
	connResult := <-mySQLChannel
	if connResult.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"detail": connResult.Error,
		})
		return
	}
	defer connResult.Conn.Close()

	// Channel
	mongoChannel := make(chan models.InsertOneResult)
	// MySql collection
	mySQLCollection := db.GetCollection(db.GetDBClient(), "mysql")

	DTNow, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if err := connResult.Conn.Ping(); err != nil {
		// mySQLFailReport is the MySQL failure report
		mySQLFailReport := models.MySQL{
			ID:               primitive.NewObjectID(),
			ConnectionString: connResult.ConnString,
			Status:           "Failed",
			Description:      err.Error(),
			CreatedAt:        DTNow,
		}
		go models.InsertOne(mySQLCollection, mySQLFailReport, mongoChannel)
		result := <-mongoChannel
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "unable to insert object",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": mySQLFailReport,
		})
	} else {
		// mySQLSuccessReport is the MySQL successful report
		mySQLSuccessReport := models.MySQL{
			ID:               primitive.NewObjectID(),
			ConnectionString: connResult.ConnString,
			Status:           "Success",
			Description:      "Ok",
			CreatedAt:        DTNow,
		}
		go models.InsertOne(mySQLCollection, mySQLSuccessReport, mongoChannel)
		result := <-mongoChannel
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "unable to insert object",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": mySQLSuccessReport,
		})
	}
}

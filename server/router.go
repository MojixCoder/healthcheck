package server

import (
	"github.com/MojixCoder/healthcheck/controllers"
	"github.com/gin-gonic/gin"
)

// NewRouter is the application router
func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/website", controllers.SiteHealthCheck)
		v1.POST("/mysql", controllers.MySQLHealthCheck)
	}
	return router
}

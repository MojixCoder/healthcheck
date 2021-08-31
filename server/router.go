package server

import (
	"github.com/MojixCoder/healthcheck/controllers"
	"github.com/gin-gonic/gin"
)

// newRouter is the application router
func newRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/website", controllers.SiteHealthCheck)
	}
	return router
}

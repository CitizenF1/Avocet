package router

import (
	"rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	return router
}

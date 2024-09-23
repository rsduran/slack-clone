package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routing for the backend
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Auth Routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Channels Routes
	router.GET("/channels", controllers.GetChannels)
	router.POST("/channels", controllers.CreateChannel)

	return router
}

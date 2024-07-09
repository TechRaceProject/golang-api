package routes

import (
	"api/src/controllers"
	"api/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/welcome", controllers.Welcome)
	}

	VehicleRoutes(router)

	return router
}

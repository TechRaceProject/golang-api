package routes

import (
	"api/src/controllers"
	"api/src/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	router = SetupCors(router)

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/welcome", controllers.Welcome)
	}

	VehicleRoutes(router)

	return router
}

func SetupCors(router *gin.Engine) *gin.Engine {
	allowedOrigins := []string{
		os.Getenv("APP_FRONTEND_URL"),
	}

	corsConfig := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	return router
}

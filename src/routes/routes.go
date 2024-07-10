package routes

import (
	"api/src/routes/protected"
	"api/src/routes/public"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

func SetupRouter(router *gin.Engine) *gin.Engine {
	router = SetupCors(router)

	apiGroup := router.Group("/api")

	public.SetupPublicRoutes(apiGroup)

	protected.SetupProtectedRoutes(apiGroup)

	return router
}

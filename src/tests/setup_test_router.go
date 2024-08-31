package tests

import (
	"api/src/config"
	"api/src/routes/protected"
	"api/src/routes/public"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupTestRouter(router *gin.Engine) *gin.Engine {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	apiGroup := router.Group("/api")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	public.SetupPublicRoutes(apiGroup, cfg)

	protected.SetupProtectedRoutes(apiGroup)

	return router
}

func GetTestRouter() *gin.Engine {
	if router == nil {
		router = setupTestRouter(gin.New())
	}

	return router
}

package tests

import (
	"api/src/routes/protected"
	"api/src/routes/public"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the test database")
	}

	return db
}

func GetTestDBConnection() *gorm.DB {
	if db == nil {
		db = setupTestDB()
	}

	return db
}

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

	public.SetupPublicRoutes(apiGroup)

	protected.SetupProtectedRoutes(apiGroup)

	return router
}

func GetTestRouter() *gin.Engine {
	if router == nil {
		router = setupTestRouter(gin.New())
	}

	return router
}

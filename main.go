package main

import (
	"api/src/models"
	"api/src/routes"
	"api/src/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server on port 8080...")

	database, err := services.InitSqlConnection()

	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}

	connection, err := database.DB()

	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}
	defer connection.Close()

	err = database.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	router := routes.SetupRouter()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8000")
}

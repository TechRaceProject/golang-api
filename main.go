package main

import (
	"api/src/models"
	"api/src/routes"
	"api/src/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting server on port 8000...")

	database := initDatabase()

	defer closeDatabaseConnection(database)

	performMigrations(database)

	initMQTT()

	startWebServer()
}

func initDatabase() *gorm.DB {
	database, err := services.InitSqlConnection()

	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}

	fmt.Println("Database connection established.")

	return database
}

func closeDatabaseConnection(database *gorm.DB) {
	connection, err := database.DB()

	if err != nil {
		log.Println("An error occurred while closing the database connection: ", err)
		return
	}

	connection.Close()

	fmt.Println("Database migrations completed.")
}

func performMigrations(database *gorm.DB) {
	err := database.AutoMigrate(
		&models.User{},
		&models.SensorData{},
		&models.Vehicle{},
		&models.Fool{},
		&models.Race{},
		&models.VehicleState{},
		&models.PrimaryLedColor{},
		&models.SecondaryLedColor{},
		&models.BuzzerVariable{},
		&models.HeadAngle{},
	)

	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	fmt.Println("Database migrations completed.")
}

func initMQTT() {
	fmt.Println("Starting mqtt connection...")

	client := services.InitMQTTClient("tcp://mosquitto:1883")

	services.Subscribe(client, "esp32/#")

	fmt.Println("MQTT connection established and subscribed to topic.")
}

func startWebServer() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router = routes.SetupRouter(router)

	router.Run(":8000")

	fmt.Println("Server started on port 8000.")
}

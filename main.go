package main

import (
	"api/src/models"
	"api/src/routes"
	"api/src/services"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting server on port 8000...")

	database, err := services.InitSqlConnection()
	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}

	connection, err := database.DB()
	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}
	defer connection.Close()

	err = database.AutoMigrate(&models.User{}, &models.SensorData{}, &models.Vehicle{})
	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	fmt.Println("Starting mqtt connection...")
	client := services.InitMQTTClient("tcp://mosquitto:1883")
	services.Subscribe(client, "esp32/#")

	router := routes.SetupRouter()
	router.Run(":8000")

	select {}
}

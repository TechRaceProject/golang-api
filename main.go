package main

import (
	"api/src/models"
	"api/src/routes"
	"api/src/services"
	"fmt"
	"log"

	"api/src/controllers"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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

	err = database.AutoMigrate(&models.User{}, &models.SensorData{})
	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	fmt.Println("Starting mqtt connection...")
	client := services.InitMQTTClient("tcp://mosquitto:1883")
	services.Subscribe(client, "esp32/#")

	router := routes.SetupRouter()
	router.Run(":8000")

	select {}
	// Initialiser la connexion à la base de données
	db, err := services.InitSqlConnection()
	if err != nil {
		log.Fatalf("Erreur d'initialisation de la base de données: %s", err)
	}
	log.Println("Connexion à la base de données réussie")

	// Initialiser le client MQTT
	mqttClient, err := services.InitMqttClient("tcp://mosquitto:1883", db)
	if err != nil {
		log.Fatalf("Erreur de connexion au broker MQTT: %s", err)
	}
	defer mqttClient.Disconnect(250)

	// Configurer le routeur Gin
	router := gin.Default()
	router.GET("/sensor-data", controllers.GetAllSensorData)

	// Lancer le serveur web
	g := new(errgroup.Group)
	g.Go(func() error {
		log.Println("Serveur web démarré sur le port 8000")
		return router.Run(":8000")
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %s", err)
	}
}

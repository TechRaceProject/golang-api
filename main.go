package main

import (
	"api/src/models"
	"api/src/routes"
	"api/src/services"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting server on port 8000...")

	database := initDatabase()

	defer closeDatabaseConnection(database)

	performMigrations(database)

	initVehicleData(database)

	seedDatabase(database)

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

func initVehicleData(database *gorm.DB) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file in initVehicleData: ", err)
	}

	for i := 1; ; i++ {
		nameKey := fmt.Sprintf("VEHICLE_%d_NAME", i)
		ipAddressKey := fmt.Sprintf("VEHICLE_%d_IPADDRESS", i)
		availableKey := fmt.Sprintf("VEHICLE_%d_IS_AVAILABLE", i)

		name := os.Getenv(nameKey)
		ip := os.Getenv(ipAddressKey)
		isAvailableStr := os.Getenv(availableKey)

		if name == "" || ip == "" || isAvailableStr == "" {
			break
		}

		isAvailable, err := strconv.ParseBool(strings.TrimSpace(isAvailableStr))

		if err != nil {
			log.Fatal("Error parsing isAvailable variable to boolean in initVehicleData: ", err)
		}

		database.FirstOrCreate(&models.Vehicle{
			Name:        name,
			IpAdress:    ip,
			IsAvailable: isAvailable,
		})
	}
}

func seedDatabase(database *gorm.DB) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file in seedDatabase: ", err)
	}

	allowDatabaseSeeding := os.Getenv("ALLOW_DATABASE_SEEDING")

	if allowDatabaseSeeding != "true" {
		return
	}

	var usernames = []string{"David", "Goliath", "Pierre"}
	var emails = []string{"l-david@test.com", "a-goliath@test.com", "q-pierre@test.com"}
	hashedPassword, _ := services.HashPassword("password")

	for i := 0; i < len(usernames); i++ {
		database.FirstOrCreate(&models.User{}, models.User{
			Username: &usernames[i],
			Email:    emails[i],
			Password: string(hashedPassword),
		})
	}

	vehicle := database.First(&models.Vehicle{})

	if vehicle.RowsAffected == 0 {
		database.FirstOrCreate(&models.Vehicle{}, models.Vehicle{
			Name:        "Seed Vehicle",
			IpAdress:    "0.0.0.0",
			IsAvailable: false,
		})
	}
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

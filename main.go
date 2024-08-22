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
	"time"

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

	var vehicle models.Vehicle
	database.FirstOrCreate(&vehicle, models.Vehicle{
		Name:        "Seed Vehicle",
		IpAdress:    "0.0.0.0",
		IsAvailable: false,
	})

	var users []models.User
	var vehicles []models.Vehicle
	var races []models.Race
	var userAlreadyHaveRaces bool
	database.Find(&users)
	database.Find(&vehicles)

	for _, user := range users {
		var userAlreadyHaveVehicleStateForThisVehicle bool

		for _, vehicle := range vehicles {
			userAlreadyHaveVehicleStateForThisVehicle = database.Where("user_id = ? AND vehicle_id = ?", user.ID, vehicle.ID).
				Find(&models.VehicleState{}).
				RowsAffected > 0

			if !userAlreadyHaveVehicleStateForThisVehicle {
				vehicle.InitVehicleState(&user, database)
			}
		}

		userAlreadyHaveRaces = database.Where("user_id = ?", user.ID).Find(&races).RowsAffected >= 3

		if userAlreadyHaveRaces {
			continue
		}

		for i := 0; i < 3; i++ {
			startTime := time.Now().Add(time.Duration(i) * time.Hour)
			endTime := startTime.Add(1 * time.Hour)
			race := models.Race{
				VehicleID:          vehicle.ID,
				StartTime:          startTime,
				EndTime:            &endTime,
				NumberOfCollisions: uint8(i + 1),
				DistanceTravelled:  100 + (i * 10),
				AverageSpeed:       120 + (i * 5),
				OutOfParcours:      uint8(i % 2),
				UserID:             user.ID,
			}
			database.Create(&race)
		}
	}
}

func initMQTT() {
	fmt.Println("Starting MQTT connection...")

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

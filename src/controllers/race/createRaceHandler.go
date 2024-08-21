package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/race"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRaceHandler(c *gin.Context) {
	fmt.Println("Start the creation of the race")
	// Récupération de l'ID utilisateur depuis les paramètres de l'URL
	userIdStr := c.Param("userId")
	if userIdStr == "" || userIdStr == "0" || userIdStr == ":userId" {
		services.SetUnprocessableEntity(c, "User not found")
		return
	}
	// Conversion du userId de string à uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		services.SetUnprocessableEntity(c, "Invalid user ID")
		return
	}

	var createRaceValidator validators.CreateRaceValidator

	// Validation de la requête JSON
	if err := c.ShouldBindJSON(&createRaceValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	// Validation des données via le validateur
	if err := createRaceValidator.Validate(); err != nil {
		services.SetValidationErrorResponse(c, err)
		return
	}

	// Création du modèle Race avec les données validées
	race := models.Race{
		StartTime:          createRaceValidator.StartTime,
		EndTime:            createRaceValidator.EndTime,
		NumberOfCollisions: *createRaceValidator.NumberOfCollisions,
		DistanceTravelled:  *createRaceValidator.DistanceTravelled,
		AverageSpeed:       *createRaceValidator.AverageSpeed,
		OutOfParcours:      *createRaceValidator.OutOfParcours,
		RaceType:           createRaceValidator.RaceType,
		VehicleID:          createRaceValidator.VehicleID,
		UserID:             uint(userId), // Conversion de uint64 à uint
	}

	// Récupération de la connexion à la base de données
	db := services.GetConnection()

	// Création de l'enregistrement dans la base de données
	if err := db.Create(&race).Error; err != nil {
		fmt.Printf("Error creating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to create Race")
		return
	}
	fmt.Println("Race created")
	// Réponse de succès avec l'objet Race créé
	services.SetCreated(c, "Race created successfully", race)
}

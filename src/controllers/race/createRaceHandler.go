package handlers

import (
	"api/src/models"
	"api/src/models/attributes"
	"api/src/services"
	validators "api/src/validators/race"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRaceHandler(c *gin.Context) {
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

	startTime := createRaceValidator.StartTime
	var endTime *attributes.CustomTime

	if createRaceValidator.EndTime != nil {
		endTime = createRaceValidator.EndTime
	}

	// Création du modèle Race avec les données validées
	race := models.Race{
		Name:               createRaceValidator.Name,
		StartTime:          startTime,
		EndTime:            endTime,
		NumberOfCollisions: *createRaceValidator.NumberOfCollisions,
		DistanceCovered:    *createRaceValidator.DistanceCovered,
		AverageSpeed:       *createRaceValidator.AverageSpeed,
		OutOfParcours:      *createRaceValidator.OutOfParcours,
		Status:             createRaceValidator.Status,
		Type:               createRaceValidator.Type,
		VehicleID:          createRaceValidator.VehicleID,
		UserID:             uint(userId),
	}

	// Récupération de la connexion à la base de données
	db := services.GetConnection()

	// Création de l'enregistrement dans la base de données
	if err := db.Create(&race).Error; err != nil {
		fmt.Printf("Error creating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to create Race")
		return
	}

	db.Preload("Vehicle").Find(&race)

	// Réponse de succès avec l'objet Race créé
	services.SetCreated(c, "Race created successfully", race)
}

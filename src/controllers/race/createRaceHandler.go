package handlers

import (
	"api/src/models"
	"api/src/models/attributes"
	"api/src/services"
	validators "api/src/validators/race"
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

	db := services.GetConnection()

	if db.Where("id = ?", userId).Find(&models.User{}).RowsAffected == 0 {
		services.SetUnprocessableEntity(c, "User not found")
		return
	}

	var startTime *attributes.CustomTime = createRaceValidator.StartTime
	var endTime *attributes.CustomTime

	if createRaceValidator.EndTime != nil {
		endTime = createRaceValidator.EndTime
	}

	// Création du modèle Race avec les données validées
	race := models.Race{
		Name:              createRaceValidator.Name,
		StartTime:         startTime,
		EndTime:           endTime,
		CollisionDuration: 0,
		DistanceCovered:   0,
		AverageSpeed:      0,
		OutOfParcours:     0,
		Status:            createRaceValidator.Status,
		Type:              createRaceValidator.Type,
		VehicleID:         createRaceValidator.VehicleID,
		UserID:            uint(userId),
	}

	// Création de l'enregistrement dans la base de données
	if err := db.Create(&race).Error; err != nil {
		services.SetInternalServerError(c, err.Error())
		return
	}

	db.Preload("Vehicle").Find(&race)

	// Réponse de succès avec l'objet Race créé
	services.SetCreated(c, "Race created successfully", race)
}

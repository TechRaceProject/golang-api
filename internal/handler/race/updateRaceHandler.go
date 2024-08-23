package handlers

import (
	errors "api/internal/errors"
	"api/internal/models"
	validators "api/internal/validators/race"
	utils "api/pkg/httputils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdateRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	db := utils.GetConnection()

	// Récupère la course existante
	var existingRace models.Race
	if err := db.First(&existingRace, raceID).Error; err != nil {
		errors.SetNotFound(c, "Race not found")
		return
	}

	var raceValidator validators.UpdateRaceValidator

	// Valide les données JSON
	if err := c.ShouldBindJSON(&raceValidator); err != nil {
		errors.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := raceValidator.Validate(); err != nil {
		errors.SetUnprocessableEntity(c, err.Error())
		return
	}

	// Mise à jour uniquement du champ End_time
	if raceValidator.EndTime != nil && !raceValidator.EndTime.IsZero() {
		if raceValidator.EndTime.Before(existingRace.StartTime) {
			errors.SetUnprocessableEntity(c, "EndTime cannot be before StartTime")
			return
		}
		existingRace.EndTime = raceValidator.EndTime
	}

	if raceValidator.Name != "" {
		existingRace.Name = raceValidator.Name
	}

	if raceValidator.Status != "" {
		existingRace.Status = raceValidator.Status
	}

	// Sauvegarde les modifications dans la base de données
	if err := db.Preload("Vehicle").Save(&existingRace).Error; err != nil {
		fmt.Printf("Error updating Race: %v\n", err)
		errors.SetInternalServerError(c, "Failed to update Race")
		return
	}

	// Réponse de succès avec l'objet Race mis à jour
	utils.SetOK(c, "Race updated successfully", existingRace)
}

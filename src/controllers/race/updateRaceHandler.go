package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/race"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdateRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	db := services.GetConnection()

	// Récupère la course existante
	var existingRace models.Race
	if err := db.First(&existingRace, raceID).Error; err != nil {
		services.SetNotFound(c, "Race not found")
		return
	}

	var raceValidator validators.UpdateRaceValidator

	// Valide les données JSON
	if err := c.ShouldBindJSON(&raceValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := raceValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	// Mise à jour uniquement du champ End_time
	if raceValidator.EndTime != nil && !raceValidator.EndTime.IsZero() {
		if raceValidator.EndTime.Before(existingRace.Start_time) {
			services.SetUnprocessableEntity(c, "EndTime cannot be before StartTime")
			return
		}
		existingRace.End_time = raceValidator.EndTime
	}

	// Sauvegarde les modifications dans la base de données
	if err := db.Save(&existingRace).Error; err != nil {
		fmt.Printf("Error updating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to update Race")
		return
	}

	// Réponse de succès avec l'objet Race mis à jour
	services.SetOK(c, "Race updated successfully", existingRace)
}

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
		if err.Error() == "record not found" {
			services.SetNotFound(c, err.Error())
			return
		}

		services.SetInternalServerError(c, err.Error())
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
		if raceValidator.EndTime.Before(existingRace.StartTime.Time) {
			services.SetUnprocessableEntity(c, "EndTime cannot be before StartTime")

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

	if err := db.Save(&existingRace).Error; err != nil {
		fmt.Printf("Error updating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to update Race : "+err.Error())
		return
	}

	db.Preload("Vehicle").First(&existingRace, raceID)

	services.SetOK(c, "Race updated successfully", existingRace)
}

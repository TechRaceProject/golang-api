package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"
	validators "api/src/validators/race"

	"github.com/gin-gonic/gin"
)

func UpdateRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	db := services.GetConnection()

	var existingRace models.Race
	if err := db.Preload("Vehicle").First(&existingRace, raceID).Error; err != nil {
		fmt.Println("Error retrieving Race from the database:", err)
		services.SetNotFound(c, "Race not found")
		return
	}

	fmt.Println("Existing Race:", existingRace)

	var raceValidator validators.CreateRaceValidator

	if err := c.ShouldBindJSON(&raceValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := raceValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	existingRace.Duration = raceValidator.Duration
	existingRace.ElapsedTime = raceValidator.ElapsedTime
	existingRace.Laps = raceValidator.Laps
	existingRace.RaceType = raceValidator.RaceType
	existingRace.AverageSpeed = raceValidator.AverageSpeed
	existingRace.TotalFaults = raceValidator.TotalFaults
	existingRace.EffectiveDuration = raceValidator.EffectiveDuration
	existingRace.UserID = raceValidator.UserID
	existingRace.VehicleID = raceValidator.VehicleID

	if err := db.Save(&existingRace).Error; err != nil {
		fmt.Printf("Error updating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to update Race")
		return
	}

	services.SetOK(c, "Race updated successfully", existingRace)
}

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

	raceValidator := validators.CreateRaceValidator{
		Duration:          existingRace.Duration,
		ElapsedTime:       existingRace.ElapsedTime,
		Laps:              existingRace.Laps,
		RaceType:          existingRace.RaceType,
		AverageSpeed:      existingRace.AverageSpeed,
		TotalFaults:       existingRace.TotalFaults,
		EffectiveDuration: existingRace.EffectiveDuration,
		UserID:            existingRace.UserID,
		VehicleID:         existingRace.VehicleID,
	}

	if err := raceValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	existingRace.Update(raceValidator)

	db.Save(&existingRace)

	services.SetCreated(c, "Race updated successfully", existingRace)
}

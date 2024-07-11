package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/race"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateRaceHandler(c *gin.Context) {
	var race models.Race

	if err := c.ShouldBindJSON(&race); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	raceValidator := validators.CreateRaceValidator{
		Duration:          race.Duration,
		ElapsedTime:       race.ElapsedTime,
		Laps:              race.Laps,
		RaceType:          race.RaceType,
		AverageSpeed:      race.AverageSpeed,
		TotalFaults:       race.TotalFaults,
		EffectiveDuration: race.EffectiveDuration,
		UserID:            race.UserID,
		VehicleID:         race.VehicleID,
	}

	if err := raceValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	db := services.GetConnection()

	if err := db.Create(&race).Error; err != nil {
		// Log the detailed error message for debugging purposes
		fmt.Printf("Error creating Race: %v\n", err)

		// Set an appropriate error response
		services.SetInternalServerError(c, "Failed to create Race")
		return
	}

	services.SetCreated(c, "Race created successfully", race)
}

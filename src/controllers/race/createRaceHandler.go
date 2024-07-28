package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/race"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateRaceHandler(c *gin.Context) {
	var createRaceValidator validators.CreateRaceValidator

	if err := c.ShouldBindJSON(&createRaceValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := createRaceValidator.Validate(); err != nil {
		services.SetValidationErrorResponse(c, err)
		return
	}

	race := models.Race{
		Duration:          createRaceValidator.Duration,
		ElapsedTime:       createRaceValidator.ElapsedTime,
		Laps:              createRaceValidator.Laps,
		RaceType:          createRaceValidator.RaceType,
		AverageSpeed:      createRaceValidator.AverageSpeed,
		TotalFaults:       createRaceValidator.TotalFaults,
		EffectiveDuration: createRaceValidator.EffectiveDuration,
		UserID:            createRaceValidator.UserID,
		VehicleID:         createRaceValidator.VehicleID,
	}

	db := services.GetConnection()

	if err := db.Create(&race).Error; err != nil {
		fmt.Printf("Error creating Race: %v\n", err)
		services.SetInternalServerError(c, "Failed to create Race")
		return
	}

	services.SetCreated(c, "Race created successfully", race)
}

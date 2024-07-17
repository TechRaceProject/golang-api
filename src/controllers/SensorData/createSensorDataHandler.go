package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/sensorData"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateSensorDataHandler(c *gin.Context) {
	var sensorData models.SensorData

	if err := c.ShouldBindJSON(&sensorData); err != nil {
		services.SetJsonBindingErrorResponse(c, err)

		return
	}

	sensorDataValidator := validators.CreateSensorDataValidator{
		Light: sensorData.Light,
		Sonar: sensorData.Sonar,
		Track: sensorData.Track,
	}

	if err := sensorDataValidator.Validate(); err != nil {
		services.SetValidationErrorResponse(c, err)

		return
	}

	db := services.GetConnection()

	if err := db.Create(&sensorData).Error; err != nil {
		// Log the detailed error message for debugging purposes
		fmt.Printf("Error creating Race: %v\n", err)

		// Set an appropriate error response
		services.SetInternalServerError(c, "Failed to create Race")
		return
	}

	services.SetCreated(c, "Race created successfully", sensorData)
}

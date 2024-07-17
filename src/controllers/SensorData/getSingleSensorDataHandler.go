package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func GetSingleSensorDataHandler(c *gin.Context) {
	db := services.GetConnection()

	var sensorData models.SensorData

	sensorDataId := c.Param("sensorDataId")

	if err := db.Preload("Vehicle").First(&sensorData, sensorDataId).Error; err != nil {
		if err.Error() == "record not found" {
			services.SetNotFound(c, "Sensor Data not found")
		} else {
			services.SetInternalServerError(c, "Failed to retrieve Sensor Data")
		}
		return
	}

	services.SetOK(c, "Sensor Data retrieved successfully", sensorData)
}

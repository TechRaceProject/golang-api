package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func GetAllSensorDataHandler(c *gin.Context) {
	db := services.GetConnection()

	var sensorData []models.SensorData

	if err := db.Find(&sensorData).Error; err != nil {
		services.SetInternalServerError(c, "Failed to retrieve Sensor Data")
		return
	}

	services.SetOK(c, "Sensor Data retrieved successfully", sensorData)
}

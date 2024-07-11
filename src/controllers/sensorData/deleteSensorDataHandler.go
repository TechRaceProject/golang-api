package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteSensorDataHandler(c *gin.Context) {
	sensorDataId := c.Param("sensorDataId")

	// Access the database connection
	db := services.GetConnection()

	var existingSensorData models.SensorData
	if err := db.First(&existingSensorData, sensorDataId).Error; err != nil {
		fmt.Println("Error retrieving Sensor Data from the database:", err)
		services.SetNotFound(c, "Sensor Data not found")
		return
	}

	fmt.Println("Existing Sensor Data to delete:", existingSensorData)

	if err := db.Delete(&existingSensorData).Error; err != nil {
		fmt.Println("Error deleting Sensor Data from the database:", err)
		services.SetInternalServerError(c, "Failed to delete Sensor Data")
		return
	}

	services.SetNoContent(c)
}

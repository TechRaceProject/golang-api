package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"
	validators "api/src/validators/sensorData"

	"github.com/gin-gonic/gin"
)

func UpdateSensorDataHandler(c *gin.Context) {
	sensorDataId := c.Param("sensorDataId")

	db := services.GetConnection()

	var existingSensorData models.SensorData
	if err := db.First(&existingSensorData, sensorDataId).Error; err != nil {
		fmt.Println("Error retrieving Sensor Data from the database:", err)
		services.SetNotFound(c, "Sensor Data not found")
		return
	}

	fmt.Println("Existing  Sensor Data:", existingSensorData)

	sensorDataValidator := validators.CreateSensorDataValidator{
		Light: existingSensorData.Light,
		Sonar: existingSensorData.Sonar,
		Track: existingSensorData.Track,
	}

	if err := sensorDataValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	existingSensorData.Update(sensorDataValidator)

	db.Save(&existingSensorData)

	services.SetCreated(c, "Sensor Data updated successfully", existingSensorData)
}

package controllers

import (
	"api/src/models"
	"api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Helper function to transform vehicle model to API response format.
func transformVehicleToResponse(vehicle models.Vehicle) gin.H {
	return gin.H{
		"data": gin.H{
			"type":       "vehicule",
			"id":         strconv.Itoa(int(vehicle.ID)),
			"attributes": vehicleToAttributes(vehicle),
		},
	}
}

// Helper function to map vehicle fields to response attributes.
func vehicleToAttributes(vehicle models.Vehicle) gin.H {
	return gin.H{
		"battery_life":   vehicle.BatteryLife,
		"vehicle_name":   vehicle.Name,
		"line_sensors":   []bool{vehicle.LineSensor1, vehicle.LineSensor2, vehicle.LineSensor3}, // Assume 3 line sensors
		"camera":         vehicle.Camera,
		"sonar_range":    vehicle.SonarRange,
		"wheel_power":    []int{vehicle.WheelPower1, vehicle.WheelPower2, vehicle.WheelPower3, vehicle.WheelPower4}, // 4 wheel powers
		"led_color":      vehicle.LedColor,
		"display_panel":  vehicle.DisplayPanel,
		"speaker_status": vehicle.SpeakerStatus,
		"sound_playing":  vehicle.SoundPlaying,
	}
}


func GetVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	result := services.GetConnection().First(&vehicle, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, transformVehicleToResponse(vehicle))
}

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	result := services.GetConnection().Find(&vehicles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var vehiclesResponse []gin.H
	for _, vehicle := range vehicles {
		vehiclesResponse = append(vehiclesResponse, transformVehicleToResponse(vehicle))
	}

	c.JSON(http.StatusOK, gin.H{"data": vehiclesResponse})
}

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.BindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := services.GetConnection().Create(&vehicle)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, transformVehicleToResponse(vehicle))
}

func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	result := services.GetConnection().First(&vehicle, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	if err := c.BindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result = services.GetConnection().Save(&vehicle)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, transformVehicleToResponse(vehicle))
}

func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	result := services.GetConnection().First(&vehicle, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	result = services.GetConnection().Delete(&vehicle)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Vehicule deleted successfully"}})
}

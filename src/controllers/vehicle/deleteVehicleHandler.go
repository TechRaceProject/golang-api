package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteVehicleHandler(c *gin.Context) {
	vehicleId := c.Param("vehicleId")

	// Access the database connection
	db := services.GetConnection()

	var existingVehicle models.Race
	if err := db.First(&existingVehicle, vehicleId).Error; err != nil {
		fmt.Println("Error retrieving vehicle from the database:", err)
		services.SetNotFound(c, "Vehicle not found")
		return
	}

	fmt.Println("Existing vehicle to delete:", existingVehicle)

	if err := db.Delete(&existingVehicle).Error; err != nil {
		fmt.Println("Error deleting vehicle from the database:", err)
		services.SetInternalServerError(c, "Failed to delete vehicle")
		return
	}

	services.SetNoContent(c)
}

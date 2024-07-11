package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteRaceHandler(c *gin.Context) {
	raceId := c.Param("raceId")

	// Access the database connection
	db := services.GetConnection()

	var existingRace models.Race
	if err := db.First(&existingRace, raceId).Error; err != nil {
		fmt.Println("Error retrieving race from the database:", err)
		services.SetNotFound(c, "Race not found")
		return
	}

	fmt.Println("Existing race to delete:", existingRace)

	if err := db.Delete(&existingRace).Error; err != nil {
		fmt.Println("Error deleting race from the database:", err)
		services.SetInternalServerError(c, "Failed to delete race")
		return
	}

	services.SetNoContent(c)
}

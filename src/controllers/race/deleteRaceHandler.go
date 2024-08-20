package handlers

import (
	"api/src/models"
	"api/src/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func DeleteRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	// Access the database connection
	db := services.GetConnection()

	var existingRace models.Race

	// Recherche de la course par ID
	if err := db.Where("id = ?", raceID).First(&existingRace).Error; err != nil {
		fmt.Println(raceID)
		services.SetNotFound(c, "Race not found")
		return
	}

	// Suppression de la course
	if err := db.Delete(&existingRace).Error; err != nil {
		services.SetInternalServerError(c, "Internal server error")
		return
	}

	services.SetNoContent(c)
}

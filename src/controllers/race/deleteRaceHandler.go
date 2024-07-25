package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	// Access the database connection
	db := services.GetConnection()

	var existingRace models.Race

	query := db.Where("id", raceID).Find(&existingRace)

	if query.RowsAffected == 0 {
		services.SetNotFound(c, "Race not found")
		return
	}

	query = db.Where("id", raceID).Delete(&existingRace)

	if query.Error != nil {
		services.SetInternalServerError(c, "Internal server error")
		return
	}

	services.SetNoContent(c)
}

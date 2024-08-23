package handlers

import (
	errors "api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
)

func DeleteRaceHandler(c *gin.Context) {
	raceID := c.Param("raceId")

	// Access the database connection
	db := utils.GetConnection()

	var existingRace models.Race

	// Recherche de la course par ID
	if err := db.Where("id = ?", raceID).First(&existingRace).Error; err != nil {
		errors.SetNotFound(c, "Race not found")
		return
	}

	// Suppression de la course
	if err := db.Delete(&existingRace).Error; err != nil {
		errors.SetInternalServerError(c, "Internal server error")
		return
	}

	errors.SetNoContent(c)
}

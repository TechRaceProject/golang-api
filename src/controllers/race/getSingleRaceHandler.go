package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func GetSingleRaceHandler(c *gin.Context) {
	db := services.GetConnection()

	var race models.Race

	raceId := c.Param("raceId")

	if err := db.Preload("Vehicle").First(&race, raceId).Error; err != nil {
		if err.Error() == "record not found" {
			services.SetNotFound(c, "Race not found")
		} else {
			services.SetInternalServerError(c, "Failed to retrieve race")
		}
		return
	}

	services.SetOK(c, "Race retrieved successfully", race)
}

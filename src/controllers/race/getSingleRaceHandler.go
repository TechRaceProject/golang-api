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

	query := db.Where("id", raceId).Find(&race)

	if query.RowsAffected == 0 {
		services.SetNotFound(c, "Race not found")
		return
	}

	services.SetOK(c, "Race retrieved successfully", race)
}

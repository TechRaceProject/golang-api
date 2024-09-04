package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetAllRaceHandler(c *gin.Context) {
	db := services.GetConnection()

	var races []models.Race

	if err := db.Preload(clause.Associations).Find(&races).Error; err != nil {
		services.SetInternalServerError(c, "Failed to retrieve races")
		return
	}

	services.SetOK(c, "Races retrieved successfully", races)
}

package handlers

import (
	errors "api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
)

func GetAllRaceHandler(c *gin.Context) {
	db := utils.GetConnection()

	var races []models.Race

	if err := db.Preload("Vehicle").Find(&races).Error; err != nil {
		errors.SetInternalServerError(c, "Failed to retrieve races")
		return
	}

	utils.SetOK(c, "Races retrieved successfully", races)
}

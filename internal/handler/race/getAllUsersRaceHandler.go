package handlers

import (
	errors "api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsersRaceHandler(c *gin.Context) {
	db := utils.GetConnection()

	var races []models.Race

	userIdStr := c.Param("userId")

	// Convertir userId de string à uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		errors.SetUnprocessableEntity(c, "Invalid user ID")
		return
	}

	db.Where("user_id = ?", uint(userId)).Preload("Vehicle").Find(&races)

	utils.SetOK(c, "User races retrieved successfully", races)
}

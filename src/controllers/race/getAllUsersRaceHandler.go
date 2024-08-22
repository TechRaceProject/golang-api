package handlers

import (
	"api/src/models"
	"api/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsersRaceHandler(c *gin.Context) {
	db := services.GetConnection()

	var races []models.Race

	userIdStr := c.Param("userId")

	// Convertir userId de string à uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		services.SetUnprocessableEntity(c, "Invalid user ID")
		return
	}

	db.Where("user_id = ?", uint(userId)).Find(&races)

	services.SetOK(c, "User races retrieved successfully", races)
}

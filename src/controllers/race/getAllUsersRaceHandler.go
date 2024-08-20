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

	// Requête pour trouver toutes les courses pour un utilisateur donné
	query := db.Where("user_id = ?", uint(userId)).Find(&races)

	if query.RowsAffected == 0 {
		services.SetNotFound(c, "No races found for this user")
		return
	}

	services.SetOK(c, "Races retrieved successfully", races)
}

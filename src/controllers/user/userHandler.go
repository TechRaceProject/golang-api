package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	db := services.GetConnection()

	var users models.User

	userId := c.Param("userId")

	if err := db.First(&users, userId).Error; err != nil {
		if err.Error() == "record not found" {
			services.SetNotFound(c, "User not found")
		} else {
			services.SetInternalServerError(c, "Failed to retrieve user")
		}
		return
	}

	services.SetOK(c, "User retrieved successfully", users)
}

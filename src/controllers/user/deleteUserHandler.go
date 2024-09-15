package handlers

import (
	"api/src/models"
	"api/src/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	db := services.GetConnection()

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		if err.Error() == "record not found" {
			services.SetNotFound(c, "User not found")
			return
		}
		services.SetInternalServerError(c, err.Error())
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		services.SetInternalServerError(c, "Failed to delete user: "+err.Error())
		return
	}

	services.SetOK(c, "User deleted successfully", nil)
}

package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	userId := c.Param("userId")

	// Access the database connection
	db := services.GetConnection()

	var existingUser models.User
	if err := db.First(&existingUser, userId).Error; err != nil {
		fmt.Println("Error retrieving user from the database:", err)
		services.SetNotFound(c, "User not found")
		return
	}

	fmt.Println("Existing user to delete:", existingUser)

	if err := db.Delete(&existingUser).Error; err != nil {
		fmt.Println("Error deleting user from the database:", err)
		services.SetInternalServerError(c, "Failed to delete user")
		return
	}

	services.SetNoContent(c)
}

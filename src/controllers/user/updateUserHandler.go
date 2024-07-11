package handlers

import (
	"fmt"
	"net/http"

	"api/src/models"
	"api/src/services"
	validators "api/src/validators/user"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(c *gin.Context) {
	userID := c.Param("userId")

	db := services.GetConnection()

	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		fmt.Println("Error retrieving user from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println("Existing User:", existingUser)

	user, exists := c.Get("user")
	if !exists {
		fmt.Println("User information not available for this request put")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not available for this request put"})
		return
	}

	updateUser, ok := user.(*models.User)
	if !ok {
		fmt.Println("User information not available in the expected format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not available in the expected format"})
		return
	}

	var userValidator validators.UpdateUserValidator

	if err := c.ShouldBindJSON(&userValidator); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userValidator.Validate(); err != nil {
		fmt.Println("Error validating user input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser.Update(userValidator)

	db.Save(updateUser)

	c.JSON(http.StatusOK, gin.H{"user": updateUser})
}

package handlers

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(c *gin.Context) {
	userID := c.Param("id")

	db := services.GetConnection()

	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		if err.Error() == "record not found" {
			services.SetNotFound(c, "User not found")
			return
		}
		services.SetInternalServerError(c, err.Error())
		return
	}

	var userValidator validators.UpdateUserValidator

	if err := c.ShouldBindJSON(&userValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := userValidator.Validate(); err != nil {
		services.SetUnprocessableEntity(c, err.Error())
		return
	}

	if userValidator.Username != "" {
		existingUser.Username = &userValidator.Username
	}

	if userValidator.Email != "" {
		existingUser.Email = userValidator.Email
	}

	// Save the updated user back to the database
	if err := db.Save(&existingUser).Error; err != nil {
		fmt.Printf("Error updating User: %v\n", err)
		services.SetInternalServerError(c, "Failed to update User: "+err.Error())
		return
	}

	// Return success response with updated user data
	services.SetOK(c, "User updated successfully", existingUser)
}

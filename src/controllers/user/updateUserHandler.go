package handlers

import (
	"fmt"
	"net/http"
	"encoding/base64"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strings"

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

	// Check if profile pic is included in request
	if userValidator.ProfilePic != "" {
		// Validate and save profile pic
		profilePic, err := parseAndValidateImage(userValidator.ProfilePic)
		if err != nil {
			fmt.Println("Error validating profile picture:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		existingUser.ProfilePic = profilePic
	}

	db.Save(updateUser)

	c.JSON(http.StatusOK, gin.H{"user": updateUser})
}

func parseAndValidateImage(encodedImage string) ([]byte, error) {
	parts := strings.SplitN(encodedImage, ",", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid base64 image format")
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}

	// Check image format and size
	if len(decoded) > 1*1024*1024 { // 1MB limit
		return nil, errors.New("image size exceeds 1MB")
	}

	switch img.(type) {
	case *image.JPEG:
		return decoded, nil
	case *image.PNG:
		return decoded, nil
	default:
		return nil, errors.New("unsupported image format")
	}
}

package handlers

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func GetAllUserHandler(c *gin.Context) {
	db := services.GetConnection()

	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		services.SetInternalServerError(c, "Failed to retrieve User")
		return
	}

	services.SetOK(c, "Users retrieved successfully", users)

}

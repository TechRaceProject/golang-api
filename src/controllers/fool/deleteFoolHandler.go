package handlers

import (
	"fmt"

	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func DeleteFoolHandler(c *gin.Context) {
	foolId := c.Param("foolId")

	// Access the database connection
	db := services.GetConnection()

	var existingFool models.Fool
	if err := db.First(&existingFool, foolId).Error; err != nil {
		services.SetNotFound(c, "Fool not found")
		return
	}

	fmt.Println("Existing fool to delete:", existingFool)

	if err := db.Delete(&existingFool).Error; err != nil {
		services.SetInternalServerError(c, "Failed to delete fool")
		return
	}

	services.SetNoContent(c)
}

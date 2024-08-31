package vehicle_state

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetUserVehicleStatesHandler(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" || userId == "0" || userId == ":id" {
		services.SetUnprocessableEntity(c, "User id is required")

		return
	}

	connection := services.GetConnection()

	if connection.Where("id = ?", userId).First(&models.User{}).RowsAffected == 0 {
		services.SetNotFound(c, "User not found")

		return
	}

	var vehicleState []models.VehicleState

	connection.Preload(clause.Associations).Where("user_id = ?", userId).Find(&vehicleState)

	services.SetOK(c, "Vehicule states successfully retrieved", vehicleState)
}

package vehicle

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetVehicleHandler(c *gin.Context) {
	vehicleId := c.Param("id")

	if vehicleId == "" || vehicleId == "0" || vehicleId == ":id" {
		services.SetUnprocessableEntity(c, "Vehicle id is required")

		return
	}

	connection := services.GetConnection()

	var vehicle models.Vehicle

	if connection.Where("id = ?", vehicleId).Preload(clause.Associations).First(&vehicle).RowsAffected == 0 {
		services.SetUnprocessableEntity(c, "Vehicle not found")

		return
	}

	services.SetOK(c, "Vehicle successfully retrieved", vehicle)
}

package vehicle_state

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetASingleVehiculeStateHandler(c *gin.Context) {
	vehiculeStateId := c.Param("id")

	if vehiculeStateId == "" || vehiculeStateId == "0" || vehiculeStateId == ":id" {
		services.SetUnprocessableEntity(c, "Vehicle state id is required")

		return
	}

	connection := services.GetConnection()

	var vehicleState models.VehicleState

	if connection.Where("id = ?", vehiculeStateId).Preload(clause.Associations).First(&vehicleState).RowsAffected == 0 {
		services.SetUnprocessableEntity(c, "Vehicle state not found")

		return
	}

	services.SetOK(c, "Vehicule state successfully retrieved", vehicleState)
}

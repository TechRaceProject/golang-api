package vehicle_state

import (
	"api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetASingleVehiculeStateHandler(c *gin.Context) {
	vehiculeStateId := c.Param("id")

	if vehiculeStateId == "" || vehiculeStateId == "0" || vehiculeStateId == ":id" {
		errors.SetUnprocessableEntity(c, "Vehicle state id is required")

		return
	}

	connection := utils.GetConnection()

	var vehicleState models.VehicleState

	if connection.Where("id = ?", vehiculeStateId).Preload(clause.Associations).First(&vehicleState).RowsAffected == 0 {
		errors.SetUnprocessableEntity(c, "Vehicle state not found")

		return
	}

	utils.SetOK(c, "Vehicule state successfully retrieved", vehicleState)
}

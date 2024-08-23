package vehicle

import (
	"api/internal/errors"
	"api/internal/models"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetVehicleHandler(c *gin.Context) {
	vehicleId := c.Param("id")

	if vehicleId == "" || vehicleId == "0" || vehicleId == ":id" {
		errors.SetUnprocessableEntity(c, "Vehicle id is required")

		return
	}

	connection := utils.GetConnection()

	var vehicle models.Vehicle

	if connection.Where("id = ?", vehicleId).Preload(clause.Associations).First(&vehicle).RowsAffected == 0 {
		errors.SetUnprocessableEntity(c, "Vehicle not found")

		return
	}

	utils.SetOK(c, "Vehicle successfully retrieved", vehicle)
}

package vehicle

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/vehicle"

	"github.com/gin-gonic/gin"
)

func UpdateVehicleHandler(c *gin.Context) {
	vehicleId := c.Param("id")

	if vehicleId == "" || vehicleId == "0" || vehicleId == ":id" {
		services.SetUnprocessableEntity(c, "Vehicle not found")
		return
	}

	var UpdateVehicleValidator validators.UpdateVehicleValidator
	if err := c.ShouldBindJSON(&UpdateVehicleValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := UpdateVehicleValidator.Validate(); err != nil {
		services.SetValidationErrorResponse(c, err)
		return
	}

	connection := services.GetConnection()
	if connection == nil {
		services.SetInternalServerError(c, "Database connection error")
		return
	}

	var vehicle models.Vehicle
	if connection.Where("id = ?", vehicleId).First(&vehicle).RowsAffected == 0 {
		services.SetUnprocessableEntity(c, "Vehicle not found")
		return
	}

	errorOccuredInTransaction := connection.Model(&vehicle).
		UpdateColumn("is_available", *UpdateVehicleValidator.IsAvailable).Error

	if errorOccuredInTransaction != nil {
		services.SetInternalServerError(c, "Error during vehicle update")
		return
	}

	services.SetOK(c, "Vehicle state successfully updated", vehicle)
}

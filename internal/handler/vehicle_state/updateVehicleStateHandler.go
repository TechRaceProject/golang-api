package vehicle_state

import (
	"api/internal/errors"
	"api/internal/models"
	validators "api/internal/validators/vehicleState"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func UpdateVehicleStateHandler(c *gin.Context) {
	vehiculeStateId := c.Param("id")

	if vehiculeStateId == "" || vehiculeStateId == "0" || vehiculeStateId == ":id" {
		errors.SetUnprocessableEntity(c, "Vehicle state not found")
		return
	}

	var UpdateVehicleStateValidator validators.UpdateVehicleStateValidator

	if err := c.ShouldBindJSON(&UpdateVehicleStateValidator); err != nil {
		errors.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := UpdateVehicleStateValidator.Validate(); err != nil {
		errors.SetValidationErrorResponse(c, err)
		return
	}

	connection := utils.GetConnection()

	var vehicleState models.VehicleState

	if connection.Where("id = ?", vehiculeStateId).First(&vehicleState).RowsAffected == 0 {
		errors.SetUnprocessableEntity(c, "Vehicle state not found")
		return
	}

	connection.Model(models.VehicleState{}).
		Where("id = ?", vehiculeStateId).
		Updates(models.VehicleState{
			Face:           UpdateVehicleStateValidator.Face,
			LedAnimation:   UpdateVehicleStateValidator.LedAnimation,
			BuzzerAlarm:    UpdateVehicleStateValidator.BuzzerAlarm,
			VideoActivated: UpdateVehicleStateValidator.VideoActivated,
		})

	connection.Model(models.PrimaryLedColor{}).
		Where("id = ?", vehicleState.PrimaryLedColorID).
		Updates(models.PrimaryLedColor{
			LedIdentifier: UpdateVehicleStateValidator.PrimaryLedColor.LedIdentifier,
			Red:           UpdateVehicleStateValidator.PrimaryLedColor.Red,
			Green:         UpdateVehicleStateValidator.PrimaryLedColor.Green,
			Blue:          UpdateVehicleStateValidator.PrimaryLedColor.Blue,
		})

	connection.Model(models.SecondaryLedColor{}).
		Where("id = ?", vehicleState.SecondaryLedColorID).
		Updates(models.SecondaryLedColor{
			BinaryRepresentation: UpdateVehicleStateValidator.SecondaryLedColor.BinaryRepresentation,
			Red:                  UpdateVehicleStateValidator.SecondaryLedColor.Red,
			Green:                UpdateVehicleStateValidator.SecondaryLedColor.Green,
			Blue:                 UpdateVehicleStateValidator.SecondaryLedColor.Blue,
		})

	connection.Model(models.BuzzerVariable{}).
		Where("id = ?", vehicleState.BuzzerVariableID).
		Updates(models.BuzzerVariable{
			Activated: UpdateVehicleStateValidator.BuzzerVariable.Activated,
			Frequency: UpdateVehicleStateValidator.BuzzerVariable.Frequency,
		})

	connection.Model(models.HeadAngle{}).
		Where("id = ?", vehicleState.HeadAngleID).
		Updates(models.HeadAngle{
			VerticalAngle:   UpdateVehicleStateValidator.HeadAngle.VerticalAngle,
			HorizontalAngle: UpdateVehicleStateValidator.HeadAngle.HorizontalAngle,
		})

	connection.Where("id = ?", vehiculeStateId).Preload(clause.Associations).First(&vehicleState)

	err := broadcastUpdatedVehicleState(vehicleState)

	if err != nil {
		errors.SetInternalServerError(c, "Error while broadcasting vehicle state update")

		return
	}

	utils.SetOK(c, "Vehicule state successfully updated", vehicleState)
}

func broadcastUpdatedVehicleState(vehicleState models.VehicleState) error {
	json, err := vehicleState.ToJson()

	if err != nil {
		return err
	}

	go utils.BroadcastMessage(json)

	return nil
}

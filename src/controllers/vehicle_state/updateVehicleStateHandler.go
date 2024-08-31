package vehicle_state

import (
	"api/src/models"
	"api/src/services"
	validators "api/src/validators/vehicleState"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdateVehicleStateHandler(c *gin.Context) {
	vehiculeStateId := c.Param("id")

	if vehiculeStateId == "" || vehiculeStateId == "0" || vehiculeStateId == ":id" {
		services.SetUnprocessableEntity(c, "Vehicle state not found")
		return
	}

	var UpdateVehicleStateValidator validators.UpdateVehicleStateValidator

	if err := c.ShouldBindJSON(&UpdateVehicleStateValidator); err != nil {
		services.SetJsonBindingErrorResponse(c, err)
		return
	}

	if err := UpdateVehicleStateValidator.Validate(); err != nil {
		services.SetValidationErrorResponse(c, err)
		return
	}

	connection := services.GetConnection()

	var vehicleState models.VehicleState

	if connection.Where("id = ?", vehiculeStateId).First(&vehicleState).RowsAffected == 0 {
		services.SetUnprocessableEntity(c, "Vehicle state not found")
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

	for _, primaryLedColor := range *UpdateVehicleStateValidator.PrimaryLedColors {
		model := models.PrimaryLedColor{
			LedIdentifier: primaryLedColor.LedIdentifier,
			Red:           primaryLedColor.Red,
			Green:         primaryLedColor.Green,
			Blue:          primaryLedColor.Blue,
		}

		updatePrimaryLedColorRelationship(connection, model, vehicleState.ID)
	}

	buzzerVariable := models.BuzzerVariable{
		ID:        *vehicleState.BuzzerVariableID,
		Activated: UpdateVehicleStateValidator.BuzzerVariable.Activated,
		Frequency: UpdateVehicleStateValidator.BuzzerVariable.Frequency,
	}

	updateBuzzerVariableRelationship(connection, buzzerVariable)

	headAngle := models.HeadAngle{
		ID:              *vehicleState.HeadAngleID,
		VerticalAngle:   UpdateVehicleStateValidator.HeadAngle.VerticalAngle,
		HorizontalAngle: UpdateVehicleStateValidator.HeadAngle.HorizontalAngle,
	}

	updateHeadAngleRelationship(connection, headAngle)

	connection.Where("id = ?", vehiculeStateId).Preload(clause.Associations).First(&vehicleState)

	err := broadcastUpdatedVehicleState(vehicleState)

	if err != nil {
		services.SetInternalServerError(c, "Error while broadcasting vehicle state update")

		return
	}

	services.SetOK(c, "Vehicule state successfully updated", vehicleState)
}

func updateBuzzerVariableRelationship(connection *gorm.DB, buzzerVariable models.BuzzerVariable) error {
	err := connection.Model(models.BuzzerVariable{}).
		Where("id = ?", buzzerVariable.ID).
		Updates(models.BuzzerVariable{
			Activated: buzzerVariable.Activated,
			Frequency: buzzerVariable.Frequency,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func updateHeadAngleRelationship(connection *gorm.DB, headAngle models.HeadAngle) error {
	err := connection.Model(models.HeadAngle{}).
		Where("id = ?", headAngle.ID).
		Updates(models.HeadAngle{
			VerticalAngle:   headAngle.VerticalAngle,
			HorizontalAngle: headAngle.HorizontalAngle,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func updatePrimaryLedColorRelationship(connection *gorm.DB, primaryLedColor models.PrimaryLedColor, vehicleStateId uint) error {
	err := connection.Model(models.PrimaryLedColor{}).
		Where("vehicle_state_id = ?", vehicleStateId).
		Where("led_identifier = ?", primaryLedColor.LedIdentifier).
		Updates(models.PrimaryLedColor{
			Red:   primaryLedColor.Red,
			Green: primaryLedColor.Green,
			Blue:  primaryLedColor.Blue,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func broadcastUpdatedVehicleState(vehicleState models.VehicleState) error {
	json, err := vehicleState.ToJson()

	if err != nil {
		return err
	}

	go services.BroadcastMessage(json)

	return nil
}

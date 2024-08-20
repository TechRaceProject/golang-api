package validators

import (
	"github.com/go-playground/validator/v10"
)

type CreateVehicleStateValidator struct {
	VehicleID      uint  `json:"vehicle_id" validate:"required"`
	Face           uint8 `json:"face" validate:"required,gte=0,lte=10"`
	LedAnimation   uint8 `json:"led_animation" validate:"required,gte=0,lte=5"`
	BuzzerAlarm    uint8 `json:"buzzer_alarm" validate:"required,gte=0,lte=1"`
	VideoActivated uint8 `json:"video_activated" validate:"required,gte=0,lte=1"`
	UserID         uint  `json:"user_id" validate:"required"`
}

func (c *CreateVehicleStateValidator) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

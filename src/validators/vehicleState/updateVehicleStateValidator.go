package validators

import "github.com/go-playground/validator/v10"

type PrimaryLedColorValidator struct {
	LedIdentifier *int   `json:"led_identifier" validate:"required,gte=0"`
	Red           *uint8 `json:"red" validate:"required,min=0,max=255"`
	Green         *uint8 `json:"green" validate:"required,min=0,max=255"`
	Blue          *uint8 `json:"blue" validate:"required,min=0,max=255"`
}

type BuzzerVariableValidator struct {
	Activated *uint8  `json:"activated" validate:"required,oneof=0 1"`
	Frequency *uint16 `json:"frequency" validate:"required,min=0,max=10000"`
}

type HeadAngleValidator struct {
	VerticalAngle   *uint `json:"vertical_angle" validate:"required,min=0,max=180"`
	HorizontalAngle *uint `json:"horizontal_angle" validate:"required,min=0,max=180"`
}

type UpdateVehicleStateValidator struct {
	Face             *uint8                      `json:"face" validate:"required,gte=0"`
	LedAnimation     *uint8                      `json:"led_animation" validate:"required,gte=0,lte=5"`
	BuzzerAlarm      *uint8                      `json:"buzzer_alarm" validate:"required,oneof=0 1"`
	VideoActivated   *uint8                      `json:"video_activated" validate:"required,oneof=0 1"`
	PrimaryLedColors *[]PrimaryLedColorValidator `json:"primary_led_colors" validate:"required,min=12,max=12,dive"`
	BuzzerVariable   *BuzzerVariableValidator    `json:"buzzer_variable" validate:"required"`
	HeadAngle        *HeadAngleValidator         `json:"head_angle" validate:"required"`
}

func (u *UpdateVehicleStateValidator) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

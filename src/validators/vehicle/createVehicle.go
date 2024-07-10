package validators

import (
	"github.com/go-playground/validator/v10"
)

type CreateVehicleValidator struct {
	Name          string  `json:"vehicle_name" validate:"required"`
	BatteryLife   float64 `json:"battery_life" validate:"required,gte=0"`
	LineSensor1   bool    `json:"line_sensor1"`
	LineSensor2   bool    `json:"line_sensor2"`
	LineSensor3   bool    `json:"line_sensor3"`
	Camera        bool    `json:"camera"`
	SonarRange    float64 `json:"sonar_range" validate:"required,gte=0"`
	WheelPower1   int     `json:"wheel_power1"`
	WheelPower2   int     `json:"wheel_power2"`
	WheelPower3   int     `json:"wheel_power3"`
	WheelPower4   int     `json:"wheel_power4"`
	LedColor      string  `json:"led_color"`
	DisplayPanel  string  `json:"display_panel"`
	SpeakerStatus bool    `json:"speaker_status"`
	SoundPlaying  string  `json:"sound_playing"`
}

func (c *CreateVehicleValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

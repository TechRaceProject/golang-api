package validators

import "github.com/go-playground/validator/v10"

type CreateRaceValidator struct {
	Duration          int    `json:"duration" validate:"required,gte=0"`
	ElapsedTime       int    `json:"elapsed_time" validate:"required,gte=0"`
	Laps              int    `json:"laps" validate:"required,gte=0"`
	RaceType          string `json:"race_type" validate:"required,oneof=VS TIME_TRIAL"`
	AverageSpeed      int    `json:"average_speed" validate:"required,gte=0"`
	TotalFaults       int    `json:"total_faults" validate:"required,gte=0"`
	EffectiveDuration int    `json:"effective_duration" validate:"required,gte=0"`
	UserID            uint   `json:"user_id" validate:"required"`
	VehicleID         uint   `json:"vehicle_id" validate:"required"`
}

func (c *CreateRaceValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

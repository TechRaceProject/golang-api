package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)


type CreateRaceValidator struct {
	StartTime          time.Time  `json:"start_time" validate:"required"`
	EndTime            *time.Time `json:"end_time" validate:"omitempty,gtefield=StartTime"`
	NumberOfCollisions *uint8      `json:"number_of_collisions" validate:"required,min=0,gte=0"`
	DistanceTravelled  *int        `json:"distance_travelled" validate:"required,min=0,gte=0"`
	AverageSpeed       *int        `json:"average_speed" validate:"required,min=0,gte=0"`
	OutOfParcours      *uint8      `json:"out_of_parcours" validate:"required,min=0,gte=0"`
	RaceType           string      `json:"race_type" validate:"required,gte=0"`
	VehicleID          uint       `json:"vehicle_id" validate:"required"`
}

func (c *CreateRaceValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}


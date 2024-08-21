package validators

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateRaceValidator struct {
	StartTime            time.Time    `json:"start_time" validate:"required"`
	EndTime             *time.Time   `json:"end_time" validate:"omitempty,gtefield=StartTime"`
	NumberOfCollisions   uint8    `json:"number_of_collisions" validate:"required,gte=0"`
	DistanceTravelled    int    `json:"distance_travelled" validate:"required,gte=0"`
	AverageSpeed         int    `json:"average_speed" validate:"required,gte=0"`
	OutOfParcours        uint8    `json:"out_of_parcours" validate:"required,gte=0"`
	VehicleID         	 uint   `json:"vehicle_id" validate:"required"`
	

}

func (c *CreateRaceValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}



type UpdateRaceValidator struct {
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time" validate:"omitempty"`
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()

	// Validation personnalisée: vérifier si EndTime n'est pas inférieur à StartTime
	if u.EndTime != nil && u.EndTime.Before(u.StartTime) {
		return errors.New("end_time cannot be before start_time")
	}

	// Si la validation personnalisée passe, alors on fait la validation standard des tags
	return validate.Struct(u)
}

package validators

import (
	"api/src/models/attributes"

	"github.com/go-playground/validator/v10"
)

type CreateRaceValidator struct {
	Name      string                 `json:"name" validate:"required"`
	StartTime *attributes.CustomTime `json:"start_time" validate:"omitempty"`
	EndTime   *attributes.CustomTime `json:"end_time" validate:"omitempty,gtefield=StartTime"`
	Status    string                 `json:"status" validate:"required,oneof='not_started' 'in_progress' 'completed'"`
	Type      string                 `json:"type" validate:"required,oneof='manual' 'auto'"`
	VehicleID uint                   `json:"vehicle_id" validate:"required"`
}

func (c *CreateRaceValidator) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

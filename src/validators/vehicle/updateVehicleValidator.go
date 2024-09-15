package validators

import "github.com/go-playground/validator/v10"

type UpdateVehicleValidator struct {
	IsAvailable *bool `json:"is_available" validate:"required"`
}

func (u *UpdateVehicleValidator) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

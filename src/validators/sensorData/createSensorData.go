package validators

import (
	"github.com/go-playground/validator/v10"
)

type CreateSensorDataValidator struct {
	Light float64 `json:"light" validate:"required,gte=0"`
	Sonar float64 `json:"sonar" validate:"required,gte=0"`
	Track float64 `json:"track" validate:"required,gte=0"`
}

func (c *CreateSensorDataValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

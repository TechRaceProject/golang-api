package validators

import (
	"github.com/go-playground/validator/v10"
)

type CreateFoolValidator struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateFoolValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

package validators

import (
	"github.com/go-playground/validator/v10"
)

type UpdateUserValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3"`
}

func (u *UpdateUserValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

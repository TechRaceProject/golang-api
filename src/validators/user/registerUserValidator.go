package validators

import (
	"github.com/go-playground/validator/v10"
)

type RegisterUserValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required,min=3"`
}

func (u *RegisterUserValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

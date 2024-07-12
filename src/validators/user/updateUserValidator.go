package validators

import (
	"github.com/go-playground/validator/v10"
)

type UpdateUserValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" vaxlidate:"omitempty,min=6"`
	FirstName string `json:"first_name" validate:"omitempty,min=2"`
	Lastname string `json:"last_name" validate:"omitempty,min=2"`
	ProfilePic string `json:"profile_pic,omitempty" validate:"omitempty,base64image"`
}

func (u *UpdateUserValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

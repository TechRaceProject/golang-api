package validators

import (
	"api/src/models/attributes"

	"github.com/go-playground/validator/v10"
)

type UpdateRaceValidator struct {
	EndTime *attributes.CustomTime `json:"end_time"`
	Name    string                 `json:"name"`
	Status  string                 `json:"status" validate:"oneof='not_started' 'in_progress' 'completed' ''"`
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

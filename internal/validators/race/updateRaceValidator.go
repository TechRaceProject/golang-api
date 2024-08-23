package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateRaceValidator struct {
	EndTime *time.Time `json:"end_time" validate:"omitempty"`
	Name    string     `json:"name"`
	Status  string     `json:"status" validate:"oneof='not_started' 'in_progress' 'completed'"`
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

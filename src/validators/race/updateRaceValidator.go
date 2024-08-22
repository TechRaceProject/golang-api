package validators

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateRaceValidator struct {
	EndTime   *time.Time `json:"end_time" validate:"omitempty"`
	Name      string     `json:"name"`
	StartTime time.Time  `json:"start_time"`
	Status    string     `json:"status" validate:"oneof='not_started' 'in_progress' 'completed'"`
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()

	if u.EndTime != nil && u.EndTime.Before(u.StartTime) {
		return errors.New("end_time cannot be before start_time")
	}

	return validate.Struct(u)
}

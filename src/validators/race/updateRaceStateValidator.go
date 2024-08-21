package validators

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateRaceValidator struct {
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time" validate:"omitempty"`
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()

	// Validation personnalisée: vérifier si EndTime n'est pas inférieur à StartTime
	if u.EndTime != nil && u.EndTime.Before(u.StartTime) {
		return errors.New("end_time cannot be before start_time")
	}

	// Si la validation personnalisée passe, alors on fait la validation standard des tags
	return validate.Struct(u)
}
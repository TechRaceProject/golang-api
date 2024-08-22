package validators

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateRaceValidator struct {
	StartTime  time.Time  `json:"start_time"`
	EndTime    *time.Time `json:"end_time" validate:"omitempty"`
	RaceStatus string     `json:"race_status" validate:"required,raceStatus"`
}

// Custom validation function for RaceStatus enum
func raceStatusValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	validStatuses := []string{"Not Started", "In Progress", "Completed"}

	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

func (u *UpdateRaceValidator) Validate() error {
	validate := validator.New()

	// Register the custom validator for RaceStatus
	validate.RegisterValidation("raceStatus", raceStatusValidator)

	// Custom validation: check if EndTime is not before StartTime
	if u.EndTime != nil && u.EndTime.Before(u.StartTime) {
		return errors.New("end_time cannot be before start_time")
	}

	// Standard validation based on struct tags
	return validate.Struct(u)
}

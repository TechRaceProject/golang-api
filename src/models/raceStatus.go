package models

import (
	"errors"
)

type RaceStatus string

const (
	RaceStatusNotStarted RaceStatus = "Not Started"
	RaceStatusInProgress RaceStatus = "In Progress"
	RaceStatusCompleted  RaceStatus = "Completed"
)

// Validates the race status value
func (rs RaceStatus) IsValid() error {
	switch rs {
	case RaceStatusNotStarted, RaceStatusInProgress, RaceStatusCompleted:
		return nil
	}
	return errors.New("invalid race status")
}

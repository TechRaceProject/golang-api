package models

import (
	validators "api/src/validators/race"
)

type Race struct {
	ID                uint `gorm:"primaryKey"`
	Duration          int
	ElapsedTime       int
	Laps              int
	RaceType          string
	AverageSpeed      int
	TotalFaults       int
	EffectiveDuration int
	UserID            uint
	VehicleID         uint
	Vehicle           Vehicle `gorm:"foreignKey:VehicleID"`
	Fool              []Fool  `gorm:"many2many:race_fool;" json:"fool"`
	Model
}

func (r *Race) Create(createRace validators.CreateRaceValidator) error {
	if err := createRace.Validate(); err != nil {
		return err
	}
	r.Duration = createRace.Duration
	r.ElapsedTime = createRace.ElapsedTime
	r.Laps = createRace.Laps
	r.RaceType = createRace.RaceType
	r.AverageSpeed = createRace.AverageSpeed
	r.TotalFaults = createRace.TotalFaults
	r.EffectiveDuration = createRace.EffectiveDuration
	r.UserID = createRace.UserID
	r.VehicleID = createRace.VehicleID
	return nil
}

func (r *Race) Update(updateRace validators.CreateRaceValidator) error {
	if err := updateRace.Validate(); err != nil {
		return err
	}
	r.Duration = updateRace.Duration
	r.ElapsedTime = updateRace.ElapsedTime
	r.Laps = updateRace.Laps
	r.RaceType = updateRace.RaceType
	r.AverageSpeed = updateRace.AverageSpeed
	r.TotalFaults = updateRace.TotalFaults
	r.EffectiveDuration = updateRace.EffectiveDuration
	r.UserID = updateRace.UserID
	r.VehicleID = updateRace.VehicleID
	return nil
}

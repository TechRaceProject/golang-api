package models

import (
	validators "api/src/validators/race"
	"time"
)

type Race struct {
	ID                		uint `gorm:"primaryKey"`
	VehicleID       			uint
	Start_time       			time.Time
	End_time						 *time.Time
	Number_of_collisions  uint8
	Distance_travelled    int
	Average_speed       	int
	Out_of_parcours 			uint8
	UserID								uint
	
	Model
}
// 



func (race *Race) Update(updateRace validators.CreateRaceValidator) error {
	// Valide les données d'entrée
	if err := updateRace.Validate(); err != nil {
		return err
	}

	// Vérifie si StartTime est fourni et met à jour s'il est différent de l'instant zéro
	if !updateRace.StartTime.IsZero() {
		race.Start_time = updateRace.StartTime
	}

	// Vérifie si EndTime est fourni et met à jour si c'est le cas
	if updateRace.EndTime != nil {
		race.End_time = updateRace.EndTime
	}

	if updateRace.NumberOfCollisions != 0 {
		race.Number_of_collisions = updateRace.NumberOfCollisions
	}

	if updateRace.DistanceTravelled != 0 {
		race.Distance_travelled = updateRace.DistanceTravelled
	}

	if updateRace.AverageSpeed != 0 {
		race.Average_speed = updateRace.AverageSpeed
	}

	if updateRace.OutOfParcours != 0 {
		race.Out_of_parcours = updateRace.OutOfParcours
	}

	// Met à jour l'ID du véhicule si différent de zéro
	if updateRace.VehicleID != 0 {
		race.VehicleID = updateRace.VehicleID
	}

	return nil
}


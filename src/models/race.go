package models

import (
	"time"
)

type Race struct {
	ID                 uint       `gorm:"primaryKey"`
	VehicleID          uint       `json:"vehicle_id"`
	StartTime          time.Time  `json:"start_time"`
	EndTime            *time.Time `json:"end_time"`
	NumberOfCollisions uint8      `json:"number_of_collisions"`
	DistanceTravelled  int        `json:"distance_travelled"`
	AverageSpeed       int        `json:"average_speed"`
	OutOfParcours      uint8      `json:"out_of_parcours"`
	RaceType      		 string     	`json:"race_type"`		
	UserID             uint       `json:"user_id"`
	Model
}

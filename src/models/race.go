package models

import (
	"api/src/models/attributes"
)

type Race struct {
	ID                 uint                   `gorm:"primaryKey"`
	Name               string                 `json:"name"`
	VehicleID          uint                   `json:"vehicle_id"`
	Vehicle            Vehicle                `gorm:"foreignKey:VehicleID" json:"vehicle"`
	StartTime          attributes.CustomTime  `gorm:"type:datetime" json:"start_time"`
	EndTime            *attributes.CustomTime `gorm:"type:datetime" json:"end_time"`
	NumberOfCollisions uint8                  `json:"number_of_collisions"`
	DistanceTravelled  int                    `json:"distance_travelled"`
	AverageSpeed       int                    `json:"average_speed"`
	OutOfParcours      uint8                  `json:"out_of_parcours"`
	Status             string                 `json:"status"`
	Type               string                 `json:"type"`
	UserID             uint                   `json:"user_id"`
	User               User                   `gorm:"foreignKey:UserID" json:"-"`
	Model
}

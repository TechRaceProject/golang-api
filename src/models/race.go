package models

import (
	"api/src/models/attributes"
)

type Race struct {
	ID                uint                   `gorm:"primaryKey"`
	Name              string                 `json:"name"`
	VehicleID         uint                   `json:"vehicle_id"`
	Vehicle           Vehicle                `gorm:"foreignKey:VehicleID" json:"vehicle"`
	StartTime         *attributes.CustomTime `gorm:"type:datetime" json:"start_time"`
	EndTime           *attributes.CustomTime `gorm:"type:datetime" json:"end_time"`
	CollisionDuration int                    `gorm:"not null" json:"collision_duration"`
	DistanceCovered   int                    `gorm:"not null" json:"distance_covered"`
	AverageSpeed      float64                `gorm:"not null" json:"average_speed"`
	OutOfParcours     int                    `gorm:"not null" json:"out_of_parcours"`
	Status            string                 `gorm:"not null" json:"status"`
	Type              string                 `gorm:"not null" json:"type"`
	UserID            uint                   `json:"user_id"`
	User              User                   `gorm:"foreignKey:UserID" json:"user"`
	Model
}

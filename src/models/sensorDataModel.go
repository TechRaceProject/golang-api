package models

import "time"

type SensorData struct {
	ID    uint `gorm:"primaryKey"`
	Light int
	Sonar int
	Track int
	CreatedAt time.Time
}
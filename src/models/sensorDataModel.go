package models

import "time"

type SensorData struct {
	ID        uint `gorm:"primaryKey"`
	Light     float64
	Sonar     float64
	Track     float64
	CreatedAt time.Time
}

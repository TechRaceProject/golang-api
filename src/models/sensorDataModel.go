package models

// SensorData model
type SensorData struct {
	ID    uint `gorm:"primaryKey"`
	Light float64
	Sonar float64
	Track float64
	Model
}

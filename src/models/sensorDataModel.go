package models

import (
	validators "api/src/validators/sensorData"
)

// SensorData model
type SensorData struct {
	ID    uint `gorm:"primaryKey"`
	Light float64
	Sonar float64
	Track float64
	Model
}

func (s *SensorData) Create(createSensorData validators.CreateSensorDataValidator) error {
	if err := createSensorData.Validate(); err != nil {
		return err
	}
	s.Light = createSensorData.Light
	s.Sonar = createSensorData.Sonar
	s.Track = createSensorData.Track
	return nil
}

func (s *SensorData) Update(updateSensorData validators.CreateSensorDataValidator) error {
	if err := updateSensorData.Validate(); err != nil {
		return err
	}
	s.Light = updateSensorData.Light
	s.Sonar = updateSensorData.Sonar
	s.Track = updateSensorData.Track
	return nil
}

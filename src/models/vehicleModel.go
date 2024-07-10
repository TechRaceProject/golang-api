package models

import (
	validators "api/src/validators/vehicle"
)

type Vehicle struct {
	ID            uint    `gorm:"primaryKey"`
	Name          string  `json:"vehicle_name"`
	BatteryLife   float64 `json:"battery_life"`
	LineSensor1   bool    `json:"line_sensor1"`
	LineSensor2   bool    `json:"line_sensor2"`
	LineSensor3   bool    `json:"line_sensor3"`
	Camera        bool    `json:"camera"`
	SonarRange    float64 `json:"sonar_range"`
	WheelPower1   int     `json:"wheel_power1"`
	WheelPower2   int     `json:"wheel_power2"`
	WheelPower3   int     `json:"wheel_power3"`
	WheelPower4   int     `json:"wheel_power4"`
	LedColor      string  `json:"led_color"`
	DisplayPanel  string  `json:"display_panel"`
	SpeakerStatus bool    `json:"speaker_status"`
	SoundPlaying  string  `json:"sound_playing"`
	Model
}

func (v *Vehicle) Create(createVehicle validators.CreateVehicleValidator) error {
	if err := createVehicle.Validate(); err != nil {
		return err
	}
	v.Name = createVehicle.Name
	v.BatteryLife = createVehicle.BatteryLife
	v.LineSensor1 = createVehicle.LineSensor1
	v.LineSensor2 = createVehicle.LineSensor2
	v.LineSensor3 = createVehicle.LineSensor3
	v.Camera = createVehicle.Camera
	v.SonarRange = createVehicle.SonarRange
	v.WheelPower1 = createVehicle.WheelPower1
	v.WheelPower2 = createVehicle.WheelPower2
	v.WheelPower3 = createVehicle.WheelPower3
	v.WheelPower4 = createVehicle.WheelPower4
	v.LedColor = createVehicle.LedColor
	v.DisplayPanel = createVehicle.DisplayPanel
	v.SpeakerStatus = createVehicle.SpeakerStatus
	v.SoundPlaying = createVehicle.SoundPlaying
	return nil
}

func (v *Vehicle) Update(updateVehicle validators.CreateVehicleValidator) error {
	if err := updateVehicle.Validate(); err != nil {
		return err
	}
	v.Name = updateVehicle.Name
	v.BatteryLife = updateVehicle.BatteryLife
	v.LineSensor1 = updateVehicle.LineSensor1
	v.LineSensor2 = updateVehicle.LineSensor2
	v.LineSensor3 = updateVehicle.LineSensor3
	v.Camera = updateVehicle.Camera
	v.SonarRange = updateVehicle.SonarRange
	v.WheelPower1 = updateVehicle.WheelPower1
	v.WheelPower2 = updateVehicle.WheelPower2
	v.WheelPower3 = updateVehicle.WheelPower3
	v.WheelPower4 = updateVehicle.WheelPower4
	v.LedColor = updateVehicle.LedColor
	v.DisplayPanel = updateVehicle.DisplayPanel
	v.SpeakerStatus = updateVehicle.SpeakerStatus
	v.SoundPlaying = updateVehicle.SoundPlaying
	return nil
}

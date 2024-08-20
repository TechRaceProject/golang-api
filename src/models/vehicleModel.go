package models

import "gorm.io/gorm"

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

func (vehicle *Vehicle) InitVehicleState(user *User, db *gorm.DB) (VehicleState, error) {
	defaultUint := uint(0)
	defaultUint8 := uint8(0)
	defaultUint16 := uint16(0)

	vehicleState := VehicleState{
		VehicleID:      vehicle.ID,
		Face:           &defaultUint8,
		LedAnimation:   &defaultUint8,
		BuzzerAlarm:    &defaultUint8,
		VideoActivated: &defaultUint8,
		UserID:         user.ID,
	}

	if err := db.Create(&vehicleState).Error; err != nil {
		return VehicleState{}, err
	}

	primaryLedColor := PrimaryLedColor{
		LedIdentifier:  new(int),
		Red:            &defaultUint8,
		Green:          &defaultUint8,
		Blue:           &defaultUint8,
		VehicleStateID: vehicleState.ID,
	}

	if err := db.Create(&primaryLedColor).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.PrimaryLedColor = &primaryLedColor

	secondaryLedColor := SecondaryLedColor{
		BinaryRepresentation: new(int),
		Red:                  &defaultUint8,
		Green:                &defaultUint8,
		Blue:                 &defaultUint8,
		VehicleStateID:       vehicleState.ID,
	}

	if err := db.Create(&secondaryLedColor).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.SecondaryLedColor = &secondaryLedColor

	buzzerVariable := BuzzerVariable{
		Activated:      &defaultUint8,
		Frequency:      &defaultUint16,
		VehicleStateID: vehicleState.ID,
	}

	if err := db.Create(&buzzerVariable).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.BuzzerVariable = &buzzerVariable

	headAngle := HeadAngle{
		VerticalAngle:   &defaultUint,
		HorizontalAngle: &defaultUint,
		VehicleStateID:  vehicleState.ID,
	}

	if err := db.Create(&headAngle).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.HeadAngle = &headAngle

	db.Save(&vehicleState)

	return vehicleState, nil
}

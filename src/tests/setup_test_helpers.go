package tests

import (
	"api/src/models"

	"gorm.io/gorm"
)

func SetupTestVehicle(db *gorm.DB) *models.Vehicle {
	vehicle := &models.Vehicle{
		Name:          "Test Vehicle",
		BatteryLife:   100.0,
		LineSensor1:   true,
		LineSensor2:   false,
		LineSensor3:   true,
		Camera:        true,
		SonarRange:    50.0,
		WheelPower1:   90,
		WheelPower2:   80,
		WheelPower3:   70,
		WheelPower4:   60,
		LedColor:      "red",
		DisplayPanel:  "LCD",
		SpeakerStatus: true,
		SoundPlaying:  "test sound",
	}

	err := db.Create(&vehicle).Error
	if err != nil {
		panic(err)
	}

	return vehicle
}

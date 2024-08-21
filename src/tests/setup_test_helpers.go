package tests

import (
	"api/src/models"

	"gorm.io/gorm"
)

func SetupTestVehicle(db *gorm.DB) *models.Vehicle {
	vehicle := &models.Vehicle{
		Name:        "Test Vehicle",
		IpAdress:    "1.1.1.1",
		IsAvailable: true,
	}

	err := db.Create(&vehicle).Error

	if err != nil {
		panic(err)
	}

	return vehicle
}

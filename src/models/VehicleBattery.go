package models

type VehicleBattery struct {
	ID        uint    `gorm:"primaryKey"`
	VehicleID uint    `gorm:"not null" json:"vehicle_id"`
	Vehicle   Vehicle `gorm:"foreignKey:VehicleID" json:"vehicle"`
	Value     int     `gorm:"not null" json:"value"`
	Model
}

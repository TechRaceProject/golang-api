package models

type VehicleHistory struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	VehicleID uint    `gorm:"not null" json:"vehicle_id"`
	Vehicle   Vehicle `gorm:"foreignKey:VehicleID" json:"vehicle"`
	Message   string  `gorm:"not null" json:"message"`
	Model
}

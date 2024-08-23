package models

type VehicleBatteryModel struct {
	ID           uint    `gorm:"primaryKey"`
	VehicleID    uint    `gorm:"not null" json:"vehicle_id"`
	Vehicle      Vehicle `gorm:"foreignKey:VehicleID" json:"vehicle"`
	BatteryValue int     `gorm:"not null" json:"battery_value"`
	Model
}

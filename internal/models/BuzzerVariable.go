package models

type BuzzerVariable struct {
	ID             uint    `gorm:"primaryKey"`
	Activated      *uint8  `gorm:"not null" json:"activated"`
	Frequency      *uint16 `gorm:"not null" json:"frequency"`
	VehicleStateID uint    `gorm:"not null"`
	Model
}

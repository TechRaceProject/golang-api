package models

type PrimaryLedColor struct {
	ID             uint   `gorm:"primaryKey"`
	LedIdentifier  *int   `gorm:"not null" json:"led_identifier"`
	Red            *uint8 `gorm:"not null" json:"red"`
	Green          *uint8 `gorm:"not null" json:"green"`
	Blue           *uint8 `gorm:"not null" json:"blue"`
	VehicleStateID uint   `gorm:"not null"`
	Model
}

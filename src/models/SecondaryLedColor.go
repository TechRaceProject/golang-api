package models

import "gorm.io/gorm"

type SecondaryLedColor struct {
	ID                   uint   `gorm:"primaryKey"`
	BinaryRepresentation *int   `gorm:"not null" json:"binary_representation"`
	Red                  *uint8 `gorm:"not null" json:"red"`
	Green                *uint8 `gorm:"not null" json:"green"`
	Blue                 *uint8 `gorm:"not null" json:"blue"`
	VehicleStateID       uint   `gorm:"not null"`
	Model
}

func CreateSecondaryLedColor(db *gorm.DB, secondaryLedColor *SecondaryLedColor) (*SecondaryLedColor, error) {
	if err := db.Create(secondaryLedColor).Error; err != nil {
		return nil, err
	}

	return secondaryLedColor, nil
}

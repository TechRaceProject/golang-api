package models

type VehicleState struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	VehicleID uint `json:"-"`
	//Vehicle             Vehicle            `gorm:"foreignKey:VehicleID"`
	Face                *uint8             `gorm:"not null" json:"face"`
	LedAnimation        *uint8             `gorm:"not null" json:"led_animation"`
	BuzzerAlarm         *uint8             `gorm:"not null" json:"buzzer_alarm"`
	VideoActivated      *uint8             `gorm:"not null" json:"video_activated"`
	PrimaryLedColorID   *uint              `json:"-"`
	PrimaryLedColor     *PrimaryLedColor   `gorm:"foreignKey:PrimaryLedColorID"`
	SecondaryLedColorID *uint              `json:"-"`
	SecondaryLedColor   *SecondaryLedColor `gorm:"foreignKey:SecondaryLedColorID"`
	BuzzerVariableID    *uint              `json:"-"`
	BuzzerVariable      *BuzzerVariable    `gorm:"foreignKey:BuzzerVariableID"`
	HeadAngleID         *uint              `json:"-"`
	HeadAngle           *HeadAngle         `gorm:"foreignKey:HeadAngleID"`
	UserID              uint               `json:"-"`
	Model
}

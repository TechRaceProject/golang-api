package models

import "encoding/json"

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

func (vehicleState *VehicleState) ToJson() (string, error) {
	type vehicleStateJson struct {
		Type       string      `json:"type"`
		ID         uint        `json:"id"`
		Attributes interface{} `json:"attributes"`
	}

	type attributes struct {
		Face              *uint8      `json:"face"`
		LedAnimation      *uint8      `json:"led_animation"`
		BuzzerAlarm       *uint8      `json:"buzzer_alarm"`
		VideoActivated    *uint8      `json:"video_activated"`
		PrimaryLedColor   interface{} `json:"primary_led_color,omitempty"`
		SecondaryLedColor interface{} `json:"secondary_led_color,omitempty"`
		BuzzerVariable    interface{} `json:"buzzer_variable,omitempty"`
		HeadAngle         interface{} `json:"head_angle,omitempty"`
	}

	attr := attributes{
		Face:              vehicleState.Face,
		LedAnimation:      vehicleState.LedAnimation,
		BuzzerAlarm:       vehicleState.BuzzerAlarm,
		VideoActivated:    vehicleState.VideoActivated,
		PrimaryLedColor:   vehicleState.PrimaryLedColor,
		SecondaryLedColor: vehicleState.SecondaryLedColor,
		BuzzerVariable:    vehicleState.BuzzerVariable,
		HeadAngle:         vehicleState.HeadAngle,
	}

	vsJson := vehicleStateJson{
		Type:       "vehicle_state",
		ID:         vehicleState.ID,
		Attributes: attr,
	}

	jsonData, err := json.Marshal(vsJson)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

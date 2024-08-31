package models

import "encoding/json"

type VehicleState struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	VehicleID uint `json:"vehicle_id"`
	//Vehicle             Vehicle            `gorm:"foreignKey:VehicleID"`
	Face             *uint8            `gorm:"not null" json:"face"`
	LedAnimation     *uint8            `gorm:"not null" json:"led_animation"`
	BuzzerAlarm      *uint8            `gorm:"not null" json:"buzzer_alarm"`
	VideoActivated   *uint8            `gorm:"not null" json:"video_activated"`
	PrimaryLedColors []PrimaryLedColor `gorm:"foreignKey:VehicleStateID" json:"primary_led_colors"`
	BuzzerVariableID *uint             `json:"-"`
	BuzzerVariable   *BuzzerVariable   `gorm:"foreignKey:BuzzerVariableID"`
	HeadAngleID      *uint             `json:"-"`
	HeadAngle        *HeadAngle        `gorm:"foreignKey:HeadAngleID"`
	UserID           uint              `json:"-"`
	Model
}

func (vehicleState *VehicleState) ToJson() (string, error) {
	type vehicleStateJson struct {
		Type       string      `json:"type"`
		ID         uint        `json:"id"`
		Attributes interface{} `json:"attributes"`
	}

	type attributes struct {
		Face             *uint8      `json:"face"`
		LedAnimation     *uint8      `json:"led_animation"`
		BuzzerAlarm      *uint8      `json:"buzzer_alarm"`
		VideoActivated   *uint8      `json:"video_activated"`
		PrimaryLedColors interface{} `json:"primary_led_colors,omitempty"`
		BuzzerVariable   interface{} `json:"buzzer_variable,omitempty"`
		HeadAngle        interface{} `json:"head_angle,omitempty"`
	}

	attr := attributes{
		Face:             vehicleState.Face,
		LedAnimation:     vehicleState.LedAnimation,
		BuzzerAlarm:      vehicleState.BuzzerAlarm,
		VideoActivated:   vehicleState.VideoActivated,
		PrimaryLedColors: vehicleState.PrimaryLedColors,
		BuzzerVariable:   vehicleState.BuzzerVariable,
		HeadAngle:        vehicleState.HeadAngle,
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

package models

import "gorm.io/gorm"

type Vehicle struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	IpAdress    string `json:"ip_adress"`
	IsAvailable bool   `json:"is_available"`
	Model
}

func (vehicle *Vehicle) InitVehicleState(user *User, db *gorm.DB) (VehicleState, error) {
	defaultUint := uint(0)
	defaultUint8 := uint8(0)
	defaultUint16 := uint16(0)

	vehicleState := VehicleState{
		VehicleID:      vehicle.ID,
		Face:           &defaultUint8,
		LedAnimation:   &defaultUint8,
		BuzzerAlarm:    &defaultUint8,
		VideoActivated: &defaultUint8,
		UserID:         user.ID,
	}

	if err := db.Create(&vehicleState).Error; err != nil {
		return VehicleState{}, err
	}

	binaryRepresentations := []int{4097, 4098, 4100, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	primaryLedColors := []PrimaryLedColor{}

	for _, binaryRepresentation := range binaryRepresentations {
		primaryLedColor := PrimaryLedColor{
			LedIdentifier:  &binaryRepresentation,
			Red:            &defaultUint8,
			Green:          &defaultUint8,
			Blue:           &defaultUint8,
			VehicleStateID: vehicleState.ID,
		}

		primaryLedColors = append(primaryLedColors, primaryLedColor)
	}

	if err := db.Create(&primaryLedColors).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.PrimaryLedColors = primaryLedColors

	buzzerVariable := BuzzerVariable{
		Activated:      &defaultUint8,
		Frequency:      &defaultUint16,
		VehicleStateID: vehicleState.ID,
	}

	if err := db.Create(&buzzerVariable).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.BuzzerVariable = &buzzerVariable

	headAngle := HeadAngle{
		VerticalAngle:   &defaultUint,
		HorizontalAngle: &defaultUint,
		VehicleStateID:  vehicleState.ID,
	}

	if err := db.Create(&headAngle).Error; err != nil {
		return VehicleState{}, err
	}

	vehicleState.HeadAngle = &headAngle

	if err := db.Save(&vehicleState).Error; err != nil {
		return VehicleState{}, err
	}

	return vehicleState, nil
}

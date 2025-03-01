package models

type HeadAngle struct {
	ID              uint  `gorm:"primaryKey" json:"-"`
	VerticalAngle   *uint `gorm:"not null" json:"vertical_angle"`
	HorizontalAngle *uint `gorm:"not null" json:"horizontal_angle"`
	VehicleStateID  uint  `gorm:"not null" json:"-"`
}

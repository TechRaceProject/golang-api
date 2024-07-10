package models

type Fool struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name"`
	RaceID uint
	Race   Race `gorm:"foreignKey:RaceID"`
	Model
}

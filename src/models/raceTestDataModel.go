package models

type RaceTestData struct {
	ID       uint `gorm:"primaryKey"`
	Timer    float64
	Distance float64
	Model
}

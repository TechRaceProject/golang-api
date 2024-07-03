package models

import "time"

type SensorData struct {
    ID        uint `gorm:"primaryKey"`
    Topic     string
    Payload   string
    Timestamp time.Time
}

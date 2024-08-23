package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `gorm:"type:datetime"`
	UpdatedAt time.Time      `gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package models

import (
	"api/src/models/attributes"

	"gorm.io/gorm"
)

type Model struct {
	// json format YYYY-MM-DD HH:MM:SS
	CreatedAt attributes.CustomTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt attributes.CustomTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt        `gorm:"type:datetime;index" json:"deleted_at"`
}

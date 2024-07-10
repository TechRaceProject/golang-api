package models

import (
	validators "api/src/validators/fool"
)

type Fool struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	Model
}

func (f *Fool) Create(createFool validators.CreateFoolValidator) error {
	if err := createFool.Validate(); err != nil {
		return err
	}
	f.Name = createFool.Name
	return nil
}

func (f *Fool) Update(updateFool validators.CreateFoolValidator) error {
	if err := updateFool.Validate(); err != nil {
		return err
	}
	f.Name = updateFool.Name
	return nil
}

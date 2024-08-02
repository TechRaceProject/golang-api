package models

import (
	validators "api/src/validators/user"
)

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username *string `gorm:"unique" json:"username"`
	Email    string  `gorm:"unique;not null" json:"email"`
	Password string  `gorm:"not null" json:"-"`
	Model
}

func (u *User) Update(updateUser validators.UpdateUserValidator) {
	u.Username = &updateUser.Username

	u.Email = updateUser.Email

	if updateUser.Password != "" {
		u.Password = updateUser.Password
	}
}

func (u *User) Create(CreateUser validators.RegisterUserValidator) {
	u.Username = &CreateUser.Username

	u.Email = CreateUser.Email
	u.Password = CreateUser.Password
}

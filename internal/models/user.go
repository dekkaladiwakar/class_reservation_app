package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	RoleID       uint   `gorm:"not null"`
	Student      Student
	Auth         Auth
	Dean         Dean
}

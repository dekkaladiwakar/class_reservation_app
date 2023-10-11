package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Role Role
	Auth Auth
}

type Auth struct {
	gorm.Model
	Token        string
	PasswordHash string
	UserID       uint
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	RoleID  uint `gorm:"not null"`
	Student Student
	Auth    Auth
	Dean    Dean
}

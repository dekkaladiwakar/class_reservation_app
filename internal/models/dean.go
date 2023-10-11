package models

import "gorm.io/gorm"

type Dean struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	Name   string
}

package models

import "gorm.io/gorm"

type Dean struct {
	gorm.Model
	Name  string
	Auth  Auth   `gorm:"foreignKey:UserID"`
	Slots []Slot `gorm:"foreignKey:DeanID"`
}

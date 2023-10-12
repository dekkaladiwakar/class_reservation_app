package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	StudyYear int    `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
}

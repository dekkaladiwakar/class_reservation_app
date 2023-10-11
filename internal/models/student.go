package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StudyYear int    `gorm:"not null"` // Renamed from YearOfStudy for brevity
	Email     string `gorm:"unique;not null"`
	Auth      Auth   `gorm:"foreignKey:UserID"`
}

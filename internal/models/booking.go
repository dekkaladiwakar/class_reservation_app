package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	TimeslotID uint `gorm:"not null"`
	StudentID  uint `gorm:"not null"`
}

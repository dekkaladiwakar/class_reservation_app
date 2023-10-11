package models

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	DeanID    uint
	StartTime time.Time
	EndTime   time.Time
	Booked    bool
	StudentID *uint    `gorm:"index"`                // nullable foreign key to allow for unbooked slots
	Student   *Student `gorm:"foreignKey:StudentID"` // the student who booked the slot, if any
}

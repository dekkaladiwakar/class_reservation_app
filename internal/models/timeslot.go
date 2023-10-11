package models

import (
	"time"

	"gorm.io/gorm"
)

type Timeslot struct {
	gorm.Model
	DeanID    uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	Booked    bool      `gorm:"default:false"`
	Bookings  Booking
}

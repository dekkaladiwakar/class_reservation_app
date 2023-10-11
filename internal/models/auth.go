package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Token  string `gorm:"not null"`
	// Expiry time.Time `gorm:"not null"`
}

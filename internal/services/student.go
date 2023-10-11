package services

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/repositories"
	"gorm.io/gorm"
)

func CreateStudent() *gorm.DB {
	return repositories.CreateStudent()
}

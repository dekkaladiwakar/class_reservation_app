package services

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/repositories"
)

func CreateStudent() (uint, error) {
	return repositories.CreateStudent()
}

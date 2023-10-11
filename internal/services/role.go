package services

import "github.com/dekkaladiwakar/class-reservation-app/internal/repositories"

func CreateRole() (uint, string, error) {
	return repositories.CreateRole()
}

func GetRole(roleType string) (uint, error) {
	return repositories.GetRole(roleType)
}

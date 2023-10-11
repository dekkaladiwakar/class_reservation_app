package services

import "github.com/dekkaladiwakar/class-reservation-app/internal/repositories"

func CreateUser(roleType string) (uint, error) {
	id, err := GetRole(roleType)
	if err != nil {
		return 0, err
	}
	return repositories.CreateUser(id)
}

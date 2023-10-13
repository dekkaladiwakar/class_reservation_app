package services

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/repositories"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
)

func CreateUser(roleType string) (uint, error) {
	id, err := GetRole(roleType)
	if err != nil {
		return 0, err
	}

	temporaryPassword := "loginpswd"
	hash, err := scripts.ConvertStringToHash(temporaryPassword)
	if err != nil {
		return 0, err
	}

	return repositories.CreateUser(id, hash)
}

func GetLastUser() (uint, error) {
	return repositories.GetLastUser()
}

func VerifyPassword(userID uint, password string) (bool, error) {
	return repositories.VerifyPassword(userID, password)
}

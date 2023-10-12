package services

import "github.com/dekkaladiwakar/class-reservation-app/internal/repositories"

func CreateDean() (uint, error) {
	id, err := GetLastUser()

	if err != nil {
		return 0, err
	}

	return repositories.CreateDean(id)

}

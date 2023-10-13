package services

import (
	"fmt"

	"github.com/dekkaladiwakar/class-reservation-app/internal/repositories"
	"github.com/google/uuid"
)

func Login(ID uint, roleType string, password string) (string, error) {
	roleID, err := GetRole(roleType)
	if err != nil {
		return "", err
	}

	var userID uint

	if roleID == 1 { // Query the ID in Student's table
		userID, err = repositories.GetUserByStudentID(ID)
		if err != nil {
			return "", err
		}
	}

	if roleID == 2 { // Query the ID in Dean's table
		userID, err = repositories.GetUserByDeanID(ID)
		if err != nil {
			return "", err
		}
	}

	if userID == 0 {
		return "", fmt.Errorf("user not found for this ID %d. contact admin", ID)
	}

	fmt.Println("UserID is %d.", userID)

	// Verify Password
	_, err = VerifyPassword(userID, password)
	if err != nil {
		return "", err
	}

	// Generate UUID (alias auth.Token)
	uuid := uuid.New()

	// Save Token
	_, err = SetToken(userID, uuid.String())
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func SetToken(userID uint, token string) (bool, error) {
	return repositories.SetToken(userID, token)
}

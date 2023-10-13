package repositories

import (
	"errors"
	"fmt"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"gorm.io/gorm"
)

func CreateDean(userID uint) (uint, error) {
	name := scripts.GenerateRandomString(10)
	dean := models.Dean{UserID: userID, Name: name}

	result := db.DB.Create(&dean)
	if result.Error != nil {
		return 0, result.Error
	}

	return dean.ID, result.Error
}

func GetUserByDeanID(ID uint) (uint, error) {
	var dean models.Dean

	result := db.DB.First(&dean, ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("no dean found with ID %d", ID)
		}
		return 0, result.Error
	}

	return dean.UserID, nil
}

package repositories

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
)

func CreateDean(userID uint) (uint, error) {
	name := scripts.GenerateRandomString(10)
	dean := models.Dean{UserID: userID, Name: name}
	result := db.DB.Create(&dean)

	return dean.ID, result.Error
}

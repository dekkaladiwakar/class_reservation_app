package repositories

import (
	"errors"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"gorm.io/gorm"
)

func GetToken(userID uint, password string) (string, error) {
	auth := models.Auth{}

	result := db.DB.Where(models.Auth{UserID: userID}).Find(&auth)

	if result.RowsAffected == 0 {
		return "", errors.New("auth not found. contact admin")
	}

	return auth.Token, nil
}

func SetToken(userID uint, token string) (bool, error) {

	auth := models.Auth{UserID: userID}
	result := db.DB.First(&auth)

	var r *gorm.DB
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Add new token
		auth.Token = token
		r = db.DB.Create(&auth)
	} else if result.Error == nil {
		// Update token
		auth.Token = token
		r = db.DB.Save(&auth)
	}

	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

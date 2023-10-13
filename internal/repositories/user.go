package repositories

import (
	"errors"
	"fmt"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"gorm.io/gorm"
)

/*
* This route is trigger when a student or dean registers for the first time.
* Mostly for internal uses
 */
func CreateUser(roleID uint, hash string) (uint, error) {

	user := models.User{RoleID: roleID, PasswordHash: hash}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, result.Error

}

func GetLastUser() (uint, error) {
	user := models.User{}

	result := db.DB.Last(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("no users found")
		}
		return 0, result.Error
	}
	return user.ID, nil
}

func UpdatePassword(userID uint, password string) error {

	hash, err := scripts.ConvertStringToHash(password)
	if err != nil {
		return err
	}

	user := models.User{PasswordHash: hash}
	user.ID = userID

	result := db.DB.Update("password_hash", &user)

	return result.Error
}

func VerifyPassword(userID uint, password string) (bool, error) {
	var user models.User

	result := db.DB.First(&user, userID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return false, fmt.Errorf("no user found with ID %d", userID)
		}
		return false, result.Error
	}

	err := scripts.CompareHashAndPassword(user.PasswordHash, password)
	if err != nil {
		return false, errors.New("password is incorrect")
	}

	return true, nil
}

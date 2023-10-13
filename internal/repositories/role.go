package repositories

import (
	"errors"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"gorm.io/gorm"
)

func CreateRole() (uint, string, error) {
	role := models.Role{Name: "Dean"}

	result := db.DB.Create(&role)
	if result.Error != nil {
		return 0, "", result.Error
	}

	return role.ID, role.Name, result.Error
}

func GetRole(roleType string) (uint, error) {
	var role models.Role

	result := db.DB.Where("name = ?", roleType).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("no role found with the specified roleType")
		}
		return 0, result.Error
	}
	return role.ID, nil

}

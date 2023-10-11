package repositories

import (
	"errors"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
)

func CreateRole() (uint, string, error) {
	role := models.Role{Name: "Dean"}

	result := db.DB.Create(&role)

	return role.ID, role.Name, result.Error
}

func GetRole(roleType string) (uint, error) {
	role := models.Role{}

	result := db.DB.Where(models.Role{Name: roleType}).Find(&role)

	if result.RowsAffected == 0 {
		return 0, errors.New("no role found with the specified roleType")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return role.ID, nil

}

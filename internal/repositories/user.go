package repositories

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
)

/*
* This route is trigger when a student or dean registers for the first time.
* Mostly for internal uses
 */
func CreateUser(roleID uint) (uint, error) {

	user := models.User{RoleID: roleID}

	result := db.DB.Create(&user)

	return user.ID, result.Error

}

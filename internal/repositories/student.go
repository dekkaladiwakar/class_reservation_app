package repositories

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"gorm.io/gorm"
)

func CreateStudent(userID uint) (uint, error) {
	name := scripts.GenerateRandomString(10)
	student := models.Student{UserID: userID, Name: name, StudyYear: rand.Intn(4) + 1, Email: fmt.Sprintf("%s@gmail.com", name)}

	result := db.DB.Create(&student)
	if result.Error != nil {
		return 0, result.Error
	}
	return student.ID, result.Error
}

func GetUserByStudentID(ID uint) (uint, error) {
	var student models.Student

	result := db.DB.First(&student, ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("no student found with ID %d", ID)
		}
		return 0, result.Error
	}

	return student.UserID, nil
}

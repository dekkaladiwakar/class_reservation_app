package repositories

import (
	"fmt"
	"math/rand"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"gorm.io/gorm"
)

func CreateStudent() *gorm.DB {
	name := scripts.GenerateRandomString(10)
	student := models.Student{Name: name, StudyYear: rand.Intn(4) + 1, Email: fmt.Sprintf("%s@gmail.com", name)}

	result := db.DB.Create(&student)
	return result
}

package db

import (
	"fmt"
	"log"
	"os"

	"github.com/dekkaladiwakar/class-reservation-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbName, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database...")
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Student{}, &models.Dean{}, &models.Auth{}, &models.Timeslot{}, &models.Booking{})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/dekkaladiwakar/class-reservation-app/internal/router"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/config"
	"github.com/dekkaladiwakar/class-reservation-app/pkg/db"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}

func init() {
	config.LoadEnv()
	db.ConnectToDB()
}

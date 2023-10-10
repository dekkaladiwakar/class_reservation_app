package main

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/initializers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.GET("/ping", pingHandler)
	r.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	return router
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

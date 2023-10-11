package router

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/ping", pingHandler)

	router.POST("/role", handlers.CreateRole)
	router.POST("/user", handlers.CreateUser)

	studentGroup := router.Group("/student")

	studentGroup.POST("/", handlers.CreateStudent)

	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

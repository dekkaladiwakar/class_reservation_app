package router

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Health Check
	router.GET("/ping", pingHandler)

	// Admin Routes
	router.POST("/role", handlers.CreateRole)
	router.POST("/user", handlers.CreateUser)

	// Student Routes
	studentGroup := router.Group("/student")
	studentGroup.POST("", handlers.CreateStudent)

	// Dean Routes
	deanGroup := router.Group("/dean")
	deanGroup.POST("", handlers.CreateDean)

	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

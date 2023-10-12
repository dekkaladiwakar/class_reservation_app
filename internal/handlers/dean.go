package handlers

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateDean(c *gin.Context) {
	id, err := services.CreateDean()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

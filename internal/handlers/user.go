package handlers

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/services"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"github.com/gin-gonic/gin"
)

type CreateUserDto struct {
	RoleType string `json:"roleType" validate:"required"`
}

func CreateUser(c *gin.Context) {

	var user CreateUserDto

	// Parse the JSON request and bind it to Struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request body
	if err := scripts.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a User
	id, error := services.CreateUser(user.RoleType)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

package handlers

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/services"
	"github.com/dekkaladiwakar/class-reservation-app/scripts"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	ID       uint   `json:"id" validate:"required"`
	RoleType string `json:"roleType" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
	login := LoginDto{}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scripts.Validate(login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(login.ID, login.RoleType, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": login.ID, "token": token})
}

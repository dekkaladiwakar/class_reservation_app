package handlers

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	id, name, error := services.CreateRole()

	if error != nil {
		c.JSON(http.StatusBadRequest, error)
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "name": name})
}

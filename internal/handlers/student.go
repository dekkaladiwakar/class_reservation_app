package handlers

import (
	"net/http"

	"github.com/dekkaladiwakar/class-reservation-app/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	student := services.CreateStudent()
	if student.Error != nil {
		c.JSON(http.StatusBadRequest, student.Error)
	}
	c.JSON(http.StatusOK, student)
}

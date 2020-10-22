package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang_backend_gin/models"
	"golang_backend_gin/respositories"
	"net/http"
)

func CreateTeacher(c *gin.Context) {
	ID := uint(uuid.ClockSequence())
	Name := c.Request.FormValue("name")
	Gender := c.Request.FormValue("gender")
	Subject := c.Request.FormValue("subject")
	teacher := models.Teacher{
		ID:      ID,
		Name:    Name,
		Gender:  Gender,
		Subject: Subject,
	}

	result, err := respositories.CreateTeacher(teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, result)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"golang_backend_gin/controllers"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.POST("teacher", controllers.CreateTeacher)
	return router
}

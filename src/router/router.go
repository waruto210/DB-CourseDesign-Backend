package router

import (
	"db_course_design_backend/src/handler"
	"db_course_design_backend/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.JWT())

	apiv1 := router.Group("/api/v1")
	{
		apiv1.POST("login", handler.Login)
	}


	return router
}
package router

import (
	"db_course_design_backend/src/handler"
	"db_course_design_backend/src/middleware"
	"flag"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	if flag.Lookup("test.v") == nil {
		// if it is not in testing
		router.Use(middleware.JWT())
	}

	apiv1 := router.Group("/api/v1")
	{
		apiv1.POST("login", handler.Login)
	}


	apiv1.POST("student", handler.StudentCreate)
	apiv1.DELETE("student", handler.StudentDelete)
	apiv1.PUT("student", handler.StudentUpdate)
	apiv1.GET("student", handler.StudentQuery)

	return router
}
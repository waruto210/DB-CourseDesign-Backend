package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	CourseNo   string `json:"id"`
	CourseName string `json:"name"`
}

func GetCourse(c *gin.Context) {
	teacherNo := c.GetString(e.KEY_USER_ID)
	var courses []model.CourseInfo
	db.GetDB().Where(&model.CourseInfo{TeaNo: teacherNo}).Find(&courses)
	var retCourses []Course
	for _, course := range courses {
		c := Course{
			CourseNo:   course.CourseNo,
			CourseName: course.CourseName,
		}
		retCourses = append(retCourses, c)
	}
	result := model.GetResutByCode(e.SUCCESS)
	result.Data = retCourses
	c.JSON(http.StatusOK, result)
}

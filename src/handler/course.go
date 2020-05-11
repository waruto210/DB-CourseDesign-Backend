package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	CourseNo   string `json:"id"`
	CourseName string `json:"name"`
}
func GetCourse(c *gin.Context) {

	code := e.SUCCESS
	claims, _ := utils.ParseToken(c.Request.Header.Get(e.HEADER_AUTHORIZATION))
	teacherNo := claims.Audience
	if db.GetDB().Where(&model.TeacherInfo{TeaNo: teacherNo}).First(&model.TeacherInfo{}).RecordNotFound() {
		code = e.ERROR_USER_TYPE
		c.JSON(http.StatusOK, model.GetResutByCode(code))
		return
	}
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
	result := model.GetResutByCode(code)
	result.Data = retCourses
	c.JSON(http.StatusOK, result)
}
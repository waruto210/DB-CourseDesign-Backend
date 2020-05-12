package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentCourseCreate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.BindJSON(&parameter) != nil || ((model.StudentCourse{}) == parameter && parameter.Score != 0) {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&parameter).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
}

func StudentCourseUpdate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.BindJSON(&parameter) != nil || ((model.StudentCourse{}) == parameter && parameter.Score != 0) {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).First(&model.StudentCourse{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.StudentCourse{}).Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).Update(e.KEY_SCORE, parameter.Score)

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func StudentCourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)
	stuNo := c.Query(e.KEY_STU_NO)

	if courseNo == "" || stuNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.StudentCourse{
		CourseNo: courseNo,
		StuNo:    stuNo,
	})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func StudentCourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	stuNo, stuNoExist := c.GetQuery(e.KEY_STU_NO)

	var studentCourses []model.StudentCourse

	query := db.GetDB()
	if courseNoExist {
		query = query.Where(&model.StudentCourse{CourseNo: courseNo})
	}
	if stuNoExist {
		query = query.Where(&model.StudentCourse{StuNo: stuNo})
	}

	query.Find(&studentCourses)

	result := model.GetResutByCode(e.SUCCESS)
	result.Data = studentCourses
	c.JSON(http.StatusOK, result)
}

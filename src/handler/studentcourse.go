package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func StudentCourseCreate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.CourseNo == "" || parameter.StuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&parameter).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_STUDENT_COURSE_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

func StudentCourseUpdate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.CourseNo == "" || parameter.StuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).First(&model.StudentCourse{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_STUDENT_COURSE_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.StudentCourse{}).Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).Update(e.KEY_SCORE, parameter.Score)

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func StudentCourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)
	stuNo := c.Query(e.KEY_STU_NO)

	if courseNo == "" || stuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.StudentCourse{
		CourseNo: courseNo,
		StuNo:    stuNo,
	})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func StudentCourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	stuNo, stuNoExist := c.GetQuery(e.KEY_STU_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)

	var studentCourses []model.StudentCourse

	query := db.GetDB()
	if courseNoExist {
		query = query.Where(&model.StudentCourse{CourseNo: courseNo})
	}
	if stuNoExist {
		query = query.Where(&model.StudentCourse{StuNo: stuNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &studentCourses)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&studentCourses)
		result := model.GetResultByCode(e.SUCCESS)
		result.Data = studentCourses
		c.JSON(http.StatusOK, result)
	}
}

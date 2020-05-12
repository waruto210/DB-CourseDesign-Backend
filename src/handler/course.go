package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CourseCreate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.BindJSON(&parameter) != nil || (model.CourseInfo{}) == parameter {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&model.CourseInfo{CourseNo: parameter.CourseNo, CourseName: parameter.CourseName, TeaNo: parameter.TeaNo}).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
}

func CourseUpdate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.BindJSON(&parameter) != nil || (model.CourseInfo{}) == parameter {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).First(&model.CourseInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.CourseInfo{}).Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).Update(&model.CourseInfo{CourseName: parameter.CourseName, TeaNo: parameter.TeaNo})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func CourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)

	if courseNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.CourseInfo{
		CourseNo: courseNo,
	})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func CourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	var courses []model.CourseInfo

	query := db.GetDB()
	if courseNoExist {
		query = query.Where(&model.StudentCourse{CourseNo: courseNo})
	}

	query.Find(&courses)

	result := model.GetResutByCode(e.SUCCESS)
	result.Data = courses
	c.JSON(http.StatusOK, result)
}

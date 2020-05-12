package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseParameter struct {
	CourseNo   string `json:"course_no"`
	CourseName string `json:"course_name"`
	TeaNo      string `json:"tea_no"`
}

func CourseCreate(c *gin.Context) {
	parameter := CourseParameter{}

	if c.BindJSON(&parameter) != nil || parameter.CourseNo == "" || parameter.CourseName == "" || parameter.TeaNo == "" {
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
	parameter := CourseParameter{}

	if c.BindJSON(&parameter) != nil || parameter.CourseNo == "" || parameter.CourseName == "" || parameter.TeaNo == "" {
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
	// more parameters
	courseNo, exist := c.GetQuery(e.KEY_COURSE_NO)
	if exist {
		// query one person
		if courseNo == "" {
			c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
			return
		}
		course := model.CourseInfo{}
		if db.GetDB().Where(&model.CourseInfo{CourseNo: courseNo}).First(&course).RecordNotFound() {
			c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
			return
		}
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = []model.CourseInfo{course}
		c.JSON(http.StatusOK, result)
	} else {
		var courses []model.CourseInfo
		db.GetDB().Find(&courses)
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = courses
		c.JSON(http.StatusOK, result)
	}
}

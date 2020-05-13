package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CourseCreate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.BindJSON(&parameter) != nil || parameter.TeaNo == "" || parameter.CourseNo == "" || parameter.CourseName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&model.CourseInfo{CourseNo: parameter.CourseNo, CourseName: parameter.CourseName, TeaNo: parameter.TeaNo}).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_COURSE_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

func CourseUpdate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.BindJSON(&parameter) != nil || parameter.TeaNo == "" || parameter.CourseNo == "" || parameter.CourseName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).First(&model.CourseInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_COURSE_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.CourseInfo{}).Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).Update(&model.CourseInfo{CourseName: parameter.CourseName, TeaNo: parameter.TeaNo})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func CourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)

	if courseNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.CourseInfo{
		CourseNo: courseNo,
	})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func CourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)

	var courses []model.CourseInfo

	query := db.GetDB()
	if courseNoExist {
		query = query.Where(&model.StudentCourse{CourseNo: courseNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &courses)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&courses)
		result := model.GetResultByCode(e.SUCCESS)
		result.Data = courses
		c.JSON(http.StatusOK, result)
	}

}

package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentCreate(c *gin.Context) {
	student := model.StudentInfo{}

	if c.BindJSON(&student) != nil || student.StuNo == "" || student.StuName == "" || student.ClassNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&student).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
}

func StudentUpdate(c *gin.Context) {
	student := model.StudentInfo{}

	if c.BindJSON(&student) != nil || student.StuNo == "" || student.StuName == "" || student.ClassNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.StudentInfo{StuNo: student.StuNo}).First(&model.StudentInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&student).Where(&model.StudentInfo{StuNo: student.StuNo}).Update(&student)

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func StudentDelete(c *gin.Context) {
	stuNo := c.Query(e.KEY_STU_NO)

	if stuNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.StudentInfo{
		StuNo: stuNo,
	})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func StudentQuery(c *gin.Context) {
	// more parameters
	stuNo, exist := c.GetQuery(e.KEY_STU_NO)
	if exist {
		// query one person
		if stuNo == "" {
			c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
			return
		}
		student := model.StudentInfo{}
		if db.GetDB().Where(&model.StudentInfo{StuNo: stuNo}).First(&student).RecordNotFound() {
			c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
			return
		}
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = []model.StudentInfo{student}
		c.JSON(http.StatusOK, result)
	} else {
		var students []model.StudentInfo
		db.GetDB().Find(&students)
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = students
		c.JSON(http.StatusOK, result)
	}
}

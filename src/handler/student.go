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

	if c.BindJSON(&student) != nil || (model.StudentInfo{} == student) {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	// TODO use Transaction?
	if err := CreateUser(student.StuNo, model.USERTYPE_STUDENT); err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
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

	if c.BindJSON(&student) != nil || (model.StudentInfo{} == student) {
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
	stuNo, stuNoExist := c.GetQuery(e.KEY_STU_NO)
	var students []model.StudentInfo

	query := db.GetDB()
	if stuNoExist {
		query = query.Where(&model.StudentInfo{StuNo: stuNo})
	}

	query.Find(&students)

	result := model.GetResutByCode(e.SUCCESS)
	result.Data = students
	c.JSON(http.StatusOK, result)
}

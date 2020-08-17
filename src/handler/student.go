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

// 创建学生
func StudentCreate(c *gin.Context) {
	student := model.StudentInfo{}

	if c.ShouldBindBodyWith(&student, binding.JSON) != nil || student.StuNo == "" || student.StuName == "" || student.ClassNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	tx := db.GetDB().Begin()
	count := 0
	tx.Model(&model.ClassInfo{}).Where(&model.ClassInfo{ClassNo: student.ClassNo}).Count(&count)
	if count == 0 {
		tx.Rollback()
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_CLASS_NOT_EXIST))
		return
	}

	if err := CreateUser(tx, student.StuNo, model.USERTYPE_STUDENT); err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_EXIST))
		return
	}

	if err := tx.Create(&student).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_EXIST))
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

// 修改学生
func StudentUpdate(c *gin.Context) {
	student := model.StudentInfo{}

	if c.ShouldBindBodyWith(&student, binding.JSON) != nil || student.StuNo == "" || student.StuName == "" || student.ClassNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.StudentInfo{StuNo: student.StuNo}).First(&model.StudentInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&student).Where(&model.StudentInfo{StuNo: student.StuNo}).Update(&student)

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

// 删除学生
func StudentDelete(c *gin.Context) {
	stuNo := c.Query(e.KEY_STU_NO)

	if stuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Where(&model.User{
		UserId: stuNo,
	}).Delete(&model.User{})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

// 学生查询
func StudentQuery(c *gin.Context) {
	stuNo, stuNoExist := c.GetQuery(e.KEY_STU_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)
	var students []model.StudentInfo

	query := db.GetDB().Model(&model.StudentInfo{})
	if stuNoExist {
		query = query.Where(&model.StudentInfo{StuNo: stuNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &students)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&students)
		result := model.GetResultByCode(e.SUCCESS)
		result.Data = students
		c.JSON(http.StatusOK, result)
	}
}

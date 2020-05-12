package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TeacherCreate(c *gin.Context) {
	teacher := model.TeacherInfo{}

	if c.BindJSON(&teacher) != nil || (model.TeacherInfo{}) == teacher {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	// TODO use Transaction?
	if err := CreateUser(teacher.TeaNo, model.USERTYPE_TEACHER); err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
		return
	}

	if err := db.GetDB().Create(&teacher).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
}

func TeacherUpdate(c *gin.Context) {
	teacher := model.TeacherInfo{}

	if c.BindJSON(&teacher) != nil || (model.TeacherInfo{}) == teacher {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.TeacherInfo{TeaNo: teacher.TeaNo}).First(&model.TeacherInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&teacher).Where(&model.TeacherInfo{TeaNo: teacher.TeaNo}).Update(&teacher)

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func TeacherDelete(c *gin.Context) {
	teaNo := c.Query(e.KEY_TEA_NO)

	if teaNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.TeacherInfo{
		TeaNo: teaNo,
	})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func TeacherQuery(c *gin.Context) {
	teaNo, teaNoExist := c.GetQuery(e.KEY_TEA_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)
	var teachers []model.TeacherInfo

	query := db.GetDB()
	if teaNoExist {
		query = query.Where(&model.TeacherInfo{TeaNo: teaNo})
	}

	if pageExist {
		result := model.GetResutByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &teachers)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&teachers)
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = teachers
		c.JSON(http.StatusOK, result)
	}
}

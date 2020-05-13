package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClassCreate(c *gin.Context) {
	parameter := model.ClassInfo{}

	if c.BindJSON(&parameter) != nil || parameter.ClassNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	if err := db.GetDB().Create(&model.ClassInfo{ClassNo: parameter.ClassNo}).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_CLASS_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
}

func ClassDelete(c *gin.Context) {
	classNo := c.Query(e.KEY_CLASS_NO)

	if classNo == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Delete(&model.ClassInfo{
		ClassNo: classNo,
	})

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func ClassQuery(c *gin.Context) {
	classNo, courseNoExist := c.GetQuery(e.KEY_CLASS_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)

	var classes []model.ClassInfo

	query := db.GetDB()
	if courseNoExist {
		query = query.Where(&model.ClassInfo{ClassNo: classNo})
	}

	if pageExist {
		result := model.GetResutByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &classes)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&classes)
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = classes
		c.JSON(http.StatusOK, result)
	}
}

package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	userid := c.PostForm(e.KEY_USERID)
	passwd := c.PostForm(e.KEY_PASSWD)

	user := model.User{}
	if err := db.GetDB().Where(&model.User{UserId: userid}).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}
	if !utils.CheckPasswd(passwd, string(user.Passwd)) {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_PASSWD_NOT_MATCH))
		return
	}

	token, err := utils.GenerateToken(userid)
	if err != nil {
		log.Printf("cannot generate token for %s, because: %s", userid, err)
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR))
		return
	}
	result := model.GetResutByCode(e.SUCCESS)
	result.Data = gin.H{
		"token": token,
	}
	c.JSON(http.StatusOK, result)
}

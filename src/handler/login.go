package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"strconv"
)

type LoginParam struct {
	UserId string `json:"user_id"`
	Passwd string `json:"passwd"`
}

// 用户登录
func Login(c *gin.Context) {
	loginParam := LoginParam{}

	if c.ShouldBindBodyWith(&loginParam, binding.JSON) != nil || loginParam.UserId == "" || loginParam.Passwd == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}
	user := model.User{}

	// special admin
	if loginParam.UserId == "2017211000" && loginParam.Passwd == "2017211000" {
		user.UserId = "2017211000"
		user.UserType = model.USERTYPE_ADMIN
	} else {
		if err := db.GetDB().Where(&model.User{UserId: loginParam.UserId}).First(&user).Error; err != nil {
			c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_NOT_EXIST))
			return
		}
		if !utils.CheckPasswd(loginParam.Passwd, string(user.Passwd)) {
			c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_PASSWD_NOT_MATCH))
			return
		}
	}
	token, err := utils.GenerateToken(loginParam.UserId, strconv.Itoa(int(user.UserType)))
	if err != nil {
		log.Printf("cannot generate token for %s, because: %s", loginParam.UserId, err)
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR))
		return
	}
	result := model.GetResultByCode(e.SUCCESS)
	result.Data = gin.H{
		"token":         token,
		e.KEY_USER_ID:   loginParam.UserId,
		e.KEY_USER_TYPE: user.UserType,
	}
	c.JSON(http.StatusOK, result)
}

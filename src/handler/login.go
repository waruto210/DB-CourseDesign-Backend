package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type LoginParam struct {
	UserId string `json:"userid"`
	Passwd string `json:"passwd"`
}

func Login(c *gin.Context) {
	loginParam := LoginParam{}

	if c.BindJSON(&loginParam) != nil || loginParam.UserId == "" || loginParam.Passwd == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}
	user := model.User{}
	if err := db.GetDB().Where(&model.User{UserId: loginParam.UserId}).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}
	if !utils.CheckPasswd(loginParam.Passwd, string(user.Passwd)) {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_PASSWD_NOT_MATCH))
		return
	}

	token, err := utils.GenerateToken(loginParam.UserId, strconv.Itoa(int(user.UserType)))
	if err != nil {
		log.Printf("cannot generate token for %s, because: %s", loginParam.UserId, err)
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR))
		return
	}
	result := model.GetResutByCode(e.SUCCESS)
	result.Data = gin.H{
		"token":         token,
		e.KEY_USERID:    loginParam.UserId,
		e.KEY_USER_TYPE: user.UserType,
	}
	c.JSON(http.StatusOK, result)
}

package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(userid string, userType model.UserType) error {
	passwd, err := utils.HashPasswd(userid) // init passwd with userid
	if err != nil {
		return err
	}
	user := model.User{UserId: userid, UserType: userType, Passwd: []byte(passwd)}
	if err = db.GetDB().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

type UserParameter struct {
	UserId    string `json:"user_id"`
	Passwd    string `json:"passwd"`
	OldPasswd string `json:"old_passwd"`
}

func UserPasswdUpdate(c *gin.Context) {
	userType := c.GetInt(e.KEY_USER_TYPE)
	parameter := UserParameter{}

	if c.BindJSON(&parameter) != nil || parameter.UserId == "" || parameter.Passwd == "" {
		c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
		return
	}
	if userType != int(model.USERTYPE_ADMIN) { // if the current user is not admin, an old_passwd is required
		if parameter.OldPasswd == "" {
			c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
			return
		}
	}
	user := model.User{}
	if db.GetDB().Where(&model.User{UserId: parameter.UserId}).First(&user).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
		return
	}
	if userType != int(model.USERTYPE_ADMIN) {
		if !utils.CheckPasswd(parameter.Passwd, string(user.Passwd)) { // old_passwd is wrong
			c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_PASSWD_NOT_MATCH))
			return
		}
	}

	passwd, err := utils.HashPasswd(parameter.Passwd)
	if err != nil {
		c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR))
		return
	}
	user.Passwd = []byte(passwd)
	db.GetDB().Model(&model.User{}).Where(&model.User{UserId: parameter.UserId}).Update(&user)

	c.JSON(http.StatusOK, model.GetResutByCode(e.SUCCESS))
	return
}

func UserQuery(c *gin.Context) {
	// more parameters
	userId, exist := c.GetQuery(e.KEY_USER_ID)
	if exist {
		// query one person
		if userId == "" {
			c.JSON(http.StatusOK, model.GetResutByCode(e.INVALID_PARAMS))
			return
		}
		user := model.User{}
		if db.GetDB().Where(&model.User{UserId: userId}).First(&user).RecordNotFound() {
			c.JSON(http.StatusOK, model.GetResutByCode(e.ERROR_USER_NOT_EXIST))
			return
		}
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = []model.User{user}
		c.JSON(http.StatusOK, result)
	} else {
		var users []model.User
		db.GetDB().Find(&users)
		result := model.GetResutByCode(e.SUCCESS)
		result.Data = users
		c.JSON(http.StatusOK, result)
	}
}

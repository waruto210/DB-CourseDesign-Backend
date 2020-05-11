package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
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

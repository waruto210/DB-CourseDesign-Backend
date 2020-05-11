package model

import (
	"db_course_design_backend/src/utils/e"
)

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"msg" example:"请求信息"`
	Data    interface{} `json:"data" `
}

func GetResutByCode(code int) Result {
	return Result{Code: code, Message: e.GetMsg(code)}
}

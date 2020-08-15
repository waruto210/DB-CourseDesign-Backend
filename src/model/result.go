package model

import (
	"db_course_design_backend/src/utils/e"
)

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}

func GetResultByCode(code int) Result {
	return Result{Code: code, Message: e.GetMsg(code)}
}

type PagingData struct {
	Size  int         `json:"size"`  // Data size
	Total int         `json:"total"` // total number of pages
	Page  int         `json:"page"`  // current page number
	Data  interface{} `json:"data"`
}

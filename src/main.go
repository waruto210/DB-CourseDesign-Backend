package main

import (
	"db_course_design_backend/src/database"
	"db_course_design_backend/src/router"
)

func main() {
	// 初始化数据库连接
	database.Init()
	// 初始化路由
	r := router.SetUpRouter()
	// 启动服务
	_ = r.Run(":8080")
}

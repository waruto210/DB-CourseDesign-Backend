package main

import (
	"db_course_design_backend/src/database"
	"db_course_design_backend/src/router"
)

func main() {
	database.Init()
	r := router.SetUpRouter()
	_ = r.Run(":8080")
}

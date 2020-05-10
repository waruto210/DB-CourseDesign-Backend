package main

import "db_course_design_backend/src/router"

func main() {
	r := router.SetUpRouter()
	_ = r.Run(":8080")
}

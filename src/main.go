package main

import "db_course_design_backend/src/initRouter"

func main() {
	router := initRouter.SetUpRouter()
	_ = router.Run(":8080")
}

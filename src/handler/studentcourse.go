package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func StudentCourseCreate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.CourseNo == "" || parameter.StuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Model(&model.StudentInfo{}).Where(&model.StudentInfo{StuNo: parameter.StuNo}).First(&model.StudentInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_STUDENT_NOT_EXIST))
		return
	}

	if db.GetDB().Model(&model.CourseInfo{}).Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).First(&model.CourseInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_COURSE_NOT_EXIST))
		return
	}

	if err := db.GetDB().Create(&parameter).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_STUDENT_COURSE_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

func StudentCourseUpdate(c *gin.Context) {
	parameter := model.StudentCourse{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.CourseNo == "" || parameter.StuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).First(&model.StudentCourse{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_STUDENT_COURSE_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.StudentCourse{}).Where(&model.StudentCourse{StuNo: parameter.StuNo, CourseNo: parameter.CourseNo}).Update(e.KEY_SCORE, parameter.Score)

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func StudentCourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)
	stuNo := c.Query(e.KEY_STU_NO)

	if courseNo == "" || stuNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Where(&model.StudentCourse{
		CourseNo: courseNo,
		StuNo:    stuNo,
	}).Delete(&model.StudentCourse{})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

type StudentCourseInfo struct {
	StuNo       string `json:"stu_no"`
	CourseNo    string `json:"course_no"`
	Score       model.NullInt64    `json:"score"`
	CourseName  string `json:"course_name"`
	TeaName     string `json:"tea_name"`
	TeaNo       string `json:"tea_no"`
	StudentName string `json:"stu_name"`
}

func StudentCourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	stuNo, stuNoExist := c.GetQuery(e.KEY_STU_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)

	var studentCourses []model.StudentCourse

	query := db.GetDB().Model(&model.StudentCourse{})
	if courseNoExist {
		query = query.Where(&model.StudentCourse{CourseNo: courseNo})
	}
	if stuNoExist {
		query = query.Where(&model.StudentCourse{StuNo: stuNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &studentCourses)
		studentCoursesInfo := make([]StudentCourseInfo, len(studentCourses))
		for index, sc := range studentCourses {
			stu := model.StudentInfo{}
			db.GetDB().Where(&model.StudentInfo{StuNo: sc.StuNo}).Select("stu_name").First(&stu)
			course := model.CourseInfo{}
			db.GetDB().Where(&model.CourseInfo{CourseNo: sc.CourseNo}).Select("course_name, tea_no").First(&course)
			tea := model.TeacherInfo{}
			db.GetDB().Where(&model.TeacherInfo{TeaNo: course.TeaNo}).Select("tea_name").First(&tea)

			studentCoursesInfo[index] = StudentCourseInfo{
				StuNo:       sc.StuNo,
				CourseNo:    sc.CourseNo,
				Score:       sc.Score,
				CourseName:  course.CourseName,
				TeaName:     tea.TeaName,
				TeaNo:       course.TeaNo,
				StudentName: stu.StuName,
			}
		}
		payload.Data = studentCoursesInfo
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&studentCourses)
		studentCoursesInfo := make([]StudentCourseInfo, len(studentCourses))
		for index, sc := range studentCourses {
			stu := model.StudentInfo{}
			db.GetDB().Where(&model.StudentInfo{StuNo: sc.StuNo}).Select("stu_name").First(&stu)
			course := model.CourseInfo{}
			db.GetDB().Where(&model.CourseInfo{CourseNo: sc.CourseNo}).Select("course_name, tea_no").First(&course)
			tea := model.TeacherInfo{}
			db.GetDB().Where(&model.TeacherInfo{TeaNo: course.TeaNo}).Select("tea_name").First(&tea)

			studentCoursesInfo[index] = StudentCourseInfo{
				StuNo:       sc.StuNo,
				CourseNo:    sc.CourseNo,
				Score:       sc.Score,
				CourseName:  course.CourseName,
				TeaName:     tea.TeaName,
				TeaNo:       course.TeaNo,
				StudentName: stu.StuName,
			}
		}
		result := model.GetResultByCode(e.SUCCESS)
		result.Data = studentCoursesInfo
		c.JSON(http.StatusOK, result)
	}
}

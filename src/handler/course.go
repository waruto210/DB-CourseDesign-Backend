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

func CourseCreate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.TeaNo == "" || parameter.CourseNo == "" || parameter.CourseName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Model(&model.TeacherInfo{}).Where(&model.TeacherInfo{TeaNo: parameter.TeaNo}).First(&model.TeacherInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_TEACHER_NOT_EXIST))
		return
	}

	if err := db.GetDB().Create(&model.CourseInfo{CourseNo: parameter.CourseNo, CourseName: parameter.CourseName, TeaNo: parameter.TeaNo}).Error; err != nil {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_COURSE_EXIST))
		return
	}
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

func CourseUpdate(c *gin.Context) {
	parameter := model.CourseInfo{}

	if c.ShouldBindBodyWith(&parameter, binding.JSON) != nil || parameter.TeaNo == "" || parameter.CourseNo == "" || parameter.CourseName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).First(&model.CourseInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_COURSE_NOT_EXIST))
		return
	}

	db.GetDB().Model(&model.CourseInfo{}).Where(&model.CourseInfo{CourseNo: parameter.CourseNo}).Update(&model.CourseInfo{CourseName: parameter.CourseName, TeaNo: parameter.TeaNo})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

func CourseDelete(c *gin.Context) {
	courseNo := c.Query(e.KEY_COURSE_NO)

	if courseNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Where(&model.CourseInfo{
		CourseNo: courseNo,
	}).Delete(&model.CourseInfo{})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

type MoreCourseInfo struct {
	CourseNo   string `json:"course_no"`
	CourseName string `json:"course_name"`
	TeaNo      string `json:"tea_no"`
	TeaName    string `json:"tea_name"`
}

func CourseQuery(c *gin.Context) {
	courseNo, courseNoExist := c.GetQuery(e.KEY_COURSE_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)

	var courses []model.CourseInfo

	query := db.GetDB().Model(&model.CourseInfo{})
	if courseNoExist {
		query = query.Where(&model.CourseInfo{CourseNo: courseNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &courses)
		courseInfo := make([]MoreCourseInfo, len(courses))
		for index, c := range courses {
			tea := model.TeacherInfo{}
			db.GetDB().Where(&model.TeacherInfo{TeaNo: c.TeaNo}).Select("tea_name").First(&tea)

			courseInfo[index] = MoreCourseInfo{
				CourseNo:   c.CourseNo,
				CourseName: c.CourseName,
				TeaNo:      c.TeaNo,
				TeaName:    tea.TeaName,
			}
		}
		payload.Data = courseInfo
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&courses)
		result := model.GetResultByCode(e.SUCCESS)
		courseInfo := make([]MoreCourseInfo, len(courses))
		for index, c := range courses {
			tea := model.TeacherInfo{}
			db.GetDB().Where(&model.TeacherInfo{TeaNo: c.TeaNo}).Select("tea_name").First(&tea)

			courseInfo[index] = MoreCourseInfo{
				CourseNo:   c.CourseNo,
				CourseName: c.CourseName,
				TeaNo:      c.TeaNo,
				TeaName:    tea.TeaName,
			}
		}
		result.Data = courseInfo
		c.JSON(http.StatusOK, result)
	}

}

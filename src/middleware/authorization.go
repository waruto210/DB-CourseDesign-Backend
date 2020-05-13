package middleware

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

type StudentBody struct {
	StuNo string `json:"stu_no"`
}

type TeacherBody struct {
	TeaNo string `json:"tea_no"`
}

type UserBody struct {
	UserId string `json:"user_id"`
}

type StudentCourseBody struct {
	StuNo    string `json:"stu_no"`
	CourseNo string `json:"course_no"`
}

// access control
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString(e.KEY_USER_ID)
		userType := c.GetInt(e.KEY_USER_TYPE)

		if userType == int(model.USERTYPE_ADMIN) {
			// admin can do anything
			c.Next()
			return
		}

		if strings.HasSuffix(c.Request.URL.Path, "/login") {
			c.Next()
			return
		} else if strings.HasSuffix(c.Request.URL.Path, "/statistics") {
			c.Next()
			return
		} else if strings.HasSuffix(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_STUDENT) && c.Query(e.KEY_STU_NO) == userId {
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodPut {
			body := StudentBody{}
			if userType == int(model.USERTYPE_STUDENT) && c.ShouldBindBodyWith(&body, binding.JSON) == nil && body.StuNo == userId {
				// TODO add test for this
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/teacher") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_TEACHER) && c.Query(e.KEY_TEA_NO) == userId {
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/teacher") && c.Request.Method == http.MethodPut {
			body := TeacherBody{}
			if userType == int(model.USERTYPE_TEACHER) && c.ShouldBindBodyWith(&body, binding.JSON) == nil && body.TeaNo == userId {
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/course") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_TEACHER) && c.Query(e.KEY_TEA_NO) == userId { // teacher can only query courses that he teaches
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/user/passwd") && c.Request.Method == http.MethodPut {
			body := UserBody{}
			if c.ShouldBindBodyWith(&body, binding.JSON) == nil && body.UserId == userId { // each user can only change their own password
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/user") && c.Request.Method == http.MethodGet {
			if c.Query(e.KEY_USER_ID) == userId {
				c.Next()
				return
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/studentcourse") && c.Request.Method == http.MethodPut {
			if userType == int(model.USERTYPE_TEACHER) {
				body := StudentCourseBody{}
				if c.ShouldBindBodyWith(&body, binding.JSON) == nil && body.CourseNo != "" {
					if count := 0; db.GetDB().Where(model.CourseInfo{CourseNo: body.CourseNo, TeaNo: userId}).Count(&count).Error == nil && count > 0 {
						// teacher can only set score of a course which he teaches
						c.Next()
						return
					}
				}
			}
		} else if strings.HasSuffix(c.Request.URL.Path, "/studentcourse") && c.Request.Method == http.MethodGet {
			courseNo := c.Query(e.KEY_COURSE_NO)
			stuNo := c.Query(e.KEY_STU_NO)
			if userType == int(model.USERTYPE_STUDENT) && stuNo == userId {
				c.Next()
				return
			}
			if userType == int(model.USERTYPE_TEACHER) {
				if count := 0; db.GetDB().Where(model.CourseInfo{CourseNo: courseNo, TeaNo: userId}).Count(&count).Error == nil && count > 0 {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, model.GetResultByCode(e.ACCESS_FORBIDDEN))
		c.Abort()
	}
}

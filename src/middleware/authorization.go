package middleware

import (
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type StudentBody struct {
	StuNo string `json:"stu_no"`
}

type TeacherBody struct {
	TeaNo string `json:"tea_no"`
}

// access control
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString(e.KEY_USERID)
		userType := c.GetInt(e.KEY_USER_TYPE)

		if userType == int(model.USERTYPE_ADMIN) {
			// admin can do anything
			c.Next()
			return
		}

		// student
		if strings.Contains(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_STUDENT) && c.Query(e.KEY_STU_NO) == userId {
				c.Next()
				return
			}
		} else if strings.Contains(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodPut {
			body := StudentBody{}
			if userType == int(model.USERTYPE_STUDENT) && c.BindJSON(&body) == nil && body.StuNo == userId { // TODO add test for this
				c.Next()
				return
			}
		} else if strings.Contains(c.Request.URL.Path, "/teacher") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_TEACHER) && c.Query(e.KEY_TEA_NO) == userId {
				c.Next()
				return
			}
		} else if strings.Contains(c.Request.URL.Path, "/teacher") && c.Request.Method == http.MethodPut {
			body := TeacherBody{}
			if userType == int(model.USERTYPE_TEACHER) && c.BindJSON(&body) == nil && body.TeaNo == userId {
				c.Next()
				return
			}
		} else if strings.Contains(c.Request.URL.Path, "/course") && c.Request.Method == http.MethodGet {
			if userType == int(model.USERTYPE_TEACHER) {
				c.Next()
				return
			}
		} else if strings.Contains(c.Request.URL.Path, "/login") {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, model.GetResutByCode(e.ACCESS_FORBIDDEN))
		c.Abort()
	}
}

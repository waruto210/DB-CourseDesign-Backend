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

		if userType == int(model.USERTYPE_STUDENT) {
			// student
			if strings.Contains(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodGet {
				if c.Query(e.KEY_STU_NO) == userId {
					c.Next()
					return
				}
			} else if strings.Contains(c.Request.URL.Path, "/student") && c.Request.Method == http.MethodGet {
				body := StudentBody{}
				if c.BindJSON(&body) != nil { // TODO add test for this
					if body.StuNo == userId {
						c.Next()
						return
					}
				}
			} else if strings.Contains(c.Request.URL.Path, "/login") {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, model.GetResutByCode(e.ACCESS_FORBIDDEN))
		c.Abort()
	}
}

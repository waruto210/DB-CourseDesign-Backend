package middleware

import (
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := model.Result{
			Code:    e.SUCCESS,
			Message: "",
			Data:    nil,
		}
		jwtToken := c.Request.Header.Get(e.HEADER_AUTHORIZATION)
		if len(jwtToken) == 0 {
			result.Code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(jwtToken)
			if err != nil {
				result.Code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				result.Code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				c.Set(e.KEY_USERID, claims.Id)
			}
		}
		if result.Code != e.SUCCESS && !strings.Contains(c.Request.URL.Path, "login") {
			result.Message = e.GetMsg(result.Code)
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

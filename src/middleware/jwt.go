package middleware

import (
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := model.Result{
			Code:    e.SUCCESS,
			Message: "",
			Data:    nil,
		}
		jwtToken := c.Request.Header.Get("Authorization")
		if len(jwtToken) == 0 {
			result.Code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(jwtToken)
			if err != nil {
				result.Code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				result.Code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if result.Code != e.SUCCESS {
			result.Message = e.GetMsg(result.Code)
			c.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
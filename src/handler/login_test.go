package handler_test

import (
	"bytes"
	"db_course_design_backend/src/database"
	"db_course_design_backend/src/router"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	database.Init()
	router := router.SetUpRouter()

	w := httptest.NewRecorder()
	body, err:=jsoniter.MarshalToString(gin.H{
		e.KEY_USER_ID: "user_id",
		e.KEY_PASSWD: "passwd",
	})
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
	assert.Equal(t, e.ERROR_USER_NOT_EXIST, jsoniter.Get([]byte(w.Body.String()), "code").ToInt())

}

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

func TestStudentCreate(t *testing.T) {
	database.Init()
	r := router.SetUpRouter()

	w := httptest.NewRecorder()
	body, _ := jsoniter.MarshalToString(gin.H{
		e.KEY_STU_NO:   "2017211000",
		e.KEY_STU_NAME: "张三",
		e.KEY_CLASS_NO: "2017211111",
	})
	req, _ := http.NewRequest("POST", "/api/v1/student", bytes.NewBufferString(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}
func TestStudentQuery(t *testing.T) {
	database.Init()
	r := router.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/student", nil)
	q := req.URL.Query()
	q.Set(e.KEY_STU_NO, "2017211000")
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}

func TestStudentUpdate(t *testing.T) {
	database.Init()
	r := router.SetUpRouter()

	w := httptest.NewRecorder()
	body, _ := jsoniter.MarshalToString(gin.H{
		e.KEY_STU_NO:   "2017211000",
		e.KEY_STU_NAME: "张四",
		e.KEY_CLASS_NO: "2017211111",
	})
	req, _ := http.NewRequest("PUT", "/api/v1/student", bytes.NewBufferString(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}


func TestStudentDelete(t *testing.T) {
	database.Init()
	r := router.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/student", nil)
	q := req.URL.Query()
	q.Set(e.KEY_STU_NO, "2017211000")
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}



package handler_test

import (
	"db_course_design_backend/src/database"
	"db_course_design_backend/src/router"
	"db_course_design_backend/src/utils/e"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStatistic(t *testing.T) {
	database.Init()
	r := router.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/statistics", nil)
	q := req.URL.Query()
	q.Set(e.KEY_COURSE_NO, "20200003")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}

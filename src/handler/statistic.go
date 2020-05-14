package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/grd/statistics"
	"net/http"
)

type ScoreMap struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

func GetStatistic(c *gin.Context) {
	code := e.SUCCESS
	courseNo := c.Query(e.KEY_COURSE_NO)
	if  courseNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
	}
	var scores []model.StudentCourse
	db.GetDB().Where(&model.StudentCourse{CourseNo: courseNo}).Where("score <> ?", nil).Find(&scores)
	var retScores []ScoreMap
	retScores = append(retScores, ScoreMap{
		Name:  "100",
		Value: 0,
	})
	retScores = append(retScores, ScoreMap{
		Name:  "90-99",
		Value: 0,
	})
	retScores = append(retScores, ScoreMap{
		Name:  "80-89",
		Value: 0,
	})
	retScores = append(retScores, ScoreMap{
		Name:  "70-79",
		Value: 0,
	})
	retScores = append(retScores, ScoreMap{
		Name:  "60-69",
		Value: 0,
	})
	retScores = append(retScores, ScoreMap{
		Name:  "under_60",
		Value: 0,
	})

	var s statistics.Int64
	for _, score := range scores {
		s = append(s, int64(score.Score.Int64))
		if score.Score.Int64 == 100 {
			retScores[0].Value += 1
		} else if score.Score.Int64 >= 90 {
			retScores[1].Value += 1
		} else if score.Score.Int64 >= 80 {
			retScores[2].Value += 1
		} else if score.Score.Int64 >= 70 {
			retScores[3].Value += 1
		} else if score.Score.Int64 >= 60 {
			retScores[4].Value += 1
		} else {
			retScores[5].Value += 1
		}
	}

	variance := statistics.Variance(&s)
	max, _ := statistics.Max(&s)
	min, _ := statistics.Min(&s)
	rangeDiff := max - min
	result := model.GetResultByCode(code)
	result.Data = retScores
	c.JSON(http.StatusOK, gin.H{
		"code": result.Code,
		"message": result.Message,
		"data": result.Data,
		"variance": variance,
		"range_diff": rangeDiff,
	})
}
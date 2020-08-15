package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grd/statistics"
)

// ScoreMap 记录对应分数段和人数的Map
type ScoreMap struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// GetStatistic 返回一门课的统计信息
func GetStatistic(c *gin.Context) {
	code := e.SUCCESS
	courseNo := c.Query(e.KEY_COURSE_NO)
	log.Println("course_no:", courseNo)
	if courseNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
	}
	var scores []model.StudentCourse

	db.GetDB().Where(&model.StudentCourse{CourseNo: courseNo}).Where("score is not null").Find(&scores)
	log.Println("scores:", scores)
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
	// 统计各分段人数
	var s statistics.Int64
	for _, score := range scores {
		s = append(s, int64(score.Score.Int64))
		if score.Score.Int64 == 100 {
			retScores[0].Value++
		} else if score.Score.Int64 >= 90 {
			retScores[1].Value++
		} else if score.Score.Int64 >= 80 {
			retScores[2].Value++
		} else if score.Score.Int64 >= 70 {
			retScores[3].Value++
		} else if score.Score.Int64 >= 60 {
			retScores[4].Value++
		} else {
			retScores[5].Value++
		}
	}
	// 极差和方差
	var rangeDiff float64 = -1
	var variance float64 = -1

	if len(scores) > 1 {
		variance = statistics.Variance(&s)
	}
	if len(scores) > 0 {
		max, _ := statistics.Max(&s)
		min, _ := statistics.Min(&s)
		rangeDiff = max - min
	}
	result := model.GetResultByCode(code)
	result.Data = retScores

	c.JSON(http.StatusOK, gin.H{
		"code":       result.Code,
		"message":    result.Message,
		"data":       result.Data,
		"variance":   variance,
		"range_diff": rangeDiff,
	})
}

package utils

import (
	"db_course_design_backend/src/config"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils/e"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	bcrypt "golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

func GenerateToken(userid, usetype string) (string, error) {
	nowTime := time.Now().Unix()
	expireTime := nowTime + int64(config.Duration)
	claims := jwt.StandardClaims{
		Audience:  usetype,
		ExpiresAt: expireTime,
		Id:        userid,
		IssuedAt:  nowTime,
		Issuer:    config.Issuer,
		NotBefore: nowTime,
		Subject:   config.Subject,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(config.Secret)
	log.Println("token: ", token)
	return token, err
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.Secret, nil
		})
	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func HashPasswd(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func CheckPasswd(passwd string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwd))
	if err != nil {
		return false
	}
	return true
}

func GenPagePayload(query *gorm.DB, page string, container interface{}) *model.PagingData {
	var count int
	query.Count(&count)
	pageSize := e.VALUE_PAGE_SIZE_DEFAULT
	total := (count + pageSize - 1) / pageSize
	pageNum, _ := strconv.Atoi(page)
	if pageNum > total {
		pageNum = total
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize
	limit := count - offset
	if limit > e.VALUE_PAGE_SIZE_DEFAULT {
		limit = e.VALUE_PAGE_SIZE_DEFAULT
	}
	query.Offset(offset).Limit(limit).Find(container)
	payload := model.PagingData{
		Size:  limit,
		Total: total,
		Page:  pageNum,
		Data:  container,
	}
	return &payload
}

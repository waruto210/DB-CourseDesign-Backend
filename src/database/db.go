package database

import (
	"db_course_design_backend/src/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

func Init() {
	var err error = nil
	// export DBURL="db_class:dbclassmm@/student_score?charset=utf8&parseTime=True&loc=Local"
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		panic("${DBURL} was not set")
	}
	db, err = gorm.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
	if !db.HasTable(&model.StudentInfo{}) {
		db.CreateTable(&model.StudentInfo{})
	}
	if !db.HasTable(&model.TeacherInfo{}) {
		db.CreateTable(&model.TeacherInfo{})
	}
	if !db.HasTable(&model.ClassInfo{}) {
		db.CreateTable(&model.ClassInfo{})
	}
	if !db.HasTable(&model.CourseInfo{}) {
		db.CreateTable(&model.CourseInfo{})
	}
	if !db.HasTable(&model.StudentCourse{}) {
		db.CreateTable(&model.StudentCourse{})
	}
	if !db.HasTable(&model.Admin{}) {
		db.CreateTable(&model.Admin{})
	}
}

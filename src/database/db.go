package database

import (
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
	// 从环境变量中获取数据库链接
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		panic("${DBURL} was not set")
	}
	db, err = gorm.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

}

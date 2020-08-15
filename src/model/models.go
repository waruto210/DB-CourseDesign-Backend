package model

type UserType byte

const (
	USERTYPE_STUDENT UserType = 1
	USERTYPE_TEACHER UserType = 2
	USERTYPE_ADMIN   UserType = 3
)

type User struct {
	UserId   string   `json:"user_id" gorm:"column:user_id;primary_key;not null"`
	UserType UserType `json:"user_type" gorm:"column:user_type"`
	Passwd   []byte   `json:"-" gorm:"column:passwd"`
}

type StudentInfo struct {
	StuNo   string `json:"stu_no" gorm:"column:stu_no;not null"`
	StuName string `json:"stu_name" gorm:"column:stu_name"`
	ClassNo string `json:"class_no" gorm:"column:class_no"`
}

type TeacherInfo struct {
	TeaNo   string `json:"tea_no" gorm:"column:tea_no;not null"`
	TeaName string `json:"tea_name" gorm:"column:tea_name"`
}

type ClassInfo struct {
	ClassNo string `json:"class_no" gorm:"column:class_no;primary_key;not null"`
}

type CourseInfo struct {
	CourseNo   string `json:"course_no" gorm:"column:course_no;primary_key;not null"`
	CourseName string `json:"course_name" gorm:"column:course_name"`
	TeaNo      string `json:"tea_no" gorm:"column:tea_no"`
}

type StudentCourse struct {
	StuNo    string    `json:"stu_no" gorm:"column:stu_no;not null"`
	CourseNo string    `json:"course_no" gorm:"column:course_no;not null"`
	Score    NullInt64 `json:"score" gorm:"column:score"`
}

type Admin struct {
	AdminNo string `gorm:"column:admin_no;not null;"`
}

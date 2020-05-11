package model

type UserType byte

const (
	USERTYPE_STUDENT UserType = 0
	USERTYPE_TEACHER UserType = 1
	USERTYPE_ADMIN   UserType = 2
)

type User struct {
	UserId   string   `gorm:"column:user_id;primary_key;not null"`
	UserType UserType `gorm:"column:user_type"`
	Passwd   []byte   `gorm:"column:passwd"`
}

type StudentInfo struct {
	StuNo   string `json:"stu_no";gorm:"column:stu_no;not null"`
	StuName string `json:"stu_name";gorm:"column:stu_name"`
	ClassNo string `json:"class_no";gorm:"column:class_no"`
}

type TeacherInfo struct {
	TeaNo   string `gorm:"column:tea_no;not null"`
	TeaName string `gorm:"column:tea_name"`
}

type ClassInfo struct {
	ClassNo string `gorm:"column:class_no;primary_key;not null"`
}

type CourseInfo struct {
	CourseNo string `gorm:"column:course_no;primary_key;not null"`
	TeaNo    string `gorm:"column:tea_no"`
	CourseName string `gorm:"column:course_name"`
}

type StudentCourse struct {
	StuNo    string `gorm:"column:stu_no;not null"`
	CourseNo string `gorm:"column:course_no;not null"`
	Score    int    `gorm:"column:score"`
}

type Admin struct {
	AdminNo string `gorm:"column:admin_no;not null;"`
}

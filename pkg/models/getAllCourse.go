package models

func GetAllCourses() []Course {
	var courses []Course
	DB.Find(&courses)
	return courses
}
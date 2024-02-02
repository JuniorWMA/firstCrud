package models


func GetAllCourse() []Course {
	var Books []Course
	db.Find(&Books)
	return Books
}
package models

func DeleteCourse(Id int) Course{
	var course Course
	db.Where("Id=?", Id).Delete(&course)
	return course
}
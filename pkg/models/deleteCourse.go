package models

func DeleteCourse(Id int) Course{
	var course Course
	DB.Where("Id=?", Id).Delete(&course)
	return course
}
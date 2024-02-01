package models

import "github.com/jinzhu/gorm"

func GetByIdCourse(Id int) (*Course, *gorm.DB) {
	var course Course
	db := DB.Where("Id=?", Id).Find(&course)
	return &course, db
}
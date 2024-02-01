package models


func (c *Course) CreateCourse() *Course {
	DB.NewRecord(c)
	DB.Create(&c)
	return c
}
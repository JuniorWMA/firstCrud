package models


func (b *Course) CreateBCourse() *Course {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Course struct {
	gorm.Model
	Name     string  `gorm:""json:"name,omitempty"`
	Category string  `json:"category,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Duration int     `json:"duration,omitempty"`
}

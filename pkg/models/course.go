package models

import (
	"github.com/jinzhu/gorm"
	"github.com/juniorWMA/firstCrud/pkg/configs"
)

var db *gorm.DB


func init()  {
	configs.Connect()
	db = configs.GetDB()
	db.AutoMigrate(&Course{})
	
}

type Course struct {
	gorm.Model
	Name     string  `gorm:""json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Duration int     `json:"duration"`
}



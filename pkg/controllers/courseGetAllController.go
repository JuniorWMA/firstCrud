package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/juniorWMA/firstCrud/pkg/models"
)

var NewCourse models.Course

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	
	courses := models.GetAllCourses()
	res, err := json.Marshal(courses)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
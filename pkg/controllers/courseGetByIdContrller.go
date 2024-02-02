package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juniorWMA/firstCrud/pkg/models"
)

func CourseById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["Id"])
	if err != nil {
		log.Fatal("Error in parse:", err)
	}
	course, _ := models.GetByIdCourse(ID)
	courseId, err := json.Marshal(course)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(courseId)

}

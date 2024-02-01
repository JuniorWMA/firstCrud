package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/juniorWMA/firstCrud/pkg/models"
	"github.com/juniorWMA/firstCrud/pkg/utils"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	course := &models.Course{}
	utils.ParseBody( r,  course)
	c := course.CreateCourse()
	res, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}
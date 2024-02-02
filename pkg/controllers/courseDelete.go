package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juniorWMA/firstCrud/pkg/models"
)

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCourse, err := strconv.Atoi(params["Id"])
	if err != nil {
		log.Fatal(err)
	}
	course := models.DeleteCourse(idCourse)
	res, _ := json.Marshal(course)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
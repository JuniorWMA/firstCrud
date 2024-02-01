package routes

import (
	"github.com/gorilla/mux"
	"github.com/juniorWMA/firstCrud/pkg/controllers"
)

var GetRoutesCourse = func(r *mux.Router) {
	r.HandleFunc("/courses", controllers.GetAllCourses).Methods("GET")
	r.HandleFunc("/course", controllers.CreateCourse).Methods("POST")
}

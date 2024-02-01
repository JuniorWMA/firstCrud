package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juniorWMA/firstCrud/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.GetRoutesCourse(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
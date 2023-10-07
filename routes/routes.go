package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusrmello/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}


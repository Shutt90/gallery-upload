package routes

import (
	"github.com/gorilla/mux"
	"github.com/shutt90/gallery-upload/controllers"
)

func Routes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeController).Methods("GET")

}

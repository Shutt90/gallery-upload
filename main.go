package main

import (
	"log"
	"net/http"

	"github.com/shutt90/gallery-upload/routes"

)

func main() {

	// http.Handle(http.NewServeMux().ServeHTTP("/", routes.Routes), nil)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
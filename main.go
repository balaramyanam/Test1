package main

import (
	"log"
	"net/http"
	"instance/controllers"

	"github.com/gorilla/mux"
)

func main() {
	ret := controllers.Createinstance()

	if ret != nil {

	}

	my := mux.NewRouter().StrictSlash(true)

	my.HandleFunc("/Idilbatch", controllers.Search).Methods("POST")

	log.Fatal(http.ListenAndServe(":5004", my))

}

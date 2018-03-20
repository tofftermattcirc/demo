package main

import (
	"log"
	"net/http"

	"circadence/demo/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	controller := controllers.NewHeaderController()
	controller.RegisterRouter(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Serving")
	server.ListenAndServe()
}

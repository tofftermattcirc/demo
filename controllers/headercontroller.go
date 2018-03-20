package controllers

import (
	"log"
	"net/http"

	//_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"circadence/demo/services"
)

const (
	DefaultRootPath    = "/"
	RegisterHeaderPath = "/api/headers"
)

type HeaderController struct {
	HeaderService *services.HeaderService
}

func NewHeaderController() *HeaderController {
	return &HeaderController{
		HeaderService: services.NewHeaderService(),
	}
}

/** POST **/
func (ref *HeaderController) RegisterHeader(w http.ResponseWriter, r *http.Request) {
	log.Printf("HeaderController: RegisterHeader: start")
	servicedData := ref.HeaderService.RegisterHeaderServicer(r.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(servicedData)

	log.Printf("HeaderController: RegisterHeader: end")

}

func (ref *HeaderController) RegisterRouter(router *mux.Router) {
	log.Printf("HeaderController: RegisterRouter: start")

	router.HandleFunc(RegisterHeaderPath, ref.RegisterHeader).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static/")))

	log.Printf("HeaderController: RegisterRouter: end")
}

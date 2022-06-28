package app

import (
	"vnia-auth-session/controller"

	"github.com/gorilla/mux"
)

func NewRouter(control controller.ControllerUsers) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", control.Home)
	router.HandleFunc("/register", control.Register).Methods("POST")
	router.HandleFunc("/login", control.Login).Methods("POST")

	return router
}

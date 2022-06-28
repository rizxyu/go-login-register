package controller

import "net/http"

type ControllerUsers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}

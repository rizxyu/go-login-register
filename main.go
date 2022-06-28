package main

import (
	"log"
	"net/http"
	"vnia-auth-session/app"
	"vnia-auth-session/controller"
)

func main() {
	db := app.NewDB()
	controllers := controller.NewControllerUsers(db)
	log.Println("Server Berjalan Di *:8000")
	http.ListenAndServe(":8000", app.NewRouter(controllers))
}

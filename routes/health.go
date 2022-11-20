package routes

import (
	"github.com/RaeedAsif/flare-go-test/controllers"
	"github.com/gorilla/mux"
)

func ServerHealthRoute(router *mux.Router) {
	router.HandleFunc("/", controllers.GetServerHealth()).Methods("GET")
}

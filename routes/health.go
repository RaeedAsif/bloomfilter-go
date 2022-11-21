package routes

import (
	"github.com/gorilla/mux"

	"github.com/RaeedAsif/flare-go-test/controllers"
)

func ServerHealthRoute(router *mux.Router) {
	router.HandleFunc("/", controllers.GetServerHealth()).Methods("GET")
}

package routes

import (
	"github.com/gorilla/mux"

	"github.com/RaeedAsif/bloomfilter-go/controllers"
)

// ServerHealthRoute serves health route
func ServerHealthRoute(router *mux.Router) {
	router.HandleFunc("/health", controllers.GetServerHealth()).Methods("GET")
}

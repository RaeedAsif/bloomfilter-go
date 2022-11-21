package routes

import (
	"github.com/gorilla/mux"

	"github.com/RaeedAsif/flare-go-test/controllers"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{username}", controllers.GetUser()).Methods("GET")
}

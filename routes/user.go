package routes

import (
	"github.com/RaeedAsif/flare-go-test/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{username}", controllers.GetUser()).Methods("GET")
}

package routes

import (
	"github.com/gorilla/mux"

	"github.com/RaeedAsif/flare-go-test/controllers"
)

// UserRoute serves user routes
func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/username/{username}", controllers.IsUsernameExists()).Methods("GET")
}

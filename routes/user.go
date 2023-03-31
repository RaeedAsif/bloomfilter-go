package routes

import (
	"github.com/gorilla/mux"

	"github.com/RaeedAsif/bloomfilter-go/controllers"
)

// UserRoute serves user routes
func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/username/{username}", controllers.IsUsernameExists()).Methods("GET")
}

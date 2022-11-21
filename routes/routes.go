package routes

import (
	"github.com/gorilla/mux"
)

// Init iniates NewRouter
func Init() *mux.Router {
	//router
	router := mux.NewRouter()

	//routes
	UserRoute(router)
	ServerHealthRoute(router)
	SwaggerRoute(router)

	return router
}

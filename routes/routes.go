package routes

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	//router
	router := mux.NewRouter()

	//routes
	UserRoute(router)
	ServerHealthRoute(router)

	return router
}
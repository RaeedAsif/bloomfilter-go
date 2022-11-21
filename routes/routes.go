package routes

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// Init iniates NewRouter
func Init() *mux.Router {
	//router
	router := mux.NewRouter()

	//routes
	UserRoute(router)
	ServerHealthRoute(router)

	// swagger documentation serve
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	router.Handle("/docs", sh)

	return router
}

package routes

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// SwaggerRoute server swagger documentation URI
func SwaggerRoute(router *mux.Router) {
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	router.Handle("/docs", sh)
}

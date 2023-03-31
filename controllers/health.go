package controllers

import (
	"net/http"

	"github.com/RaeedAsif/bloomfilter-go/json"
)

// swagger:route GET /health Server GetServerHealth
// GetServerHealth serves http.HandlerFunc for server health
//
// security:
// - basic
//
// responses:
//
//	200: SuccessResponseHealth
func GetServerHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.Health(w)
	}
}

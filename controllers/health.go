package controllers

import (
	"net/http"

	"github.com/RaeedAsif/flare-go-test/json"
)

// GetServerHealth serves http.HandlerFunc for server health
func GetServerHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.Health(w)
	}
}

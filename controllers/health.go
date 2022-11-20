package controllers

import (
	"net/http"

	"github.com/RaeedAsif/flare-go-test/json"
)

func GetServerHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.Health(w)
	}
}

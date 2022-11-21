package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RaeedAsif/flare-go-test/json"
	"github.com/RaeedAsif/flare-go-test/storage"
)

// IsUsernameExists serves http.HandlerFunc for username avaibality
func IsUsernameExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		isExists := storage.IsUsernameExists(username)

		json.Success(w, isExists)
	}
}

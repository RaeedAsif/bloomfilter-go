package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RaeedAsif/flare-go-test/json"
	"github.com/RaeedAsif/flare-go-test/storage"
)

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		user, err := storage.GetUser(username)
		if err != nil {
			json.Error(w, err)
			return
		}

		json.Success(w, user)
	}
}

package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RaeedAsif/bloomfilter-go/json"
	"github.com/RaeedAsif/bloomfilter-go/storage"
)

// swagger:route GET /user/username/{username} User GetUsernameExists
// IsUsernameExists serves http.HandlerFunc for username avaibality
//
// parameters:
//   + name: username
//     in: path
//     description: username of the user
//     type: string
//     required: true
//
// security:
// - basic
//
// responses:
//
//	200: SuccessResponse

func IsUsernameExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		isExists := storage.IsUsernameExists(username)

		json.Success(w, isExists)
	}
}

package json

import (
	"encoding/json"
	"net/http"
)

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Success(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	response := ResponseSuccess{Status: http.StatusOK, Message: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	response := ResponseError{Status: http.StatusInternalServerError, Message: "error", Error: err.Error()}
	json.NewEncoder(w).Encode(response)
}

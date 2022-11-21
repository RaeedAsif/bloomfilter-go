package json

import (
	"encoding/json"
	"net/http"
)

// ResponseSuccess for success message response
type ResponseSuccess struct {
	Status  int         `json:"status"`  // status code
	Message string      `json:"message"` // respone message
	Data    interface{} `json:"data"`    // response data
}

// ResponseError for error message response
type ResponseError struct {
	Status  int    `json:"status"`  // status code
	Message string `json:"message"` // respone message
	Error   string `json:"error"`   // response error message
}

// Success writes success header and send json response
func Success(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	response := ResponseSuccess{Status: http.StatusOK, Message: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

// Error writes error header and send json response
func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	response := ResponseError{Status: http.StatusInternalServerError, Message: "error", Error: err.Error()}
	json.NewEncoder(w).Encode(response)
}

// Health serves server health
func Health(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]string{"server_health": "good"})
}

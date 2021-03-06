package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int
	Message string
}

func NewErrorNoMessage(status int) ErrorResponse {
	return ErrorResponse{
		Status: status,
	}
}

func NewInternalServerError(err error) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func handleError(w http.ResponseWriter, err ErrorResponse) {
	var body struct {
		Message string `json:"message"`
	}

	body.Message = err.Message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(body)
}

func handleJSONResponse(w http.ResponseWriter, v interface{}) {
	message, err := json.Marshal(v)
	if err != nil {
		handleError(w, NewErrorNoMessage(500))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(message)
}

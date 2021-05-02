package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrIntervalServer = "interval server error"
	ErrNotFound       = "not found"
	ErrDecode         = "invalid params"
)

func Decode(req *http.Request, payload interface{}) error {
	return json.NewDecoder(req.Body).Decode(payload)
}

func SendError(w http.ResponseWriter, errorMessage string, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return encode(w, ErrorResponse{Message: errorMessage})
}

func Send(w http.ResponseWriter, payload interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return encode(w, payload)
}

func encode(w http.ResponseWriter, payload interface{}) error {
	return json.NewEncoder(w).Encode(payload)
}

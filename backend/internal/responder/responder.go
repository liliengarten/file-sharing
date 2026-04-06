package responder

import (
	"encoding/json"
	"liliengarten/filesharing/internal/models"
	"net/http"
)

func Response(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := models.Response{
		Message: message,
	}

	json.NewEncoder(w).Encode(resp)
}

func DataResponse[T any](w http.ResponseWriter, message string, data []T, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := models.DataResponse[T]{
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}

func ErrorResponse(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := models.ErrorResponse{
		Message: "Error",
		Error:   err,
	}

	json.NewEncoder(w).Encode(resp)
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Generic response handler
func respondWithJSON(w http.ResponseWriter, status int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Generic error handler
func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, ApiResponse{Error: message})
}

// Generic request decoder
func decodeRequest[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// Generic ID parser from URL
func parseIDParam(r *http.Request, param string) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars[param])
}

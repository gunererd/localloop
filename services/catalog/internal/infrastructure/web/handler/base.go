package handler

import (
	"encoding/json"
	catalog "localloop/services/catalog/internal/domain"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func respondWithJSON(w http.ResponseWriter, status int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, ApiResponse{Error: message})
}

func decodeRequest[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func parseIDParam(r *http.Request, param string) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars[param])
}

type HandlerFunc[Req any] func(Req, *http.Request) (any, error)

func HandleRequest[Req any](handler HandlerFunc[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req

		// Skip body decoding for GET/DELETE requests
		if r.Method != http.MethodGet && r.Method != http.MethodDelete {
			decodedReq, err := decodeRequest[Req](r)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
			req = decodedReq
		}

		data, err := handler(req, r)
		if err != nil {
			status := http.StatusInternalServerError
			switch err {
			case catalog.ErrCategoryNotFound, catalog.ErrFieldNotFound:
				status = http.StatusNotFound
			case catalog.ErrInvalidInput:
				status = http.StatusBadRequest
			}
			respondWithError(w, status, err.Error())
			return
		}

		message := ""
		switch r.Method {
		case http.MethodPost:
			message = "created successfully"
		case http.MethodPut:
			message = "updated successfully"
		case http.MethodDelete:
			message = "deleted successfully"
		}

		respondWithJSON(w, getSuccessStatus(r.Method), ApiResponse{
			Message: message,
			Data:    data,
		})
	}
}

func getSuccessStatus(method string) int {
	switch method {
	case http.MethodPost:
		return http.StatusCreated
	default:
		return http.StatusOK
	}
}

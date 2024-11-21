package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, status int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, ApiResponse{Error: message})
}

func DecodeRequest[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func ParseIDParam(r *http.Request, param string) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars[param])
}

type HandlerFunc[Req any] func(Req, *http.Request) (any, error)

func HandleRequest[Req any](handler HandlerFunc[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req

		if r.Method != http.MethodGet && r.Method != http.MethodDelete {
			decodedReq, err := DecodeRequest[Req](r)
			if err != nil {
				RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
			req = decodedReq
		}

		data, err := handler(req, r)
		if err != nil {
			status := http.StatusInternalServerError
			RespondWithError(w, status, err.Error())
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

		RespondWithJSON(w, getSuccessStatus(r.Method), ApiResponse{
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

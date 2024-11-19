package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CatalogHandler struct{}

func NewCatalogHandler() *CatalogHandler {
	return &CatalogHandler{}
}

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (h *CatalogHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	// Placeholder response
	categories := []Category{
		{ID: "1", Name: "Electronics", Description: "Electronic devices and accessories"},
		{ID: "2", Name: "Books", Description: "Physical and digital books"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{Data: categories})
}

func (h *CatalogHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid request payload"})
		return
	}

	// Placeholder response
	category.ID = "3" // In real implementation, this would be generated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "Category created successfully",
		Data:    category,
	})
}

func (h *CatalogHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Placeholder response
	category := Category{
		ID:          id,
		Name:        "Electronics",
		Description: "Electronic devices and accessories",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{Data: category})
}

func (h *CatalogHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var category Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid request payload"})
		return
	}

	category.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "Category updated successfully",
		Data:    category,
	})
}

func (h *CatalogHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "Category deleted successfully",
		Data: map[string]string{
			"id": id,
		},
	})
}

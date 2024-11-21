package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"localloop/libs/pkg/errorbuilder"
	apperror "localloop/services/web/internal/shared/error"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId"`
	ParentName  string     `json:"parentName,omitempty"`
}

type createCategoryRequest struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
}

type updateCategoryRequest struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
}

type CatalogRepository interface {
	ListCategories() ([]Category, error)
	GetCategory(id uuid.UUID) (*Category, error)
	CreateCategory(name, description string, parentID *uuid.UUID) error
	UpdateCategory(id uuid.UUID, name, description string, parentID *uuid.UUID) error
	DeleteCategory(id uuid.UUID) error
}

type catalogRepository struct {
	baseURL string
	client  *http.Client
}

func NewCatalogRepository(baseURL string) *catalogRepository {
	return &catalogRepository{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type apiResponse struct {
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

func (r *catalogRepository) ListCategories() ([]Category, error) {
	resp, err := r.client.Get(r.baseURL + "/categories")
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Error != "" {
		return nil, fmt.Errorf(apiResp.Error)
	}

	var categories []Category
	if err := json.Unmarshal(apiResp.Data, &categories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal categories: %w", err)
	}

	return categories, nil
}

func (r *catalogRepository) GetCategory(id uuid.UUID) (*Category, error) {
	resp, err := r.client.Get(fmt.Sprintf("%s/categories/%s", r.baseURL, id))
	if err != nil {
		return nil, apperror.ErrCatalogService(
			apperror.WithService("catalog"),
			errorbuilder.WithOriginal(err),
		)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, apperror.ErrInvalidJSON(
			errorbuilder.WithOriginal(err),
		)
	}

	if apiResp.Error != "" {
		if resp.StatusCode == http.StatusNotFound {
			return nil, apperror.ErrCategoryNotFound(
				apperror.WithResource("category", id.String()),
			)
		}
		return nil, apperror.ErrCatalogService(
			apperror.WithService("catalog"),
			errorbuilder.WithContext(map[string]any{
				"error": apiResp.Error,
			}),
		)
	}

	var category Category
	if err := json.Unmarshal(apiResp.Data, &category); err != nil {
		return nil, apperror.ErrInvalidJSON(
			errorbuilder.WithOriginal(err),
		)
	}

	return &category, nil
}

func (r *catalogRepository) CreateCategory(name, description string, parentID *uuid.UUID) error {
	reqBody := createCategoryRequest{
		Name:        name,
		Description: description,
		ParentID:    parentID,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := r.client.Post(
		r.baseURL+"/categories",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Error != "" {
		return fmt.Errorf(apiResp.Error)
	}

	return nil
}

func (r *catalogRepository) UpdateCategory(id uuid.UUID, name, description string, parentID *uuid.UUID) error {
	reqBody := updateCategoryRequest{
		Name:        name,
		Description: description,
		ParentID:    parentID,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/categories/%s", r.baseURL, id),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Error != "" {
		return fmt.Errorf(apiResp.Error)
	}

	return nil
}

func (r *catalogRepository) DeleteCategory(id uuid.UUID) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/categories/%s", r.baseURL, id),
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Error != "" {
		return fmt.Errorf(apiResp.Error)
	}

	return nil
}

package handler

import (
	"encoding/json"
	catalog "localloop/services/catalog/internal/domain"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CatalogHandler struct {
	catalogService *catalog.Service
}

func NewCatalogHandler(catalogService *catalog.Service) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
	}
}

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type CreateCategoryRequest struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
}

type UpdateCategoryRequest struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
}

type CategoryResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// Field request/response types
type CreateFieldRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldTypeID uuid.UUID `json:"fieldTypeId"`
}

type UpdateFieldRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldTypeID uuid.UUID `json:"fieldTypeId"`
}

type FieldResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldTypeID uuid.UUID `json:"fieldTypeId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Field Type request/response types
type CreateFieldTypeRequest struct {
	Name                string                 `json:"name"`
	TypeDiscriminatorID uuid.UUID              `json:"typeDiscriminatorId"`
	Properties          map[string]interface{} `json:"properties"`
}

type FieldTypeResponse struct {
	ID                  uuid.UUID              `json:"id"`
	Name                string                 `json:"name"`
	TypeDiscriminatorID uuid.UUID              `json:"typeDiscriminatorId"`
	Properties          map[string]interface{} `json:"properties"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
}

// Field Type Discriminator request/response types
type CreateFieldTypeDiscriminatorRequest struct {
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ValidationSchema map[string]interface{} `json:"validationSchema"`
}

type FieldTypeDiscriminatorResponse struct {
	ID               uuid.UUID              `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ValidationSchema map[string]interface{} `json:"validationSchema"`
	CreatedAt        time.Time              `json:"createdAt"`
}

type AssignFieldRequest struct {
	IsRequired   bool  `json:"isRequired"`
	DisplayOrder int32 `json:"displayOrder"`
}

func (h *CatalogHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.catalogService.ListCategories(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: categories})
}

func (h *CatalogHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[CreateCategoryRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.CreateCategoryParams{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
	}

	category, err := h.catalogService.CreateCategory(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrInvalidInput {
			status = http.StatusBadRequest
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ApiResponse{
		Message: "Category created successfully",
		Data:    toCategoryResponse(category),
	})
}

func (h *CatalogHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.catalogService.GetCategory(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrCategoryNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: toCategoryResponse(category)})
}

func (h *CatalogHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	req, err := decodeRequest[UpdateCategoryRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.UpdateCategoryParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
	}

	category, err := h.catalogService.UpdateCategory(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrCategoryNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Category updated successfully",
		Data:    toCategoryResponse(category),
	})
}

func (h *CatalogHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = h.catalogService.DeleteCategory(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrCategoryNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Category deleted successfully",
		Data: map[string]string{
			"id": id.String(),
		},
	})
}

// Helper function to convert domain model to response
func toCategoryResponse(c *catalog.Category) CategoryResponse {
	return CategoryResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		ParentID:    c.ParentID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (h *CatalogHandler) ListFields(w http.ResponseWriter, r *http.Request) {
	fields, err := h.catalogService.ListFields(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch fields")
		return
	}

	var fieldResponses []FieldResponse
	for _, field := range fields {
		fieldResponses = append(fieldResponses, toFieldResponse(field))
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: fieldResponses})
}

func (h *CatalogHandler) CreateField(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[CreateFieldRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.CreateFieldParams{
		Name:        req.Name,
		Description: req.Description,
		FieldTypeID: req.FieldTypeID,
	}

	field, err := h.catalogService.CreateField(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrInvalidInput {
			status = http.StatusBadRequest
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ApiResponse{
		Message: "Field created successfully",
		Data:    toFieldResponse(field),
	})
}

func (h *CatalogHandler) GetField(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field ID")
		return
	}

	field, err := h.catalogService.GetField(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: toFieldResponse(field)})
}

func (h *CatalogHandler) UpdateField(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field ID")
		return
	}

	req, err := decodeRequest[UpdateFieldRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.UpdateFieldParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		FieldTypeID: req.FieldTypeID,
	}

	field, err := h.catalogService.UpdateField(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Field updated successfully",
		Data:    toFieldResponse(field),
	})
}

func (h *CatalogHandler) DeleteField(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field ID")
		return
	}

	err = h.catalogService.DeleteField(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Field deleted successfully",
		Data: map[string]string{
			"id": id.String(),
		},
	})
}

func toFieldResponse(f *catalog.Field) FieldResponse {
	return FieldResponse{
		ID:          f.ID,
		Name:        f.Name,
		Description: f.Description,
		FieldTypeID: f.FieldTypeID,
		CreatedAt:   f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
	}
}

func (h *CatalogHandler) ListFieldTypes(w http.ResponseWriter, r *http.Request) {
	fieldTypes, err := h.catalogService.ListFieldTypes(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch field types")
		return
	}

	var fieldTypeResponses []FieldTypeResponse
	for _, fieldType := range fieldTypes {
		fieldTypeResponses = append(fieldTypeResponses, toFieldTypeResponse(fieldType))
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: fieldTypeResponses})
}

func (h *CatalogHandler) CreateFieldType(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[CreateFieldTypeRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.CreateFieldTypeParams{
		Name:                req.Name,
		TypeDiscriminatorID: req.TypeDiscriminatorID,
		Properties:          req.Properties,
	}

	fieldType, err := h.catalogService.CreateFieldType(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrInvalidInput {
			status = http.StatusBadRequest
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ApiResponse{
		Message: "Field type created successfully",
		Data:    toFieldTypeResponse(fieldType),
	})
}

func (h *CatalogHandler) GetFieldType(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field type ID")
		return
	}

	fieldType, err := h.catalogService.GetFieldType(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: toFieldTypeResponse(fieldType)})
}

func toFieldTypeResponse(ft *catalog.FieldType) FieldTypeResponse {
	return FieldTypeResponse{
		ID:                  ft.ID,
		Name:                ft.Name,
		TypeDiscriminatorID: ft.TypeDiscriminatorID,
		Properties:          ft.Properties,
		CreatedAt:           ft.CreatedAt,
		UpdatedAt:           ft.UpdatedAt,
	}
}

// Field Type Discriminator handlers
func (h *CatalogHandler) ListFieldTypeDiscriminators(w http.ResponseWriter, r *http.Request) {
	discriminators, err := h.catalogService.ListFieldTypeDiscriminators(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch field type discriminators")
		return
	}

	var discriminatorResponses []FieldTypeDiscriminatorResponse
	for _, disc := range discriminators {
		discriminatorResponses = append(discriminatorResponses, toFieldTypeDiscriminatorResponse(disc))
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: discriminatorResponses})
}

func (h *CatalogHandler) CreateFieldTypeDiscriminator(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[CreateFieldTypeDiscriminatorRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.CreateFieldTypeDiscriminatorParams{
		Name:             req.Name,
		Description:      req.Description,
		ValidationSchema: req.ValidationSchema,
	}

	discriminator, err := h.catalogService.CreateFieldTypeDiscriminator(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrInvalidInput {
			status = http.StatusBadRequest
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ApiResponse{
		Message: "Field type discriminator created successfully",
		Data:    toFieldTypeDiscriminatorResponse(discriminator),
	})
}

func (h *CatalogHandler) GetFieldTypeDiscriminator(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field type discriminator ID")
		return
	}

	discriminator, err := h.catalogService.GetFieldTypeDiscriminator(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: toFieldTypeDiscriminatorResponse(discriminator)})
}

func toFieldTypeDiscriminatorResponse(d *catalog.FieldTypeDiscriminator) FieldTypeDiscriminatorResponse {
	return FieldTypeDiscriminatorResponse{
		ID:               d.ID,
		Name:             d.Name,
		Description:      d.Description,
		ValidationSchema: d.ValidationSchema,
		CreatedAt:        d.CreatedAt,
	}
}

func (h *CatalogHandler) GetCategoryFields(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := uuid.Parse(vars["categoryId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid category ID"})
		return
	}

	fields, err := h.catalogService.GetCategoryFields(r.Context(), categoryID)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrCategoryNotFound {
			status = http.StatusNotFound
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(ApiResponse{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{Data: fields})
}

func (h *CatalogHandler) AssignFieldToCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := uuid.Parse(vars["categoryId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid category ID"})
		return
	}

	fieldID, err := uuid.Parse(vars["fieldId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid field ID"})
		return
	}

	var req AssignFieldRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{Error: "Invalid request payload"})
		return
	}

	params := catalog.AssignFieldParams{
		CategoryID:   categoryID,
		FieldID:      fieldID,
		IsRequired:   req.IsRequired,
		DisplayOrder: req.DisplayOrder,
	}

	if err := h.catalogService.AssignFieldToCategory(r.Context(), params); err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrCategoryNotFound || err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		} else if err == catalog.ErrInvalidInput {
			status = http.StatusBadRequest
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(ApiResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "Field assigned to category successfully",
		Data: map[string]interface{}{
			"categoryId": categoryID,
			"fieldId":    fieldID,
		},
	})
}

func (h *CatalogHandler) UpdateFieldType(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field type ID")
		return
	}

	req, err := decodeRequest[CreateFieldTypeRequest](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := catalog.UpdateFieldTypeParams{
		ID:                  id,
		Name:                req.Name,
		TypeDiscriminatorID: req.TypeDiscriminatorID,
		Properties:          req.Properties,
	}

	fieldType, err := h.catalogService.UpdateFieldType(r.Context(), params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Field type updated successfully",
		Data:    toFieldTypeResponse(fieldType),
	})
}

func (h *CatalogHandler) DeleteFieldType(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid field type ID")
		return
	}

	err = h.catalogService.DeleteFieldType(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == catalog.ErrFieldNotFound {
			status = http.StatusNotFound
		}
		respondWithError(w, status, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Message: "Field type deleted successfully",
		Data: map[string]string{
			"id": id.String(),
		},
	})
}

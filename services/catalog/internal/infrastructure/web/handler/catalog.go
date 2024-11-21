package handler

import (
	"fmt"
	catalog "localloop/services/catalog/internal/domain"
	"net/http"
	"time"

	"github.com/google/uuid"
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

type UpdateFieldTypeRequest struct {
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

type UpdateFieldTypeDiscriminatorRequest struct {
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ValidationSchema map[string]interface{} `json:"validationSchema"`
}

// Request/Response types
type GetCategoryFieldsRequest struct {
	CategoryID uuid.UUID `param:"categoryId"`
}

type CategoryFieldResponse struct {
	Field        FieldResponse `json:"field"`
	IsRequired   bool          `json:"isRequired"`
	DisplayOrder int32         `json:"displayOrder"`
}

type AssignFieldToCategoryRequest struct {
	CategoryID   uuid.UUID `param:"categoryId"`
	FieldID      uuid.UUID `param:"fieldId"`
	IsRequired   bool      `json:"isRequired"`
	DisplayOrder int32     `json:"displayOrder"`
}

func (h *CatalogHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.catalogService.ListCategories(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{Data: categories})
}

func (h *CatalogHandler) CreateCategory(req CreateCategoryRequest, r *http.Request) (any, error) {
	params := catalog.CreateCategoryParams{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
	}

	category, err := h.catalogService.CreateCategory(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toCategoryResponse(category), nil
}

func (h *CatalogHandler) GetCategory(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	category, err := h.catalogService.GetCategory(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return toCategoryResponse(category), nil
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

func (h *CatalogHandler) CreateField(req CreateFieldRequest, r *http.Request) (any, error) {
	params := catalog.CreateFieldParams{
		Name:        req.Name,
		Description: req.Description,
		FieldTypeID: req.FieldTypeID,
	}

	field, err := h.catalogService.CreateField(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldResponse(field), nil
}

func (h *CatalogHandler) GetField(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field ID: %w", err)
	}

	field, err := h.catalogService.GetField(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return toFieldResponse(field), nil
}

func (h *CatalogHandler) UpdateField(req UpdateFieldRequest, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field ID: %w", err)
	}

	params := catalog.UpdateFieldParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		FieldTypeID: req.FieldTypeID,
	}

	field, err := h.catalogService.UpdateField(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldResponse(field), nil
}

func (h *CatalogHandler) DeleteField(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field ID: %w", err)
	}

	if err := h.catalogService.DeleteField(r.Context(), id); err != nil {
		return nil, err
	}

	return map[string]string{"id": id.String()}, nil
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

func (h *CatalogHandler) ListFieldTypes(_ struct{}, r *http.Request) (any, error) {
	fieldTypes, err := h.catalogService.ListFieldTypes(r.Context())
	if err != nil {
		return nil, err
	}

	var responses []FieldTypeResponse
	for _, ft := range fieldTypes {
		responses = append(responses, toFieldTypeResponse(ft))
	}
	return responses, nil
}

func (h *CatalogHandler) CreateFieldType(req CreateFieldTypeRequest, r *http.Request) (any, error) {
	params := catalog.CreateFieldTypeParams{
		Name:                req.Name,
		TypeDiscriminatorID: req.TypeDiscriminatorID,
		Properties:          req.Properties,
	}

	fieldType, err := h.catalogService.CreateFieldType(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldTypeResponse(fieldType), nil
}

func (h *CatalogHandler) GetFieldType(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type ID: %w", err)
	}

	fieldType, err := h.catalogService.GetFieldType(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return toFieldTypeResponse(fieldType), nil
}

func (h *CatalogHandler) UpdateFieldType(req UpdateFieldTypeRequest, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type ID: %w", err)
	}

	params := catalog.UpdateFieldTypeParams{
		ID:                  id,
		Name:                req.Name,
		TypeDiscriminatorID: req.TypeDiscriminatorID,
		Properties:          req.Properties,
	}

	fieldType, err := h.catalogService.UpdateFieldType(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldTypeResponse(fieldType), nil
}

func (h *CatalogHandler) DeleteFieldType(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type ID: %w", err)
	}

	if err := h.catalogService.DeleteFieldType(r.Context(), id); err != nil {
		return nil, err
	}

	return map[string]string{"id": id.String()}, nil
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
func (h *CatalogHandler) ListFieldTypeDiscriminators(_ struct{}, r *http.Request) (any, error) {
	discriminators, err := h.catalogService.ListFieldTypeDiscriminators(r.Context())
	if err != nil {
		return nil, err
	}

	var responses []FieldTypeDiscriminatorResponse
	for _, disc := range discriminators {
		responses = append(responses, toFieldTypeDiscriminatorResponse(disc))
	}
	return responses, nil
}

func (h *CatalogHandler) CreateFieldTypeDiscriminator(req CreateFieldTypeDiscriminatorRequest, r *http.Request) (any, error) {
	params := catalog.CreateFieldTypeDiscriminatorParams{
		Name:             req.Name,
		Description:      req.Description,
		ValidationSchema: req.ValidationSchema,
	}

	discriminator, err := h.catalogService.CreateFieldTypeDiscriminator(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldTypeDiscriminatorResponse(discriminator), nil
}

func (h *CatalogHandler) GetFieldTypeDiscriminator(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type discriminator ID: %w", err)
	}

	discriminator, err := h.catalogService.GetFieldTypeDiscriminator(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return toFieldTypeDiscriminatorResponse(discriminator), nil
}

func (h *CatalogHandler) UpdateFieldTypeDiscriminator(req UpdateFieldTypeDiscriminatorRequest, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type discriminator ID: %w", err)
	}

	params := catalog.UpdateFieldTypeDiscriminatorParams{
		ID:               id,
		Name:             req.Name,
		Description:      req.Description,
		ValidationSchema: req.ValidationSchema,
	}

	discriminator, err := h.catalogService.UpdateFieldTypeDiscriminator(r.Context(), params)
	if err != nil {
		return nil, err
	}

	return toFieldTypeDiscriminatorResponse(discriminator), nil
}

func (h *CatalogHandler) DeleteFieldTypeDiscriminator(_ struct{}, r *http.Request) (any, error) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid field type discriminator ID: %w", err)
	}

	if err := h.catalogService.DeleteFieldTypeDiscriminator(r.Context(), id); err != nil {
		return nil, err
	}

	return map[string]string{"id": id.String()}, nil
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

func (h *CatalogHandler) GetCategoryFields(_ GetCategoryFieldsRequest, r *http.Request) (any, error) {
	categoryID, err := parseIDParam(r, "categoryId")
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	fields, err := h.catalogService.GetCategoryFields(r.Context(), categoryID)
	if err != nil {
		return nil, err
	}

	var responses []CategoryFieldResponse
	for _, fieldInfo := range fields {
		responses = append(responses, CategoryFieldResponse{
			Field:        toFieldResponse(fieldInfo.Field),
			IsRequired:   fieldInfo.IsRequired,
			DisplayOrder: fieldInfo.DisplayOrder,
		})
	}
	return responses, nil
}

func (h *CatalogHandler) AssignFieldToCategory(req AssignFieldToCategoryRequest, r *http.Request) (any, error) {
	categoryID, err := parseIDParam(r, "categoryId")
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	fieldID, err := parseIDParam(r, "fieldId")
	if err != nil {
		return nil, fmt.Errorf("invalid field ID: %w", err)
	}

	params := catalog.AssignFieldParams{
		CategoryID:   categoryID,
		FieldID:      fieldID,
		IsRequired:   req.IsRequired,
		DisplayOrder: req.DisplayOrder,
	}

	if err := h.catalogService.AssignFieldToCategory(r.Context(), params); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"categoryId": categoryID,
		"fieldId":    fieldID,
	}, nil
}

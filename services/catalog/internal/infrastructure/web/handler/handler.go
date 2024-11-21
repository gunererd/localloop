package handler

import (
	"context"
	catalog "localloop/services/catalog/internal/domain"
	"net/http"

	"localloop/libs/pkg/web/crud"

	"github.com/google/uuid"
)

type CatalogHandler struct {
	catalogService *catalog.Service
	Category       *crud.CRUDHandler[catalog.Category, CreateCategoryRequest, UpdateCategoryRequest, CategoryResponse]
	Field          *crud.CRUDHandler[catalog.Field, CreateFieldRequest, UpdateFieldRequest, FieldResponse]
	FieldType      *crud.CRUDHandler[catalog.FieldType, CreateFieldTypeRequest, UpdateFieldTypeRequest, FieldTypeResponse]
	Discriminator  *crud.CRUDHandler[catalog.FieldTypeDiscriminator, CreateFieldTypeDiscriminatorRequest, UpdateFieldTypeDiscriminatorRequest, FieldTypeDiscriminatorResponse]
}

func NewCatalogHandler(catalogService *catalog.Service) *CatalogHandler {
	h := &CatalogHandler{
		catalogService: catalogService,
	}

	// Initialize category handler
	h.Category = crud.NewCRUDHandler(
		func(ctx context.Context, req CreateCategoryRequest) (*catalog.Category, error) {
			return h.catalogService.CreateCategory(ctx, catalog.CreateCategoryParams{
				Name:        req.Name,
				Description: req.Description,
				ParentID:    req.ParentID,
			})
		},
		h.catalogService.GetCategory,
		func(ctx context.Context, id uuid.UUID, req UpdateCategoryRequest) (*catalog.Category, error) {
			return h.catalogService.UpdateCategory(ctx, catalog.UpdateCategoryParams{
				ID:          id,
				Name:        req.Name,
				Description: req.Description,
				ParentID:    req.ParentID,
			})
		},
		h.catalogService.DeleteCategory,
		h.catalogService.ListCategories,
		toCategoryResponse,
	)

	h.Field = crud.NewCRUDHandler(
		func(ctx context.Context, req CreateFieldRequest) (*catalog.Field, error) {
			return h.catalogService.CreateField(ctx, catalog.CreateFieldParams{
				Name:        req.Name,
				Description: req.Description,
				FieldTypeID: req.FieldTypeID,
			})
		},
		h.catalogService.GetField,
		func(ctx context.Context, id uuid.UUID, req UpdateFieldRequest) (*catalog.Field, error) {
			return h.catalogService.UpdateField(ctx, catalog.UpdateFieldParams{
				ID:          id,
				Name:        req.Name,
				Description: req.Description,
				FieldTypeID: req.FieldTypeID,
			})
		},
		h.catalogService.DeleteField,
		h.catalogService.ListFields,
		toFieldResponse,
	)

	h.FieldType = crud.NewCRUDHandler(
		func(ctx context.Context, req CreateFieldTypeRequest) (*catalog.FieldType, error) {
			return h.catalogService.CreateFieldType(ctx, catalog.CreateFieldTypeParams{
				Name:                req.Name,
				TypeDiscriminatorID: req.TypeDiscriminatorID,
				Properties:          req.Properties,
			})
		},
		h.catalogService.GetFieldType,
		func(ctx context.Context, id uuid.UUID, req UpdateFieldTypeRequest) (*catalog.FieldType, error) {
			return h.catalogService.UpdateFieldType(ctx, catalog.UpdateFieldTypeParams{
				ID:                  id,
				Name:                req.Name,
				TypeDiscriminatorID: req.TypeDiscriminatorID,
				Properties:          req.Properties,
			})
		},
		h.catalogService.DeleteFieldType,
		h.catalogService.ListFieldTypes,
		toFieldTypeResponse,
	)

	h.Discriminator = crud.NewCRUDHandler(
		func(ctx context.Context, req CreateFieldTypeDiscriminatorRequest) (*catalog.FieldTypeDiscriminator, error) {
			return h.catalogService.CreateFieldTypeDiscriminator(ctx, catalog.CreateFieldTypeDiscriminatorParams{
				Name:             req.Name,
				Description:      req.Description,
				ValidationSchema: req.ValidationSchema,
			})
		},
		h.catalogService.GetFieldTypeDiscriminator,
		func(ctx context.Context, id uuid.UUID, req UpdateFieldTypeDiscriminatorRequest) (*catalog.FieldTypeDiscriminator, error) {
			return h.catalogService.UpdateFieldTypeDiscriminator(ctx, catalog.UpdateFieldTypeDiscriminatorParams{
				ID:               id,
				Name:             req.Name,
				Description:      req.Description,
				ValidationSchema: req.ValidationSchema,
			})
		},
		h.catalogService.DeleteFieldTypeDiscriminator,
		h.catalogService.ListFieldTypeDiscriminators,
		toFieldTypeDiscriminatorResponse,
	)

	return h
}

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (h *CatalogHandler) GetCategoryFields(req GetCategoryFieldsRequest, r *http.Request) (any, error) {
	fields, err := h.catalogService.GetCategoryFields(r.Context(), req.CategoryID)
	if err != nil {
		return nil, err
	}

	var responses []CategoryFieldResponse
	for _, fieldInfo := range fields {
		responses = append(responses, toCategoryFieldResponse(fieldInfo))
	}
	return responses, nil
}

func (h *CatalogHandler) AssignFieldToCategory(req AssignFieldToCategoryRequest, r *http.Request) (any, error) {
	params := catalog.AssignFieldParams{
		CategoryID:   req.CategoryID,
		FieldID:      req.FieldID,
		IsRequired:   req.IsRequired,
		DisplayOrder: req.DisplayOrder,
	}

	if err := h.catalogService.AssignFieldToCategory(r.Context(), params); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"categoryId": req.CategoryID,
		"fieldId":    req.FieldID,
	}, nil
}

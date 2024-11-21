// Response types for handlers
package handler

import (
	catalog "localloop/services/catalog/internal/domain"
	"time"

	"github.com/google/uuid"
)

// Category Response
type CategoryResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func toCategoryResponse(category *catalog.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

// Field Response
type FieldResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldTypeID uuid.UUID `json:"fieldTypeId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func toFieldResponse(field *catalog.Field) FieldResponse {
	return FieldResponse{
		ID:          field.ID,
		Name:        field.Name,
		Description: field.Description,
		FieldTypeID: field.FieldTypeID,
		CreatedAt:   field.CreatedAt,
		UpdatedAt:   field.UpdatedAt,
	}
}

// FieldType Response
type FieldTypeResponse struct {
	ID                  uuid.UUID              `json:"id"`
	Name                string                 `json:"name"`
	TypeDiscriminatorID uuid.UUID              `json:"typeDiscriminatorId"`
	Properties          map[string]interface{} `json:"properties"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
}

func toFieldTypeResponse(fieldType *catalog.FieldType) FieldTypeResponse {
	return FieldTypeResponse{
		ID:                  fieldType.ID,
		Name:                fieldType.Name,
		TypeDiscriminatorID: fieldType.TypeDiscriminatorID,
		Properties:          fieldType.Properties,
		CreatedAt:           fieldType.CreatedAt,
		UpdatedAt:           fieldType.UpdatedAt,
	}
}

// FieldTypeDiscriminator Response
type FieldTypeDiscriminatorResponse struct {
	ID               uuid.UUID              `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ValidationSchema map[string]interface{} `json:"validationSchema"`
	CreatedAt        time.Time              `json:"createdAt"`
}

func toFieldTypeDiscriminatorResponse(discriminator *catalog.FieldTypeDiscriminator) FieldTypeDiscriminatorResponse {
	return FieldTypeDiscriminatorResponse{
		ID:               discriminator.ID,
		Name:             discriminator.Name,
		Description:      discriminator.Description,
		ValidationSchema: discriminator.ValidationSchema,
		CreatedAt:        discriminator.CreatedAt,
	}
}

// CategoryField Response
type CategoryFieldResponse struct {
	Field        FieldResponse `json:"field"`
	IsRequired   bool          `json:"isRequired"`
	DisplayOrder int32         `json:"displayOrder"`
}

func toCategoryFieldResponse(fieldInfo *catalog.CategoryFieldInfo) CategoryFieldResponse {
	return CategoryFieldResponse{
		Field:        toFieldResponse(fieldInfo.Field),
		IsRequired:   fieldInfo.IsRequired,
		DisplayOrder: fieldInfo.DisplayOrder,
	}
}

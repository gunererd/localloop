// Request types for handlers
package handler

import "github.com/google/uuid"

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

// Field Type Discriminator request/response types
type CreateFieldTypeDiscriminatorRequest struct {
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ValidationSchema map[string]interface{} `json:"validationSchema"`
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

type AssignFieldToCategoryRequest struct {
	CategoryID   uuid.UUID `param:"categoryId"`
	FieldID      uuid.UUID `param:"fieldId"`
	IsRequired   bool      `json:"isRequired"`
	DisplayOrder int32     `json:"displayOrder"`
}

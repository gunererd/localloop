package catalog

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID
	Name        string
	Description string
	ParentID    *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Field struct {
	ID          uuid.UUID
	Name        string
	Description string
	FieldTypeID uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FieldType struct {
	ID                  uuid.UUID
	Name                string
	TypeDiscriminatorID uuid.UUID
	Properties          map[string]interface{}
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type CreateCategoryParams struct {
	Name        string `validate:"required"`
	Description string
	ParentID    *uuid.UUID
}

type UpdateCategoryParams struct {
	ID          uuid.UUID `validate:"required"`
	Name        string    `validate:"required"`
	Description string
	ParentID    *uuid.UUID
}

type AssignFieldParams struct {
	CategoryID   uuid.UUID `validate:"required"`
	FieldID      uuid.UUID `validate:"required"`
	IsRequired   bool
	DisplayOrder int32 `validate:"required"`
}

type CategoryField struct {
	CategoryID   uuid.UUID
	FieldID      uuid.UUID
	IsRequired   bool
	DisplayOrder int32
}

type FieldTypeDiscriminator struct {
	ID               uuid.UUID
	Name             string
	Description      string
	ValidationSchema map[string]interface{}
	CreatedAt        time.Time
}

type CreateFieldParams struct {
	Name        string `validate:"required"`
	Description string
	FieldTypeID uuid.UUID `validate:"required"`
}

type UpdateFieldParams struct {
	ID          uuid.UUID `validate:"required"`
	Name        string    `validate:"required"`
	Description string
	FieldTypeID uuid.UUID `validate:"required"`
}

type CreateFieldTypeParams struct {
	Name                string                 `validate:"required"`
	TypeDiscriminatorID uuid.UUID              `validate:"required"`
	Properties          map[string]interface{} `validate:"required"`
}

type CreateFieldTypeDiscriminatorParams struct {
	Name             string `validate:"required"`
	Description      string
	ValidationSchema map[string]interface{} `validate:"required"`
}

type UpdateFieldTypeParams struct {
	ID                  uuid.UUID              `validate:"required"`
	Name                string                 `validate:"required"`
	TypeDiscriminatorID uuid.UUID              `validate:"required"`
	Properties          map[string]interface{} `validate:"required"`
}

type UpdateFieldTypeDiscriminatorParams struct {
	ID               uuid.UUID `validate:"required"`
	Name             string    `validate:"required"`
	Description      string
	ValidationSchema map[string]interface{} `validate:"required"`
}

type CategoryFieldInfo struct {
	Field        *Field
	IsRequired   bool
	DisplayOrder int32
}

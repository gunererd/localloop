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

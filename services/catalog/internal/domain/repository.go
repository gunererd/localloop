package catalog

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// Category operations
	CreateCategory(ctx context.Context, category *Category) error
	GetCategory(ctx context.Context, id uuid.UUID) (*Category, error)
	UpdateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	ListCategories(ctx context.Context) ([]*Category, error)

	// Field operations
	CreateField(ctx context.Context, field *Field) error
	GetField(ctx context.Context, id uuid.UUID) (*Field, error)
	UpdateField(ctx context.Context, field *Field) error
	DeleteField(ctx context.Context, id uuid.UUID) error
	ListFields(ctx context.Context) ([]*Field, error)
	AssignFieldToCategory(ctx context.Context, params AssignFieldParams) error
	GetCategoryFields(ctx context.Context, categoryID uuid.UUID) ([]*CategoryFieldInfo, error)

	// Field Type operations
	CreateFieldType(ctx context.Context, fieldType *FieldType) error
	GetFieldType(ctx context.Context, id uuid.UUID) (*FieldType, error)
	UpdateFieldType(ctx context.Context, fieldType *FieldType) error
	DeleteFieldType(ctx context.Context, id uuid.UUID) error
	ListFieldTypes(ctx context.Context) ([]*FieldType, error)

	// Field Type Discriminator operations
	CreateFieldTypeDiscriminator(ctx context.Context, discriminator *FieldTypeDiscriminator) error
	GetFieldTypeDiscriminator(ctx context.Context, id uuid.UUID) (*FieldTypeDiscriminator, error)
	ListFieldTypeDiscriminators(ctx context.Context) ([]*FieldTypeDiscriminator, error)

	// Transaction support
	// WithTx(ctx context.Context, fn func(repo Repository) error) error
	UpdateFieldTypeDiscriminator(ctx context.Context, discriminator *FieldTypeDiscriminator) error
	DeleteFieldTypeDiscriminator(ctx context.Context, id uuid.UUID) error
}

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
}

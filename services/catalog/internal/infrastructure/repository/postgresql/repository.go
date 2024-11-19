package postgresql

import (
	"context"
	"database/sql"
	"localloop/services/catalog/internal/domain/catalog"
	"localloop/services/catalog/internal/infrastructure/repository/postgresql/sqlc"

	"github.com/google/uuid"
)

type CatalogRepository struct {
	q *sqlc.Queries
}

func NewCatalogRepository(db *sql.DB) *CatalogRepository {
	return &CatalogRepository{
		q: sqlc.New(db),
	}
}

func (r *CatalogRepository) CreateCategory(ctx context.Context, category *catalog.Category) error {
	params := sqlc.CreateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: sql.NullString{String: category.Description, Valid: category.Description != ""},
		ParentID:    uuid.NullUUID{UUID: uuid.Nil, Valid: false},
	}

	if category.ParentID != nil {
		params.ParentID = uuid.NullUUID{UUID: *category.ParentID, Valid: true}
	}

	result, err := r.q.CreateCategory(ctx, params)
	if err != nil {
		return err
	}

	category.CreatedAt = result.CreatedAt
	category.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) GetCategory(ctx context.Context, id uuid.UUID) (*catalog.Category, error) {
	result, err := r.q.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	var parentID *uuid.UUID
	if result.ParentID.Valid {
		id := result.ParentID.UUID
		parentID = &id
	}

	return &catalog.Category{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description.String,
		ParentID:    parentID,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (r *CatalogRepository) UpdateCategory(ctx context.Context, category *catalog.Category) error {
	params := sqlc.UpdateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: sql.NullString{String: category.Description, Valid: category.Description != ""},
		ParentID:    uuid.NullUUID{UUID: uuid.Nil, Valid: false},
	}

	if category.ParentID != nil {
		params.ParentID = uuid.NullUUID{UUID: *category.ParentID, Valid: true}
	}

	result, err := r.q.UpdateCategory(ctx, params)
	if err != nil {
		return err
	}

	category.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCategory(ctx, id)
}

func (r *CatalogRepository) ListCategories(ctx context.Context) ([]*catalog.Category, error) {
	results, err := r.q.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	categories := make([]*catalog.Category, len(results))
	for i, result := range results {
		var parentID *uuid.UUID
		if result.ParentID.Valid {
			id := result.ParentID.UUID
			parentID = &id
		}

		categories[i] = &catalog.Category{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description.String,
			ParentID:    parentID,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		}
	}

	return categories, nil
}

func (r *CatalogRepository) CreateField(ctx context.Context, field *catalog.Field) error {
	params := sqlc.CreateFieldParams{
		ID:          field.ID,
		Name:        field.Name,
		Description: sql.NullString{String: field.Description, Valid: field.Description != ""},
		FieldTypeID: field.FieldTypeID,
	}

	result, err := r.q.CreateField(ctx, params)
	if err != nil {
		return err
	}

	field.CreatedAt = result.CreatedAt
	field.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) GetField(ctx context.Context, id uuid.UUID) (*catalog.Field, error) {
	result, err := r.q.GetField(ctx, id)
	if err != nil {
		return nil, err
	}

	return &catalog.Field{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description.String,
		FieldTypeID: result.FieldTypeID,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (r *CatalogRepository) UpdateField(ctx context.Context, field *catalog.Field) error {
	params := sqlc.UpdateFieldParams{
		ID:          field.ID,
		Name:        field.Name,
		Description: sql.NullString{String: field.Description, Valid: field.Description != ""},
		FieldTypeID: field.FieldTypeID,
	}

	result, err := r.q.UpdateField(ctx, params)
	if err != nil {
		return err
	}

	field.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) DeleteField(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteField(ctx, id)
}

func (r *CatalogRepository) ListFields(ctx context.Context) ([]*catalog.Field, error) {
	results, err := r.q.ListFields(ctx)
	if err != nil {
		return nil, err
	}

	fields := make([]*catalog.Field, len(results))
	for i, result := range results {
		fields[i] = &catalog.Field{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description.String,
			FieldTypeID: result.FieldTypeID,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		}
	}

	return fields, nil
}

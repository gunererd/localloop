package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	catalog "localloop/services/catalog/internal/domain"
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

func unmarshalJSON[T any](data json.RawMessage) (T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return result, nil
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

func (r *CatalogRepository) AssignFieldToCategory(ctx context.Context, params catalog.AssignFieldParams) error {
	return r.q.AssignFieldToCategory(ctx, sqlc.AssignFieldToCategoryParams{
		CategoryID: params.CategoryID,
		FieldID:    params.FieldID,
	})
}

func (r *CatalogRepository) CreateFieldType(ctx context.Context, fieldType *catalog.FieldType) error {
	// Convert map to json.RawMessage
	jsonProperties, err := json.Marshal(fieldType.Properties)
	if err != nil {
		return fmt.Errorf("failed to marshal properties: %w", err)
	}

	params := sqlc.CreateFieldTypeParams{
		ID:                  fieldType.ID,
		Name:                fieldType.Name,
		TypeDiscriminatorID: fieldType.TypeDiscriminatorID,
		Properties:          jsonProperties,
	}

	result, err := r.q.CreateFieldType(ctx, params)
	if err != nil {
		return err
	}

	fieldType.CreatedAt = result.CreatedAt
	fieldType.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) GetFieldType(ctx context.Context, id uuid.UUID) (*catalog.FieldType, error) {
	result, err := r.q.GetFieldType(ctx, id)
	if err != nil {
		return nil, err
	}

	properties, err := unmarshalJSON[map[string]interface{}](result.Properties)
	if err != nil {
		return nil, err
	}

	return &catalog.FieldType{
		ID:                  result.ID,
		Name:                result.Name,
		TypeDiscriminatorID: result.TypeDiscriminatorID,
		Properties:          properties,
		CreatedAt:           result.CreatedAt,
		UpdatedAt:           result.UpdatedAt,
	}, nil
}

func (r *CatalogRepository) UpdateFieldType(ctx context.Context, fieldType *catalog.FieldType) error {
	// Convert map to json.RawMessage
	jsonProperties, err := json.Marshal(fieldType.Properties)
	if err != nil {
		return fmt.Errorf("failed to marshal properties: %w", err)
	}

	params := sqlc.UpdateFieldTypeParams{
		ID:                  fieldType.ID,
		Name:                fieldType.Name,
		TypeDiscriminatorID: fieldType.TypeDiscriminatorID,
		Properties:          jsonProperties,
	}
	result, err := r.q.UpdateFieldType(ctx, params)
	if err != nil {
		return err
	}

	fieldType.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *CatalogRepository) DeleteFieldType(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteFieldType(ctx, id)
}

func (r *CatalogRepository) ListFieldTypes(ctx context.Context) ([]*catalog.FieldType, error) {
	results, err := r.q.ListFieldTypes(ctx)
	if err != nil {
		return nil, err
	}

	fieldTypes := make([]*catalog.FieldType, len(results))
	for i, result := range results {
		properties, err := unmarshalJSON[map[string]interface{}](result.Properties)
		if err != nil {
			return nil, err
		}

		fieldTypes[i] = &catalog.FieldType{
			ID:                  result.ID,
			Name:                result.Name,
			TypeDiscriminatorID: result.TypeDiscriminatorID,
			Properties:          properties,
			CreatedAt:           result.CreatedAt,
			UpdatedAt:           result.UpdatedAt,
		}
	}

	return fieldTypes, nil
}

func (r *CatalogRepository) CreateFieldTypeDiscriminator(ctx context.Context, discriminator *catalog.FieldTypeDiscriminator) error {
	jsonSchema, err := json.Marshal(discriminator.ValidationSchema)
	if err != nil {
		return fmt.Errorf("failed to marshal validation schema: %w", err)
	}

	params := sqlc.CreateFieldTypeDiscriminatorParams{
		ID:               discriminator.ID,
		Name:             discriminator.Name,
		Description:      sql.NullString{String: discriminator.Description, Valid: discriminator.Description != ""},
		ValidationSchema: jsonSchema,
	}

	result, err := r.q.CreateFieldTypeDiscriminator(ctx, params)
	if err != nil {
		return err
	}

	discriminator.CreatedAt = result.CreatedAt
	return nil
}

func (r *CatalogRepository) GetFieldTypeDiscriminator(ctx context.Context, id uuid.UUID) (*catalog.FieldTypeDiscriminator, error) {
	result, err := r.q.GetFieldTypeDiscriminator(ctx, id)
	if err != nil {
		return nil, err
	}

	validationSchema, err := unmarshalJSON[map[string]interface{}](result.ValidationSchema)
	if err != nil {
		return nil, err
	}

	return &catalog.FieldTypeDiscriminator{
		ID:               result.ID,
		Name:             result.Name,
		Description:      result.Description.String,
		ValidationSchema: validationSchema,
		CreatedAt:        result.CreatedAt,
	}, nil
}

func (r *CatalogRepository) ListFieldTypeDiscriminators(ctx context.Context) ([]*catalog.FieldTypeDiscriminator, error) {
	results, err := r.q.ListFieldTypeDiscriminators(ctx)
	if err != nil {
		return nil, err
	}

	discriminators := make([]*catalog.FieldTypeDiscriminator, len(results))
	for i, result := range results {
		validationSchema, err := unmarshalJSON[map[string]interface{}](result.ValidationSchema)
		if err != nil {
			return nil, err
		}

		discriminators[i] = &catalog.FieldTypeDiscriminator{
			ID:               result.ID,
			Name:             result.Name,
			Description:      result.Description.String,
			ValidationSchema: validationSchema,
			CreatedAt:        result.CreatedAt,
		}
	}

	return discriminators, nil
}

func (r *CatalogRepository) GetCategoryFields(ctx context.Context, categoryID uuid.UUID) ([]*catalog.Field, error) {
	results, err := r.q.GetCategoryFields(ctx, categoryID)
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

// func (r *CatalogRepository) WithTx(ctx context.Context, fn func(repo catalog.Repository) error) error {
// 	tx, err := r.q.db.(*sql.DB).BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	txRepo := &CatalogRepository{
// 		q: r.q.WithTx(tx),
// 	}

// 	if err := fn(txRepo); err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }

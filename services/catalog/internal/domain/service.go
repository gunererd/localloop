package catalog

import (
	"context"
	"database/sql"

	"localloop/libs/pkg/errorbuilder"
	apperror "localloop/services/catalog/internal/shared/error"

	"github.com/google/uuid"
)

type ServiceConfig struct {
	// Add any service-level configuration here
}

type Service struct {
	repo Repository
	cfg  ServiceConfig
}

func NewService(repo Repository, cfg ServiceConfig) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}

// Category operations
func (s *Service) CreateCategory(ctx context.Context, params CreateCategoryParams) (*Category, error) {
	if params.Name == "" {
		return nil, apperror.ErrInvalidCategoryName(
			apperror.WithValidation("name", "name cannot be empty"),
		)
	}

	if params.ParentID != nil {
		// Verify parent exists
		_, err := s.GetCategory(ctx, *params.ParentID)
		if err != nil {
			return nil, apperror.ErrCategoryNotFound(
				apperror.WithCategory(params.ParentID.String()),
				errorbuilder.WithOriginal(err),
			)
		}
	}

	category := &Category{
		ID:          uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		ParentID:    params.ParentID,
	}

	if err := s.repo.CreateCategory(ctx, category); err != nil {
		return nil, apperror.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
		)
	}

	return category, nil
}

func (s *Service) GetCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	category, err := s.repo.GetCategory(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrCategoryNotFound(
				apperror.WithCategory(id.String()),
			)
		}
		return nil, apperror.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
		)
	}
	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, params UpdateCategoryParams) (*Category, error) {
	category := &Category{
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
		ParentID:    params.ParentID,
	}

	if err := s.repo.UpdateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCategory(ctx, id)
}

func (s *Service) ListCategories(ctx context.Context) ([]*Category, error) {
	return s.repo.ListCategories(ctx)
}

// Field operations
func (s *Service) CreateField(ctx context.Context, params CreateFieldParams) (*Field, error) {
	field := &Field{
		ID:          uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		FieldTypeID: params.FieldTypeID,
	}

	if err := s.repo.CreateField(ctx, field); err != nil {
		return nil, err
	}

	return field, nil
}

func (s *Service) GetField(ctx context.Context, id uuid.UUID) (*Field, error) {
	field, err := s.repo.GetField(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrFieldNotFound(
				apperror.WithField(id.String()),
			)
		}
		return nil, apperror.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
		)
	}
	return field, nil
}

func (s *Service) UpdateField(ctx context.Context, params UpdateFieldParams) (*Field, error) {
	field := &Field{
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
		FieldTypeID: params.FieldTypeID,
	}

	if err := s.repo.UpdateField(ctx, field); err != nil {
		return nil, err
	}

	return field, nil
}

func (s *Service) DeleteField(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteField(ctx, id)
}

func (s *Service) ListFields(ctx context.Context) ([]*Field, error) {
	return s.repo.ListFields(ctx)
}

func (s *Service) AssignFieldToCategory(ctx context.Context, params AssignFieldParams) error {
	// Verify category exists
	if _, err := s.GetCategory(ctx, params.CategoryID); err != nil {
		return apperror.ErrCategoryNotFound(
			apperror.WithCategory(params.CategoryID.String()),
			errorbuilder.WithOriginal(err),
		)
	}

	// Verify field exists
	if _, err := s.GetField(ctx, params.FieldID); err != nil {
		return apperror.ErrFieldNotFound(
			apperror.WithField(params.FieldID.String()),
			errorbuilder.WithOriginal(err),
		)
	}

	if err := s.repo.AssignFieldToCategory(ctx, params); err != nil {
		return apperror.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
		)
	}

	return nil
}

func (s *Service) GetCategoryFields(ctx context.Context, categoryID uuid.UUID) ([]*CategoryFieldInfo, error) {
	return s.repo.GetCategoryFields(ctx, categoryID)
}

// Field Type operations
func (s *Service) CreateFieldType(ctx context.Context, params CreateFieldTypeParams) (*FieldType, error) {
	fieldType := &FieldType{
		ID:                  uuid.New(),
		Name:                params.Name,
		TypeDiscriminatorID: params.TypeDiscriminatorID,
		Properties:          params.Properties,
	}

	if err := s.repo.CreateFieldType(ctx, fieldType); err != nil {
		return nil, err
	}

	return fieldType, nil
}

func (s *Service) GetFieldType(ctx context.Context, id uuid.UUID) (*FieldType, error) {
	return s.repo.GetFieldType(ctx, id)
}

func (s *Service) ListFieldTypes(ctx context.Context) ([]*FieldType, error) {
	return s.repo.ListFieldTypes(ctx)
}

func (s *Service) UpdateFieldType(ctx context.Context, params UpdateFieldTypeParams) (*FieldType, error) {
	fieldType := &FieldType{
		ID:                  params.ID,
		Name:                params.Name,
		TypeDiscriminatorID: params.TypeDiscriminatorID,
		Properties:          params.Properties,
	}

	if err := s.repo.UpdateFieldType(ctx, fieldType); err != nil {
		return nil, err
	}

	return fieldType, nil
}

func (s *Service) DeleteFieldType(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteFieldType(ctx, id)
}

// Field Type Discriminator operations
func (s *Service) CreateFieldTypeDiscriminator(ctx context.Context, params CreateFieldTypeDiscriminatorParams) (*FieldTypeDiscriminator, error) {
	discriminator := &FieldTypeDiscriminator{
		ID:               uuid.New(),
		Name:             params.Name,
		Description:      params.Description,
		ValidationSchema: params.ValidationSchema,
	}

	if err := s.repo.CreateFieldTypeDiscriminator(ctx, discriminator); err != nil {
		return nil, err
	}

	return discriminator, nil
}

func (s *Service) GetFieldTypeDiscriminator(ctx context.Context, id uuid.UUID) (*FieldTypeDiscriminator, error) {
	return s.repo.GetFieldTypeDiscriminator(ctx, id)
}

func (s *Service) ListFieldTypeDiscriminators(ctx context.Context) ([]*FieldTypeDiscriminator, error) {
	return s.repo.ListFieldTypeDiscriminators(ctx)
}

func (s *Service) UpdateFieldTypeDiscriminator(ctx context.Context, params UpdateFieldTypeDiscriminatorParams) (*FieldTypeDiscriminator, error) {
	// Check if discriminator exists
	existing, err := s.repo.GetFieldTypeDiscriminator(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	discriminator := &FieldTypeDiscriminator{
		ID:               params.ID,
		Name:             params.Name,
		Description:      params.Description,
		ValidationSchema: params.ValidationSchema,
		CreatedAt:        existing.CreatedAt,
	}

	if err := s.repo.UpdateFieldTypeDiscriminator(ctx, discriminator); err != nil {
		return nil, err
	}

	return discriminator, nil
}

func (s *Service) DeleteFieldTypeDiscriminator(ctx context.Context, id uuid.UUID) error {
	// Check if discriminator exists
	if _, err := s.repo.GetFieldTypeDiscriminator(ctx, id); err != nil {
		return err
	}

	return s.repo.DeleteFieldTypeDiscriminator(ctx, id)
}

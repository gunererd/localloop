package catalog

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrFieldNotFound    = errors.New("field not found")
	ErrInvalidInput     = errors.New("invalid input")
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
	category := &Category{
		ID:          uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		ParentID:    params.ParentID,
	}

	if err := s.repo.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) GetCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	return s.repo.GetCategory(ctx, id)
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
	return s.repo.GetField(ctx, id)
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
	return s.repo.AssignFieldToCategory(ctx, params)
}

func (s *Service) GetCategoryFields(ctx context.Context, categoryID uuid.UUID) ([]*Field, error) {
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

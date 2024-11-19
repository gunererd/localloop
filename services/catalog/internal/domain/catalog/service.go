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

type Service struct {
	repo Repository
}

type ServiceConfig struct {
	// Add any service-specific configuration here
}

func NewService(repo Repository, cfg ServiceConfig) *Service {
	return &Service{
		repo: repo,
	}
}

// Category operations
func (s *Service) CreateCategory(ctx context.Context, name, description string, parentID *uuid.UUID) (*Category, error) {
	if name == "" {
		return nil, ErrInvalidInput
	}

	category := &Category{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		ParentID:    parentID,
	}

	if err := s.repo.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) GetCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	category, err := s.repo.GetCategory(ctx, id)
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, id uuid.UUID, name, description string) (*Category, error) {
	category, err := s.repo.GetCategory(ctx, id)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	if name != "" {
		category.Name = name
	}
	if description != "" {
		category.Description = description
	}

	if err := s.repo.UpdateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return ErrCategoryNotFound
	}
	return nil
}

func (s *Service) ListCategories(ctx context.Context) ([]*Category, error) {
	return s.repo.ListCategories(ctx)
}

// Field operations
func (s *Service) CreateField(ctx context.Context, name, description string, fieldTypeID uuid.UUID) (*Field, error) {
	if name == "" || fieldTypeID == uuid.Nil {
		return nil, ErrInvalidInput
	}

	field := &Field{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		FieldTypeID: fieldTypeID,
	}

	if err := s.repo.CreateField(ctx, field); err != nil {
		return nil, err
	}

	return field, nil
}

func (s *Service) GetField(ctx context.Context, id uuid.UUID) (*Field, error) {
	field, err := s.repo.GetField(ctx, id)
	if err != nil {
		return nil, ErrFieldNotFound
	}
	return field, nil
}

func (s *Service) ListFields(ctx context.Context) ([]*Field, error) {
	return s.repo.ListFields(ctx)
}

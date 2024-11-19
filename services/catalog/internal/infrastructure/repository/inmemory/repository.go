package inmemory

import (
	"context"
	"errors"
	"localloop/services/catalog/internal/domain/catalog"
	"sync"

	"github.com/google/uuid"
)

type CatalogRepository struct {
	categories map[uuid.UUID]*catalog.Category
	fields     map[uuid.UUID]*catalog.Field
	mu         sync.RWMutex
}

func NewCatalogRepository() *CatalogRepository {
	return &CatalogRepository{
		categories: make(map[uuid.UUID]*catalog.Category),
		fields:     make(map[uuid.UUID]*catalog.Field),
	}
}

// Category operations
func (r *CatalogRepository) CreateCategory(ctx context.Context, category *catalog.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[category.ID]; exists {
		return errors.New("category already exists")
	}

	r.categories[category.ID] = category
	return nil
}

func (r *CatalogRepository) GetCategory(ctx context.Context, id uuid.UUID) (*catalog.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, exists := r.categories[id]
	if !exists {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (r *CatalogRepository) UpdateCategory(ctx context.Context, category *catalog.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[category.ID]; !exists {
		return errors.New("category not found")
	}

	r.categories[category.ID] = category
	return nil
}

func (r *CatalogRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[id]; !exists {
		return errors.New("category not found")
	}

	delete(r.categories, id)
	return nil
}

func (r *CatalogRepository) ListCategories(ctx context.Context) ([]*catalog.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*catalog.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

// Field operations
func (r *CatalogRepository) CreateField(ctx context.Context, field *catalog.Field) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.fields[field.ID]; exists {
		return errors.New("field already exists")
	}

	r.fields[field.ID] = field
	return nil
}

func (r *CatalogRepository) GetField(ctx context.Context, id uuid.UUID) (*catalog.Field, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	field, exists := r.fields[id]
	if !exists {
		return nil, errors.New("field not found")
	}

	return field, nil
}

func (r *CatalogRepository) UpdateField(ctx context.Context, field *catalog.Field) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.fields[field.ID]; !exists {
		return errors.New("field not found")
	}

	r.fields[field.ID] = field
	return nil
}

func (r *CatalogRepository) DeleteField(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.fields[id]; !exists {
		return errors.New("field not found")
	}

	delete(r.fields, id)
	return nil
}

func (r *CatalogRepository) ListFields(ctx context.Context) ([]*catalog.Field, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	fields := make([]*catalog.Field, 0, len(r.fields))
	for _, field := range r.fields {
		fields = append(fields, field)
	}

	return fields, nil
}

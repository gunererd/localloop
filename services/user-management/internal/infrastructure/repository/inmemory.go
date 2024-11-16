package repository

import (
	"context"
	"errors"
	"localloop/services/user-management/internal/domain/user"
	"sync"
)

// InMemoryRepository implements the Repository interface using a map for storage.
type InMemoryRepository struct {
	data map[string]user.User
	mu   sync.RWMutex
}

// NewInMemoryRepository creates a new instance of InMemoryRepository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string]user.User),
	}
}

// Create adds a new user to the in-memory storage.
func (r *InMemoryRepository) Create(ctx context.Context, user user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[user.Email]; exists {
		return errors.New("user already exists")
	}

	r.data[user.Email] = user
	return nil
}

// FindByEmail retrieves a user by email from the in-memory storage.
func (r *InMemoryRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, exists := r.data[email]
	if !exists {
		return user.User{}, errors.New("user not found")
	}

	return u, nil
}

// Update modifies user data in the in-memory storage.
func (r *InMemoryRepository) Update(ctx context.Context, email string, updates user.UpdateData) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.data[email]
	if !exists {
		return errors.New("user not found")
	}

	if updates.Name != "" {
		user.Name = updates.Name
	}
	if updates.Hash != "" {
		user.Hash = updates.Hash
		user.Salt = updates.Salt
	}

	r.data[email] = user
	return nil
}
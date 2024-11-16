package user

import (
	"context"
)

// UpdateData represents the fields that can be updated for a user
type UpdateData struct {
	Hash string
	Salt string
	Name string
}

// Repository defines the methods that any user repository implementation must have
type Repository interface {
	Create(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, email string, updates UpdateData) error
}

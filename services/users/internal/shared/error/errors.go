package apperror

import (
	"localloop/libs/pkg/errorbuilder"
)

var (
	// Authentication errors
	ErrInvalidCredentials = errorbuilder.NewError("invalid credentials", errorbuilder.ErrAuth)
	ErrTokenExpired       = errorbuilder.NewError("token has expired", errorbuilder.ErrAuth)
	ErrInvalidToken       = errorbuilder.NewError("invalid token", errorbuilder.ErrAuth)

	// Resource errors
	ErrUserNotFound = errorbuilder.NewError("user not found", errorbuilder.ErrNotFound)
	ErrUserExists   = errorbuilder.NewError("user already exists", errorbuilder.ErrConflict)
	ErrEmailExists  = errorbuilder.NewError("email is already in use", errorbuilder.ErrConflict)

	// Validation errors
	ErrInvalidEmail    = errorbuilder.NewError("invalid email format", errorbuilder.ErrValidation)
	ErrInvalidPassword = errorbuilder.NewError("invalid password format", errorbuilder.ErrValidation)
	ErrInvalidName     = errorbuilder.NewError("invalid name format", errorbuilder.ErrValidation)

	// Internal errors
	ErrHashingPassword   = errorbuilder.NewError("failed to hash password", errorbuilder.ErrInternal)
	ErrGeneratingSalt    = errorbuilder.NewError("failed to generate salt", errorbuilder.ErrInternal)
	ErrDatabaseOperation = errorbuilder.NewError("database operation failed", errorbuilder.ErrInternal)
)

func WithUser(email string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"email": email,
	})
}

func WithValidation(field, reason string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"field":  field,
		"reason": reason,
	})
}

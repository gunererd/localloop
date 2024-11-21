package apperror

import (
	"localloop/libs/pkg/errorbuilder"
)

var (
	// Authentication errors
	ErrInvalidSession = errorbuilder.NewError("invalid session", errorbuilder.ErrAuth)
	ErrSessionExpired = errorbuilder.NewError("session has expired", errorbuilder.ErrAuth)
	ErrUnauthorized   = errorbuilder.NewError("unauthorized access", errorbuilder.ErrAuth)

	// Resource errors
	ErrCategoryNotFound = errorbuilder.NewError("category not found", errorbuilder.ErrNotFound)
	ErrListingNotFound  = errorbuilder.NewError("listing not found", errorbuilder.ErrNotFound)
	ErrUserNotFound     = errorbuilder.NewError("user not found", errorbuilder.ErrNotFound)

	// Service errors
	ErrCatalogService = errorbuilder.NewError("catalog service error", errorbuilder.ErrInternal)
	ErrUserService    = errorbuilder.NewError("user service error", errorbuilder.ErrInternal)
	ErrListingService = errorbuilder.NewError("listing service error", errorbuilder.ErrInternal)

	// Validation errors
	ErrInvalidRequest = errorbuilder.NewError("invalid request", errorbuilder.ErrValidation)
	ErrInvalidForm    = errorbuilder.NewError("invalid form data", errorbuilder.ErrValidation)

	// Internal errors
	ErrTemplateRender = errorbuilder.NewError("template rendering failed", errorbuilder.ErrInternal)
	ErrInvalidJSON    = errorbuilder.NewError("invalid JSON data", errorbuilder.ErrInternal)
)

func WithResource(resourceType, id string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"resourceType": resourceType,
		"resourceId":   id,
	})
}

func WithService(service string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"service": service,
	})
}

func WithValidation(field, reason string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"field":  field,
		"reason": reason,
	})
}

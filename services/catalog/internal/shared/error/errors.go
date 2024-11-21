package apperror

import (
	"localloop/libs/pkg/errorbuilder"
)

var (
	// Resource errors
	ErrCategoryNotFound      = errorbuilder.NewError("category not found", errorbuilder.ErrNotFound)
	ErrFieldNotFound         = errorbuilder.NewError("field not found", errorbuilder.ErrNotFound)
	ErrFieldTypeNotFound     = errorbuilder.NewError("field type not found", errorbuilder.ErrNotFound)
	ErrDiscriminatorNotFound = errorbuilder.NewError("field type discriminator not found", errorbuilder.ErrNotFound)

	// Validation errors
	ErrInvalidCategoryName = errorbuilder.NewError("invalid category name", errorbuilder.ErrValidation)
	ErrInvalidFieldName    = errorbuilder.NewError("invalid field name", errorbuilder.ErrValidation)
	ErrInvalidFieldType    = errorbuilder.NewError("invalid field type", errorbuilder.ErrValidation)
	ErrInvalidProperties   = errorbuilder.NewError("invalid properties", errorbuilder.ErrValidation)

	// Conflict errors
	ErrCategoryExists  = errorbuilder.NewError("category already exists", errorbuilder.ErrConflict)
	ErrFieldExists     = errorbuilder.NewError("field already exists", errorbuilder.ErrConflict)
	ErrFieldTypeExists = errorbuilder.NewError("field type already exists", errorbuilder.ErrConflict)

	// Internal errors
	ErrDatabaseOperation = errorbuilder.NewError("database operation failed", errorbuilder.ErrInternal)
	ErrInvalidJSON       = errorbuilder.NewError("invalid JSON data", errorbuilder.ErrInternal)
)

func WithCategory(id string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"categoryId": id,
	})
}

func WithField(id string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"fieldId": id,
	})
}

func WithValidation(field, reason string) errorbuilder.ErrorOption {
	return errorbuilder.WithContext(map[string]any{
		"field":  field,
		"reason": reason,
	})
}

package errorbuilder

import (
	"fmt"
)

type ErrorCode int

const (
	ErrValidation ErrorCode = 400
	ErrAuth       ErrorCode = 401
	ErrNotFound   ErrorCode = 404
	ErrConflict   ErrorCode = 409
	ErrInternal   ErrorCode = 500
)

type ErrorBuilder interface {
	Error() string
	Unwrap() error
}

type CustomError struct {
	Code    ErrorCode      `json:"code"`
	Context map[string]any `json:"context,omitempty"`
	Message string         `json:"message"`
	Orig    error          `json:"-"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *CustomError) Unwrap() error {
	return e.Orig
}

func NewError(message string, code ErrorCode) func(opts ...ErrorOption) *CustomError {
	return func(opts ...ErrorOption) *CustomError {
		err := &CustomError{
			Code:    code,
			Message: message,
			Context: make(map[string]any),
		}

		for _, opt := range opts {
			opt(err)
		}

		return err
	}
}

type ErrorOption func(*CustomError)

func WithContext(ctx map[string]any) ErrorOption {
	return func(e *CustomError) {
		e.Context = ctx
	}
}

func WithOriginal(err error) ErrorOption {
	return func(e *CustomError) {
		e.Orig = err
	}
}

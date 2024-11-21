package crud

import (
	"context"
	"fmt"
	"localloop/libs/pkg/web/handler"
	"net/http"

	"github.com/google/uuid"
)

type EmptyRequest struct{}

type CRUDHandler[D any, C any, U any, R any] struct {
	create     func(context.Context, C) (*D, error)
	get        func(context.Context, uuid.UUID) (*D, error)
	update     func(context.Context, uuid.UUID, U) (*D, error)
	delete     func(context.Context, uuid.UUID) error
	list       func(context.Context) ([]*D, error)
	toResponse func(*D) R
}

func NewCRUDHandler[D any, C any, U any, R any](
	create func(context.Context, C) (*D, error),
	get func(context.Context, uuid.UUID) (*D, error),
	update func(context.Context, uuid.UUID, U) (*D, error),
	delete func(context.Context, uuid.UUID) error,
	list func(context.Context) ([]*D, error),
	toResponse func(*D) R,
) *CRUDHandler[D, C, U, R] {
	return &CRUDHandler[D, C, U, R]{
		create:     create,
		get:        get,
		update:     update,
		delete:     delete,
		list:       list,
		toResponse: toResponse,
	}
}

func (h *CRUDHandler[D, C, U, R]) Create(req C, r *http.Request) (any, error) {
	result, err := h.create(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return h.toResponse(result), nil
}

func (h *CRUDHandler[D, C, U, R]) Get(_ EmptyRequest, r *http.Request) (any, error) {
	id, err := handler.ParseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	result, err := h.get(r.Context(), id)
	if err != nil {
		return nil, err
	}
	return h.toResponse(result), nil
}

func (h *CRUDHandler[D, C, U, R]) Update(req U, r *http.Request) (any, error) {
	id, err := handler.ParseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	result, err := h.update(r.Context(), id, req)
	if err != nil {
		return nil, err
	}
	return h.toResponse(result), nil
}

func (h *CRUDHandler[D, C, U, R]) Delete(_ EmptyRequest, r *http.Request) (any, error) {
	id, err := handler.ParseIDParam(r, "id")
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	if err := h.delete(r.Context(), id); err != nil {
		return nil, err
	}
	return map[string]string{"id": id.String()}, nil
}

func (h *CRUDHandler[D, C, U, R]) List(_ EmptyRequest, r *http.Request) (any, error) {
	results, err := h.list(r.Context())
	if err != nil {
		return nil, err
	}

	responses := make([]R, len(results))
	for i, result := range results {
		responses[i] = h.toResponse(result)
	}
	return responses, nil
}

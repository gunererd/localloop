// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: field.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createField = `-- name: CreateField :one
INSERT INTO fields (
    id, name, description, field_type_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, name, description, field_type_id, created_at, updated_at
`

type CreateFieldParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	FieldTypeID uuid.UUID      `json:"fieldTypeId"`
}

func (q *Queries) CreateField(ctx context.Context, arg CreateFieldParams) (Field, error) {
	row := q.db.QueryRowContext(ctx, createField,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.FieldTypeID,
	)
	var i Field
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.FieldTypeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteField = `-- name: DeleteField :exec
DELETE FROM fields
WHERE id = $1
`

func (q *Queries) DeleteField(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteField, id)
	return err
}

const getField = `-- name: GetField :one
SELECT id, name, description, field_type_id, created_at, updated_at FROM fields
WHERE id = $1
`

func (q *Queries) GetField(ctx context.Context, id uuid.UUID) (Field, error) {
	row := q.db.QueryRowContext(ctx, getField, id)
	var i Field
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.FieldTypeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listFields = `-- name: ListFields :many
SELECT id, name, description, field_type_id, created_at, updated_at FROM fields
`

func (q *Queries) ListFields(ctx context.Context) ([]Field, error) {
	rows, err := q.db.QueryContext(ctx, listFields)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Field{}
	for rows.Next() {
		var i Field
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.FieldTypeID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateField = `-- name: UpdateField :one
UPDATE fields
SET name = $2,
    description = $3,
    field_type_id = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, description, field_type_id, created_at, updated_at
`

type UpdateFieldParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	FieldTypeID uuid.UUID      `json:"fieldTypeId"`
}

func (q *Queries) UpdateField(ctx context.Context, arg UpdateFieldParams) (Field, error) {
	row := q.db.QueryRowContext(ctx, updateField,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.FieldTypeID,
	)
	var i Field
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.FieldTypeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
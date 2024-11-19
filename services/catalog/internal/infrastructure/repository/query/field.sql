-- name: CreateField :one
INSERT INTO fields (
    id, name, description, field_type_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetField :one
SELECT * FROM fields
WHERE id = $1;

-- name: ListFields :many
SELECT * FROM fields;

-- name: UpdateField :one
UPDATE fields
SET name = $2,
    description = $3,
    field_type_id = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteField :exec
DELETE FROM fields
WHERE id = $1; 
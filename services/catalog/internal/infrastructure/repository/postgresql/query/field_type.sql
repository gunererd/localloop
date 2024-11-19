-- name: CreateFieldType :one
INSERT INTO field_types (
    id, 
    name, 
    type_discriminator_id,
    properties
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetFieldType :one
SELECT * FROM field_types
WHERE id = $1;

-- name: ListFieldTypes :many
SELECT * FROM field_types;

-- name: UpdateFieldType :one
UPDATE field_types
SET name = $2,
    type_discriminator_id = $3,
    properties = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteFieldType :exec
DELETE FROM field_types
WHERE id = $1; 
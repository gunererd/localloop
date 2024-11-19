-- name: CreateCategory :one
INSERT INTO categories (
    id, name, description, parent_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    description = $3,
    parent_id = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1; 
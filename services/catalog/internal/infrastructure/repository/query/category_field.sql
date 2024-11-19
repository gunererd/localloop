-- name: AssignFieldToCategory :exec
INSERT INTO category_fields (
    category_id, field_id, is_required, display_order
) VALUES (
    $1, $2, $3, $4
);

-- name: GetCategoryFields :many
SELECT f.*, cf.is_required, cf.display_order
FROM fields f
JOIN category_fields cf ON f.id = cf.field_id
WHERE cf.category_id = $1
ORDER BY cf.display_order; 
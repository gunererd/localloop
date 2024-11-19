-- name: CreateFieldTypeDiscriminator :one
INSERT INTO field_type_discriminators (
    id,
    name,
    description,
    validation_schema
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetFieldTypeDiscriminator :one
SELECT * FROM field_type_discriminators
WHERE id = $1;

-- name: ListFieldTypeDiscriminators :many
SELECT * FROM field_type_discriminators; 
-- name: CreateCategory :one
INSERT INTO categories (id, name)
VALUES (?, ?)
    RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM categories ORDER BY name;

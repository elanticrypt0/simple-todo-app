-- name: GetAllTodos :many
SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT id, title, completed, created_at FROM todos WHERE id = ? LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (id, title, completed, category_id) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos SET title = ?, completed = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;

-- name: GetTodosWithCategory :many
SELECT
    t.id,
    t.title,
    t.completed,
    t.created_at,
    c.name as category_name
FROM todos t
         LEFT JOIN categories c ON t.category_id = c.id
ORDER BY t.created_at DESC;

-- name: GetTodoWithCategory :one
SELECT
    t.id,
    t.title,
    t.completed,
    c.name as category_name
FROM todos t
         LEFT JOIN categories c ON t.category_id = c.id
WHERE t.id = ?;
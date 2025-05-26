-- name: GetAllTodos :many
SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT id, title, completed, created_at FROM todos WHERE id = ? LIMIT 1;

-- name: CreateTodo :exec
INSERT INTO todos (id, title, completed, created_at) VALUES (?, ?, ?, ?);

-- name: UpdateTodo :exec
UPDATE todos SET title = ?, completed = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;

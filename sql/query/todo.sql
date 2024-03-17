-- name: GetTodos :many
SELECT * FROM todo;

-- name: GetTodo :one
SELECT * FROM todo WHERE id=$1;

-- name: InsertTodo :one
INSERT INTO todo (title, description, completed, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING *;

-- name: UpdateTodo :one
UPDATE todo  SET title=$1, description=$2, completed=$3, updated_at=CURRENT_TIMESTAMP WHERE id=$4 RETURNING updated_at;

-- name: DeleteTodo :exec
DELETE FROM todo WHERE id=$1;

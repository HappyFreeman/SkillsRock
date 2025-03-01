-- name: CreateTask :one
INSERT INTO tasks (title, description, status)
VALUES ($1, $2, $3)
RETURNING *;
--

-- name: GetTasks :many
SELECT * FROM tasks;
--

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
--

-- name: GetTaskById :one
 SELECT * FROM tasks WHERE id = $1;
--

-- name: UpdateTask :one
UPDATE tasks
SET 
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    status = COALESCE($4, status),
    updated_at = now()
WHERE id = $1
RETURNING *;
--
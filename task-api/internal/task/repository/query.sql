-- name: saveTask :one
INSERT INTO tasks (id, title, status, created_at, updated_at)
VALUES ($1, $2, $3, now(), now())
RETURNING id, title, status, created_at, updated_at;

-- name: getAll :many
SELECT id, title, status, created_at, updated_at
FROM tasks
WHERE deleted_at IS NULL
order by created_at desc, updated_at desc
OFFSET $1 LIMIT $2;

-- name: GetCount :one
SELECT count(*)
FROM tasks
WHERE deleted_at IS NULL;

-- name: getByID :one
SELECT id, title, status, created_at, updated_at
FROM tasks
WHERE id = $1;

-- name: updateTask :one
UPDATE tasks
SET title  = $1,
    status = $2
WHERE id = $3
RETURNING id, title, status, created_at, updated_at;

-- name: deleteTask :one
UPDATE tasks
SET deleted_at = now()
WHERE id = $1
RETURNING id, title, status, created_at, updated_at;
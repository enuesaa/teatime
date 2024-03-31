-- name: ListTeas :many
SELECT * FROM teas
WHERE teapod = ?;

-- name: GetTea :one
SELECT * FROM teas
WHERE teapod = ? and rid = ? LIMIT 1;

-- name: CreateTea :one
INSERT INTO teas (
  teapod, collection, value
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateTea :exec
UPDATE teas
set value = ?
WHERE teapod = ? and rid = ?;

-- name: DeleteTea :exec
DELETE FROM teas
WHERE teapod = ? and rid = ?;

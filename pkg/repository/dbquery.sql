-- name: ListTeas :many
SELECT * FROM teas
WHERE teapod = ?;

-- name: GetTea :one
SELECT * FROM teas
WHERE teapod = ? and teaid = ? LIMIT 1;

-- name: CreateTea :one
INSERT INTO teas (
  teapod, collection, teaid, value
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteTea :exec
DELETE FROM teas
WHERE teapod = ? and teaid = ?;

-- name: ListTeas :many
SELECT * FROM teas
WHERE teapod = ?;

-- name: ListTeasByTeaboxName :many
SELECT * FROM teas
WHERE teapod = ? and teabox = ?;

-- name: GetTea :one
SELECT * FROM teas
WHERE teapod = ? and teaid = ? LIMIT 1;

-- name: CreateTea :one
INSERT INTO teas (
  teapod, teabox, teaid, value
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteTea :exec
DELETE FROM teas
WHERE teapod = ? and teaid = ?;

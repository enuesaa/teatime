-- name: ListKvsOfTeapod :many
SELECT * FROM kvs
WHERE teapod = ?;

-- name: GetKv :one
SELECT * FROM kvs
WHERE id = ? LIMIT 1;

-- name: CreateKv :one
INSERT INTO kvs (
  teapod, path, value
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateKv :exec
UPDATE kvs
set value = ?
WHERE id = ? and path = ?;

-- name: DeleteKv :exec
DELETE FROM kvs
WHERE id = ?;

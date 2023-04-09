-- name: CreateString :one
INSERT INTO tsarka (
    string,
    max_substring
) VALUES (
          $1, $2
         ) RETURNING *;

-- name: GetString :one
SELECT max_substring FROM tsarka
WHERE id = $1 LIMIT 1;

-- name: ListStrings :many
SELECT * FROM tsarka
ORDER BY id LIMIT $1
OFFSET $2;

-- name: UpdateString :one
UPDATE tsarka SET max_substring = $1
WHERE id = $2
RETURNING *;

-- name: DeleteString :one
DELETE FROM tsarka WHERE id = $1
    RETURNING *;

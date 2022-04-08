-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = ? LIMIT 1;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = ?;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY owner LIMIT $1 OFFSET $2;

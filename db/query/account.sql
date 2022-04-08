-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdatebBalance :exec
UPDATE accounts SET balance = $2
WHERE id = $1;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner LIMIT $1 OFFSET $2;

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

-- name: CreateEntries :one
INSERT INTO entries (
    account_id, amount
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = ? LIMIT 1;

-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = ?;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id LIMIT $1 OFFSET $2;
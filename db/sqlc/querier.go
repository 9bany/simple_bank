// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Accounts, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entries, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntries(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Accounts, error)
	GetEntries(ctx context.Context, id int64) (Entries, error)
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetTransfer(ctx context.Context, id int64) (Transfers, error)
	GetUser(ctx context.Context, username string) (Users, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error)
	ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfers, error)
	UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (Accounts, error)
}

var _ Querier = (*Queries)(nil)

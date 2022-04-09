package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

type TransfersTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransfersResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *Store) TransfersTx(ctx context.Context, arg TransfersTxParams) (TransfersResult, error) {

	var result TransfersResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}
		
		// get send account and update it
		fromAccount, errGetFromAcc := q.GetAccount(context.Background(), result.FromEntry.AccountID)
		
		if errGetFromAcc != nil {
			return errGetFromAcc
		}

		result.FromAccount, err = q.UpdateBalance(context.Background(), UpdateBalanceParams{
			ID: result.FromEntry.AccountID,
			Balance: fromAccount.Balance + result.FromEntry.Amount,
		})

		if err != nil {
			return err
		}

		// get recive account and update it
		toAccount, errGetToAcc := q.GetAccount(context.Background(), result.ToEntry.AccountID)
		
		if errGetToAcc != nil {
			return errGetToAcc
		}

		result.ToAccount, err = q.UpdateBalance(context.Background(), UpdateBalanceParams{
			ID: result.ToEntry.AccountID,
			Balance: toAccount.Balance + result.ToEntry.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

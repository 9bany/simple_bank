package db

import (
	"9bany/simple_bank/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entries {
	account := createRandomAccount(t)

	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entries, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, account.ID, entries.AccountID)
	require.Equal(t, arg.Amount, entries.Amount)

	require.NotZero(t, entries.ID)
	require.NotZero(t, entries.CreatedAt)

	return entries
}
func TestCreateEntries(t *testing.T) {
	createRandomEntries(t)
}

func TestGetEntries(t *testing.T) {
	entries := createRandomEntries(t)
	entries2, err := testQueries.GetEntries(context.Background(), entries.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entries2)

	require.Equal(t, entries.ID, entries2.ID)
	require.Equal(t, entries.Amount, entries2.Amount)

	require.NotZero(t, entries2.ID)
	require.NotZero(t, entries2.CreatedAt)
}

func TestDeleteEntries(t *testing.T) {
	entries := createRandomEntries(t)
	err := testQueries.DeleteEntries(context.Background(), entries.ID)
	
	require.NoError(t, err)

	entries2, err := testQueries.GetEntries(context.Background(), entries.ID)
	require.Error(t, err)
	require.Empty(t, entries2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntries(t)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entrieses, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrieses, 5)

	for _, entries := range entrieses {
		require.NotEmpty(t, entries)
	}
}
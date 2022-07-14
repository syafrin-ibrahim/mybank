package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/syafrin-ibrahim/mybank/util"
)

func createRandomTransfer(t *testing.T, acc1, acc2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        util.RandomMoney(),
	}
	trs, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trs)

	require.Equal(t, arg.FromAccountID, trs.FromAccountID)
	require.Equal(t, arg.ToAccountID, trs.ToAccountID)
	require.Equal(t, arg.Amount, trs.Amount)

	require.NotZero(t, trs.ID)
	require.NotZero(t, trs.CreatedAt)
	return trs
}

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	createRandomTransfer(t, acc1, acc2)
}

func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	trs1 := createRandomTransfer(t, acc1, acc2)
	trs2, err := testQueries.GetTransfer(context.Background(), trs1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trs2)

	require.Equal(t, trs1.ID, trs2.ID)
	require.Equal(t, trs1.FromAccountID, trs2.FromAccountID)
	require.Equal(t, trs1.ToAccountID, trs2.ToAccountID)
	require.Equal(t, trs1.Amount, trs2.Amount)
	require.WithinDuration(t, trs1.CreatedAt, trs2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, acc1, acc2)
		createRandomTransfer(t, acc2, acc1)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc1.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == acc1.ID || transfer.ToAccountID == acc1.ID)
	}

}

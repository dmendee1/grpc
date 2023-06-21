package db

import (
	"context"
	"github.com/dmendee1/simpletest/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	NotNullAccount(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: account.Balance,
	}

	testQueries.UpdateAccount(context.Background(), arg)
	updatedAccount, err := testQueries.GetAccount(context.Background(), arg.ID)

	require.NoError(t, err)
	require.NotNil(t, updatedAccount)

	NotNullAccount(t, updatedAccount)

	require.Equal(t, arg.ID, updatedAccount.ID)
	require.Equal(t, arg.Balance, updatedAccount.Balance)

}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	NotNullAccount(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account.ID)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		NotNullAccount(t, account)
	}
}

func NotNullAccount(t *testing.T, account Accounts) {
	require.NotNil(t, account)

	require.NotZero(t, account.ID)
	require.NotNil(t, account.Owner)
	require.NotNil(t, account.Balance)
	require.NotNil(t, account.Currency)
	require.NotZero(t, account.CreatedAt)
}

package transactions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
)

func TestAccount_Deposit(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	_ = am.Deposit(3.34)
	assert.Equal(t, 34.34, am.balance)
}

func TestAccount_Withdraw(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	_ = am.Withdraw(3.34)
	assert.Equal(t, 27.66, am.balance)
}

func TestAccount_GetBalance(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	assert.Equal(t, 31.0, am.GetBalance())
}

func TestAccount_GetStatement(t *testing.T) {
	am := NewAccountManager().(*account)

	_ = am.Deposit(123)
	_ = am.Withdraw(11)

	expectedInBal, expectedOutBal, expectedOps := 0.0, 112.0, []models.Operation{
		{Date: time.Now().Truncate(time.Hour * 24), Amount: 123},
		{Date: time.Now().Truncate(time.Hour * 24), Amount: -11},
	}
	gotInBal, gotOutBal, gotOps := am.GetStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
	assert.Equal(t, expectedInBal, gotInBal)
	assert.Equal(t, expectedOutBal, gotOutBal)
	assert.Equal(t, expectedOps, gotOps)
}

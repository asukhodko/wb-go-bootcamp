package facade

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/notification"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions/restrictions"
)

func TestFacade_Deposit(t *testing.T) {
	amm := &transactions.AccountManagerMock{}
	amm.On("Deposit", 123.0).Return(nil)

	nm := &notification.NotifierMock{}
	nm.On("Notify", "", "Account refilled by 123.00, balance: 0.00").Return()
	nm.On("Notify", "", "Account not refilled: account is restricted").Return()

	rm := &restrictions.CheckerMock{}
	rm.On("IsRestricted").Return(false)

	f := NewAccountManager(&models.Person{}, amm, rm, nm).(*facade)

	tearDown := func() {
		amm.Calls = nil
		nm.Calls = nil
		rm.Calls = nil
	}

	t.Run("able to perform simple deposit", func(t *testing.T) {
		defer tearDown()
		f.Deposit(123)
		amm.AssertCalled(t, "Deposit", 123.0)
	})

	t.Run("have notification", func(t *testing.T) {
		defer tearDown()
		f.Deposit(123)
		nm.AssertCalled(t, "Notify", "", "Account refilled by 123.00, balance: 0.00")
	})

	t.Run("is restricted", func(t *testing.T) {
		defer tearDown()
		rm.Mock.ExpectedCalls = nil
		rm.On("IsRestricted").Return(true)
		f.Deposit(123)
		amm.AssertNotCalled(t, "Deposit", 123.0)
		nm.AssertCalled(t, "Notify", "", "Account not refilled: account is restricted")
	})
}

func TestFacade_Withdraw(t *testing.T) {
	amm := &transactions.AccountManagerMock{}
	amm.On("Withdraw", 123.0).Return(nil)
	amm.On("Withdraw", 124.0).Return(errors.New("insufficient founds"))

	nm := &notification.NotifierMock{}
	nm.On("Notify", "", "Successful withdrawal from the am by 123.00, balance: 0.00").Return()
	nm.On("Notify", "", "Successful withdrawal from the am by 124.00, balance: 0.00").Return()
	nm.On("Notify", "", "Account withdrawal failed: insufficient founds").Return()
	nm.On("Notify", "", "Account withdrawal failed: account is restricted").Return()

	rm := &restrictions.CheckerMock{}
	rm.On("IsRestricted").Return(false)

	f := NewAccountManager(&models.Person{}, amm, rm, nm).(*facade)

	tearDown := func() {
		amm.Calls = nil
		nm.Calls = nil
		rm.Calls = nil
	}

	t.Run("able to perform simple withdraw", func(t *testing.T) {
		defer tearDown()
		f.Withdraw(123)
		amm.AssertCalled(t, "Withdraw", 123.0)
	})

	t.Run("have notification", func(t *testing.T) {
		defer tearDown()
		f.Withdraw(123)
		nm.AssertCalled(t, "Notify", "", "Successful withdrawal from the am by 123.00, balance: 0.00")
	})

	t.Run("check negative balance", func(t *testing.T) {
		defer tearDown()
		f.Withdraw(124)
		nm.AssertCalled(t, "Notify", "", "Account withdrawal failed: insufficient founds")
	})

	t.Run("is restricted", func(t *testing.T) {
		defer tearDown()
		rm.Mock.ExpectedCalls = nil
		rm.On("IsRestricted").Return(true)
		f.Withdraw(123)
		amm.AssertNotCalled(t, "Withdraw", 123.0)
		nm.AssertCalled(t, "Notify", "", "Account withdrawal failed: account is restricted")
	})
}

func TestFacade_PrintStatement(t *testing.T) {
	amm := &transactions.AccountManagerMock{}
	amm.On(
		"GetStatement",
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now().Truncate(time.Hour*24),
	).Return(
		11.0,
		12.2,
		[]models.Operation{
			{
				Date:   time.Now().Truncate(time.Hour * 24),
				Amount: 123,
			},
			{
				Date:   time.Now().Truncate(time.Hour * 24),
				Amount: 11,
			},
		},
	)
	nm := &notification.NotifierMock{}
	rm := &restrictions.CheckerMock{}
	f := NewAccountManager(&models.Person{}, amm, rm, nm).(*facade)

	var buf bytes.Buffer
	_, _ = f.PrintStatement(
		&buf,
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now().Truncate(time.Hour*24),
	)

	assert.Equal(t,
		"\tStatement from 2020-01-01 to 2020-02-21\n"+
			"\t\tIn balance: 11.00, Out balance: 12.20, Current balance: 0.00\n"+
			"\t\tOperations:\t\t\tDate: 2020-02-21, Amount: +123.00\n\t\t\tDate: 2020-02-21, Amount: +11.00\n",
		buf.String(),
	)
}

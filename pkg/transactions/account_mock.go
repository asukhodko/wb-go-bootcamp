package transactions

import (
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
)

// AccountManagerMock - мок для интерфейса AccountManager
type AccountManagerMock struct {
	mock.Mock
}

// Deposit имитирует пополнение счёта
func (m *AccountManagerMock) Deposit(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}

// Deposit имитирует снятие со счёта
func (m *AccountManagerMock) Withdraw(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}

// Deposit имитирует получение выписки
func (m *AccountManagerMock) GetStatement(from, to time.Time) (inBal, outBal float64, ops []models.Operation) {
	args := m.Called(from, to)
	inBal = args.Get(0).(float64)
	outBal = args.Get(1).(float64)
	ops = args.Get(2).([]models.Operation)
	return
}

// Deposit имитирует получение остатка
func (m *AccountManagerMock) GetBalance() float64 {
	return 0
}

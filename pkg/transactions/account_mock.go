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
func (m *AccountManagerMock) Deposit(amount float32) error {
	return nil
}

// Deposit имитирует снятие со счёта
func (m *AccountManagerMock) Withdraw(amount float32) error {
	return nil
}

// Deposit имитирует получение выписки
func (m *AccountManagerMock) GetStatement(from, to time.Time) (inBal, outBal float32, ops []models.Operation) {
	inBal = 0
	outBal = 0
	ops = nil
	return
}

// Deposit имитирует получение остатка
func (m *AccountManagerMock) GetBalance() float32 {
	return 0
}

package facade

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// AccountManagerMock мок для интерфейса AccountManager
type AccountManagerMock struct {
	mock.Mock
}

// Deposit имитирует пополнение счёта
func (m *AccountManagerMock) Deposit(amount float64) {
	_ = amount
}

// Deposit имитирует снятие со счёта
func (m *AccountManagerMock) Withdraw(amount float64) {
	_ = amount
}

// Deposit имитирует рапечатку выписки
func (m *AccountManagerMock) PrintStatement(from, to time.Time) {
	_, _ = from, to
}

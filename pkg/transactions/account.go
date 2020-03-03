package transactions

import (
	"errors"
	"time"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
)

// AccountManager предоставляет операции для работы со счётом
type AccountManager interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetStatement(from, to time.Time) (inBal, outBal float64, ops []models.Operation)
	GetBalance() float64
}

type account struct {
	ops     []models.Operation
	balance float64
}

// Deposit пополняет счёт
func (a *account) Deposit(amount float64) (err error) {
	if amount < 0 {
		err = errors.New("deposit: amount out of range")
		return
	}
	a.ops = append(a.ops, models.Operation{
		Date:   time.Now().Truncate(time.Hour * 24),
		Amount: amount,
	})
	a.balance += amount
	return
}

// Withdraw снимает со счёта
func (a *account) Withdraw(amount float64) (err error) {
	if amount < 0 {
		err = errors.New("withdraw: amount out of range")
		return
	}
	if a.balance-amount < 0 {
		err = errors.New("withdraw: insufficient funds")
		return
	}
	a.ops = append(a.ops, models.Operation{
		Date:   time.Now().Truncate(time.Hour * 24),
		Amount: -amount,
	})
	a.balance -= amount
	return
}

// GetStatement возвращает выписку по счёту за период
func (a *account) GetStatement(from, to time.Time) (inBal, outBal float64, ops []models.Operation) {
	for _, op := range a.ops {
		if op.Date.Before(from) {
			inBal += op.Amount
			outBal += op.Amount
		} else if !op.Date.After(to) {
			ops = append(ops, op)
			outBal += op.Amount
		}
	}
	return
}

// GetBalance возвращает текущий остаток
func (a *account) GetBalance() float64 {
	return a.balance
}

// NewAccountManager создаёт новый счёт
func NewAccountManager() AccountManager {
	return &account{
		balance: 0,
	}
}

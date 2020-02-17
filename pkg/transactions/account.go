package transactions

import (
	"errors"
	"time"
)

// AccountManager предоставляет операции для работы со счётом
type AccountManager interface {
	Deposit(amount float32) error
	Withdraw(amount float32) error
	GetStatement(from, to time.Time) (inBal, outBal float32, ops []Operation)
	GetBalance() float32
}

type account struct {
	AccountManager
	ops     []Operation
	balance float32
}

// Deposit пополняет счёт
func (a *account) Deposit(amount float32) error {
	if amount < 0 {
		return errors.New("deposit: amount out of range")
	}
	a.ops = append(a.ops, Operation{
		Date:   time.Now(),
		Amount: amount,
	})
	a.balance += amount
	return nil
}

// Withdraw снимает со счёта
func (a *account) Withdraw(amount float32) error {
	if amount < 0 {
		return errors.New("withdraw: amount out of range")
	}
	if a.balance-amount < 0 {
		return errors.New("withdraw: insufficient funds")
	}
	a.ops = append(a.ops, Operation{
		Date:   time.Now(),
		Amount: -amount,
	})
	a.balance -= amount
	return nil
}

// GetStatement возвращает выписку по счёту за период
func (a *account) GetStatement(from, to time.Time) (inBal, outBal float32, ops []Operation) {
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
func (a *account) GetBalance() float32 {
	return a.balance
}

// NewAccountManager создаёт новый счёт
func NewAccountManager() AccountManager {
	return &account{balance: 0}
}

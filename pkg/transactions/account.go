package transactions

import (
	"errors"
	"time"
)

// Account представляет счёт
type Account struct {
	person  *Person
	ops     []Operation
	balance float32
}

// NewAccount создаёт новый счёт
func NewAccount(person *Person) *Account {
	return &Account{person: person, balance: 0}
}

// Deposit пополняет счёт
func (a *Account) Deposit(amount float32) error {
	if amount < 0 {
		return errors.New("deposit: amount out of range")
	}
	a.ops = append(a.ops, Operation{
		date:   time.Now(),
		amount: amount,
	})
	a.balance += amount
	return nil
}

// Withdraw снимает со счёта
func (a *Account) Withdraw(amount float32) error {
	if amount < 0 {
		return errors.New("withdraw: amount out of range")
	}
	if a.balance-amount < 0 {
		return errors.New("withdraw: insufficient funds")
	}
	a.ops = append(a.ops, Operation{
		date:   time.Now(),
		amount: -amount,
	})
	a.balance -= amount
	return nil
}

// GetStatement возвращает выписку по счёту за период
func (a *Account) GetStatement(from, to time.Time) (inBal, outBal float32, ops []Operation) {
	for _, op := range a.ops {
		if op.date.Before(from) {
			inBal += op.amount
			outBal += op.amount
		} else if !op.date.After(to) {
			ops = append(ops, op)
			outBal += op.amount
		}
	}
	return
}

// GetBalance возвращает текущий остаток
func (a *Account) GetBalance() float32 {
	return a.balance
}

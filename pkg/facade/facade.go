package facade

import (
	"errors"
	"fmt"
	"time"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/notification"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions/restrictions"
)

// AccountManager - фасад для работы со счётом
type AccountManager interface {
	Seed(hasRestrictions bool)
	PrintStatement(from, to time.Time)
	Deposit(amount float32)
	Withdraw(amount float32)
}

type facade struct {
	AccountManager
	person       *transactions.Person
	account      *transactions.Account
	restrictions restrictions.Checker
	notifier     notification.Notifier
}

// Seed заполняет начальными данными
func (f *facade) Seed(hasRestrictions bool) {
	f.restrictions.SetupRestrictions(hasRestrictions)
	_ = f.account.Deposit(1.22)
	_ = f.account.Deposit(5)
	_ = f.account.Deposit(12.8)
	_ = f.account.Withdraw(7)
	_ = f.account.Withdraw(7.5)
	_ = f.account.Deposit(22)
}

// PrintStatement печатает выписку по счёту
func (f *facade) PrintStatement(from, to time.Time) {
	fmt.Printf("\tStatement from %s to %s\n", from.Format("2006-01-02"), to.Format("2006-01-02"))
	inBal, outBal, ops := f.account.GetStatement(from, to)
	fmt.Printf("\t\tIn balance: %.2f, Out balance: %.2f, Current balance: %.2f\n", inBal, outBal, f.account.GetBalance())
	fmt.Println("\t\tOperations:")
	for _, op := range ops {
		fmt.Printf("\t\t\t%s\n", op.String())
	}
}

// Deposit осуществляет пополнение счёта, если нет ограничений, и уведомляет владельца счёта об операции
func (f *facade) Deposit(amount float32) {
	var message string
	var err error
	if f.restrictions.IsRestricted() {
		err = errors.New("Account restricted")
	}
	if err == nil {
		err = f.account.Deposit(amount)
	}
	if err == nil {
		message = fmt.Sprintf("Account refilled by %.2f, balance: %.2f", amount, f.account.GetBalance())
	} else {
		message = fmt.Sprintf("Account not refilled: %s", err.Error())
	}
	f.notifier.Notify(f.person.GetPhoneNumber(), message)
}

// Withdraw осуществляет снятие со счёта, если нет ограничений и достаточно средств, и уведомляет владельца счёта об операции
func (f *facade) Withdraw(amount float32) {
	var message string
	var err error
	if f.restrictions.IsRestricted() {
		err = errors.New("Account restricted")
	}
	if err == nil {
		err = f.account.Withdraw(amount)
	}
	if err == nil {
		message = fmt.Sprintf("Successful withdrawal from the account by %.2f, balance: %.2f", amount, f.account.GetBalance())
	} else {
		message = fmt.Sprintf("Account withdrawal failed: %s", err.Error())
	}
	f.notifier.Notify(f.person.GetPhoneNumber(), message)
}

// NewAccountManager конструирует новый фасад
func NewAccountManager(personName string, phoneNumber string) AccountManager {
	person := transactions.NewPerson(personName, phoneNumber)
	return &facade{
		person:       person,
		account:      transactions.NewAccount(person),
		restrictions: restrictions.NewChecker(),
		notifier:     notification.NewNotifier(),
	}
}

package facade

import (
	"errors"
	"fmt"
	"time"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
)

type accountManager interface {
	Deposit(amount float64) (err error)
	Withdraw(amount float64) (err error)
	GetStatement(from, to time.Time) (inBal, outBal float64, ops []models.Operation)
	GetBalance() float64
}

type checker interface {
	SetupRestrictions(hasRestrictions bool)
	IsRestricted() bool
}

type notifier interface {
	Notify(phoneNumber, message string)
}

// AccountManager - фасад для работы со счётом
type AccountManager interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	PrintStatement(from, to time.Time)
}

type facade struct {
	person       *models.Person
	am           accountManager
	restrictions checker
	notifier     notifier
}

// PrintStatement печатает выписку по счёту
func (f *facade) PrintStatement(from, to time.Time) {
	fmt.Printf("\tStatement from %s to %s\n", from.Format("2006-01-02"), to.Format("2006-01-02"))
	inBal, outBal, ops := f.am.GetStatement(from, to)
	fmt.Printf("\t\tIn balance: %.2f, Out balance: %.2f, Current balance: %.2f\n", inBal, outBal, f.am.GetBalance())
	fmt.Println("\t\tOperations:")
	for _, op := range ops {
		fmt.Printf("\t\t\t%s\n", op.String())
	}
}

// Deposit осуществляет пополнение счёта, если нет ограничений, и уведомляет владельца счёта об операции
func (f *facade) Deposit(amount float64) {
	var (
		message string
		err     error
	)
	if f.restrictions.IsRestricted() {
		err = errors.New("am is restricted")
	}
	if err == nil {
		err = f.am.Deposit(amount)
	}
	if err == nil {
		message = fmt.Sprintf("Account refilled by %.2f, balance: %.2f", amount, f.am.GetBalance())
	} else {
		message = fmt.Sprintf("Account not refilled: %s", err.Error())
	}
	f.notifier.Notify(f.person.PhoneNumber, message)
}

// Withdraw осуществляет снятие со счёта, если нет ограничений и достаточно средств, и уведомляет владельца счёта об операции
func (f *facade) Withdraw(amount float64) {
	var (
		message string
		err     error
	)
	if f.restrictions.IsRestricted() {
		err = errors.New("am is restricted")
	}
	if err == nil {
		err = f.am.Withdraw(amount)
	}
	if err == nil {
		message = fmt.Sprintf("Successful withdrawal from the am by %.2f, balance: %.2f", amount, f.am.GetBalance())
	} else {
		message = fmt.Sprintf("Account withdrawal failed: %s", err.Error())
	}
	f.notifier.Notify(f.person.PhoneNumber, message)
}

// NewAccountManager конструирует новый фасад
func NewAccountManager(person *models.Person, am accountManager, checker checker, notifier notifier) AccountManager {
	return &facade{
		person:       person,
		am:           am,
		restrictions: checker,
		notifier:     notifier,
	}
}

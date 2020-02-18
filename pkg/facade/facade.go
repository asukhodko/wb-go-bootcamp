package facade

import (
	"errors"
	"fmt"
	"time"

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/notification"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions/restrictions"
)

type accountManager interface {
	Deposit(amount float32) (err error)
	Withdraw(amount float32) (err error)
	GetStatement(from, to time.Time) (inBal, outBal float32, ops []models.Operation)
	GetBalance() float32
}

// AccountManager - фасад для работы со счётом
type AccountManager interface {
	Seed(hasRestrictions bool)
	Deposit(amount float32)
	Withdraw(amount float32)
	PrintStatement(from, to time.Time)
}

type facade struct {
	person       *models.Person
	am           accountManager
	restrictions restrictions.Checker
	notifier     notification.Notifier
}

// Seed заполняет начальными данными
func (f *facade) Seed(hasRestrictions bool) {
	f.restrictions.SetupRestrictions(hasRestrictions)
	_ = f.am.Deposit(1.22)
	_ = f.am.Deposit(5)
	_ = f.am.Deposit(12.8)
	_ = f.am.Withdraw(7)
	_ = f.am.Withdraw(7.5)
	_ = f.am.Deposit(22)
}

// PrintStatement печатает выписку по счёту
func (f *facade) PrintStatement(from, to time.Time) {
	fmt.Printf("\tStatement from %s to %s\n", from.Format("2006-01-02"), to.Format("2006-01-02"))
	inBal, outBal, ops := f.am.GetStatement(from, to)
	fmt.Printf("\t\tIn balance: %.2f, Out balance: %.2f, Current balance: %.2f\n", inBal, outBal, f.am.GetBalance())
	fmt.Println("\t\tOperations:")
	for _, op := range ops {
		fmt.Printf("\t\t\tDate: %s, Amount: %+.2f\n", op.Date.Format("2006-01-02"), op.Amount)
	}
}

// Deposit осуществляет пополнение счёта, если нет ограничений, и уведомляет владельца счёта об операции
func (f *facade) Deposit(amount float32) {
	var message string
	var err error
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
func (f *facade) Withdraw(amount float32) {
	var message string
	var err error
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
func NewAccountManager(person *models.Person, am accountManager) AccountManager {
	return &facade{
		person:       person,
		am:           am,
		restrictions: restrictions.NewChecker(),
		notifier:     notification.NewNotifier(),
	}
}

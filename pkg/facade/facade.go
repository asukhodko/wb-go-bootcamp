package facade

import (
	"fmt"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions/restrictions"
	"time"
)

type Facade struct {
	person       *transactions.Person
	account      *transactions.Account
	restrictions *restrictions.AccountRestrictions
}

func NewFacade(personName string) Facade {
	person := transactions.NewPerson(personName)
	return Facade{
		person:       person,
		account:      transactions.NewAccount(person),
		restrictions: restrictions.NewAccountRestrictions(),
	}
}

func (f *Facade) Seed(hasRestrictions bool) {
	f.restrictions.SetupRestrictions(hasRestrictions)
	_ = f.account.Deposit(1.22)
	_ = f.account.Deposit(5)
	_ = f.account.Deposit(12.8)
	_ = f.account.Withdraw(7)
	_ = f.account.Withdraw(7.5)
	_ = f.account.Deposit(22)
}

func (f *Facade) PrintStatement(from, to time.Time) error {
	fmt.Printf("\tStatement from %s to %s\n", from.Format("2006-01-02"), to.Format("2006-01-02"))
	inBal, outBal, ops := f.account.GetStatement(from, to)
	fmt.Printf("\t\tIn balance: %.2f, Out balance: %.2f, Current balance: %.2f\n", inBal, outBal, f.account.GetBalance())
	fmt.Println("\t\tOperations:")
	for _, op := range ops {
		fmt.Printf("\t\t\t%s\n", op.String())
	}
	return nil
}

func (f *Facade) Deposit(amount float32) error {
	return f.account.Deposit(amount)
}

func (f *Facade) Withdraw(amount float32) error {
	return f.account.Withdraw(amount)
}

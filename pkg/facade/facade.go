package facade

import (
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"time"
)

type Facade struct {
	person  *transactions.Person
	account *transactions.Account
}

func NewFacade(personName string) Facade {
	person := transactions.NewPerson(personName)
	return Facade{
		person:  person,
		account: transactions.NewAccount(person),
	}
}

func (f *Facade) Seed() {

}

func (f *Facade) PrintStatement(from, to time.Time) error {
	return nil
}

func (f *Facade) Deposit(amount float32) error {
	return nil
}

func (f *Facade) Withdraw(amount float32) error {
	return nil
}

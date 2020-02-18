package main

import (
	"fmt"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"time"

	// поскольку пакеты в `pkg`, а не в `internal`, приходится указывать весь путь до пакета, а не относительно проекта
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/facade"
)

func main() {
	demoFacadeForPerson("SomePerson", "+79161234567", false)
	demoFacadeForPerson("SomeStrangePerson", "+19993216547", true)
}

func demoFacadeForPerson(personName string, phoneNumber string, hasRestrictions bool) {
	fmt.Printf("Person: %s\n", personName)

	f := facade.NewAccountManager(
		&models.Person{
			Name:        personName,
			PhoneNumber: phoneNumber,
		},
		transactions.NewAccountManager(),
	)
	f.Seed(hasRestrictions)

	f.PrintStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)

	f.Deposit(123)
	f.Withdraw(10)
	f.Withdraw(200)

	f.PrintStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
}

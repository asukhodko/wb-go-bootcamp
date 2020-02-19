package main

import (
	"fmt"
	"time"

	// поскольку пакеты в `pkg`, а не в `internal`, приходится указывать весь путь до пакета, а не относительно проекта

	"github.com/asukhodko/wb-go-bootcamp-1/pkg/facade"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/notification"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/transactions/restrictions"
)

func main() {
	demoFacadeForPerson("SomePerson", "+79161234567", false)
	demoFacadeForPerson("SomeStrangePerson", "+19993216547", true)
}

func demoFacadeForPerson(personName string, phoneNumber string, hasRestrictions bool) {
	fmt.Printf("Person: %s\n", personName)

	am := transactions.NewAccountManager()
	checker := restrictions.NewChecker()
	f := facade.NewAccountManager(
		&models.Person{
			Name:        personName,
			PhoneNumber: phoneNumber,
		},
		am,
		checker,
		notification.NewNotifier(),
	)
	seed(am, checker, hasRestrictions)

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

func seed(am transactions.AccountManager, checker restrictions.Checker, hasRestrictions bool) {
	checker.SetupRestrictions(hasRestrictions)
	_ = am.Deposit(1.22)
	_ = am.Deposit(5)
	_ = am.Deposit(12.8)
	_ = am.Withdraw(7)
	_ = am.Withdraw(7.5)
	_ = am.Deposit(22)
}

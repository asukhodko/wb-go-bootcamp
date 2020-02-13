package main

import (
	"fmt"
	"time"

	// поскольку пакеты в `pkg`, а не в `internal`, приходится указывать весь путь до пакета, а не относительно проекта
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/facade"
)

// Тема: 1. Массивы и срезы
// Задание: Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern  в соответствии с конвенцией

func main() {
	demoFacadeForPerson("SomePerson", "+79161234567", false)
	demoFacadeForPerson("SomeStrangePerson", "+19993216547", true)
}

func demoFacadeForPerson(pesonName string, phoneNumber string, hasRestrictions bool) {
	fmt.Printf("Person: %s\n", pesonName)

	f := facade.NewFacade(pesonName, phoneNumber)
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

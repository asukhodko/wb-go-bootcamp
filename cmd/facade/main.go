package main

import (
	"fmt"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/facade"
	"time"
)

// Тема: 1. Массивы и срезы
// Задание: Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern  в соответствии с конвенцией

func main() {
	demoFacadeForPerson("SomePerson", false)
	demoFacadeForPerson("SomeStrangePerson", true)
}

func demoFacadeForPerson(pesonName string, hasRestrictions bool) {
	fmt.Printf("Person: %s\n", pesonName)

	f := facade.NewFacade(pesonName)
	f.Seed(hasRestrictions)

	err := f.PrintStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
	if err != nil {
		fmt.Printf("f.PrintStatement(1) error: %s\n", err.Error())
	}

	err = f.Deposit(123)
	if err != nil {
		fmt.Printf("f.Deposit error: %s\n", err.Error())
	}

	err = f.Withdraw(10)
	if err != nil {
		fmt.Printf("f.Withdraw(1) error: %s\n", err.Error())
	}

	err = f.Withdraw(200)
	if err != nil {
		fmt.Printf("f.Withdraw(2) error: %s\n", err.Error())
	}

	err = f.PrintStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
	if err != nil {
		fmt.Printf("f.PrintStatement(2) error: %s\n", err.Error())
	}
}

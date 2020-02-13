package main

import (
	"fmt"
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/facade"
	"time"
)

// Тема: 1. Массивы и срезы
// Задание: Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern  в соответствии с конвенцией

func main() {
	demoFacadeForPerson("SomePerson")
	demoFacadeForPerson("SomeStrangePerson")
}

func demoFacadeForPerson(pesonName string) {
	fmt.Printf("Person: %s\n", pesonName)
	f := facade.NewFacade(pesonName)
	f.Seed()

	err := f.PrintStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
	if err != nil {
		fmt.Printf("f.PrintStatement error: %s\n", err.Error())
	}

	err = f.Deposit(123)
	if err != nil {
		fmt.Printf("f.Deposit error: %s\n", err.Error())
	}
	err = f.Withdraw(10)
	if err != nil {
		fmt.Printf("f.Withdraw error: %s\n", err.Error())
	}
}

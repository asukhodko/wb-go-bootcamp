Практическое задание 1 по го.

Тема:
1. Массивы и срезы

Задание:

Реализовать паттерн фасад https://en.wikipedia.org/wiki/Facade_pattern  в соответствии с конвенцией


Вывод программы:
```
Person: SomePerson
	Statement from 2020-01-01 to 2020-02-13
		In balance: 0.00, Out balance: 26.52, Current balance: 26.52
		Operations:
			Date: 2020-02-13, Amount: +1.22
			Date: 2020-02-13, Amount: +5.00
			Date: 2020-02-13, Amount: +12.80
			Date: 2020-02-13, Amount: -7.00
			Date: 2020-02-13, Amount: -7.50
			Date: 2020-02-13, Amount: +22.00
[NOTIFICATION TO +79161234567]: Account refilled by 123.00, balance: 149.52
[NOTIFICATION TO +79161234567]: Successful withdrawal from the account by 10.00, balance: 139.52
[NOTIFICATION TO +79161234567]: Account withdrawal failed: withdraw: insufficient funds
	Statement from 2020-01-01 to 2020-02-13
		In balance: 0.00, Out balance: 139.52, Current balance: 139.52
		Operations:
			Date: 2020-02-13, Amount: +1.22
			Date: 2020-02-13, Amount: +5.00
			Date: 2020-02-13, Amount: +12.80
			Date: 2020-02-13, Amount: -7.00
			Date: 2020-02-13, Amount: -7.50
			Date: 2020-02-13, Amount: +22.00
			Date: 2020-02-13, Amount: +123.00
			Date: 2020-02-13, Amount: -10.00
Person: SomeStrangePerson
	Statement from 2020-01-01 to 2020-02-13
		In balance: 0.00, Out balance: 26.52, Current balance: 26.52
		Operations:
			Date: 2020-02-13, Amount: +1.22
			Date: 2020-02-13, Amount: +5.00
			Date: 2020-02-13, Amount: +12.80
			Date: 2020-02-13, Amount: -7.00
			Date: 2020-02-13, Amount: -7.50
			Date: 2020-02-13, Amount: +22.00
[NOTIFICATION TO +19993216547]: Account not refilled: Account restricted
[NOTIFICATION TO +19993216547]: Account withdrawal failed: Account restricted
[NOTIFICATION TO +19993216547]: Account withdrawal failed: Account restricted
	Statement from 2020-01-01 to 2020-02-13
		In balance: 0.00, Out balance: 26.52, Current balance: 26.52
		Operations:
			Date: 2020-02-13, Amount: +1.22
			Date: 2020-02-13, Amount: +5.00
			Date: 2020-02-13, Amount: +12.80
			Date: 2020-02-13, Amount: -7.00
			Date: 2020-02-13, Amount: -7.50
			Date: 2020-02-13, Amount: +22.00
```

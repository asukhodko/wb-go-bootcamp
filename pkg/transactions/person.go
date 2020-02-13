package transactions

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

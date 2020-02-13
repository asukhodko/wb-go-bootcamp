package transactions

// Person представляет некоторого субъекта, владеющего счётом
type Person struct {
	name string
}

// NewPerson конструирует эклемпляр Person
func NewPerson(name string) *Person {
	return &Person{name: name}
}

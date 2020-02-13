package transactions

// Person представляет некоторого субъекта, владеющего счётом
type Person struct {
	name        string
	phoneNumber string
}

// NewPerson конструирует эклемпляр Person
func NewPerson(name, phoneNumber string) *Person {
	return &Person{name: name, phoneNumber: phoneNumber}
}

// GetPhoneNumber возвращает телефон для отправки смс-уведомлений
func (p *Person) GetPhoneNumber() string {
	return p.phoneNumber
}

package notification

import "fmt"

// Notifier - уведомлятель
type Notifier interface {
	Notify(phoneNumber, message string)
}

type notifier struct {
	Notifier
}

// NewNotifier конструирует уведомлятеля
func NewNotifier() Notifier {
	return &notifier{}
}

// Notify отправляет уведомление
func (n *notifier) Notify(phoneNumber, message string) {
	fmt.Printf("[NOTIFICATION TO %s]: %s\n", phoneNumber, message)
}

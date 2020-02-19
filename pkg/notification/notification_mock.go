package notification

import "github.com/stretchr/testify/mock"

// NotifierMock - мок для интерфейса Notifier
type NotifierMock struct {
	mock.Mock
}

// Notify имитирует отправку уведомления
func (m *NotifierMock) Notify(phoneNumber, message string) {

}

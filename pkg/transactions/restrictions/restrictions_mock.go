package restrictions

import "github.com/stretchr/testify/mock"

// CheckerMock мок для интерфейса Checker
type CheckerMock struct {
	mock.Mock
}

// SetupRestrictions имитирует установку ограничений
func (m *CheckerMock) SetupRestrictions(hasRestrictions bool) {
	hasRestrictions = hasRestrictions
}

// IsRestricted имитирует получение ограничений
func (m *CheckerMock) IsRestricted() bool {
	args := m.Called()
	return args.Bool(0)
}

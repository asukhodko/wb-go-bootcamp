package restrictions

type Checker interface {
	SetupRestrictions(hasRestrictions bool)
	IsRestricted() bool
}

type accountRestrictions struct {
	hasRestrictions bool
}

// SetupRestrictions выполняет конфигурирование ограничений
func (r *accountRestrictions) SetupRestrictions(hasRestrictions bool) {
	r.hasRestrictions = hasRestrictions
}

// IsRestricted возвращает информацию о наличии ограничений
func (r *accountRestrictions) IsRestricted() bool {
	return r.hasRestrictions
}

// NewChecker конструирует экземпляр для Checker
func NewChecker() Checker {
	return &accountRestrictions{}
}

package restrictions

// AccountRestrictions представляет ограничения по счёту
type AccountRestrictions struct {
	hasRestrictions bool
}

// NewAccountRestrictions конструирует экземпляр AccountRestrictions
func NewAccountRestrictions() *AccountRestrictions {
	return &AccountRestrictions{}
}

// SetupRestrictions выполняет конфигурирование ограничений
func (r *AccountRestrictions) SetupRestrictions(hasRestrictions bool) {
	r.hasRestrictions = hasRestrictions
}

// IsRestricted возвращает информацию о наличии ограничений
func (r *AccountRestrictions) IsRestricted() bool {
	return r.hasRestrictions
}

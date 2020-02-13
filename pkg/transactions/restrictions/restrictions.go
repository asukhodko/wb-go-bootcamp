package restrictions

type AccountRestrictions struct {
	hasRestrictions bool
}

func NewAccountRestrictions() *AccountRestrictions {
	return &AccountRestrictions{}
}

func (r *AccountRestrictions) SetupRestrictions(hasRestrictions bool) {
	r.hasRestrictions = hasRestrictions
}

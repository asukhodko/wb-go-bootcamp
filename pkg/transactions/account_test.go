package transactions

import "testing"

func TestAccount_Deposit(t *testing.T) {
	am := NewAccountManager().(*account)
	_ = am.Deposit(10)
	_ = am.Deposit(3.34)
	var expected float32 = 13.34
	if am.balance != expected {
		t.Errorf("Expected %f, got %f", expected, am.balance)
	}
}

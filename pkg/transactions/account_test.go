package transactions

import (
	"github.com/asukhodko/wb-go-bootcamp-1/pkg/models"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAccount_Deposit(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	_ = am.Deposit(3.34)
	expected := 34.34
	got := am.balance
	if got != expected {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

func TestAccount_Withdraw(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	_ = am.Withdraw(3.34)
	expected := 27.66
	got := am.balance
	if got != expected {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

func TestAccount_GetBalance(t *testing.T) {
	am := NewAccountManager().(*account)
	am.balance = 31
	expected := 31.0
	got := am.GetBalance()
	if got != expected {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

func TestAccount_GetStatement(t *testing.T) {
	am := NewAccountManager().(*account)

	_ = am.Deposit(123)
	_ = am.Withdraw(11)

	expectedInBal, expectedOutBal, expectedOps := 0.0, 112.0, []models.Operation{
		{Date: time.Now().Truncate(time.Hour * 24), Amount: 123},
		{Date: time.Now().Truncate(time.Hour * 24), Amount: -11},
	}
	gotInBal, gotOutBal, gotOps := am.GetStatement(
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Now(),
	)
	if gotInBal != expectedInBal {
		t.Errorf("Expected %f, got %f", expectedInBal, gotInBal)
	}
	if gotOutBal != expectedOutBal {
		t.Errorf("Expected %f, got %f", expectedOutBal, gotOutBal)
	}
	if !reflect.DeepEqual(gotOps, expectedOps) {
		t.Errorf("Expected %s, got %s", prettyFormatOps(expectedOps), prettyFormatOps(gotOps))
	}
}

func prettyFormatOps(ops []models.Operation) string {
	b := strings.Builder{}
	b.WriteString("[\n")
	for _, o := range ops {
		b.WriteString("\t")
		b.WriteString(o.String())
		b.WriteString("\n")
	}
	b.WriteString("]")
	return b.String()
}

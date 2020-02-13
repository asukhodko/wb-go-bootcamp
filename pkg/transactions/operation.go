package transactions

import (
	"fmt"
	"time"
)

type Operation struct {
	fmt.Stringer
	date   time.Time
	amount float32
}

func (o *Operation) String() string {
	return fmt.Sprintf("Date: %s, Amount: %+.2f", o.date.Format("2006-01-02"), o.amount)
}

package transactions

import (
	"fmt"
	"time"
)

// Operation представляет операцию по счёту
type Operation struct {
	fmt.Stringer

	date   time.Time
	amount float32
}

// String форматирует представление об операции в виде строки
func (o *Operation) String() string {
	return fmt.Sprintf("Date: %s, Amount: %+.2f", o.date.Format("2006-01-02"), o.amount)
}

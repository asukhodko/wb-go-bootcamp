package models

import (
	"fmt"
	"time"
)

// Operation представляет операцию по счёту
type Operation struct {
	Date   time.Time
	Amount float64
}

func (o *Operation) String() string {
	return fmt.Sprintf("Date: %s, Amount: %+.2f", o.Date.Format("2006-01-02"), o.Amount)
}

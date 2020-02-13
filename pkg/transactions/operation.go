package transactions

import "time"

type Operation struct {
	date   time.Time
	amount float32
}

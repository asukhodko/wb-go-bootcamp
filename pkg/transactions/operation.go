package transactions

import (
	"time"
)

// Operation представляет операцию по счёту
// Структура для данных
type Operation struct {
	Date   time.Time `json:"date"`
	Amount float32   `json:"amount"`
}

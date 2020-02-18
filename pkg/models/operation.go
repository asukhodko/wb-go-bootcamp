package models

import "time"

// Operation представляет операцию по счёту
type Operation struct {
	Date   time.Time
	Amount float32
}

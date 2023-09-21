package model

import "time"

// Date is a struct with day, month and year of the task
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

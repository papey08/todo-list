package model

import "time"

// Date is a struct with day, month and year of the task
type Date struct {
	Day   int
	Month time.Month
	Year  int
}

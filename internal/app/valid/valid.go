package valid

import (
	"errors"
	"time"
	"todo-list/internal/model"
)

const (
	maxTitleLen       = 100
	maxDescriptionLen = 500
)

var (
	noTitle            = errors.New("no title of the task")
	titleTooLong       = errors.New("title of task is very long")
	descriptionTooLong = errors.New("description of task is very long")
	dateInvalid        = errors.New("date is invalid")
	dateExpired        = errors.New("planning date of the task is expired")
)

// isLater checks if given date is later or equal than current date
func isLater(d model.Date) bool {
	year, month, day := time.Now().UTC().Date()
	return d.Year > year || (d.Year == year && d.Month > month) || (d.Year == year && d.Month == month && d.Day >= day)
}

// isLeapYear returns true if year y is leap
func isLeapYear(y int) bool {
	return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
}

// Date returns true if date is valid
func Date(d model.Date) error {
	if d.Month == time.February && d.Day == 29 && isLeapYear(d.Year) {
		return dateInvalid
	} else {
		daysInMonth := map[time.Month]int{
			time.January:   31,
			time.February:  28,
			time.March:     31,
			time.April:     30,
			time.May:       31,
			time.June:      30,
			time.July:      31,
			time.August:    31,
			time.September: 30,
			time.October:   31,
			time.November:  30,
			time.December:  31,
		}
		if d.Year > 0 && d.Day > 0 && d.Day <= daysInMonth[d.Month] {
			return nil
		} else {
			return dateInvalid
		}
	}
}

// TodoTask checks if all fields of task struct are valid
func TodoTask(t model.TodoTask) error {
	errs := make([]error, 0, 3)

	if t.Title == "" { // check if task has a title
		errs = append(errs, noTitle)
	} else if len(t.Title) > maxTitleLen { // check length of the title of task
		errs = append(errs, titleTooLong)
	}

	if len(t.Description) > maxDescriptionLen { // check length of the description of task
		errs = append(errs, descriptionTooLong)
	}

	if err := Date(t.PlanningDate); err != nil { // check if date is valid
		errs = append(errs, err)
	} else if !isLater(t.PlanningDate) { // check if date is expired
		errs = append(errs, dateExpired)
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errors.Join(errs...)
	}
}

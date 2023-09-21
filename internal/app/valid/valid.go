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
	datePassed         = errors.New("planning date of the task is passed")
)

// isLater checks if given date is later or equal than current date
func isLater(d model.Date) bool {
	year, month, day := time.Now().UTC().Date()
	return d.Year > year || (d.Year == year && d.Month > month) || (d.Year == year && d.Month == month && d.Day >= day)
}

// TodoTask checks if all fields of task struct are valid
func TodoTask(t model.TodoTask) error {
	errs := make([]error, 0, 4)

	if t.Title == "" { // check if task has a title
		errs = append(errs, noTitle)
	}

	if len(t.Title) > maxTitleLen { // check length of the title of task
		errs = append(errs, titleTooLong)
	}

	if len(t.Description) > maxDescriptionLen { // check length of the description of task
		errs = append(errs, descriptionTooLong)
	}

	if !isLater(t.PlanningDate) { // check date of the task
		errs = append(errs, datePassed)
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errors.Join(errs...)
	}
}

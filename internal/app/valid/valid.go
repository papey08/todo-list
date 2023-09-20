package valid

import (
	"time"
	"todo-list/internal/model"
)

const (
	maxTitleLen       = 100
	MaxDescriptionLen = 500
)

// IsLater checks if given date is later or equal than current date
func IsLater(d model.Date) bool {
	year, month, day := time.Now().UTC().Date()
	return d.Year > year || (d.Year == year && d.Month > month) || (d.Year == year && d.Month == month && d.Day >= day)
}

// TodoTask checks if all fields of task struct are valid
func TodoTask(t model.TodoTask) bool {
	return len(t.Title) <= maxTitleLen && len(t.Description) <= MaxDescriptionLen && IsLater(t.PlanningDate)
}

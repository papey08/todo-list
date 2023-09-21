package valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-list/internal/model"
)

type TodoTaskTest struct {
	description  string
	givenTask    model.TodoTask
	expectedErrs []error
}

func TestTodoTask(t *testing.T) {
	tests := []TodoTaskTest{
		{
			description: "validation of valid task",
			givenTask: model.TodoTask{
				Title:       "Title of my task",
				Description: "Description of my task",
				PlanningDate: model.Date{
					Year:  2024,
					Month: 1,
					Day:   1,
				},
				Status: false,
			},
			expectedErrs: []error{},
		},
		{
			description: "validation of task without title",
			givenTask: model.TodoTask{
				Title:       "",
				Description: "Description of task without title",
				PlanningDate: model.Date{
					Year:  2024,
					Month: 1,
					Day:   1,
				},
				Status: false,
			},
			expectedErrs: []error{noTitle},
		},
		{
			description: "validation of task with invalid date",
			givenTask: model.TodoTask{
				Title:       "Title of invalid task",
				Description: "Description of invalid task",
				PlanningDate: model.Date{
					Year:  2023,
					Month: 2,
					Day:   31,
				},
				Status: false,
			},
			expectedErrs: []error{dateInvalid},
		},
		{
			description: "validation of task with expired date",
			givenTask: model.TodoTask{
				Title:       "Title of expired task",
				Description: "Description of expired task",
				PlanningDate: model.Date{
					Year:  2023,
					Month: 1,
					Day:   1,
				},
				Status: false,
			},
			expectedErrs: []error{dateExpired},
		},
		{
			description: "validation of task with very long title",
			givenTask: model.TodoTask{
				Title:       "V5ZidDlMxou0aJaQf1VhBgWD9AMxFlF3ChnpK6av3YPFkIhzYULJq2gG8zM6A00Bi9yla5bNZe1oUpH0ixhFmNnPvG67uxx306RFAMrYqBL9PuQFq4LjG6chDTT0GGvT",
				Description: "Description of task with very long title",
				PlanningDate: model.Date{
					Year:  2024,
					Month: 1,
					Day:   1,
				},
				Status: false,
			},
			expectedErrs: []error{titleTooLong},
		},
		{
			description: "validation of expired task without title",
			givenTask: model.TodoTask{
				Title:       "",
				Description: "Description of expired task without title",
				PlanningDate: model.Date{
					Year:  2023,
					Month: 1,
					Day:   1,
				},
				Status: false,
			},
			expectedErrs: []error{noTitle, dateExpired},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			givenErr := TodoTask(test.givenTask)
			for _, err := range test.expectedErrs {
				assert.ErrorIs(t, givenErr, err)
			}
		})
	}
}

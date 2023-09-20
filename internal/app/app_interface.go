package app

import (
	"context"
	"todo-list/internal/model"
)

type App interface {
	TaskRepo
}

type TaskRepo interface {
	// AddTask adds task to database
	AddTask(ctx context.Context, t model.TodoTask) (model.TodoTask, error)

	// GetTaskById searches task in database with given id
	GetTaskById(ctx context.Context, id int) (model.TodoTask, error)

	// GetTaskByText returns slice of tasks with given text in title or description
	GetTaskByText(ctx context.Context, text string) ([]model.TodoTask, error)

	// UpdateTask updates fields of task with given id
	UpdateTask(ctx context.Context, id int, t model.TodoTask) (model.TodoTask, error)

	// DeleteTask deletes task with given id from database
	DeleteTask(ctx context.Context, id int) error

	// GetTasksByStatus returns slice of tasks filtered by status with pagination
	GetTasksByStatus(ctx context.Context, status bool, limit int, offset int) ([]model.TodoTask, error)

	// GetTasksByDate returns slice of tasks filtered by planning date with pagination
	GetTasksByDate(ctx context.Context, date model.Date, limit int, offset int) ([]model.TodoTask, error)
}

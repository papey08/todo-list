package app

import (
	"context"
	"todo-list/internal/model"
)

type app struct {
	TaskRepo
}

func (a *app) AddTask(ctx context.Context, t model.TodoTask) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (a *app) GetTaskById(ctx context.Context, id int) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (a *app) GetTaskByText(ctx context.Context, text string) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

func (a *app) UpdateTask(ctx context.Context, id int, t model.TodoTask) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (a *app) DeleteTask(ctx context.Context, id int) error {
	// TODO: implement
	return nil
}

func (a *app) GetTasksByStatus(ctx context.Context, status bool, limit int, offset int) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

func (a *app) GetTasksByDate(ctx context.Context, date model.Date, limit int, offset int) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

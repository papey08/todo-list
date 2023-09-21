package app

import (
	"context"
	"errors"
	"todo-list/internal/app/valid"
	"todo-list/internal/model"
)

type app struct {
	TaskRepo
}

func (a *app) AddTask(ctx context.Context, t model.TodoTask) (model.TodoTask, error) {
	if err := valid.TodoTask(t); err != nil {
		return model.TodoTask{}, errors.Join(model.ErrInvalidTask, err)
	}
	return a.TaskRepo.AddTask(ctx, t)
}

func (a *app) GetTaskById(ctx context.Context, id int) (model.TodoTask, error) {
	return a.TaskRepo.GetTaskById(ctx, id)
}

func (a *app) GetTaskByText(ctx context.Context, text string) ([]model.TodoTask, error) {
	return a.TaskRepo.GetTaskByText(ctx, text)
}

func (a *app) UpdateTask(ctx context.Context, id int, t model.TodoTask) (model.TodoTask, error) {
	if err := valid.TodoTask(t); err != nil {
		return model.TodoTask{}, errors.Join(model.ErrInvalidTask, err)
	}
	return a.TaskRepo.UpdateTask(ctx, id, t)
}

func (a *app) DeleteTask(ctx context.Context, id int) error {
	return a.TaskRepo.DeleteTask(ctx, id)
}

func (a *app) GetTasksByStatus(ctx context.Context, status bool, limit int, offset int) ([]model.TodoTask, error) {
	return a.TaskRepo.GetTasksByStatus(ctx, status, limit, offset)
}

func (a *app) GetTasksByDate(ctx context.Context, date model.Date, limit int, offset int) ([]model.TodoTask, error) {
	return a.TaskRepo.GetTasksByDate(ctx, date, limit, offset)
}

func New(tr TaskRepo) App {
	return &app{
		TaskRepo: tr,
	}
}

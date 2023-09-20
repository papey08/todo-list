package repo

import (
	"context"
	"todo-list/internal/model"
)

type Repo struct {
	// TODO: implement
}

func (r *Repo) AddTask(ctx context.Context, t model.TodoTask) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (r *Repo) GetTaskById(ctx context.Context, id int) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (r *Repo) GetTaskByText(ctx context.Context, text string) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

func (r *Repo) UpdateTask(ctx context.Context, id int, t model.TodoTask) (model.TodoTask, error) {
	// TODO: implement
	return model.TodoTask{}, nil
}

func (r *Repo) DeleteTask(ctx context.Context, id int) error {
	// TODO: implement
	return nil
}

func (r *Repo) GetTasksByStatus(ctx context.Context, status bool, limit int, offset int) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

func (r *Repo) GetTasksByDate(ctx context.Context, date model.Date, limit int, offset int) ([]model.TodoTask, error) {
	// TODO: implement
	return nil, nil
}

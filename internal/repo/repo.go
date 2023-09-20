package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
	"todo-list/internal/app"
	"todo-list/internal/model"
)

const (
	addTaskQuery = `
		INSERT INTO tasks (title, description, planning_date, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;`

	getTaskByIdQuery = `
		SELECT * FROM tasks
		WHERE id = $1;`

	getTaskByTextQuery = `
		SELECT * FROM tasks
		WHERE (title LIKE $1 OR description LIKE $1);`

	updateTaskQuery = `
		UPDATE tasks
		SET title = $2,
		    description = $3,
		    planning_date = $4,
		    status = $5
		WHERE id = $1;`

	deleteTaskQuery = `
		DELETE FROM tasks
		WHERE id = $1;`

	getTasksByStatusQuery = `
		SELECT * FROM tasks
		WHERE status = $1
		OFFSET $2 LIMIT $3;`

	getTasksByDateQuery = `
		SELECT * FROM tasks
		WHERE planning_date = $1
		OFFSET $2 LIMIT $3;`
)

type repo struct {
	pgx.Conn
}

func (r *repo) AddTask(ctx context.Context, t model.TodoTask) (model.TodoTask, error) {
	var id int
	err := r.QueryRow(ctx, addTaskQuery,
		t.Title,
		t.Description,
		fmt.Sprintf("%d-%d-%d", t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day),
		t.Status).Scan(&id)
	if err != nil {
		return model.TodoTask{}, errors.Join(model.ErrTaskRepo, err)
	}
	t.Id = id
	return t, nil
}

func (r *repo) GetTaskById(ctx context.Context, id int) (model.TodoTask, error) {
	row := r.QueryRow(ctx, getTaskByIdQuery, id)
	var t model.TodoTask
	var d time.Time
	if err := row.Scan(&t.Id, &t.Title, &t.Description, &d, &t.Status); errors.Is(err, pgx.ErrNoRows) {
		return model.TodoTask{}, model.ErrTaskNotFound
	} else if err != nil {
		return model.TodoTask{}, errors.Join(model.ErrTaskRepo, err)
	} else {
		t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day = d.UTC().Date()
		return t, nil
	}
}

func (r *repo) GetTaskByText(ctx context.Context, text string) ([]model.TodoTask, error) {
	rows, err := r.Query(ctx, getTaskByTextQuery, text+"%")
	if err != nil {
		return nil, errors.Join(model.ErrTaskRepo, err)
	}
	defer rows.Close()

	tasks := make([]model.TodoTask, 0)
	for rows.Next() {
		var t model.TodoTask
		var d time.Time
		_ = rows.Scan(&t.Id, &t.Title, &t.Description, &d, &t.Status)
		t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day = d.UTC().Date()
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *repo) UpdateTask(ctx context.Context, id int, t model.TodoTask) (model.TodoTask, error) {
	e, err := r.Exec(ctx, updateTaskQuery,
		id,
		t.Title,
		t.Description,
		fmt.Sprintf("%d-%d-%d", t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day),
		t.Status)
	if err != nil {
		return model.TodoTask{}, errors.Join(model.ErrTaskRepo, err)
	} else if e.RowsAffected() == 0 {
		return model.TodoTask{}, model.ErrTaskNotFound
	} else {
		return model.TodoTask{
			Id:           id,
			Title:        t.Title,
			Description:  t.Description,
			PlanningDate: t.PlanningDate,
			Status:       t.Status,
		}, nil
	}
}

func (r *repo) DeleteTask(ctx context.Context, id int) error {
	e, err := r.Exec(ctx, deleteTaskQuery, id)
	if err != nil {
		return errors.Join(model.ErrTaskRepo, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrTaskNotFound
	} else {
		return nil
	}
}

func (r *repo) GetTasksByStatus(ctx context.Context, status bool, offset int, limit int) ([]model.TodoTask, error) {
	rows, err := r.Query(ctx, getTasksByStatusQuery, status, offset, limit)
	if err != nil {
		return nil, errors.Join(model.ErrTaskRepo, err)
	}
	defer rows.Close()

	tasks := make([]model.TodoTask, 0)
	for rows.Next() {
		var t model.TodoTask
		var d time.Time
		_ = rows.Scan(&t.Id, &t.Title, &t.Description, &d, &t.Status)
		t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day = d.UTC().Date()
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *repo) GetTasksByDate(ctx context.Context, date model.Date, offset int, limit int) ([]model.TodoTask, error) {
	rows, err := r.Query(ctx, getTasksByStatusQuery,
		fmt.Sprintf("%d-%d-%d", date.Year, date.Month, date.Day),
		offset, limit)
	if err != nil {
		return nil, errors.Join(model.ErrTaskRepo, err)
	}
	defer rows.Close()

	tasks := make([]model.TodoTask, 0)
	for rows.Next() {
		var t model.TodoTask
		var d time.Time
		_ = rows.Scan(&t.Id, &t.Title, &t.Description, &d, &t.Status)
		t.PlanningDate.Year, t.PlanningDate.Month, t.PlanningDate.Day = d.UTC().Date()
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func New(conn *pgx.Conn) app.TaskRepo {
	return &repo{
		Conn: *conn,
	}
}
